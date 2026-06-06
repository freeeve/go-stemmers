package stemmers

import (
	"bufio"
	"os"
	"path/filepath"
	"testing"
)

// vectorCase pairs an algorithm with its canonical Snowball vocabulary and the
// expected stems, both vendored under testdata as one word per line.
type vectorCase struct {
	algo Algorithm
	voc  string
	res  string
}

// vectorCases lists the committed parity oracles. Each voc/res pair is a
// vocabulary for that language; the expected stems are byte-for-byte what
// rust-stemmers 1.2.0 produces (see testdata/NOTICE.md for provenance).
var vectorCases = []vectorCase{
	{Arabic, "voc_ar.txt", "res_ar.txt"},
	{Danish, "voc_da.txt", "res_da.txt"},
	{Dutch, "voc_nl.txt", "res_nl.txt"},
	{English, "voc_en.txt", "res_en.txt"},
	{Finnish, "voc_fi.txt", "res_fi.txt"},
	{French, "voc_fr.txt", "res_fr.txt"},
	{German, "voc_ger.txt", "res_ger.txt"},
	{Greek, "voc_el.txt", "res_el.txt"},
	{Hungarian, "voc_hu.txt", "res_hu.txt"},
	{Italian, "voc_it.txt", "res_it.txt"},
	{Norwegian, "voc_no.txt", "res_no.txt"},
	{Portuguese, "voc_pt.txt", "res_pt.txt"},
	{Romanian, "voc_ro.txt", "res_ro.txt"},
	{Russian, "voc_ru.txt", "res_ru.txt"},
	{Spanish, "voc_es.txt", "res_es.txt"},
	{Swedish, "voc_sv.txt", "res_sv.txt"},
	{Tamil, "voc_ta.txt", "res_ta.txt"},
	{Turkish, "voc_tr.txt", "res_tr.txt"},
}

// TestVectors is the conformance teeth: every word in each vendored vocabulary
// must stem to exactly the expected output. A single divergence fails the build,
// which is the build/query-symmetry guarantee a stemmed term index depends on.
func TestVectors(t *testing.T) {
	for _, tc := range vectorCases {
		t.Run(tc.voc, func(t *testing.T) {
			voc := readLines(t, filepath.Join("testdata", tc.voc))
			res := readLines(t, filepath.Join("testdata", tc.res))
			if len(voc) != len(res) {
				t.Fatalf("%s has %d lines but %s has %d", tc.voc, len(voc), tc.res, len(res))
			}
			s := New(tc.algo)
			mismatch := 0
			for i := range voc {
				got := s.Stem(voc[i])
				if got != res[i] {
					mismatch++
					if mismatch <= 20 {
						t.Errorf("Stem(%q) = %q, want %q", voc[i], got, res[i])
					}
				}
			}
			if mismatch > 0 {
				t.Fatalf("%d/%d stems diverged from %s", mismatch, len(voc), tc.res)
			}
		})
	}
}

// TestEnglishKnown spot-checks representative English stems independent of the
// vendored vectors: each step's rewrite, an exception, and invariants.
func TestEnglishKnown(t *testing.T) {
	s := New(English)
	cases := map[string]string{
		"fruitlessly":  "fruitless", // step 2 (-li) + step 4
		"consign":      "consign",   // unchanged
		"consigned":    "consign",   // step 1b
		"consignment":  "consign",   // step 4 (-ment)
		"generously":   "generous",  // a0 prefix region + -li
		"national":     "nation",    // step 2 (-ational -> -ation -> step 3/4)
		"sky":          "sky",       // exception1 invariant
		"skis":         "ski",       // exception1 rewrite
		"dying":        "die",       // exception1 rewrite
		"early":        "earli",     // exception1 rewrite
		"news":         "news",      // exception1 invariant
		"succeed":      "succeed",   // exception2 invariant
		"herring":      "herring",   // exception2 invariant
		"":             "",          // empty
		"y":            "y",         // too short to stem
		"agreed":       "agre",      // step 1b (-eed in R1) -> agree -> step 5 e? -> agre
		"feed":         "feed",      // -eed not in R1, unchanged
		"organization": "organ",     // step 2 (-ization) + step 4
	}
	for in, want := range cases {
		if got := s.Stem(in); got != want {
			t.Errorf("Stem(%q) = %q, want %q", in, got, want)
		}
	}
}

// readLines reads a file into a slice of lines (without trailing newlines).
func readLines(t *testing.T, path string) []string {
	t.Helper()
	f, err := os.Open(path)
	if err != nil {
		t.Fatalf("open %s: %v", path, err)
	}
	defer f.Close()
	var lines []string
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	if err := sc.Err(); err != nil {
		t.Fatalf("scan %s: %v", path, err)
	}
	return lines
}
