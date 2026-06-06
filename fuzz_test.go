package stemmers

import (
	"bufio"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"unicode/utf8"
)

// allAlgorithms is every algorithm New accepts, used to exercise the whole port.
var allAlgorithms = []Algorithm{
	Arabic, Danish, Dutch, English, Finnish, French, German, Greek, Hungarian,
	Italian, Norwegian, Portuguese, Romanian, Russian, Spanish, Swedish, Tamil,
	Turkish,
}

// algoNames is the Snowball language name for each algorithm, matching the
// rustgen oracle's CLI names and used to label subtests/benchmarks.
var algoNames = map[Algorithm]string{
	Arabic: "arabic", Danish: "danish", Dutch: "dutch", English: "english",
	Finnish: "finnish", French: "french", German: "german", Greek: "greek",
	Hungarian: "hungarian", Italian: "italian", Norwegian: "norwegian",
	Portuguese: "portuguese", Romanian: "romanian", Russian: "russian",
	Spanish: "spanish", Swedish: "swedish", Tamil: "tamil", Turkish: "turkish",
}

// algoName returns the Snowball language name for a.
func algoName(a Algorithm) string { return algoNames[a] }

// FuzzStem checks robustness: no algorithm may panic on arbitrary input, and a
// stem must remain valid UTF-8. (Length is not invariant — Turkish, for one,
// appends vowels — so byte-for-byte output is verified by TestVectors and the
// differential TestRustgenParity, not here.) It is fuzz coverage on top of the
// fixed vectors.
func FuzzStem(f *testing.F) {
	for _, w := range []string{"", "a", "running", "fruitlessly", "naïveté", "Ω", "بالكتاب", "größer"} {
		f.Add(w)
	}
	// Seed from the committed corpus so the fuzz reflects real words too.
	for _, w := range firstN(readSeed(f), 500) {
		f.Add(w)
	}
	stemmers := make([]*Stemmer, len(allAlgorithms))
	for i, a := range allAlgorithms {
		stemmers[i] = New(a)
	}
	f.Fuzz(func(t *testing.T, word string) {
		// Snowball's contract is valid (lowercased) UTF-8, which Rust's &str
		// guarantees by construction; Go strings can hold arbitrary bytes. Every
		// input must still stem without panicking, but the valid-UTF-8-out
		// invariant only holds for in-contract (valid UTF-8) input.
		inputValid := utf8.ValidString(word)
		for i, s := range stemmers {
			got := s.Stem(word)
			if inputValid && !utf8.ValidString(got) {
				t.Fatalf("algo %d Stem(%q) = invalid UTF-8 %q", allAlgorithms[i], word, got)
			}
		}
	})
}

// TestRustgenParity is the cross-language differential check against the live
// rust-stemmers oracle: it stems the corpus with both Go and rustgen and asserts
// they agree for every algorithm. It runs only when the rustgen binary is built
// (`cargo build --release` in rustgen/), so the pure-Go test stays hermetic; CI's
// oracle job builds rustgen and runs it. It uses a larger streamed bench.corpus
// when present, else the committed corpus.
func TestRustgenParity(t *testing.T) {
	bin := filepath.Join("rustgen", "target", "release", "rustgen")
	if _, err := os.Stat(bin); err != nil {
		t.Skip("rustgen not built (cargo build --release in rustgen/) — skipping live oracle parity")
	}
	words, err := readLinesFile(filepath.Join("testdata", "bench.corpus"))
	if err != nil || len(words) == 0 {
		words, err = readLinesFile(filepath.Join("testdata", "corpus.txt"))
		if err != nil {
			t.Fatalf("read corpus: %v", err)
		}
	}
	abs, err := filepath.Abs(bin)
	if err != nil {
		t.Fatalf("abs: %v", err)
	}
	input := strings.Join(words, "\n") + "\n"
	for _, a := range allAlgorithms {
		t.Run(algoName(a), func(t *testing.T) {
			cmd := exec.Command(abs, algoName(a))
			cmd.Stdin = strings.NewReader(input)
			out, err := cmd.Output()
			if err != nil {
				t.Fatalf("rustgen %s: %v", algoName(a), err)
			}
			want := strings.Split(strings.TrimRight(string(out), "\n"), "\n")
			if len(want) != len(words) {
				t.Fatalf("%s: rustgen returned %d lines for %d words", algoName(a), len(want), len(words))
			}
			s := New(a)
			mismatch := 0
			for i, w := range words {
				if got := s.Stem(w); got != want[i] {
					mismatch++
					if mismatch <= 10 {
						t.Errorf("%s Stem(%q) = %q, rustgen = %q", algoName(a), w, got, want[i])
					}
				}
			}
			if mismatch > 0 {
				t.Fatalf("%s: %d/%d stems diverged from rustgen", algoName(a), mismatch, len(words))
			}
		})
	}
}

// readSeed loads the committed corpus for fuzz seeding, tolerating its absence.
func readSeed(f *testing.F) []string {
	f.Helper()
	w, err := readLinesFile(filepath.Join("testdata", "corpus.txt"))
	if err != nil {
		return nil
	}
	return w
}

// firstN returns up to n elements of s.
func firstN(s []string, n int) []string {
	if len(s) < n {
		return s
	}
	return s[:n]
}

// readLinesFile reads a file into lines without trailing newlines.
func readLinesFile(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var lines []string
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 0, 64*1024), 4*1024*1024)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines, sc.Err()
}
