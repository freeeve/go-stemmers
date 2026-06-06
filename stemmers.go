// Package stemmers is a Go port of the rust-stemmers crate: the Snowball
// stemming algorithms with a small Algorithm/Stemmer API. Every algorithm
// produces output identical to rust-stemmers 1.2.0 (and thus to the reference
// Snowball implementation), verified against the canonical Snowball vocabularies
// in testdata.
//
// Input is expected to be already lowercased, matching rust-stemmers' contract.
//
//	s := stemmers.New(stemmers.English)
//	s.Stem("fruitlessly") // "fruitless"
package stemmers

import (
	"sync"

	"github.com/freeeve/go-stemmers/internal/arabic"
	"github.com/freeeve/go-stemmers/internal/danish"
	"github.com/freeeve/go-stemmers/internal/dutch"
	"github.com/freeeve/go-stemmers/internal/english"
	"github.com/freeeve/go-stemmers/internal/finnish"
	"github.com/freeeve/go-stemmers/internal/french"
	"github.com/freeeve/go-stemmers/internal/german"
	"github.com/freeeve/go-stemmers/internal/greek"
	"github.com/freeeve/go-stemmers/internal/hungarian"
	"github.com/freeeve/go-stemmers/internal/italian"
	"github.com/freeeve/go-stemmers/internal/norwegian"
	"github.com/freeeve/go-stemmers/internal/portuguese"
	"github.com/freeeve/go-stemmers/internal/romanian"
	"github.com/freeeve/go-stemmers/internal/russian"
	"github.com/freeeve/go-stemmers/internal/snowball"
	"github.com/freeeve/go-stemmers/internal/spanish"
	"github.com/freeeve/go-stemmers/internal/swedish"
	"github.com/freeeve/go-stemmers/internal/tamil"
	"github.com/freeeve/go-stemmers/internal/turkish"
)

// Algorithm selects a stemming algorithm. The values mirror rust-stemmers'
// Algorithm enum order so the two libraries line up one-to-one.
type Algorithm int

const (
	Arabic Algorithm = iota
	Danish
	Dutch
	English
	Finnish
	French
	German
	Greek
	Hungarian
	Italian
	Norwegian
	Portuguese
	Romanian
	Russian
	Spanish
	Swedish
	Tamil
	Turkish
)

// algorithms maps each Algorithm to its Snowball stem entry point.
var algorithms = map[Algorithm]func(*snowball.Env) bool{
	Arabic:     arabic.Stem,
	Danish:     danish.Stem,
	Dutch:      dutch.Stem,
	English:    english.Stem,
	Finnish:    finnish.Stem,
	French:     french.Stem,
	German:     german.Stem,
	Greek:      greek.Stem,
	Hungarian:  hungarian.Stem,
	Italian:    italian.Stem,
	Norwegian:  norwegian.Stem,
	Portuguese: portuguese.Stem,
	Romanian:   romanian.Stem,
	Russian:    russian.Stem,
	Spanish:    spanish.Stem,
	Swedish:    swedish.Stem,
	Tamil:      tamil.Stem,
	Turkish:    turkish.Stem,
}

// Stemmer stems single words with one configured algorithm. It is safe for
// concurrent use: a pool hands each call its own working state, so a Stemmer can
// be shared across goroutines.
type Stemmer struct {
	stem func(*snowball.Env) bool
	pool sync.Pool
}

// New returns a Stemmer for the given algorithm. It panics if a is not a known
// Algorithm.
func New(a Algorithm) *Stemmer {
	stem, ok := algorithms[a]
	if !ok {
		panic("stemmers: unknown algorithm")
	}
	s := &Stemmer{stem: stem}
	s.pool.New = func() any { return new(snowball.Env) }
	return s
}

// Stem returns the stem of word. The input is expected to be already lowercased
// (Snowball's contract); a word the algorithm leaves unchanged is returned as-is.
// The working Env is taken from a pool and returned after use; the result string
// is independent of it (a fresh copy when the word changed, the input otherwise).
func (s *Stemmer) Stem(word string) string {
	e := s.pool.Get().(*snowball.Env)
	e.Reset(word)
	s.stem(e)
	out := e.GetCurrent()
	s.pool.Put(e)
	return out
}
