// Command oacorpus streams the public OpenAlex Works snapshot and derives a
// realistic word corpus for the go-stemmers benchmark, without vendoring any
// source data into the repo.
//
// Per Work it builds the indexed text (title + reconstructed abstract + authors
// + venue), tokenizes it into lowercased Unicode letter/digit terms, and
// accumulates the distinct vocabulary until a target budget is reached. The
// distinct terms are written one per line (sorted) — the input shape the bench
// (and ad-hoc rustgen parity runs) consume.
//
// The OpenAlex Works snapshot is public (CC0) and reachable over HTTPS with no
// AWS credentials; this tool streams each gzip JSON Lines part directly and
// discards the raw works, keeping only the derived corpus.
//
// Usage:
//
//	go run ./cmd/oacorpus -budget 200000 -out testdata/openalex.corpus
//	go run ./cmd/oacorpus -budget 50000 -maxworks 50000   # quick smoke corpus
package main

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
	"unicode"
)

// httpsBase is the public OpenAlex S3 bucket reached over HTTPS (no credentials).
const httpsBase = "https://openalex.s3.amazonaws.com/"

// s3Prefix is the scheme+bucket prefix manifest entry URLs carry; it is rewritten
// to httpsBase so each part streams without the AWS SDK.
const s3Prefix = "s3://openalex/"

// abstractCharCap bounds the reconstructed abstract length so a few very long
// works do not dominate the vocabulary.
const abstractCharCap = 2000

// work is the subset of the OpenAlex Work JSON the text pipeline reads.
type work struct {
	DisplayName string           `json:"display_name"`
	AbstractIdx map[string][]int `json:"abstract_inverted_index"`
	Authorships []struct {
		Author struct {
			DisplayName string `json:"display_name"`
		} `json:"author"`
	} `json:"authorships"`
	PrimaryLocation struct {
		Source struct {
			DisplayName string `json:"display_name"`
		} `json:"source"`
	} `json:"primary_location"`
}

// manifest is the OpenAlex snapshot manifest: one entry per gzip part file.
type manifest struct {
	Entries []struct {
		URL string `json:"url"`
	} `json:"entries"`
}

func main() {
	manifestURL := flag.String("manifest", httpsBase+"data/works/manifest", "OpenAlex Works snapshot manifest URL")
	out := flag.String("out", "testdata/openalex.corpus", "output corpus file (one term per line, sorted)")
	budget := flag.Int("budget", 200_000, "target number of distinct terms; streaming stops once reached")
	maxWorks := flag.Int("maxworks", 0, "safety cap on works processed (0 = unlimited)")
	minLen := flag.Int("minlen", 2, "minimum term byte length to keep")
	progress := flag.Duration("progress", 5*time.Second, "progress log interval")
	flag.Parse()

	parts, err := fetchManifest(*manifestURL)
	if err != nil {
		log.Fatalf("manifest: %v", err)
	}
	log.Printf("manifest: %d part files", len(parts))

	terms := make(map[string]struct{}, *budget+*budget/4)
	var works, downloaded int64
	start := time.Now()
	last := start

	emit := func(term string) { terms[term] = struct{}{} }

	for pi, url := range parts {
		if len(terms) >= *budget {
			break
		}
		if *maxWorks > 0 && works >= int64(*maxWorks) {
			break
		}
		n, err := streamPart(url, func(w *work) bool {
			works++
			tokenize(buildText(w), *minLen, emit)
			if time.Since(last) >= *progress {
				last = time.Now()
				log.Printf("part %d/%d  works=%d  distinct=%d  (%.0f works/s)",
					pi+1, len(parts), works, len(terms),
					float64(works)/time.Since(start).Seconds())
			}
			return len(terms) < *budget && (*maxWorks == 0 || works < int64(*maxWorks))
		})
		downloaded += n
		if err != nil {
			log.Printf("part %d (%s): %v — continuing", pi+1, url, err)
		}
	}

	log.Printf("streamed %d works, %.1f MiB, %d distinct terms in %s",
		works, float64(downloaded)/(1<<20), len(terms), time.Since(start).Round(time.Second))

	if err := writeCorpus(*out, terms); err != nil {
		log.Fatalf("write corpus: %v", err)
	}
	log.Printf("wrote %s (%d terms)", *out, len(terms))
}

// fetchManifest downloads the snapshot manifest and returns the part URLs in
// listed order, each rewritten from s3:// to HTTPS.
func fetchManifest(url string) ([]string, error) {
	resp, err := httpGet(url)
	if err != nil {
		return nil, err
	}
	defer resp.Close()
	var m manifest
	if err := json.NewDecoder(resp).Decode(&m); err != nil {
		return nil, err
	}
	parts := make([]string, 0, len(m.Entries))
	for _, e := range m.Entries {
		parts = append(parts, httpsBase+strings.TrimPrefix(e.URL, s3Prefix))
	}
	return parts, nil
}

// streamPart downloads one gzip JSON Lines part and calls fn for each decoded
// work, stopping early when fn returns false. It returns compressed bytes read.
func streamPart(url string, fn func(*work) bool) (int64, error) {
	resp, err := httpGet(url)
	if err != nil {
		return 0, err
	}
	defer resp.Close()
	counter := &countingReader{r: resp}
	gz, err := gzip.NewReader(counter)
	if err != nil {
		return counter.n, err
	}
	defer gz.Close()

	sc := bufio.NewScanner(gz)
	sc.Buffer(make([]byte, 0, 1<<20), 64<<20)
	for sc.Scan() {
		line := sc.Bytes()
		if len(line) == 0 {
			continue
		}
		var w work
		if json.Unmarshal(line, &w) != nil || w.DisplayName == "" {
			continue
		}
		if !fn(&w) {
			return counter.n, nil
		}
	}
	return counter.n, sc.Err()
}

// buildText assembles a work's indexed text: title + abstract + authors + venue.
func buildText(w *work) string {
	var sb strings.Builder
	sb.WriteString(w.DisplayName)
	appendField(&sb, reconstructAbstract(w.AbstractIdx))
	appendField(&sb, authorNames(w))
	appendField(&sb, w.PrimaryLocation.Source.DisplayName)
	return sb.String()
}

// appendField appends " " + s to b when s is non-empty.
func appendField(b *strings.Builder, s string) {
	if s != "" {
		b.WriteByte(' ')
		b.WriteString(s)
	}
}

// authorNames joins the work's author display names with "; ".
func authorNames(w *work) string {
	names := make([]string, 0, len(w.Authorships))
	for _, a := range w.Authorships {
		if a.Author.DisplayName != "" {
			names = append(names, a.Author.DisplayName)
		}
	}
	return strings.Join(names, "; ")
}

// reconstructAbstract rebuilds abstract text from an OpenAlex
// abstract_inverted_index (word -> positions), capped at abstractCharCap bytes.
func reconstructAbstract(idx map[string][]int) string {
	if len(idx) == 0 {
		return ""
	}
	maxPos := -1
	for _, ps := range idx {
		for _, p := range ps {
			if p > maxPos {
				maxPos = p
			}
		}
	}
	if maxPos < 0 {
		return ""
	}
	words := make([]string, maxPos+1)
	for word, ps := range idx {
		for _, p := range ps {
			if p >= 0 && p <= maxPos {
				words[p] = word
			}
		}
	}
	abstract := strings.Join(words, " ")
	if len(abstract) > abstractCharCap {
		abstract = abstract[:abstractCharCap]
	}
	return abstract
}

// tokenize splits text into lowercased terms of Unicode letters and digits,
// emitting each term of at least minLen bytes.
func tokenize(text string, minLen int, emit func(string)) {
	var sb strings.Builder
	flush := func() {
		if sb.Len() >= minLen {
			emit(sb.String())
		}
		sb.Reset()
	}
	for _, r := range text {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			sb.WriteRune(unicode.ToLower(r))
		} else {
			flush()
		}
	}
	flush()
}

// writeCorpus writes the distinct terms one per line, sorted bytewise.
func writeCorpus(path string, terms map[string]struct{}) error {
	sorted := make([]string, 0, len(terms))
	for t := range terms {
		sorted = append(sorted, t)
	}
	sort.Strings(sorted)

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	w := bufio.NewWriterSize(f, 1<<20)
	for _, t := range sorted {
		if _, err := fmt.Fprintln(w, t); err != nil {
			f.Close()
			return err
		}
	}
	if err := w.Flush(); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}

// httpGet issues a GET with a couple of retries and returns the response body.
func httpGet(url string) (io.ReadCloser, error) {
	var lastErr error
	for attempt := 0; attempt < 3; attempt++ {
		if attempt > 0 {
			time.Sleep(time.Duration(attempt) * time.Second)
		}
		resp, err := http.Get(url)
		if err != nil {
			lastErr = err
			continue
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			lastErr = fmt.Errorf("GET %s: status %s", url, resp.Status)
			continue
		}
		return resp.Body, nil
	}
	return nil, lastErr
}

// countingReader tallies the bytes read through it.
type countingReader struct {
	r io.Reader
	n int64
}

func (c *countingReader) Read(p []byte) (int, error) {
	n, err := c.r.Read(p)
	c.n += int64(n)
	return n, err
}
