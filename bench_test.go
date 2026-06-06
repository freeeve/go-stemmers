package stemmers

import (
	"path/filepath"
	"testing"
)

// benchCorpus returns the words to benchmark over: a larger streamed corpus if
// one is present (go run ./cmd/oacorpus -out testdata/bench.corpus), else the
// committed CC0 corpus the conformance test uses.
func benchCorpus(b *testing.B) []string {
	b.Helper()
	if w, err := readLinesFile(filepath.Join("testdata", "bench.corpus")); err == nil && len(w) > 0 {
		return w
	}
	w, err := readLinesFile(filepath.Join("testdata", "corpus.txt"))
	if err != nil {
		b.Fatalf("read corpus: %v", err)
	}
	if len(w) == 0 {
		b.Fatal("empty corpus")
	}
	return w
}

// stemAll stems every word in words, b.N times, reporting words/sec. The cursor
// cycles the corpus so a large b.N keeps measuring real work.
func stemAll(b *testing.B, a Algorithm, words []string) {
	s := New(a)
	b.ReportAllocs()
	b.ResetTimer()
	n := 0
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

// BenchmarkEnglish is the headline throughput number, compared against
// rust-stemmers via `rustgen bench english <corpus>`.
func BenchmarkEnglish(b *testing.B) {
	stemAll(b, English, benchCorpus(b))
}

// BenchmarkAll measures every algorithm's throughput over the same corpus, so a
// regression in any port shows up.
func BenchmarkAll(b *testing.B) {
	words := benchCorpus(b)
	for _, a := range allAlgorithms {
		b.Run(algoName(a), func(b *testing.B) { stemAll(b, a, words) })
	}
}
