package stemmers

import (
	"path/filepath"
	"testing"
)

// benchAlgoFiles maps each algorithm to the vocabulary the benchmark stems when
// no streamed OpenAlex corpus is present.
var benchAlgoFiles = map[Algorithm]string{
	Arabic: "voc_ar.txt", Danish: "voc_da.txt", Dutch: "voc_nl.txt",
	English: "voc_en.txt", Finnish: "voc_fi.txt", French: "voc_fr.txt",
	German: "voc_ger.txt", Greek: "voc_el.txt", Hungarian: "voc_hu.txt",
	Italian: "voc_it.txt", Norwegian: "voc_no.txt", Portuguese: "voc_pt.txt",
	Romanian: "voc_ro.txt", Russian: "voc_ru.txt", Spanish: "voc_es.txt",
	Swedish: "voc_sv.txt", Tamil: "voc_ta.txt", Turkish: "voc_tr.txt",
}

// BenchmarkEnglish stems the OpenAlex corpus if present (go run ./cmd/oacorpus),
// otherwise the English vocabulary, reporting words/sec. It is the headline
// throughput number for comparison against rust-stemmers.
func BenchmarkEnglish(b *testing.B) {
	words := loadBenchCorpus(b, English)
	s := New(English)
	b.ReportAllocs()
	b.ResetTimer()
	var n int
	for i := 0; i < b.N; i++ {
		_ = s.Stem(words[n])
		n++
		if n == len(words) {
			n = 0
		}
	}
	b.StopTimer()
	b.ReportMetric(float64(b.N)/b.Elapsed().Seconds(), "words/s")
}

// BenchmarkAll measures per-algorithm throughput over each language's vocabulary,
// so a regression in any port shows up.
func BenchmarkAll(b *testing.B) {
	for _, a := range allAlgorithms {
		a := a
		b.Run(benchAlgoFiles[a], func(b *testing.B) {
			words := loadVoc(b, benchAlgoFiles[a])
			s := New(a)
			b.ResetTimer()
			var n int
			for i := 0; i < b.N; i++ {
				_ = s.Stem(words[n])
				n++
				if n == len(words) {
					n = 0
				}
			}
			b.StopTimer()
			b.ReportMetric(float64(b.N)/b.Elapsed().Seconds(), "words/s")
		})
	}
}

// loadBenchCorpus prefers a streamed OpenAlex corpus (testdata/openalex.corpus),
// falling back to the algorithm's vocabulary so the benchmark always has input.
func loadBenchCorpus(b *testing.B, a Algorithm) []string {
	b.Helper()
	if words, err := readLinesFile(filepath.Join("testdata", "openalex.corpus")); err == nil && len(words) > 0 {
		return words
	}
	return loadVoc(b, benchAlgoFiles[a])
}

// loadVoc loads a vocabulary file from testdata, failing the benchmark if absent.
func loadVoc(b *testing.B, name string) []string {
	b.Helper()
	words, err := readLinesFile(filepath.Join("testdata", name))
	if err != nil {
		b.Fatalf("read %s: %v", name, err)
	}
	if len(words) == 0 {
		b.Fatalf("%s is empty", name)
	}
	return words
}
