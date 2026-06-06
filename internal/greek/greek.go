// Package greek is a byte-faithful Go port of rust-stemmers' generated
// Snowball "greek" stemmer. It produces output identical to rust-stemmers
// 1.2.0's Greek algorithm; the canonical Snowball greek vocabulary is the
// conformance oracle.
//
// The control flow mirrors the generated Rust: each Snowball labelled block
// (`'lab: loop { … break 'lab }`) becomes a Go labelled `for { … break LABEL }`,
// and the result-code dispatch (`else if among_var == N`) becomes a `switch`.
// The among search strings are written with Go `\uXXXX` escapes for the same
// Greek codepoints the Rust source encodes as `\u{XXXX}`.
package greek

import "github.com/freeeve/go-stemmers/internal/snowball"

// context holds the per-run state the generated algorithm threads through its
// routines: the captured suffix text reinserted by the `<+ s` operations, and
// the test1 flag that gates the final step6 deletion.
type context struct {
	sS     string
	bTest1 bool
}

var a0 = []snowball.Among[context]{
	{Str: "", SubstringI: -1, Result: 25},
	{Str: "\u03C2", SubstringI: 0, Result: 18},
	{Str: "\u0386", SubstringI: 0, Result: 1},
	{Str: "\u0388", SubstringI: 0, Result: 5},
	{Str: "\u0389", SubstringI: 0, Result: 7},
	{Str: "\u038A", SubstringI: 0, Result: 9},
	{Str: "\u03CA", SubstringI: 0, Result: 7},
	{Str: "\u03CB", SubstringI: 0, Result: 20},
	{Str: "\u038C", SubstringI: 0, Result: 15},
	{Str: "\u03CC", SubstringI: 0, Result: 15},
	{Str: "\u03CD", SubstringI: 0, Result: 20},
	{Str: "\u038E", SubstringI: 0, Result: 20},
	{Str: "\u03CE", SubstringI: 0, Result: 24},
	{Str: "\u038F", SubstringI: 0, Result: 24},
	{Str: "\u0390", SubstringI: 0, Result: 7},
	{Str: "\u0391", SubstringI: 0, Result: 1},
	{Str: "\u0392", SubstringI: 0, Result: 2},
	{Str: "\u0393", SubstringI: 0, Result: 3},
	{Str: "\u0394", SubstringI: 0, Result: 4},
	{Str: "\u0395", SubstringI: 0, Result: 5},
	{Str: "\u0396", SubstringI: 0, Result: 6},
	{Str: "\u0397", SubstringI: 0, Result: 7},
	{Str: "\u0398", SubstringI: 0, Result: 8},
	{Str: "\u0399", SubstringI: 0, Result: 9},
	{Str: "\u039A", SubstringI: 0, Result: 10},
	{Str: "\u039B", SubstringI: 0, Result: 11},
	{Str: "\u039C", SubstringI: 0, Result: 12},
	{Str: "\u039D", SubstringI: 0, Result: 13},
	{Str: "\u039E", SubstringI: 0, Result: 14},
	{Str: "\u039F", SubstringI: 0, Result: 15},
	{Str: "\u03A0", SubstringI: 0, Result: 16},
	{Str: "\u03A1", SubstringI: 0, Result: 17},
	{Str: "\u03A3", SubstringI: 0, Result: 18},
	{Str: "\u03A4", SubstringI: 0, Result: 19},
	{Str: "\u03A5", SubstringI: 0, Result: 20},
	{Str: "\u03A6", SubstringI: 0, Result: 21},
	{Str: "\u03A7", SubstringI: 0, Result: 22},
	{Str: "\u03A8", SubstringI: 0, Result: 23},
	{Str: "\u03A9", SubstringI: 0, Result: 24},
	{Str: "\u03AA", SubstringI: 0, Result: 9},
	{Str: "\u03AB", SubstringI: 0, Result: 20},
	{Str: "\u03AC", SubstringI: 0, Result: 1},
	{Str: "\u03AD", SubstringI: 0, Result: 5},
	{Str: "\u03AE", SubstringI: 0, Result: 7},
	{Str: "\u03AF", SubstringI: 0, Result: 9},
	{Str: "\u03B0", SubstringI: 0, Result: 20},
}

var a1 = []snowball.Among[context]{
	{Str: "\u03BA\u03B1\u03B8\u03B5\u03C3\u03C4\u03C9\u03C3", SubstringI: -1, Result: 10},
	{Str: "\u03C6\u03C9\u03C3", SubstringI: -1, Result: 9},
	{Str: "\u03C0\u03B5\u03C1\u03B1\u03C3", SubstringI: -1, Result: 7},
	{Str: "\u03C4\u03B5\u03C1\u03B1\u03C3", SubstringI: -1, Result: 8},
	{Str: "\u03BA\u03C1\u03B5\u03B1\u03C3", SubstringI: -1, Result: 6},
	{Str: "\u03BA\u03B1\u03B8\u03B5\u03C3\u03C4\u03C9\u03C4\u03BF\u03C3", SubstringI: -1, Result: 10},
	{Str: "\u03C6\u03C9\u03C4\u03BF\u03C3", SubstringI: -1, Result: 9},
	{Str: "\u03C0\u03B5\u03C1\u03B1\u03C4\u03BF\u03C3", SubstringI: -1, Result: 7},
	{Str: "\u03C4\u03B5\u03C1\u03B1\u03C4\u03BF\u03C3", SubstringI: -1, Result: 8},
	{Str: "\u03BA\u03C1\u03B5\u03B1\u03C4\u03BF\u03C3", SubstringI: -1, Result: 6},
	{Str: "\u03B3\u03B5\u03B3\u03BF\u03BD\u03BF\u03C4\u03BF\u03C3", SubstringI: -1, Result: 11},
	{Str: "\u03B3\u03B5\u03B3\u03BF\u03BD\u03BF\u03C3", SubstringI: -1, Result: 11},
	{Str: "\u03C6\u03B1\u03B3\u03B9\u03BF\u03C5", SubstringI: -1, Result: 1},
	{Str: "\u03C3\u03BA\u03B1\u03B3\u03B9\u03BF\u03C5", SubstringI: -1, Result: 2},
	{Str: "\u03C3\u03BF\u03B3\u03B9\u03BF\u03C5", SubstringI: -1, Result: 4},
	{Str: "\u03C4\u03B1\u03C4\u03BF\u03B3\u03B9\u03BF\u03C5", SubstringI: -1, Result: 5},
	{Str: "\u03BF\u03BB\u03BF\u03B3\u03B9\u03BF\u03C5", SubstringI: -1, Result: 3},
	{Str: "\u03BA\u03B1\u03B8\u03B5\u03C3\u03C4\u03C9\u03C4\u03B1", SubstringI: -1, Result: 10},
	{Str: "\u03C6\u03C9\u03C4\u03B1", SubstringI: -1, Result: 9},
	{Str: "\u03C0\u03B5\u03C1\u03B1\u03C4\u03B1", SubstringI: -1, Result: 7},
	{Str: "\u03C4\u03B5\u03C1\u03B1\u03C4\u03B1", SubstringI: -1, Result: 8},
	{Str: "\u03BA\u03C1\u03B5\u03B1\u03C4\u03B1", SubstringI: -1, Result: 6},
	{Str: "\u03B3\u03B5\u03B3\u03BF\u03BD\u03BF\u03C4\u03B1", SubstringI: -1, Result: 11},
	{Str: "\u03C6\u03B1\u03B3\u03B9\u03B1", SubstringI: -1, Result: 1},
	{Str: "\u03C3\u03BA\u03B1\u03B3\u03B9\u03B1", SubstringI: -1, Result: 2},
	{Str: "\u03C3\u03BF\u03B3\u03B9\u03B1", SubstringI: -1, Result: 4},
	{Str: "\u03C4\u03B1\u03C4\u03BF\u03B3\u03B9\u03B1", SubstringI: -1, Result: 5},
	{Str: "\u03BF\u03BB\u03BF\u03B3\u03B9\u03B1", SubstringI: -1, Result: 3},
	{Str: "\u03C0\u03B5\u03C1\u03B1\u03C4\u03B7", SubstringI: -1, Result: 7},
	{Str: "\u03BA\u03B1\u03B8\u03B5\u03C3\u03C4\u03C9\u03C4\u03C9\u03BD", SubstringI: -1, Result: 10},
	{Str: "\u03C6\u03C9\u03C4\u03C9\u03BD", SubstringI: -1, Result: 9},
	{Str: "\u03C0\u03B5\u03C1\u03B1\u03C4\u03C9\u03BD", SubstringI: -1, Result: 7},
	{Str: "\u03C4\u03B5\u03C1\u03B1\u03C4\u03C9\u03BD", SubstringI: -1, Result: 8},
	{Str: "\u03BA\u03C1\u03B5\u03B1\u03C4\u03C9\u03BD", SubstringI: -1, Result: 6},
	{Str: "\u03B3\u03B5\u03B3\u03BF\u03BD\u03BF\u03C4\u03C9\u03BD", SubstringI: -1, Result: 11},
	{Str: "\u03C6\u03B1\u03B3\u03B9\u03C9\u03BD", SubstringI: -1, Result: 1},
	{Str: "\u03C3\u03BA\u03B1\u03B3\u03B9\u03C9\u03BD", SubstringI: -1, Result: 2},
	{Str: "\u03C3\u03BF\u03B3\u03B9\u03C9\u03BD", SubstringI: -1, Result: 4},
	{Str: "\u03C4\u03B1\u03C4\u03BF\u03B3\u03B9\u03C9\u03BD", SubstringI: -1, Result: 5},
	{Str: "\u03BF\u03BB\u03BF\u03B3\u03B9\u03C9\u03BD", SubstringI: -1, Result: 3},
}

var a2 = []snowball.Among[context]{
	{Str: "\u03C0\u03B1", SubstringI: -1, Result: 1},
	{Str: "\u03BE\u03B1\u03BD\u03B1\u03C0\u03B1", SubstringI: 0, Result: 1},
	{Str: "\u03B5\u03C0\u03B1", SubstringI: 0, Result: 1},
	{Str: "\u03C0\u03B5\u03C1\u03B9\u03C0\u03B1", SubstringI: 0, Result: 1},
	{Str: "\u03B1\u03BD\u03B1\u03BC\u03C0\u03B1", SubstringI: 0, Result: 1},
	{Str: "\u03B5\u03BC\u03C0\u03B1", SubstringI: 0, Result: 1},
	{Str: "\u03B4\u03B1\u03BD\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03B8\u03C1\u03BF", SubstringI: -1, Result: 1},
	{Str: "\u03C3\u03C5\u03BD\u03B1\u03B8\u03C1\u03BF", SubstringI: 7, Result: 1},
}

var a3 = []snowball.Among[context]{
	{Str: "\u03C0", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03BC\u03C0", SubstringI: 0, Result: 1},
	{Str: "\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03C1", SubstringI: 2, Result: 1},
	{Str: "\u03BC\u03C0\u03C1", SubstringI: 3, Result: 1},
	{Str: "\u03B1\u03C1\u03C1", SubstringI: 2, Result: 1},
	{Str: "\u03B3\u03BB\u03C5\u03BA\u03C5\u03C1", SubstringI: 2, Result: 1},
	{Str: "\u03C0\u03BF\u03BB\u03C5\u03C1", SubstringI: 2, Result: 1},
	{Str: "\u03B1\u03BC\u03C0\u03B1\u03C1", SubstringI: 2, Result: 1},
	{Str: "\u03BC\u03B1\u03C1", SubstringI: 2, Result: 1},
	{Str: "\u03B3\u03BA\u03C1", SubstringI: 2, Result: 1},
	{Str: "\u03C0\u03B9\u03C0\u03B5\u03C1\u03BF\u03C1", SubstringI: 2, Result: 1},
	{Str: "\u03B2\u03BF\u03BB\u03B2\u03BF\u03C1", SubstringI: 2, Result: 1},
	{Str: "\u03B3\u03BB\u03C5\u03BA\u03BF\u03C1", SubstringI: 2, Result: 1},
	{Str: "\u03BB\u03BF\u03C5", SubstringI: -1, Result: 1},
	{Str: "\u03B2", SubstringI: -1, Result: 1},
	{Str: "\u03B2\u03B1\u03B8\u03C5\u03C1\u03B9", SubstringI: -1, Result: 1},
	{Str: "\u03B2\u03B1\u03C1\u03BA", SubstringI: -1, Result: 1},
	{Str: "\u03BC\u03B1\u03C1\u03BA", SubstringI: -1, Result: 1},
	{Str: "\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03BC", SubstringI: -1, Result: 1},
	{Str: "\u03BA\u03BF\u03C1\u03BD", SubstringI: -1, Result: 1},
}

var a4 = []snowball.Among[context]{
	{Str: "\u03B9\u03B6\u03B5\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03B6\u03B5\u03B9\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03B6\u03C9", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03B6\u03B1", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03B6\u03B1\u03C4\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03B6\u03B5\u03C4\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03B6\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03B6\u03BF\u03C5\u03BC\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03B6\u03B1\u03BC\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03B6\u03BF\u03C5\u03BD\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03B6\u03B1\u03BD\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03B6\u03B5\u03B9", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03B6\u03BF\u03C5\u03BD", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03B6\u03B1\u03BD", SubstringI: -1, Result: 1},
}

var a5 = []snowball.Among[context]{
	{Str: "\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03C7", SubstringI: -1, Result: 1},
	{Str: "\u03C5\u03C8", SubstringI: -1, Result: 1},
	{Str: "\u03B6\u03C9", SubstringI: -1, Result: 1},
	{Str: "\u03B2\u03B9", SubstringI: -1, Result: 1},
	{Str: "\u03BB\u03B9", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03B5\u03BD", SubstringI: -1, Result: 1},
}

var a6 = []snowball.Among[context]{
	{Str: "\u03C9\u03B8\u03B7\u03BA\u03B5\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03C9\u03B8\u03B7\u03BA\u03B1", SubstringI: -1, Result: 1},
	{Str: "\u03C9\u03B8\u03B7\u03BA\u03B1\u03C4\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03C9\u03B8\u03B7\u03BA\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03C9\u03B8\u03B7\u03BA\u03B1\u03BC\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03C9\u03B8\u03B7\u03BA\u03B1\u03BD\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03C9\u03B8\u03B7\u03BA\u03B1\u03BD", SubstringI: -1, Result: 1},
}

var a7 = []snowball.Among[context]{
	{Str: "\u03BE\u03B1\u03BD\u03B1\u03C0\u03B1", SubstringI: -1, Result: 1},
	{Str: "\u03B5\u03C0\u03B1", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03B5\u03C1\u03B9\u03C0\u03B1", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03BD\u03B1\u03BC\u03C0\u03B1", SubstringI: -1, Result: 1},
	{Str: "\u03B5\u03BC\u03C0\u03B1", SubstringI: -1, Result: 1},
	{Str: "\u03C7\u03B1\u03C1\u03C4\u03BF\u03C0\u03B1", SubstringI: -1, Result: 1},
	{Str: "\u03B5\u03BE\u03B1\u03C1\u03C7\u03B1", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B5\u03C0\u03B5", SubstringI: 7, Result: 1},
	{Str: "\u03BC\u03B5\u03C4\u03B5\u03C0\u03B5", SubstringI: 8, Result: 1},
	{Str: "\u03B5\u03C3\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03BA\u03BB\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B5\u03C3\u03C9\u03BA\u03BB\u03B5", SubstringI: 11, Result: 1},
	{Str: "\u03B5\u03BA\u03BB\u03B5", SubstringI: 11, Result: 1},
	{Str: "\u03B1\u03C0\u03B5\u03BA\u03BB\u03B5", SubstringI: 13, Result: 1},
	{Str: "\u03B1\u03C0\u03BF\u03BA\u03BB\u03B5", SubstringI: 11, Result: 1},
	{Str: "\u03B4\u03B1\u03BD\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03B8\u03C1\u03BF", SubstringI: -1, Result: 1},
	{Str: "\u03C3\u03C5\u03BD\u03B1\u03B8\u03C1\u03BF", SubstringI: 17, Result: 1},
}

var a8 = []snowball.Among[context]{
	{Str: "\u03C0", SubstringI: -1, Result: 1},
	{Str: "\u03BB\u03B1\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03B4\u03B7\u03BC\u03BF\u03BA\u03C1\u03B1\u03C4", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03C6", SubstringI: -1, Result: 1},
	{Str: "\u03B3\u03B9\u03B3\u03B1\u03BD\u03C4\u03BF\u03B1\u03C6", SubstringI: 3, Result: 1},
	{Str: "\u03B3\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B3\u03BA\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B3\u03BA", SubstringI: -1, Result: 1},
	{Str: "\u03BC", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03BF\u03C5\u03BA\u03B1\u03BC", SubstringI: 8, Result: 1},
	{Str: "\u03BA\u03BF\u03BC", SubstringI: 8, Result: 1},
	{Str: "\u03B1\u03BD", SubstringI: -1, Result: 1},
	{Str: "\u03BF\u03BB\u03BF", SubstringI: -1, Result: 1},
}

var a9 = []snowball.Among[context]{
	{Str: "\u03B9\u03C3\u03B5\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03C3\u03B1", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03C3\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03C3\u03B1\u03C4\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03C3\u03B1\u03BC\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03C3\u03B1\u03BD\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03C3\u03B1\u03BD", SubstringI: -1, Result: 1},
}

var a10 = []snowball.Among[context]{
	{Str: "\u03BE\u03B1\u03BD\u03B1\u03C0\u03B1", SubstringI: -1, Result: 1},
	{Str: "\u03B5\u03C0\u03B1", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03B5\u03C1\u03B9\u03C0\u03B1", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03BD\u03B1\u03BC\u03C0\u03B1", SubstringI: -1, Result: 1},
	{Str: "\u03B5\u03BC\u03C0\u03B1", SubstringI: -1, Result: 1},
	{Str: "\u03C7\u03B1\u03C1\u03C4\u03BF\u03C0\u03B1", SubstringI: -1, Result: 1},
	{Str: "\u03B5\u03BE\u03B1\u03C1\u03C7\u03B1", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B5\u03C0\u03B5", SubstringI: 7, Result: 1},
	{Str: "\u03BC\u03B5\u03C4\u03B5\u03C0\u03B5", SubstringI: 8, Result: 1},
	{Str: "\u03B5\u03C3\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03BA\u03BB\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B5\u03C3\u03C9\u03BA\u03BB\u03B5", SubstringI: 11, Result: 1},
	{Str: "\u03B5\u03BA\u03BB\u03B5", SubstringI: 11, Result: 1},
	{Str: "\u03B1\u03C0\u03B5\u03BA\u03BB\u03B5", SubstringI: 13, Result: 1},
	{Str: "\u03B1\u03C0\u03BF\u03BA\u03BB\u03B5", SubstringI: 11, Result: 1},
	{Str: "\u03B4\u03B1\u03BD\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03B8\u03C1\u03BF", SubstringI: -1, Result: 1},
	{Str: "\u03C3\u03C5\u03BD\u03B1\u03B8\u03C1\u03BF", SubstringI: 17, Result: 1},
}

var a11 = []snowball.Among[context]{
	{Str: "\u03B9\u03C3\u03B5\u03B9\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03C3\u03C9", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03C3\u03B5\u03C4\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03C3\u03BF\u03C5\u03BC\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03C3\u03BF\u03C5\u03BD\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03C3\u03B5\u03B9", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03C3\u03BF\u03C5\u03BD", SubstringI: -1, Result: 1},
}

var a12 = []snowball.Among[context]{
	{Str: "\u03C3\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03C3\u03B5", SubstringI: 0, Result: 1},
	{Str: "\u03C0\u03BB\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03BA\u03BB\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B5\u03C3\u03C9\u03BA\u03BB\u03B5", SubstringI: 3, Result: 1},
	{Str: "\u03B4\u03B1\u03BD\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03C3\u03C5\u03BD\u03B1\u03B8\u03C1\u03BF", SubstringI: -1, Result: 1},
}

var a13 = []snowball.Among[context]{
	{Str: "\u03C0", SubstringI: -1, Result: 1},
	{Str: "\u03B5\u03C5\u03C0", SubstringI: 0, Result: 1},
	{Str: "\u03B1\u03C0", SubstringI: 0, Result: 1},
	{Str: "\u03B5\u03BC\u03C0", SubstringI: 0, Result: 1},
	{Str: "\u03B3\u03C5\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03C7\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03C7\u03C9\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03BF\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03C7\u03C4", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03C7\u03C4", SubstringI: 9, Result: 1},
	{Str: "\u03BA\u03C4", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03BA\u03C4", SubstringI: 11, Result: 1},
	{Str: "\u03C3\u03C7", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03C3\u03C7", SubstringI: 13, Result: 1},
	{Str: "\u03C4\u03B1\u03C7", SubstringI: -1, Result: 1},
	{Str: "\u03C5\u03C8", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03C4\u03B1", SubstringI: -1, Result: 1},
	{Str: "\u03C6\u03B1", SubstringI: -1, Result: 1},
	{Str: "\u03B7\u03C6\u03B1", SubstringI: 18, Result: 1},
	{Str: "\u03BB\u03C5\u03B3", SubstringI: -1, Result: 1},
	{Str: "\u03BC\u03B5\u03B3", SubstringI: -1, Result: 1},
	{Str: "\u03B7\u03B4", SubstringI: -1, Result: 1},
	{Str: "\u03B5\u03C7\u03B8", SubstringI: -1, Result: 1},
	{Str: "\u03BA\u03B1\u03B8", SubstringI: -1, Result: 1},
	{Str: "\u03C3\u03BA", SubstringI: -1, Result: 1},
	{Str: "\u03BA\u03B1\u03BA", SubstringI: -1, Result: 1},
	{Str: "\u03BC\u03B1\u03BA", SubstringI: -1, Result: 1},
	{Str: "\u03BA\u03C5\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03C6\u03B9\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03BC", SubstringI: -1, Result: 1},
	{Str: "\u03B3\u03B5\u03BC", SubstringI: 30, Result: 1},
	{Str: "\u03B1\u03C7\u03BD", SubstringI: -1, Result: 1},
}

var a14 = []snowball.Among[context]{
	{Str: "\u03B9\u03C3\u03C4\u03BF\u03C5\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03C3\u03C4\u03B5\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03C3\u03C4\u03B7\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03C3\u03C4\u03BF\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03C3\u03C4\u03BF\u03C5", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03C3\u03C4\u03B1", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03C3\u03C4\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03C3\u03C4\u03B7", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03C3\u03C4\u03BF\u03B9", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03C3\u03C4\u03C9\u03BD", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03C3\u03C4\u03BF", SubstringI: -1, Result: 1},
}

var a15 = []snowball.Among[context]{
	{Str: "\u03C3\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03BC\u03B5\u03C4\u03B1\u03C3\u03B5", SubstringI: 0, Result: 1},
	{Str: "\u03BC\u03B9\u03BA\u03C1\u03BF\u03C3\u03B5", SubstringI: 0, Result: 1},
	{Str: "\u03B5\u03B3\u03BA\u03BB\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03C0\u03BF\u03BA\u03BB\u03B5", SubstringI: -1, Result: 1},
}

var a16 = []snowball.Among[context]{
	{Str: "\u03B4\u03B1\u03BD\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03BD\u03C4\u03B9\u03B4\u03B1\u03BD\u03B5", SubstringI: 0, Result: 1},
}

var a17 = []snowball.Among[context]{
	{Str: "\u03C4\u03BF\u03C0\u03B9\u03BA", SubstringI: -1, Result: 7},
	{Str: "\u03C3\u03BA\u03B5\u03C0\u03C4\u03B9\u03BA", SubstringI: -1, Result: 6},
	{Str: "\u03B3\u03BD\u03C9\u03C3\u03C4\u03B9\u03BA", SubstringI: -1, Result: 3},
	{Str: "\u03B1\u03B3\u03BD\u03C9\u03C3\u03C4\u03B9\u03BA", SubstringI: 2, Result: 1},
	{Str: "\u03B5\u03BA\u03BB\u03B5\u03BA\u03C4\u03B9\u03BA", SubstringI: -1, Result: 5},
	{Str: "\u03B1\u03C4\u03BF\u03BC\u03B9\u03BA", SubstringI: -1, Result: 2},
	{Str: "\u03B5\u03B8\u03BD\u03B9\u03BA", SubstringI: -1, Result: 4},
	{Str: "\u03B8\u03B5\u03B1\u03C4\u03C1\u03B9\u03BD", SubstringI: -1, Result: 10},
	{Str: "\u03B1\u03BB\u03B5\u03BE\u03B1\u03BD\u03B4\u03C1\u03B9\u03BD", SubstringI: -1, Result: 8},
	{Str: "\u03B2\u03C5\u03B6\u03B1\u03BD\u03C4\u03B9\u03BD", SubstringI: -1, Result: 9},
}

var a18 = []snowball.Among[context]{
	{Str: "\u03B9\u03C3\u03BC\u03BF\u03C5\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03C3\u03BC\u03BF\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03C3\u03BC\u03BF\u03C5", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03C3\u03BC\u03BF\u03B9", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03C3\u03BC\u03C9\u03BD", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03C3\u03BC\u03BF", SubstringI: -1, Result: 1},
}

var a19 = []snowball.Among[context]{
	{Str: "\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03C7", SubstringI: -1, Result: 1},
}

var a20 = []snowball.Among[context]{
	{Str: "\u03B1\u03C1\u03B1\u03BA\u03B9\u03B1", SubstringI: -1, Result: 1},
	{Str: "\u03BF\u03C5\u03B4\u03B1\u03BA\u03B9\u03B1", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03C1\u03B1\u03BA\u03B9", SubstringI: -1, Result: 1},
	{Str: "\u03BF\u03C5\u03B4\u03B1\u03BA\u03B9", SubstringI: -1, Result: 1},
}

var a21 = []snowball.Among[context]{
	{Str: "\u03BA\u03B1\u03C4\u03C1\u03B1\u03C0", SubstringI: -1, Result: 1},
	{Str: "\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03B2\u03C1", SubstringI: 1, Result: 1},
	{Str: "\u03BB\u03B1\u03B2\u03C1", SubstringI: 2, Result: 1},
	{Str: "\u03B1\u03BC\u03B2\u03C1", SubstringI: 2, Result: 1},
	{Str: "\u03BC\u03B5\u03C1", SubstringI: 1, Result: 1},
	{Str: "\u03B1\u03BD\u03B8\u03C1", SubstringI: 1, Result: 1},
	{Str: "\u03BA\u03BF\u03C1", SubstringI: 1, Result: 1},
	{Str: "\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03BD\u03B1\u03B3\u03BA\u03B1\u03C3", SubstringI: 8, Result: 1},
	{Str: "\u03BC\u03BF\u03C5\u03C3\u03C4", SubstringI: -1, Result: 1},
	{Str: "\u03C1\u03C5", SubstringI: -1, Result: 1},
	{Str: "\u03C6", SubstringI: -1, Result: 1},
	{Str: "\u03C3\u03C6", SubstringI: 12, Result: 1},
	{Str: "\u03B1\u03BB\u03B9\u03C3\u03C6", SubstringI: 13, Result: 1},
	{Str: "\u03C7", SubstringI: -1, Result: 1},
	{Str: "\u03B2\u03B1\u03BC\u03B2", SubstringI: -1, Result: 1},
	{Str: "\u03C3\u03BB\u03BF\u03B2", SubstringI: -1, Result: 1},
	{Str: "\u03C4\u03C3\u03B5\u03C7\u03BF\u03C3\u03BB\u03BF\u03B2", SubstringI: 17, Result: 1},
	{Str: "\u03C4\u03B6", SubstringI: -1, Result: 1},
	{Str: "\u03BA", SubstringI: -1, Result: 1},
	{Str: "\u03C3\u03BA", SubstringI: 20, Result: 1},
	{Str: "\u03BA\u03B1\u03C0\u03B1\u03BA", SubstringI: 20, Result: 1},
	{Str: "\u03C3\u03BF\u03BA", SubstringI: 20, Result: 1},
	{Str: "\u03C0\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03C6\u03C5\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03BB\u03BF\u03C5\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03BC\u03B1\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03C6\u03B1\u03C1\u03BC", SubstringI: -1, Result: 1},
	{Str: "\u03BA\u03B1\u03B9\u03BC", SubstringI: -1, Result: 1},
	{Str: "\u03BA\u03BB\u03B9\u03BC", SubstringI: -1, Result: 1},
	{Str: "\u03C3\u03C0\u03B1\u03BD", SubstringI: -1, Result: 1},
	{Str: "\u03BA\u03BF\u03BD", SubstringI: -1, Result: 1},
}

var a22 = []snowball.Among[context]{
	{Str: "\u03C0", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03B1\u03C4\u03B5\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03C4\u03BF\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03BD\u03C5\u03C6", SubstringI: -1, Result: 1},
	{Str: "\u03B2", SubstringI: -1, Result: 1},
	{Str: "\u03BA\u03B1\u03C1\u03B4", SubstringI: -1, Result: 1},
	{Str: "\u03B6", SubstringI: -1, Result: 1},
	{Str: "\u03C3\u03BA", SubstringI: -1, Result: 1},
	{Str: "\u03B2\u03B1\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03B3\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03C4\u03C1\u03B9\u03C0\u03BF\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03BC\u03B1\u03BA\u03C1\u03C5\u03BD", SubstringI: -1, Result: 1},
	{Str: "\u03B3\u03B9\u03B1\u03BD", SubstringI: -1, Result: 1},
	{Str: "\u03B7\u03B3\u03BF\u03C5\u03BC\u03B5\u03BD", SubstringI: -1, Result: 1},
	{Str: "\u03BA\u03BF\u03BD", SubstringI: -1, Result: 1},
}

var a23 = []snowball.Among[context]{
	{Str: "\u03B9\u03C4\u03C3\u03B1\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03C4\u03C3\u03B5\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03C4\u03C3\u03B1", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03BA\u03B9\u03B1", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03C1\u03B1\u03BA\u03B9\u03B1", SubstringI: 3, Result: 1},
	{Str: "\u03B1\u03BA\u03B9", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03C1\u03B1\u03BA\u03B9", SubstringI: 5, Result: 1},
	{Str: "\u03B9\u03C4\u03C3\u03C9\u03BD", SubstringI: -1, Result: 1},
}

var a24 = []snowball.Among[context]{
	{Str: "\u03B9\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03C8\u03B1\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03B9\u03C6\u03BD", SubstringI: -1, Result: 1},
	{Str: "\u03BF\u03BB\u03BF", SubstringI: -1, Result: 1},
}

var a25 = []snowball.Among[context]{
	{Str: "\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03B1\u03B9\u03C7\u03BD", SubstringI: -1, Result: 1},
}

var a26 = []snowball.Among[context]{
	{Str: "\u03B9\u03B4\u03B9\u03B1", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03B4\u03B9\u03C9\u03BD", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03B4\u03B9\u03BF", SubstringI: -1, Result: 1},
}

var a27 = []snowball.Among[context]{
	{Str: "\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03B2", SubstringI: -1, Result: 1},
	{Str: "\u03B4", SubstringI: -1, Result: 1},
	{Str: "\u03BB\u03C5\u03BA", SubstringI: -1, Result: 1},
	{Str: "\u03C6\u03C1\u03B1\u03B3\u03BA", SubstringI: -1, Result: 1},
	{Str: "\u03BF\u03B2\u03B5\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03BC\u03B7\u03BD", SubstringI: -1, Result: 1},
}

var a28 = []snowball.Among[context]{
	{Str: "\u03B9\u03C3\u03BA\u03BF\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03C3\u03BA\u03BF\u03C5", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03C3\u03BA\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03C3\u03BA\u03BF", SubstringI: -1, Result: 1},
}

var a29 = []snowball.Among[context]{
	{Str: "\u03B1\u03B4\u03B5\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03B4\u03C9\u03BD", SubstringI: -1, Result: 1},
}

var a30 = []snowball.Among[context]{
	{Str: "\u03BC\u03C0\u03B1\u03BC\u03C0", SubstringI: -1, Result: -1},
	{Str: "\u03BA\u03C5\u03C1", SubstringI: -1, Result: -1},
	{Str: "\u03C0\u03B1\u03C4\u03B5\u03C1", SubstringI: -1, Result: -1},
	{Str: "\u03C0\u03B5\u03B8\u03B5\u03C1", SubstringI: -1, Result: -1},
	{Str: "\u03BD\u03C4\u03B1\u03BD\u03C4", SubstringI: -1, Result: -1},
	{Str: "\u03B3\u03B9\u03B1\u03B3\u03B9", SubstringI: -1, Result: -1},
	{Str: "\u03B8\u03B5\u03B9", SubstringI: -1, Result: -1},
	{Str: "\u03BF\u03BA", SubstringI: -1, Result: -1},
	{Str: "\u03BC\u03B1\u03BC", SubstringI: -1, Result: -1},
	{Str: "\u03BC\u03B1\u03BD", SubstringI: -1, Result: -1},
}

var a31 = []snowball.Among[context]{
	{Str: "\u03B5\u03B4\u03B5\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03B5\u03B4\u03C9\u03BD", SubstringI: -1, Result: 1},
}

var a32 = []snowball.Among[context]{
	{Str: "\u03BA\u03C1\u03B1\u03C3\u03C0", SubstringI: -1, Result: 1},
	{Str: "\u03C5\u03C0", SubstringI: -1, Result: 1},
	{Str: "\u03B4\u03B1\u03C0", SubstringI: -1, Result: 1},
	{Str: "\u03B3\u03B7\u03C0", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03C0", SubstringI: -1, Result: 1},
	{Str: "\u03B5\u03BC\u03C0", SubstringI: -1, Result: 1},
	{Str: "\u03BF\u03C0", SubstringI: -1, Result: 1},
	{Str: "\u03BC\u03B9\u03BB", SubstringI: -1, Result: 1},
}

var a33 = []snowball.Among[context]{
	{Str: "\u03BF\u03C5\u03B4\u03B5\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03BF\u03C5\u03B4\u03C9\u03BD", SubstringI: -1, Result: 1},
}

var a34 = []snowball.Among[context]{
	{Str: "\u03C3\u03C0", SubstringI: -1, Result: 1},
	{Str: "\u03C6\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03BB\u03B9\u03C7", SubstringI: -1, Result: 1},
	{Str: "\u03C4\u03C1\u03B1\u03B3", SubstringI: -1, Result: 1},
	{Str: "\u03C6\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03C1\u03BA", SubstringI: -1, Result: 1},
	{Str: "\u03C3\u03BA", SubstringI: -1, Result: 1},
	{Str: "\u03BA\u03B1\u03BB\u03B9\u03B1\u03BA", SubstringI: -1, Result: 1},
	{Str: "\u03BB\u03BF\u03C5\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03C6\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03B5\u03C4\u03B1\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03B2\u03B5\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03C7\u03BD", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03BB\u03B5\u03BE", SubstringI: -1, Result: 1},
}

var a35 = []snowball.Among[context]{
	{Str: "\u03B5\u03C9\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03B5\u03C9\u03BD", SubstringI: -1, Result: 1},
}

var a36 = []snowball.Among[context]{
	{Str: "\u03C0", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03B1\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03B4", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03B4", SubstringI: 2, Result: 1},
	{Str: "\u03B8", SubstringI: -1, Result: 1},
	{Str: "\u03B3\u03B1\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03B5\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03BD", SubstringI: -1, Result: 1},
}

var a37 = []snowball.Among[context]{
	{Str: "\u03B9\u03BF\u03C5", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03B1", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03C9\u03BD", SubstringI: -1, Result: 1},
}

var a38 = []snowball.Among[context]{
	{Str: "\u03B9\u03BA\u03BF\u03C5", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03BA\u03B1", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03BA\u03C9\u03BD", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03BA\u03BF", SubstringI: -1, Result: 1},
}

var a39 = []snowball.Among[context]{
	{Str: "\u03BA\u03B1\u03BB\u03C0", SubstringI: -1, Result: 1},
	{Str: "\u03B3\u03B5\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03BB\u03B9\u03B1\u03C4\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03B5\u03C4\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03B9\u03C4\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03C6\u03C5\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03C7\u03B1\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03BC\u03C0\u03BF\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03C3\u03B5\u03C1\u03C4", SubstringI: -1, Result: 1},
	{Str: "\u03BC\u03C0\u03B1\u03B3\u03B9\u03B1\u03C4", SubstringI: -1, Result: 1},
	{Str: "\u03BD\u03B9\u03C4", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03B9\u03BA\u03B1\u03BD\u03C4", SubstringI: -1, Result: 1},
	{Str: "\u03B5\u03BE\u03C9\u03B4", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03B4", SubstringI: -1, Result: 1},
	{Str: "\u03BA\u03B1\u03C4\u03B1\u03B4", SubstringI: 13, Result: 1},
	{Str: "\u03C3\u03C5\u03BD\u03B1\u03B4", SubstringI: 13, Result: 1},
	{Str: "\u03B1\u03BD\u03C4\u03B9\u03B4", SubstringI: -1, Result: 1},
	{Str: "\u03B5\u03BD\u03B4", SubstringI: -1, Result: 1},
	{Str: "\u03C5\u03C0\u03BF\u03B4", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03C1\u03C9\u03C4\u03BF\u03B4", SubstringI: -1, Result: 1},
	{Str: "\u03C6\u03C5\u03BB\u03BF\u03B4", SubstringI: -1, Result: 1},
	{Str: "\u03B7\u03B8", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03BD\u03B7\u03B8", SubstringI: 21, Result: 1},
	{Str: "\u03BE\u03B9\u03BA", SubstringI: -1, Result: 1},
	{Str: "\u03BC\u03BF\u03C5\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03BC\u03BC\u03BF\u03C7\u03B1\u03BB", SubstringI: 25, Result: 1},
	{Str: "\u03C3\u03C5\u03BD\u03BF\u03BC\u03B7\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03BC\u03C0\u03BF\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03B2\u03C1\u03C9\u03BC", SubstringI: -1, Result: 1},
	{Str: "\u03C4\u03C3\u03B1\u03BC", SubstringI: -1, Result: 1},
	{Str: "\u03BC\u03C0\u03B1\u03BD", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03BC\u03B1\u03BD", SubstringI: -1, Result: 1},
	{Str: "\u03BA\u03B1\u03BB\u03BB\u03B9\u03BD", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03BF\u03C3\u03C4\u03B5\u03BB\u03BD", SubstringI: -1, Result: 1},
	{Str: "\u03C6\u03B9\u03BB\u03BF\u03BD", SubstringI: -1, Result: 1},
}

var a40 = []snowball.Among[context]{
	{Str: "\u03BF\u03C5\u03C3\u03B1\u03BC\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B7\u03C3\u03B1\u03BC\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03B3\u03B1\u03BC\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B7\u03BA\u03B1\u03BC\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B7\u03B8\u03B7\u03BA\u03B1\u03BC\u03B5", SubstringI: 3, Result: 1},
}

var a41 = []snowball.Among[context]{
	{Str: "\u03B1\u03BD\u03B1\u03C0", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03B9\u03BA\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03C0\u03BF\u03C3\u03C4", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03BF\u03C4", SubstringI: -1, Result: 1},
	{Str: "\u03C7", SubstringI: -1, Result: 1},
	{Str: "\u03C3\u03B9\u03C7", SubstringI: 4, Result: 1},
	{Str: "\u03B2\u03BF\u03C5\u03B2", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03B5\u03B8", SubstringI: -1, Result: 1},
	{Str: "\u03BE\u03B5\u03B8", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03C0\u03BF\u03B8", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03C0\u03BF\u03BA", SubstringI: -1, Result: 1},
	{Str: "\u03BF\u03C5\u03BB", SubstringI: -1, Result: 1},
}

var a42 = []snowball.Among[context]{
	{Str: "\u03C4\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03C4\u03C3", SubstringI: -1, Result: 1},
}

var a43 = []snowball.Among[context]{
	{Str: "\u03BF\u03C5\u03C3\u03B1\u03BD\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B7\u03C3\u03B1\u03BD\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03BF\u03C5\u03BD\u03C4\u03B1\u03BD\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03BF\u03C5\u03BD\u03C4\u03B1\u03BD\u03B5", SubstringI: 2, Result: 1},
	{Str: "\u03BF\u03BD\u03C4\u03B1\u03BD\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03BF\u03BD\u03C4\u03B1\u03BD\u03B5", SubstringI: 4, Result: 1},
	{Str: "\u03BF\u03C4\u03B1\u03BD\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03BF\u03C4\u03B1\u03BD\u03B5", SubstringI: 6, Result: 1},
	{Str: "\u03B1\u03B3\u03B1\u03BD\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B7\u03BA\u03B1\u03BD\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B7\u03B8\u03B7\u03BA\u03B1\u03BD\u03B5", SubstringI: 9, Result: 1},
}

var a44 = []snowball.Among[context]{
	{Str: "\u03C0", SubstringI: -1, Result: 1},
	{Str: "\u03C3\u03C0", SubstringI: 0, Result: 1},
	{Str: "\u03C0\u03BF\u03BB\u03C5\u03B4\u03B1\u03C0", SubstringI: 0, Result: 1},
	{Str: "\u03B1\u03B4\u03B1\u03C0", SubstringI: 0, Result: 1},
	{Str: "\u03C7\u03B1\u03BC\u03B7\u03BB\u03BF\u03B4\u03B1\u03C0", SubstringI: 0, Result: 1},
	{Str: "\u03C4\u03C3\u03BF\u03C0", SubstringI: 0, Result: 1},
	{Str: "\u03BA\u03BF\u03C0", SubstringI: 0, Result: 1},
	{Str: "\u03C5\u03C0\u03BF\u03BA\u03BF\u03C0", SubstringI: 6, Result: 1},
	{Str: "\u03C0\u03B5\u03C1\u03B9\u03C4\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03BF\u03C5\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03B5\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03B2\u03B5\u03C4\u03B5\u03C1", SubstringI: 10, Result: 1},
	{Str: "\u03B3\u03B5\u03C1", SubstringI: 10, Result: 1},
	{Str: "\u03BB\u03BF\u03C5\u03B8\u03B7\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03BA\u03BF\u03C1\u03BC\u03BF\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03C3\u03B1\u03C1\u03B1\u03BA\u03B1\u03C4\u03C3", SubstringI: 15, Result: 1},
	{Str: "\u03B8\u03C5\u03C3", SubstringI: 15, Result: 1},
	{Str: "\u03B2\u03B1\u03C3", SubstringI: 15, Result: 1},
	{Str: "\u03C0\u03BF\u03BB\u03B9\u03C3", SubstringI: 15, Result: 1},
	{Str: "\u03BA\u03B1\u03C3\u03C4", SubstringI: -1, Result: 1},
	{Str: "\u03B4\u03B9\u03B1\u03C4", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03BB\u03B1\u03C4", SubstringI: -1, Result: 1},
	{Str: "\u03C4\u03C3\u03B1\u03C1\u03BB\u03B1\u03C4", SubstringI: -1, Result: 1},
	{Str: "\u03C4\u03B5\u03C4", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03BF\u03C5\u03C1\u03B9\u03C4", SubstringI: -1, Result: 1},
	{Str: "\u03C3\u03BF\u03C5\u03BB\u03C4", SubstringI: -1, Result: 1},
	{Str: "\u03B6\u03C9\u03BD\u03C4", SubstringI: -1, Result: 1},
	{Str: "\u03BC\u03B1\u03B9\u03BD\u03C4", SubstringI: -1, Result: 1},
	{Str: "\u03C6", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03B5\u03BD\u03C4\u03B1\u03C1\u03C6", SubstringI: 29, Result: 1},
	{Str: "\u03BA\u03BF\u03B9\u03BB\u03B1\u03C1\u03C6", SubstringI: 29, Result: 1},
	{Str: "\u03BF\u03C1\u03C6", SubstringI: 29, Result: 1},
	{Str: "\u03B4\u03B9\u03B1\u03C6", SubstringI: 29, Result: 1},
	{Str: "\u03C3\u03C4\u03B5\u03C6", SubstringI: 29, Result: 1},
	{Str: "\u03C6\u03C9\u03C4\u03BF\u03C3\u03C4\u03B5\u03C6", SubstringI: 34, Result: 1},
	{Str: "\u03C0\u03B5\u03C1\u03B7\u03C6", SubstringI: 29, Result: 1},
	{Str: "\u03C5\u03C0\u03B5\u03C1\u03B7\u03C6", SubstringI: 36, Result: 1},
	{Str: "\u03C7", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03BF\u03BB\u03C5\u03BC\u03B7\u03C7", SubstringI: 38, Result: 1},
	{Str: "\u03B1\u03BC\u03B7\u03C7", SubstringI: 38, Result: 1},
	{Str: "\u03B2\u03B9\u03BF\u03BC\u03B7\u03C7", SubstringI: 38, Result: 1},
	{Str: "\u03BC\u03B9\u03BA\u03C1\u03BF\u03B2\u03B9\u03BF\u03BC\u03B7\u03C7", SubstringI: 41, Result: 1},
	{Str: "\u03BC\u03B5\u03B3\u03BB\u03BF\u03B2\u03B9\u03BF\u03BC\u03B7\u03C7", SubstringI: 41, Result: 1},
	{Str: "\u03BA\u03B1\u03C0\u03BD\u03BF\u03B2\u03B9\u03BF\u03BC\u03B7\u03C7", SubstringI: 41, Result: 1},
	{Str: "\u03BB\u03B9\u03C7", SubstringI: 38, Result: 1},
	{Str: "\u03C4\u03B1\u03B2", SubstringI: -1, Result: 1},
	{Str: "\u03BD\u03C4\u03B1\u03B2", SubstringI: 46, Result: 1},
	{Str: "\u03C8\u03B7\u03BB\u03BF\u03C4\u03B1\u03B2", SubstringI: 46, Result: 1},
	{Str: "\u03BB\u03B9\u03B2", SubstringI: -1, Result: 1},
	{Str: "\u03BA\u03BB\u03B9\u03B2", SubstringI: 49, Result: 1},
	{Str: "\u03BE\u03B7\u03C1\u03BF\u03BA\u03BB\u03B9\u03B2", SubstringI: 50, Result: 1},
	{Str: "\u03B3", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03BD\u03BF\u03C1\u03B3", SubstringI: 52, Result: 1},
	{Str: "\u03B5\u03BD\u03BF\u03C1\u03B3", SubstringI: 52, Result: 1},
	{Str: "\u03B1\u03B3", SubstringI: 52, Result: 1},
	{Str: "\u03C4\u03C1\u03B1\u03B3", SubstringI: 55, Result: 1},
	{Str: "\u03C4\u03C3\u03B1\u03B3", SubstringI: 55, Result: 1},
	{Str: "\u03C4\u03C3\u03B9\u03B3\u03B3", SubstringI: 52, Result: 1},
	{Str: "\u03B1\u03C4\u03C3\u03B9\u03B3\u03B3", SubstringI: 58, Result: 1},
	{Str: "\u03B1\u03B8\u03B9\u03B3\u03B3", SubstringI: 52, Result: 1},
	{Str: "\u03C3\u03C4\u03B5\u03B3", SubstringI: 52, Result: 1},
	{Str: "\u03B1\u03C0\u03B7\u03B3", SubstringI: 52, Result: 1},
	{Str: "\u03C3\u03B9\u03B3", SubstringI: 52, Result: 1},
	{Str: "\u03BA\u03B1\u03BB\u03C0\u03BF\u03C5\u03B6", SubstringI: -1, Result: 1},
	{Str: "\u03B8", SubstringI: -1, Result: 1},
	{Str: "\u03BC\u03C9\u03B1\u03BC\u03B5\u03B8", SubstringI: 65, Result: 1},
	{Str: "\u03C0\u03B9\u03B8", SubstringI: 65, Result: 1},
	{Str: "\u03B1\u03C0\u03B9\u03B8", SubstringI: 67, Result: 1},
	{Str: "\u03B2\u03B1\u03C3\u03BA", SubstringI: -1, Result: 1},
	{Str: "\u03B2\u03C1\u03B1\u03C7\u03C5\u03BA", SubstringI: -1, Result: 1},
	{Str: "\u03B4\u03B5\u03BA", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03B5\u03BB\u03B5\u03BA", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03BA", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03BD\u03B9\u03BA", SubstringI: 73, Result: 1},
	{Str: "\u03B2\u03BF\u03C5\u03BB\u03BA", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03B4\u03B9\u03C0\u03BB", SubstringI: 76, Result: 1},
	{Str: "\u03C8\u03C5\u03C7\u03BF\u03C0\u03BB", SubstringI: 76, Result: 1},
	{Str: "\u03BB\u03B1\u03BF\u03C0\u03BB", SubstringI: 76, Result: 1},
	{Str: "\u03BF\u03C5\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03B3\u03B1\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03B2\u03B1\u03B8\u03C5\u03B3\u03B1\u03BB", SubstringI: 81, Result: 1},
	{Str: "\u03BA\u03B1\u03C4\u03B1\u03B3\u03B1\u03BB", SubstringI: 81, Result: 1},
	{Str: "\u03BF\u03BB\u03BF\u03B3\u03B1\u03BB", SubstringI: 81, Result: 1},
	{Str: "\u03BA\u03B1\u03C3\u03C4\u03B5\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03BC\u03B5\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03BF\u03C1\u03C4\u03BF\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03BC", SubstringI: -1, Result: 1},
	{Str: "\u03B4\u03C1\u03B1\u03B4\u03BF\u03C5\u03BC", SubstringI: 88, Result: 1},
	{Str: "\u03B2\u03C1\u03B1\u03C7\u03BC", SubstringI: 88, Result: 1},
	{Str: "\u03BF\u03BB\u03B9\u03B3\u03BF\u03B4\u03B1\u03BC", SubstringI: 88, Result: 1},
	{Str: "\u03BC\u03BF\u03C5\u03C3\u03BF\u03C5\u03BB\u03BC", SubstringI: 88, Result: 1},
	{Str: "\u03BD", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03BC\u03B5\u03C1\u03B9\u03BA\u03B1\u03BD", SubstringI: 93, Result: 1},
}

var a45 = []snowball.Among[context]{
	{Str: "\u03B7\u03C3\u03B5\u03C4\u03B5", SubstringI: -1, Result: 1},
}

var a46 = []snowball.Among[context]{
	{Str: "\u03C0\u03C5\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03B5\u03C5\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03C7\u03C9\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03B2\u03B1\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03B2\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03B9\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03C6\u03BF\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03BD\u03B5\u03C4", SubstringI: -1, Result: 1},
	{Str: "\u03C3\u03C7", SubstringI: -1, Result: 1},
	{Str: "\u03C3\u03C5\u03BD\u03B4", SubstringI: -1, Result: 1},
	{Str: "\u03B5\u03BD\u03B4", SubstringI: -1, Result: 1},
	{Str: "\u03BF\u03B4", SubstringI: -1, Result: 1},
	{Str: "\u03C5\u03C0\u03B5\u03C1\u03B8", SubstringI: -1, Result: 1},
	{Str: "\u03C3\u03B8", SubstringI: -1, Result: 1},
	{Str: "\u03B5\u03C5\u03B8", SubstringI: -1, Result: 1},
	{Str: "\u03C1\u03B1\u03B8", SubstringI: -1, Result: 1},
	{Str: "\u03C4\u03B1\u03B8", SubstringI: -1, Result: 1},
	{Str: "\u03B4\u03B9\u03B1\u03B8", SubstringI: -1, Result: 1},
	{Str: "\u03BA\u03B1\u03B8", SubstringI: -1, Result: 1},
	{Str: "\u03C4\u03B9\u03B8", SubstringI: -1, Result: 1},
	{Str: "\u03B5\u03BA\u03B8", SubstringI: -1, Result: 1},
	{Str: "\u03C3\u03C5\u03BD\u03B8", SubstringI: -1, Result: 1},
	{Str: "\u03B5\u03BD\u03B8", SubstringI: -1, Result: 1},
	{Str: "\u03C1\u03BF\u03B8", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03C1\u03BA", SubstringI: -1, Result: 1},
	{Str: "\u03C9\u03C6\u03B5\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03B2\u03BF\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03C3\u03C5\u03BD", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03B9\u03BD", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03BF\u03BD", SubstringI: -1, Result: 1},
	{Str: "\u03C1\u03BF\u03BD", SubstringI: -1, Result: 1},
}

var a47 = []snowball.Among[context]{
	{Str: "\u03C3\u03B5\u03C1\u03C0", SubstringI: -1, Result: 1},
	{Str: "\u03BA\u03BF\u03C0", SubstringI: -1, Result: 1},
	{Str: "\u03B8\u03B1\u03C1\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03BD\u03C4\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03B2\u03B1\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03B5\u03BD\u03B1\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03B2\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03BC\u03C0\u03BF\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03C5", SubstringI: -1, Result: 1},
	{Str: "\u03C3\u03C5\u03C1\u03C6", SubstringI: -1, Result: 1},
	{Str: "\u03BD\u03B9\u03C6", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03B1\u03B3", SubstringI: -1, Result: 1},
	{Str: "\u03B4", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03B4", SubstringI: 12, Result: 1},
	{Str: "\u03B8", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03B8", SubstringI: 14, Result: 1},
	{Str: "\u03C3\u03BA", SubstringI: -1, Result: 1},
	{Str: "\u03C4\u03BF\u03BA", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03C0\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03B1\u03C1\u03B1\u03BA\u03B1\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03C3\u03BA\u03B5\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03B5\u03BC", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03BD", SubstringI: -1, Result: 1},
	{Str: "\u03B2\u03B5\u03BD", SubstringI: -1, Result: 1},
	{Str: "\u03B2\u03B1\u03C1\u03BF\u03BD", SubstringI: -1, Result: 1},
}

var a48 = []snowball.Among[context]{
	{Str: "\u03C9\u03BD\u03C4\u03B1\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03BF\u03BD\u03C4\u03B1\u03C3", SubstringI: -1, Result: 1},
}

var a49 = []snowball.Among[context]{
	{Str: "\u03BF\u03BC\u03B1\u03C3\u03C4\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03BF\u03BC\u03B1\u03C3\u03C4\u03B5", SubstringI: 0, Result: 1},
}

var a50 = []snowball.Among[context]{
	{Str: "\u03C0", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03C0", SubstringI: 0, Result: 1},
	{Str: "\u03B1\u03BA\u03B1\u03C4\u03B1\u03C0", SubstringI: 1, Result: 1},
	{Str: "\u03C3\u03C5\u03BC\u03C0", SubstringI: 0, Result: 1},
	{Str: "\u03B1\u03C3\u03C5\u03BC\u03C0", SubstringI: 3, Result: 1},
	{Str: "\u03B1\u03BC\u03B5\u03C4\u03B1\u03BC\u03C6", SubstringI: -1, Result: 1},
}

var a51 = []snowball.Among[context]{
	{Str: "\u03B1\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03BD\u03B9\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03B6", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03B1\u03C1\u03B1\u03BA\u03B1\u03BB", SubstringI: 3, Result: 1},
	{Str: "\u03B5\u03BA\u03C4\u03B5\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03BC", SubstringI: -1, Result: 1},
	{Str: "\u03BE", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03C1\u03BF", SubstringI: -1, Result: 1},
}

var a52 = []snowball.Among[context]{
	{Str: "\u03B7\u03B8\u03B7\u03BA\u03B5\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03B7\u03B8\u03B7\u03BA\u03B1", SubstringI: -1, Result: 1},
	{Str: "\u03B7\u03B8\u03B7\u03BA\u03B5", SubstringI: -1, Result: 1},
}

var a53 = []snowball.Among[context]{
	{Str: "\u03C3\u03C6", SubstringI: -1, Result: 1},
	{Str: "\u03BD\u03B1\u03C1\u03B8", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03B9\u03B8", SubstringI: -1, Result: 1},
	{Str: "\u03BF\u03B8", SubstringI: -1, Result: 1},
	{Str: "\u03C3\u03BA\u03BF\u03C5\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03C3\u03BA\u03C9\u03BB", SubstringI: -1, Result: 1},
}

var a54 = []snowball.Among[context]{
	{Str: "\u03B8", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03C1\u03BF\u03C3\u03B8", SubstringI: 0, Result: 1},
	{Str: "\u03C0\u03B1\u03C1\u03B1\u03BA\u03B1\u03C4\u03B1\u03B8", SubstringI: 0, Result: 1},
	{Str: "\u03B4\u03B9\u03B1\u03B8", SubstringI: 0, Result: 1},
	{Str: "\u03C3\u03C5\u03BD\u03B8", SubstringI: 0, Result: 1},
}

var a55 = []snowball.Among[context]{
	{Str: "\u03B7\u03BA\u03B5\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03B7\u03BA\u03B1", SubstringI: -1, Result: 1},
	{Str: "\u03B7\u03BA\u03B5", SubstringI: -1, Result: 1},
}

var a56 = []snowball.Among[context]{
	{Str: "\u03B2\u03BB\u03B5\u03C0", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03BF\u03B4\u03B1\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03C1\u03C9\u03C4", SubstringI: -1, Result: 1},
	{Str: "\u03BA\u03C5\u03BC\u03B1\u03C4", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03B1\u03BD\u03C4\u03B1\u03C7", SubstringI: -1, Result: 1},
	{Str: "\u03BB\u03B1\u03C7", SubstringI: -1, Result: 1},
	{Str: "\u03C6\u03B1\u03B3", SubstringI: -1, Result: 1},
	{Str: "\u03BB\u03B7\u03B3", SubstringI: -1, Result: 1},
	{Str: "\u03C6\u03C1\u03C5\u03B4", SubstringI: -1, Result: 1},
	{Str: "\u03BC\u03B1\u03BD\u03C4\u03B9\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03BC\u03B1\u03BB\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03BF\u03BC", SubstringI: -1, Result: 1},
}

var a57 = []snowball.Among[context]{
	{Str: "\u03B5\u03BA\u03BB\u03B9\u03C0", SubstringI: -1, Result: 1},
	{Str: "\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03BD\u03B1\u03C1\u03C1", SubstringI: 1, Result: 1},
	{Str: "\u03B5\u03BD\u03B4\u03B9\u03B1\u03C6\u03B5\u03C1", SubstringI: 1, Result: 1},
	{Str: "\u03C0\u03B1\u03C4", SubstringI: -1, Result: 1},
	{Str: "\u03BA\u03B1\u03B8\u03B1\u03C1\u03B5\u03C5", SubstringI: -1, Result: 1},
	{Str: "\u03B4\u03B5\u03C5\u03C4\u03B5\u03C1\u03B5\u03C5", SubstringI: -1, Result: 1},
	{Str: "\u03BB\u03B5\u03C7", SubstringI: -1, Result: 1},
	{Str: "\u03C4\u03C3\u03B1", SubstringI: -1, Result: 1},
	{Str: "\u03C7\u03B1\u03B4", SubstringI: -1, Result: 1},
	{Str: "\u03BC\u03B5\u03B4", SubstringI: -1, Result: 1},
	{Str: "\u03BB\u03B1\u03BC\u03C0\u03B9\u03B4", SubstringI: -1, Result: 1},
	{Str: "\u03B4\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03BB\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03BC\u03B5\u03C3\u03B1\u03B6", SubstringI: -1, Result: 1},
	{Str: "\u03B4\u03B5\u03C3\u03C0\u03BF\u03B6", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03B9\u03B8", SubstringI: -1, Result: 1},
	{Str: "\u03C6\u03B1\u03C1\u03BC\u03B1\u03BA", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03B3\u03BA", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03BD\u03B7\u03BA", SubstringI: -1, Result: 1},
	{Str: "\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03BC", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03BC", SubstringI: 21, Result: 1},
	{Str: "\u03B2\u03C1\u03BF\u03BC", SubstringI: 21, Result: 1},
	{Str: "\u03C5\u03C0\u03BF\u03C4\u03B5\u03B9\u03BD", SubstringI: -1, Result: 1},
}

var a58 = []snowball.Among[context]{
	{Str: "\u03BF\u03C5\u03C3\u03B5\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03BF\u03C5\u03C3\u03B1", SubstringI: -1, Result: 1},
	{Str: "\u03BF\u03C5\u03C3\u03B5", SubstringI: -1, Result: 1},
}

var a59 = []snowball.Among[context]{
	{Str: "\u03C8\u03BF\u03C6", SubstringI: -1, Result: -1},
	{Str: "\u03BD\u03B1\u03C5\u03BB\u03BF\u03C7", SubstringI: -1, Result: -1},
}

var a60 = []snowball.Among[context]{
	{Str: "\u03C1\u03C0", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03C6\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03C7\u03BF\u03C1\u03C4", SubstringI: -1, Result: 1},
	{Str: "\u03C3\u03C6", SubstringI: -1, Result: 1},
	{Str: "\u03BF\u03C6", SubstringI: -1, Result: 1},
	{Str: "\u03BB\u03BF\u03C7", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03B5\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03BB\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03C3\u03BC\u03B7\u03BD", SubstringI: -1, Result: 1},
}

var a61 = []snowball.Among[context]{
	{Str: "\u03C0", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03C3\u03C0", SubstringI: 0, Result: 1},
	{Str: "\u03B1\u03BD\u03C5\u03C0", SubstringI: 0, Result: 1},
	{Str: "\u03B1\u03C1\u03C4\u03B9\u03C0", SubstringI: 0, Result: 1},
	{Str: "\u03B1\u03B5\u03B9\u03C0", SubstringI: 0, Result: 1},
	{Str: "\u03C3\u03C5\u03BC\u03C0", SubstringI: 0, Result: 1},
	{Str: "\u03C0\u03C1\u03BF\u03C3\u03C9\u03C0\u03BF\u03C0", SubstringI: 0, Result: 1},
	{Str: "\u03C3\u03B9\u03B4\u03B7\u03C1\u03BF\u03C0", SubstringI: 0, Result: 1},
	{Str: "\u03B4\u03C1\u03BF\u03C3\u03BF\u03C0", SubstringI: 0, Result: 1},
	{Str: "\u03BD\u03B5\u03BF\u03C0", SubstringI: 0, Result: 1},
	{Str: "\u03BA\u03C1\u03BF\u03BA\u03B1\u03BB\u03BF\u03C0", SubstringI: 0, Result: 1},
	{Str: "\u03BF\u03BB\u03BF\u03C0", SubstringI: 0, Result: 1},
	{Str: "\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03C4\u03C1", SubstringI: 12, Result: 1},
	{Str: "\u03BF\u03C5\u03C1", SubstringI: 12, Result: 1},
	{Str: "\u03B1\u03C3\u03C0\u03B1\u03C1", SubstringI: 12, Result: 1},
	{Str: "\u03C7\u03B1\u03C1", SubstringI: 12, Result: 1},
	{Str: "\u03B1\u03C7\u03B1\u03C1", SubstringI: 16, Result: 1},
	{Str: "\u03B1\u03C0\u03B5\u03C1", SubstringI: 12, Result: 1},
	{Str: "\u03C4", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03BD\u03C5\u03C3\u03C4", SubstringI: 19, Result: 1},
	{Str: "\u03B1\u03B2\u03B1\u03C3\u03C4", SubstringI: 19, Result: 1},
	{Str: "\u03C0\u03C1\u03BF\u03C3\u03C4", SubstringI: 19, Result: 1},
	{Str: "\u03B1\u03B9\u03BC\u03BF\u03C3\u03C4", SubstringI: 19, Result: 1},
	{Str: "\u03B4\u03B9\u03B1\u03C4", SubstringI: 19, Result: 1},
	{Str: "\u03B5\u03C0\u03B9\u03C4", SubstringI: 19, Result: 1},
	{Str: "\u03C3\u03C5\u03BD\u03C4", SubstringI: 19, Result: 1},
	{Str: "\u03C5\u03C0\u03BF\u03C4", SubstringI: 19, Result: 1},
	{Str: "\u03B1\u03C0\u03BF\u03C4", SubstringI: 19, Result: 1},
	{Str: "\u03BF\u03BC\u03BF\u03C4", SubstringI: 19, Result: 1},
	{Str: "\u03BD\u03BF\u03BC\u03BF\u03C4", SubstringI: 29, Result: 1},
	{Str: "\u03BD\u03B1\u03C5", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03BF\u03BB\u03C5\u03C6", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03C6", SubstringI: -1, Result: 1},
	{Str: "\u03BE\u03B5\u03C6", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03B4\u03B7\u03C6", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03B1\u03BC\u03C6", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03BC\u03B1\u03BB\u03BB\u03B9", SubstringI: -1, Result: 1},
	{Str: "\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03BC\u03B1\u03BB", SubstringI: 38, Result: 1},
	{Str: "\u03BC", SubstringI: -1, Result: 1},
	{Str: "\u03BF\u03C5\u03BB\u03B1\u03BC", SubstringI: 40, Result: 1},
	{Str: "\u03B5\u03BD", SubstringI: -1, Result: 1},
	{Str: "\u03B4\u03B5\u03C1\u03B2\u03B5\u03BD", SubstringI: 42, Result: 1},
}

var a62 = []snowball.Among[context]{
	{Str: "\u03B1\u03B3\u03B5\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03B3\u03B1", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03B3\u03B5", SubstringI: -1, Result: 1},
}

var a63 = []snowball.Among[context]{
	{Str: "\u03B7\u03C3\u03BF\u03C5", SubstringI: -1, Result: 1},
	{Str: "\u03B7\u03C3\u03B1", SubstringI: -1, Result: 1},
	{Str: "\u03B7\u03C3\u03B5", SubstringI: -1, Result: 1},
}

var a64 = []snowball.Among[context]{
	{Str: "\u03BD", SubstringI: -1, Result: 1},
	{Str: "\u03B5\u03C0\u03C4\u03B1\u03BD", SubstringI: 0, Result: 1},
	{Str: "\u03B4\u03C9\u03B4\u03B5\u03BA\u03B1\u03BD", SubstringI: 0, Result: 1},
	{Str: "\u03C7\u03B5\u03C1\u03C3\u03BF\u03BD", SubstringI: 0, Result: 1},
	{Str: "\u03BC\u03B5\u03B3\u03B1\u03BB\u03BF\u03BD", SubstringI: 0, Result: 1},
	{Str: "\u03B5\u03C1\u03B7\u03BC\u03BF\u03BD", SubstringI: 0, Result: 1},
}

var a65 = []snowball.Among[context]{
	{Str: "\u03B7\u03C3\u03C4\u03B5", SubstringI: -1, Result: 1},
}

var a66 = []snowball.Among[context]{
	{Str: "\u03C7\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03B4\u03C5\u03C3\u03C7\u03C1", SubstringI: 0, Result: 1},
	{Str: "\u03B5\u03C5\u03C7\u03C1", SubstringI: 0, Result: 1},
	{Str: "\u03B1\u03C7\u03C1", SubstringI: 0, Result: 1},
	{Str: "\u03BA\u03BF\u03B9\u03BD\u03BF\u03C7\u03C1", SubstringI: 0, Result: 1},
	{Str: "\u03C0\u03B1\u03BB\u03B9\u03BC\u03C8", SubstringI: -1, Result: 1},
	{Str: "\u03C3\u03B2", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03C3\u03B2", SubstringI: 6, Result: 1},
	{Str: "\u03B1\u03C0\u03BB", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03B5\u03B9\u03BC\u03BD", SubstringI: -1, Result: 1},
}

var a67 = []snowball.Among[context]{
	{Str: "\u03BF\u03C5\u03BD\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B7\u03C3\u03BF\u03C5\u03BD\u03B5", SubstringI: 0, Result: 1},
	{Str: "\u03B7\u03B8\u03BF\u03C5\u03BD\u03B5", SubstringI: 0, Result: 1},
}

var a68 = []snowball.Among[context]{
	{Str: "\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03C3\u03C4\u03C1\u03B1\u03B2\u03BF\u03BC\u03BF\u03C5\u03C4\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03BA\u03B1\u03BA\u03BF\u03BC\u03BF\u03C5\u03C4\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03C3\u03C0\u03B9", SubstringI: -1, Result: 1},
	{Str: "\u03BD", SubstringI: -1, Result: 1},
	{Str: "\u03B5\u03BE\u03C9\u03BD", SubstringI: 4, Result: 1},
}

var a69 = []snowball.Among[context]{
	{Str: "\u03BF\u03C5\u03BC\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B7\u03C3\u03BF\u03C5\u03BC\u03B5", SubstringI: 0, Result: 1},
	{Str: "\u03B7\u03B8\u03BF\u03C5\u03BC\u03B5", SubstringI: 0, Result: 1},
}

var a70 = []snowball.Among[context]{
	{Str: "\u03B1\u03C3\u03BF\u03C5\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03C0\u03B1\u03C1\u03B1\u03C3\u03BF\u03C5\u03C3", SubstringI: 0, Result: 1},
	{Str: "\u03B1\u03BB\u03BB\u03BF\u03C3\u03BF\u03C5\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03C6", SubstringI: -1, Result: 1},
	{Str: "\u03C7", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03B6", SubstringI: -1, Result: 1},
	{Str: "\u03C9\u03C1\u03B9\u03BF\u03C0\u03BB", SubstringI: -1, Result: 1},
}

var a71 = []snowball.Among[context]{
	{Str: "\u03BC\u03B1\u03C4\u03BF\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03BC\u03B1\u03C4\u03B1", SubstringI: -1, Result: 1},
	{Str: "\u03BC\u03B1\u03C4\u03C9\u03BD", SubstringI: -1, Result: 1},
}

var a72 = []snowball.Among[context]{
	{Str: "\u03C5\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03BF\u03C5\u03C3", SubstringI: 0, Result: 1},
	{Str: "\u03B1\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03B5\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03B7\u03C3\u03B5\u03C3", SubstringI: 3, Result: 1},
	{Str: "\u03B7\u03B4\u03B5\u03C3", SubstringI: 3, Result: 1},
	{Str: "\u03B7\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03B5\u03B9\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03B7\u03B8\u03B5\u03B9\u03C3", SubstringI: 7, Result: 1},
	{Str: "\u03BF\u03C3", SubstringI: -1, Result: 1},
	{Str: "\u03C5", SubstringI: -1, Result: 1},
	{Str: "\u03BF\u03C5", SubstringI: 10, Result: 1},
	{Str: "\u03C9", SubstringI: -1, Result: 1},
	{Str: "\u03B7\u03C3\u03C9", SubstringI: 12, Result: 1},
	{Str: "\u03B1\u03C9", SubstringI: 12, Result: 1},
	{Str: "\u03B7\u03B8\u03C9", SubstringI: 12, Result: 1},
	{Str: "\u03B1", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03BF\u03C5\u03BC\u03B1", SubstringI: 16, Result: 1},
	{Str: "\u03BF\u03C3\u03BF\u03C5\u03BD\u03B1", SubstringI: 16, Result: 1},
	{Str: "\u03B9\u03BF\u03C3\u03BF\u03C5\u03BD\u03B1", SubstringI: 18, Result: 1},
	{Str: "\u03BF\u03BC\u03BF\u03C5\u03BD\u03B1", SubstringI: 16, Result: 1},
	{Str: "\u03B9\u03BF\u03BC\u03BF\u03C5\u03BD\u03B1", SubstringI: 20, Result: 1},
	{Str: "\u03B5", SubstringI: -1, Result: 1},
	{Str: "\u03B9\u03B5\u03C3\u03B1\u03C3\u03C4\u03B5", SubstringI: 22, Result: 1},
	{Str: "\u03BF\u03C3\u03B1\u03C3\u03C4\u03B5", SubstringI: 22, Result: 1},
	{Str: "\u03B9\u03BF\u03C3\u03B1\u03C3\u03C4\u03B5", SubstringI: 24, Result: 1},
	{Str: "\u03BF\u03C5\u03BC\u03B1\u03C3\u03C4\u03B5", SubstringI: 22, Result: 1},
	{Str: "\u03B9\u03BF\u03C5\u03BC\u03B1\u03C3\u03C4\u03B5", SubstringI: 26, Result: 1},
	{Str: "\u03B9\u03B5\u03BC\u03B1\u03C3\u03C4\u03B5", SubstringI: 22, Result: 1},
	{Str: "\u03BF\u03C5\u03C3\u03B1\u03C4\u03B5", SubstringI: 22, Result: 1},
	{Str: "\u03B7\u03C3\u03B1\u03C4\u03B5", SubstringI: 22, Result: 1},
	{Str: "\u03B1\u03B3\u03B1\u03C4\u03B5", SubstringI: 22, Result: 1},
	{Str: "\u03B7\u03BA\u03B1\u03C4\u03B5", SubstringI: 22, Result: 1},
	{Str: "\u03B7\u03B8\u03B7\u03BA\u03B1\u03C4\u03B5", SubstringI: 32, Result: 1},
	{Str: "\u03B5\u03B9\u03C4\u03B5", SubstringI: 22, Result: 1},
	{Str: "\u03B7\u03B8\u03B5\u03B9\u03C4\u03B5", SubstringI: 34, Result: 1},
	{Str: "\u03B7", SubstringI: -1, Result: 1},
	{Str: "\u03B9", SubstringI: -1, Result: 1},
	{Str: "\u03B1\u03C3\u03B1\u03B9", SubstringI: 37, Result: 1},
	{Str: "\u03B5\u03C3\u03B1\u03B9", SubstringI: 37, Result: 1},
	{Str: "\u03B9\u03B5\u03C3\u03B1\u03B9", SubstringI: 39, Result: 1},
	{Str: "\u03B1\u03C4\u03B1\u03B9", SubstringI: 37, Result: 1},
	{Str: "\u03B5\u03C4\u03B1\u03B9", SubstringI: 37, Result: 1},
	{Str: "\u03B9\u03B5\u03C4\u03B1\u03B9", SubstringI: 42, Result: 1},
	{Str: "\u03BF\u03C5\u03BD\u03C4\u03B1\u03B9", SubstringI: 37, Result: 1},
	{Str: "\u03B9\u03BF\u03C5\u03BD\u03C4\u03B1\u03B9", SubstringI: 44, Result: 1},
	{Str: "\u03BF\u03BD\u03C4\u03B1\u03B9", SubstringI: 37, Result: 1},
	{Str: "\u03BF\u03C5\u03BC\u03B1\u03B9", SubstringI: 37, Result: 1},
	{Str: "\u03B1\u03BC\u03B1\u03B9", SubstringI: 37, Result: 1},
	{Str: "\u03B9\u03B5\u03BC\u03B1\u03B9", SubstringI: 37, Result: 1},
	{Str: "\u03BF\u03BC\u03B1\u03B9", SubstringI: 37, Result: 1},
	{Str: "\u03B5\u03B9", SubstringI: 37, Result: 1},
	{Str: "\u03B7\u03C3\u03B5\u03B9", SubstringI: 51, Result: 1},
	{Str: "\u03B1\u03B5\u03B9", SubstringI: 51, Result: 1},
	{Str: "\u03B7\u03B8\u03B5\u03B9", SubstringI: 51, Result: 1},
	{Str: "\u03BF\u03B9", SubstringI: 37, Result: 1},
	{Str: "\u03BF\u03C5\u03BD", SubstringI: -1, Result: 1},
	{Str: "\u03B7\u03C3\u03BF\u03C5\u03BD", SubstringI: 56, Result: 1},
	{Str: "\u03BF\u03C3\u03BF\u03C5\u03BD", SubstringI: 56, Result: 1},
	{Str: "\u03B9\u03BF\u03C3\u03BF\u03C5\u03BD", SubstringI: 58, Result: 1},
	{Str: "\u03B7\u03B8\u03BF\u03C5\u03BD", SubstringI: 56, Result: 1},
	{Str: "\u03BF\u03BC\u03BF\u03C5\u03BD", SubstringI: 56, Result: 1},
	{Str: "\u03B9\u03BF\u03BC\u03BF\u03C5\u03BD", SubstringI: 61, Result: 1},
	{Str: "\u03C9\u03BD", SubstringI: -1, Result: 1},
	{Str: "\u03B7\u03B4\u03C9\u03BD", SubstringI: 63, Result: 1},
	{Str: "\u03B1\u03BD", SubstringI: -1, Result: 1},
	{Str: "\u03BF\u03C5\u03C3\u03B1\u03BD", SubstringI: 65, Result: 1},
	{Str: "\u03BF\u03BD\u03C4\u03BF\u03C5\u03C3\u03B1\u03BD", SubstringI: 66, Result: 1},
	{Str: "\u03B9\u03BF\u03BD\u03C4\u03BF\u03C5\u03C3\u03B1\u03BD", SubstringI: 67, Result: 1},
	{Str: "\u03B7\u03C3\u03B1\u03BD", SubstringI: 65, Result: 1},
	{Str: "\u03BF\u03C3\u03B1\u03C3\u03C4\u03B1\u03BD", SubstringI: 65, Result: 1},
	{Str: "\u03B9\u03BF\u03C3\u03B1\u03C3\u03C4\u03B1\u03BD", SubstringI: 70, Result: 1},
	{Str: "\u03BF\u03BC\u03B1\u03C3\u03C4\u03B1\u03BD", SubstringI: 65, Result: 1},
	{Str: "\u03B9\u03BF\u03BC\u03B1\u03C3\u03C4\u03B1\u03BD", SubstringI: 72, Result: 1},
	{Str: "\u03BF\u03C5\u03BD\u03C4\u03B1\u03BD", SubstringI: 65, Result: 1},
	{Str: "\u03B9\u03BF\u03C5\u03BD\u03C4\u03B1\u03BD", SubstringI: 74, Result: 1},
	{Str: "\u03BF\u03BD\u03C4\u03B1\u03BD", SubstringI: 65, Result: 1},
	{Str: "\u03B9\u03BF\u03BD\u03C4\u03B1\u03BD", SubstringI: 76, Result: 1},
	{Str: "\u03BF\u03C4\u03B1\u03BD", SubstringI: 65, Result: 1},
	{Str: "\u03B9\u03BF\u03C4\u03B1\u03BD", SubstringI: 78, Result: 1},
	{Str: "\u03B1\u03B3\u03B1\u03BD", SubstringI: 65, Result: 1},
	{Str: "\u03B7\u03BA\u03B1\u03BD", SubstringI: 65, Result: 1},
	{Str: "\u03B7\u03B8\u03B7\u03BA\u03B1\u03BD", SubstringI: 81, Result: 1},
	{Str: "\u03BF", SubstringI: -1, Result: 1},
}

var a73 = []snowball.Among[context]{
	{Str: "\u03B5\u03C3\u03C4\u03B5\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03C5\u03C4\u03B5\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03C9\u03C4\u03B5\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03BF\u03C4\u03B5\u03C1", SubstringI: -1, Result: 1},
	{Str: "\u03B5\u03C3\u03C4\u03B1\u03C4", SubstringI: -1, Result: 1},
	{Str: "\u03C5\u03C4\u03B1\u03C4", SubstringI: -1, Result: 1},
	{Str: "\u03C9\u03C4\u03B1\u03C4", SubstringI: -1, Result: 1},
	{Str: "\u03BF\u03C4\u03B1\u03C4", SubstringI: -1, Result: 1},
}

var gV = []byte{81, 65, 16, 1}
var gV2 = []byte{81, 65, 0, 1}

// hasMinLength reports whether the word has at least three characters, the
// length gate that must hold before any suffix stripping runs.
func hasMinLength(env *snowball.Env, ctx *context) bool {
	return env.RuneCount() >= 3
}

// tolower folds the trailing run of Greek capitals (and final sigma) to their
// lowercase forms, walking backwards over the whole word.
func tolower(env *snowball.Env, ctx *context) bool {
	var amongVar int32
replab0:
	for {
		v1 := env.Limit - env.Cursor
	lab1:
		for once := 0; once < 1; once++ {
			env.Ket = env.Cursor
			amongVar = snowball.FindAmongB(env, a0, ctx)
			if amongVar == 0 {
				break lab1
			}
			env.Bra = env.Cursor
			switch amongVar {
			case 1:
				if !env.SliceFrom("α") {
					return false
				}
			case 2:
				if !env.SliceFrom("β") {
					return false
				}
			case 3:
				if !env.SliceFrom("γ") {
					return false
				}
			case 4:
				if !env.SliceFrom("δ") {
					return false
				}
			case 5:
				if !env.SliceFrom("ε") {
					return false
				}
			case 6:
				if !env.SliceFrom("ζ") {
					return false
				}
			case 7:
				if !env.SliceFrom("η") {
					return false
				}
			case 8:
				if !env.SliceFrom("θ") {
					return false
				}
			case 9:
				if !env.SliceFrom("ι") {
					return false
				}
			case 10:
				if !env.SliceFrom("κ") {
					return false
				}
			case 11:
				if !env.SliceFrom("λ") {
					return false
				}
			case 12:
				if !env.SliceFrom("μ") {
					return false
				}
			case 13:
				if !env.SliceFrom("ν") {
					return false
				}
			case 14:
				if !env.SliceFrom("ξ") {
					return false
				}
			case 15:
				if !env.SliceFrom("ο") {
					return false
				}
			case 16:
				if !env.SliceFrom("π") {
					return false
				}
			case 17:
				if !env.SliceFrom("ρ") {
					return false
				}
			case 18:
				if !env.SliceFrom("σ") {
					return false
				}
			case 19:
				if !env.SliceFrom("τ") {
					return false
				}
			case 20:
				if !env.SliceFrom("υ") {
					return false
				}
			case 21:
				if !env.SliceFrom("φ") {
					return false
				}
			case 22:
				if !env.SliceFrom("χ") {
					return false
				}
			case 23:
				if !env.SliceFrom("ψ") {
					return false
				}
			case 24:
				if !env.SliceFrom("ω") {
					return false
				}
			case 25:
				if env.Cursor <= env.LimitBackward {
					break lab1
				}
				env.PreviousChar()
			}
			continue replab0
		}
		env.Cursor = env.Limit - v1
		break replab0
	}
	return true
}

// step1 rewrites the irregular nominal stems in a1 to their canonical forms and
// clears test1 so the later steps are skipped.
func step1(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a1, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	switch amongVar {
	case 1:
		if !env.SliceFrom("φα") {
			return false
		}
	case 2:
		if !env.SliceFrom("σκα") {
			return false
		}
	case 3:
		if !env.SliceFrom("ολο") {
			return false
		}
	case 4:
		if !env.SliceFrom("σο") {
			return false
		}
	case 5:
		if !env.SliceFrom("τατο") {
			return false
		}
	case 6:
		if !env.SliceFrom("κρε") {
			return false
		}
	case 7:
		if !env.SliceFrom("περ") {
			return false
		}
	case 8:
		if !env.SliceFrom("τερ") {
			return false
		}
	case 9:
		if !env.SliceFrom("φω") {
			return false
		}
	case 10:
		if !env.SliceFrom("καθεστ") {
			return false
		}
	case 11:
		if !env.SliceFrom("γεγον") {
			return false
		}
	}
	ctx.bTest1 = false
	return true
}

// steps1 strips the -ιζ verbal family (a4), then restores an -ι or -ιζ ending
// for the prefix stems captured in a2/a3.
func steps1(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a4, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !env.SliceDel() {
		return false
	}
	ctx.bTest1 = false
lab0:
	for {
		v1 := env.Limit - env.Cursor
	lab1:
		for {
			env.Ket = env.Cursor
			if snowball.FindAmongB(env, a2, ctx) == 0 {
				break lab1
			}
			env.Bra = env.Cursor
			if env.Cursor > env.LimitBackward {
				break lab1
			}
			ctx.sS = env.SliceTo()
			if ctx.sS == "" {
				return false
			}
			if !env.SliceFrom("ι") {
				return false
			}
			c := env.Cursor
			env.Insert(env.Cursor, env.Cursor, ctx.sS)
			env.Cursor = c
			break lab0
		}
		env.Cursor = env.Limit - v1
		env.Ket = env.Cursor
		if snowball.FindAmongB(env, a3, ctx) == 0 {
			return false
		}
		env.Bra = env.Cursor
		if env.Cursor > env.LimitBackward {
			return false
		}
		ctx.sS = env.SliceTo()
		if ctx.sS == "" {
			return false
		}
		if !env.SliceFrom("ιζ") {
			return false
		}
		c := env.Cursor
		env.Insert(env.Cursor, env.Cursor, ctx.sS)
		env.Cursor = c
		break lab0
	}
	return true
}

// steps2 strips the -ωσ verbal family (a6) and restores -ων for the captured
// prefix stems in a5.
func steps2(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a6, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !env.SliceDel() {
		return false
	}
	ctx.bTest1 = false
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a5, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if env.Cursor > env.LimitBackward {
		return false
	}
	ctx.sS = env.SliceTo()
	if ctx.sS == "" {
		return false
	}
	if !env.SliceFrom("ων") {
		return false
	}
	c := env.Cursor
	env.Insert(env.Cursor, env.Cursor, ctx.sS)
	env.Cursor = c
	return true
}

// steps3 strips the -ισ verbal family (a9), then handles the -ισα atlimit case
// and the captured-stem restorations for a7/a8.
func steps3(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a9, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !env.SliceDel() {
		return false
	}
	ctx.bTest1 = false
lab0:
	for {
		v1 := env.Limit - env.Cursor
	lab1:
		for {
			if !env.EqSB("ισα") {
				break lab1
			}
			if env.Cursor > env.LimitBackward {
				break lab1
			}
			if !env.SliceFrom("ισ") {
				return false
			}
			break lab0
		}
		env.Cursor = env.Limit - v1
	lab2:
		for {
			env.Ket = env.Cursor
			if snowball.FindAmongB(env, a7, ctx) == 0 {
				break lab2
			}
			env.Bra = env.Cursor
			if env.Cursor > env.LimitBackward {
				break lab2
			}
			ctx.sS = env.SliceTo()
			if ctx.sS == "" {
				return false
			}
			if !env.SliceFrom("ι") {
				return false
			}
			c := env.Cursor
			env.Insert(env.Cursor, env.Cursor, ctx.sS)
			env.Cursor = c
			break lab0
		}
		env.Cursor = env.Limit - v1
		env.Ket = env.Cursor
		if snowball.FindAmongB(env, a8, ctx) == 0 {
			return false
		}
		env.Bra = env.Cursor
		if env.Cursor > env.LimitBackward {
			return false
		}
		ctx.sS = env.SliceTo()
		if ctx.sS == "" {
			return false
		}
		if !env.SliceFrom("ισ") {
			return false
		}
		c := env.Cursor
		env.Insert(env.Cursor, env.Cursor, ctx.sS)
		env.Cursor = c
		break lab0
	}
	return true
}

// steps4 strips the -ισ family in a11 and restores -ι for the captured prefix
// stems in a10.
func steps4(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a11, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !env.SliceDel() {
		return false
	}
	ctx.bTest1 = false
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a10, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if env.Cursor > env.LimitBackward {
		return false
	}
	ctx.sS = env.SliceTo()
	if ctx.sS == "" {
		return false
	}
	if !env.SliceFrom("ι") {
		return false
	}
	c := env.Cursor
	env.Insert(env.Cursor, env.Cursor, ctx.sS)
	env.Cursor = c
	return true
}

// steps5 strips the -ιστ adjectival family (a14) and restores -ι or -ιστ for the
// captured prefix stems in a12/a13.
func steps5(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a14, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !env.SliceDel() {
		return false
	}
	ctx.bTest1 = false
lab0:
	for {
		v1 := env.Limit - env.Cursor
	lab1:
		for {
			env.Ket = env.Cursor
			if snowball.FindAmongB(env, a12, ctx) == 0 {
				break lab1
			}
			env.Bra = env.Cursor
			if env.Cursor > env.LimitBackward {
				break lab1
			}
			ctx.sS = env.SliceTo()
			if ctx.sS == "" {
				return false
			}
			if !env.SliceFrom("ι") {
				return false
			}
			c := env.Cursor
			env.Insert(env.Cursor, env.Cursor, ctx.sS)
			env.Cursor = c
			break lab0
		}
		env.Cursor = env.Limit - v1
		env.Ket = env.Cursor
		if snowball.FindAmongB(env, a13, ctx) == 0 {
			return false
		}
		env.Bra = env.Cursor
		if env.Cursor > env.LimitBackward {
			return false
		}
		ctx.sS = env.SliceTo()
		if ctx.sS == "" {
			return false
		}
		if !env.SliceFrom("ιστ") {
			return false
		}
		c := env.Cursor
		env.Insert(env.Cursor, env.Cursor, ctx.sS)
		env.Cursor = c
		break lab0
	}
	return true
}

// steps6 strips the -ισμ family (a18), restoring -ισμ or -ι for the captured
// prefix stems in a15/a16, and rewriting the irregular -ιστικ stems in a17.
func steps6(env *snowball.Env, ctx *context) bool {
	var amongVar int32
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a18, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !env.SliceDel() {
		return false
	}
	ctx.bTest1 = false
lab0:
	for {
		v1 := env.Limit - env.Cursor
	lab1:
		for {
			env.Ket = env.Cursor
			if snowball.FindAmongB(env, a15, ctx) == 0 {
				break lab1
			}
			env.Bra = env.Cursor
			if env.Cursor > env.LimitBackward {
				break lab1
			}
			ctx.sS = env.SliceTo()
			if ctx.sS == "" {
				return false
			}
			if !env.SliceFrom("ισμ") {
				return false
			}
			c := env.Cursor
			env.Insert(env.Cursor, env.Cursor, ctx.sS)
			env.Cursor = c
			break lab0
		}
		env.Cursor = env.Limit - v1
	lab2:
		for {
			env.Ket = env.Cursor
			if snowball.FindAmongB(env, a16, ctx) == 0 {
				break lab2
			}
			env.Bra = env.Cursor
			if env.Cursor > env.LimitBackward {
				break lab2
			}
			ctx.sS = env.SliceTo()
			if ctx.sS == "" {
				return false
			}
			if !env.SliceFrom("ι") {
				return false
			}
			c := env.Cursor
			env.Insert(env.Cursor, env.Cursor, ctx.sS)
			env.Cursor = c
			break lab0
		}
		env.Cursor = env.Limit - v1
		env.Ket = env.Cursor
		amongVar = snowball.FindAmongB(env, a17, ctx)
		if amongVar == 0 {
			return false
		}
		env.Bra = env.Cursor
		switch amongVar {
		case 1:
			if !env.SliceFrom("αγνωστ") {
				return false
			}
		case 2:
			if !env.SliceFrom("ατομ") {
				return false
			}
		case 3:
			if !env.SliceFrom("γνωστ") {
				return false
			}
		case 4:
			if !env.SliceFrom("εθν") {
				return false
			}
		case 5:
			if !env.SliceFrom("εκλεκτ") {
				return false
			}
		case 6:
			if !env.SliceFrom("σκεπτ") {
				return false
			}
		case 7:
			if !env.SliceFrom("τοπ") {
				return false
			}
		case 8:
			if !env.SliceFrom("αλεξανδρ") {
				return false
			}
		case 9:
			if !env.SliceFrom("βυζαντ") {
				return false
			}
		case 10:
			if !env.SliceFrom("θεατρ") {
				return false
			}
		}
		break lab0
	}
	return true
}

// steps7 strips the -αρακι diminutive family (a20) and restores -αρακ for the
// captured prefix stems in a19.
func steps7(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a20, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !env.SliceDel() {
		return false
	}
	ctx.bTest1 = false
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a19, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if env.Cursor > env.LimitBackward {
		return false
	}
	ctx.sS = env.SliceTo()
	if ctx.sS == "" {
		return false
	}
	if !env.SliceFrom("αρακ") {
		return false
	}
	c := env.Cursor
	env.Insert(env.Cursor, env.Cursor, ctx.sS)
	env.Cursor = c
	return true
}

// steps8 strips the -ιτσ diminutive family (a23), restoring -ακ or -ιτσ for the
// captured prefix stems in a21/a22 and the literal κορ stem.
func steps8(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a23, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !env.SliceDel() {
		return false
	}
	ctx.bTest1 = false
lab0:
	for {
		v1 := env.Limit - env.Cursor
	lab1:
		for {
			env.Ket = env.Cursor
			if snowball.FindAmongB(env, a21, ctx) == 0 {
				break lab1
			}
			env.Bra = env.Cursor
			if env.Cursor > env.LimitBackward {
				break lab1
			}
			ctx.sS = env.SliceTo()
			if ctx.sS == "" {
				return false
			}
			if !env.SliceFrom("ακ") {
				return false
			}
			c := env.Cursor
			env.Insert(env.Cursor, env.Cursor, ctx.sS)
			env.Cursor = c
			break lab0
		}
		env.Cursor = env.Limit - v1
	lab2:
		for {
			env.Ket = env.Cursor
			if snowball.FindAmongB(env, a22, ctx) == 0 {
				break lab2
			}
			env.Bra = env.Cursor
			if env.Cursor > env.LimitBackward {
				break lab2
			}
			ctx.sS = env.SliceTo()
			if ctx.sS == "" {
				return false
			}
			if !env.SliceFrom("ιτσ") {
				return false
			}
			c := env.Cursor
			env.Insert(env.Cursor, env.Cursor, ctx.sS)
			env.Cursor = c
			break lab0
		}
		env.Cursor = env.Limit - v1
		env.Ket = env.Cursor
		if !env.EqSB("κορ") {
			return false
		}
		env.Bra = env.Cursor
		ctx.sS = env.SliceTo()
		if ctx.sS == "" {
			return false
		}
		if !env.SliceFrom("ιτσ") {
			return false
		}
		c := env.Cursor
		env.Insert(env.Cursor, env.Cursor, ctx.sS)
		env.Cursor = c
		break lab0
	}
	return true
}

// steps9 strips the -ιδι diminutive family (a26) and restores -ιδ for the
// captured prefix stems in a24/a25.
func steps9(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a26, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !env.SliceDel() {
		return false
	}
	ctx.bTest1 = false
lab0:
	for {
		v1 := env.Limit - env.Cursor
	lab1:
		for {
			env.Ket = env.Cursor
			if snowball.FindAmongB(env, a24, ctx) == 0 {
				break lab1
			}
			env.Bra = env.Cursor
			if env.Cursor > env.LimitBackward {
				break lab1
			}
			ctx.sS = env.SliceTo()
			if ctx.sS == "" {
				return false
			}
			if !env.SliceFrom("ιδ") {
				return false
			}
			c := env.Cursor
			env.Insert(env.Cursor, env.Cursor, ctx.sS)
			env.Cursor = c
			break lab0
		}
		env.Cursor = env.Limit - v1
		env.Ket = env.Cursor
		if snowball.FindAmongB(env, a25, ctx) == 0 {
			return false
		}
		env.Bra = env.Cursor
		ctx.sS = env.SliceTo()
		if ctx.sS == "" {
			return false
		}
		if !env.SliceFrom("ιδ") {
			return false
		}
		c := env.Cursor
		env.Insert(env.Cursor, env.Cursor, ctx.sS)
		env.Cursor = c
		break lab0
	}
	return true
}

// steps10 strips the -ισκ diminutive family (a28) and restores -ισκ for the
// captured prefix stems in a27.
func steps10(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a28, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !env.SliceDel() {
		return false
	}
	ctx.bTest1 = false
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a27, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if env.Cursor > env.LimitBackward {
		return false
	}
	ctx.sS = env.SliceTo()
	if ctx.sS == "" {
		return false
	}
	if !env.SliceFrom("ισκ") {
		return false
	}
	c := env.Cursor
	env.Insert(env.Cursor, env.Cursor, ctx.sS)
	env.Cursor = c
	return true
}

// step2a deletes an -αδες/-αδων ending (a29) and, unless the stem is one of the
// kinship exceptions in a30, reinserts -αδ.
func step2a(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a29, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !env.SliceDel() {
		return false
	}
	v1 := env.Limit - env.Cursor
lab0:
	for {
		env.Ket = env.Cursor
		if snowball.FindAmongB(env, a30, ctx) == 0 {
			break lab0
		}
		env.Bra = env.Cursor
		return false
	}
	env.Cursor = env.Limit - v1
	c := env.Cursor
	env.Insert(env.Cursor, env.Cursor, "αδ")
	env.Cursor = c
	return true
}

// step2b deletes an -εδες/-εδων ending (a31) and restores -εδ for the captured
// prefix stems in a32.
func step2b(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a31, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !env.SliceDel() {
		return false
	}
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a32, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	ctx.sS = env.SliceTo()
	if ctx.sS == "" {
		return false
	}
	if !env.SliceFrom("εδ") {
		return false
	}
	c := env.Cursor
	env.Insert(env.Cursor, env.Cursor, ctx.sS)
	env.Cursor = c
	return true
}

// step2c deletes an -ουδες/-ουδων ending (a33) and restores -ουδ for the
// captured prefix stems in a34.
func step2c(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a33, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !env.SliceDel() {
		return false
	}
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a34, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	ctx.sS = env.SliceTo()
	if ctx.sS == "" {
		return false
	}
	if !env.SliceFrom("ουδ") {
		return false
	}
	c := env.Cursor
	env.Insert(env.Cursor, env.Cursor, ctx.sS)
	env.Cursor = c
	return true
}

// step2d deletes an -εως/-εων ending (a35) and restores -ε for the captured
// prefix stems in a36.
func step2d(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a35, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !env.SliceDel() {
		return false
	}
	ctx.bTest1 = false
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a36, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if env.Cursor > env.LimitBackward {
		return false
	}
	ctx.sS = env.SliceTo()
	if ctx.sS == "" {
		return false
	}
	if !env.SliceFrom("ε") {
		return false
	}
	c := env.Cursor
	env.Insert(env.Cursor, env.Cursor, ctx.sS)
	env.Cursor = c
	return true
}

// step3 deletes an -ια/-ιου/-ιων ending (a37) and, when the stem now ends in a
// vowel, restores a final -ι.
func step3(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a37, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !env.SliceDel() {
		return false
	}
	ctx.bTest1 = false
	env.Ket = env.Cursor
	if !env.InGroupingB(gV, 945, 969) {
		return false
	}
	env.Bra = env.Cursor
	ctx.sS = env.SliceTo()
	if ctx.sS == "" {
		return false
	}
	if !env.SliceFrom("ι") {
		return false
	}
	c := env.Cursor
	env.Insert(env.Cursor, env.Cursor, ctx.sS)
	env.Cursor = c
	return true
}

// step4 deletes an -ικα/-ικου/-ικων ending (a38) and restores -ικ either after a
// vowel or for the captured prefix stems in a39.
func step4(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a38, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !env.SliceDel() {
		return false
	}
	ctx.bTest1 = false
lab0:
	for {
		v1 := env.Limit - env.Cursor
	lab1:
		for {
			env.Ket = env.Cursor
			if !env.InGroupingB(gV, 945, 969) {
				break lab1
			}
			env.Bra = env.Cursor
			ctx.sS = env.SliceTo()
			if ctx.sS == "" {
				return false
			}
			if !env.SliceFrom("ικ") {
				return false
			}
			c := env.Cursor
			env.Insert(env.Cursor, env.Cursor, ctx.sS)
			env.Cursor = c
			break lab0
		}
		env.Cursor = env.Limit - v1
		env.Ket = env.Cursor
		break lab0
	}
	if snowball.FindAmongB(env, a39, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if env.Cursor > env.LimitBackward {
		return false
	}
	ctx.sS = env.SliceTo()
	if ctx.sS == "" {
		return false
	}
	if !env.SliceFrom("ικ") {
		return false
	}
	c := env.Cursor
	env.Insert(env.Cursor, env.Cursor, ctx.sS)
	env.Cursor = c
	return true
}

// step5a handles the -αγαμε/-αμε past-tense endings, deleting the optional
// extensions in a40 and restoring -αμ for the captured prefix stems in a41.
func step5a(env *snowball.Env, ctx *context) bool {
	v1 := env.Limit - env.Cursor
lab0:
	for {
		if !env.EqSB("αγαμε") {
			break lab0
		}
		if env.Cursor > env.LimitBackward {
			break lab0
		}
		if !env.SliceFrom("αγαμ") {
			return false
		}
		break lab0
	}
	env.Cursor = env.Limit - v1
	v2 := env.Limit - env.Cursor
lab1:
	for {
		env.Ket = env.Cursor
		if snowball.FindAmongB(env, a40, ctx) == 0 {
			break lab1
		}
		env.Bra = env.Cursor
		if !env.SliceDel() {
			return false
		}
		ctx.bTest1 = false
		break lab1
	}
	env.Cursor = env.Limit - v2
	env.Ket = env.Cursor
	if !env.EqSB("αμε") {
		return false
	}
	env.Bra = env.Cursor
	if !env.SliceDel() {
		return false
	}
	ctx.bTest1 = false
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a41, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if env.Cursor > env.LimitBackward {
		return false
	}
	ctx.sS = env.SliceTo()
	if ctx.sS == "" {
		return false
	}
	if !env.SliceFrom("αμ") {
		return false
	}
	c := env.Cursor
	env.Insert(env.Cursor, env.Cursor, ctx.sS)
	env.Cursor = c
	return true
}

// step5b handles the -αγανε/-ανε past-tense endings, restoring -αγαν for the
// captured prefix stems in a43/a42 and -αν after a vowel or the stems in a44.
func step5b(env *snowball.Env, ctx *context) bool {
	v1 := env.Limit - env.Cursor
lab0:
	for {
		env.Ket = env.Cursor
		if snowball.FindAmongB(env, a43, ctx) == 0 {
			break lab0
		}
		env.Bra = env.Cursor
		if !env.SliceDel() {
			return false
		}
		ctx.bTest1 = false
		env.Ket = env.Cursor
		if snowball.FindAmongB(env, a42, ctx) == 0 {
			break lab0
		}
		env.Bra = env.Cursor
		if env.Cursor > env.LimitBackward {
			break lab0
		}
		ctx.sS = env.SliceTo()
		if ctx.sS == "" {
			return false
		}
		if !env.SliceFrom("αγαν") {
			return false
		}
		c := env.Cursor
		env.Insert(env.Cursor, env.Cursor, ctx.sS)
		env.Cursor = c
		break lab0
	}
	env.Cursor = env.Limit - v1
	env.Ket = env.Cursor
	if !env.EqSB("ανε") {
		return false
	}
	env.Bra = env.Cursor
	if !env.SliceDel() {
		return false
	}
	ctx.bTest1 = false
lab1:
	for {
		v2 := env.Limit - env.Cursor
	lab2:
		for {
			env.Ket = env.Cursor
			if !env.InGroupingB(gV2, 945, 969) {
				break lab2
			}
			env.Bra = env.Cursor
			ctx.sS = env.SliceTo()
			if ctx.sS == "" {
				return false
			}
			if !env.SliceFrom("αν") {
				return false
			}
			c := env.Cursor
			env.Insert(env.Cursor, env.Cursor, ctx.sS)
			env.Cursor = c
			break lab1
		}
		env.Cursor = env.Limit - v2
		env.Ket = env.Cursor
		break lab1
	}
	if snowball.FindAmongB(env, a44, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if env.Cursor > env.LimitBackward {
		return false
	}
	ctx.sS = env.SliceTo()
	if ctx.sS == "" {
		return false
	}
	if !env.SliceFrom("αν") {
		return false
	}
	c := env.Cursor
	env.Insert(env.Cursor, env.Cursor, ctx.sS)
	env.Cursor = c
	return true
}

// step5c handles the -ετε past-tense ending, deleting the optional -η”σετε
// extension in a45 and restoring -ετ after a vowel or for the captured prefix
// stems in a46/a47.
func step5c(env *snowball.Env, ctx *context) bool {
	v1 := env.Limit - env.Cursor
lab0:
	for {
		env.Ket = env.Cursor
		if snowball.FindAmongB(env, a45, ctx) == 0 {
			break lab0
		}
		env.Bra = env.Cursor
		if !env.SliceDel() {
			return false
		}
		ctx.bTest1 = false
		break lab0
	}
	env.Cursor = env.Limit - v1
	env.Ket = env.Cursor
	if !env.EqSB("ετε") {
		return false
	}
	env.Bra = env.Cursor
	if !env.SliceDel() {
		return false
	}
	ctx.bTest1 = false
lab1:
	for {
		v2 := env.Limit - env.Cursor
	lab2:
		for {
			env.Ket = env.Cursor
			if !env.InGroupingB(gV2, 945, 969) {
				break lab2
			}
			env.Bra = env.Cursor
			ctx.sS = env.SliceTo()
			if ctx.sS == "" {
				return false
			}
			if !env.SliceFrom("ετ") {
				return false
			}
			c := env.Cursor
			env.Insert(env.Cursor, env.Cursor, ctx.sS)
			env.Cursor = c
			break lab1
		}
		env.Cursor = env.Limit - v2
	lab3:
		for {
			env.Ket = env.Cursor
			if snowball.FindAmongB(env, a46, ctx) == 0 {
				break lab3
			}
			env.Bra = env.Cursor
			ctx.sS = env.SliceTo()
			if ctx.sS == "" {
				return false
			}
			if !env.SliceFrom("ετ") {
				return false
			}
			c := env.Cursor
			env.Insert(env.Cursor, env.Cursor, ctx.sS)
			env.Cursor = c
			break lab1
		}
		env.Cursor = env.Limit - v2
		env.Ket = env.Cursor
		break lab1
	}
	if snowball.FindAmongB(env, a47, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if env.Cursor > env.LimitBackward {
		return false
	}
	ctx.sS = env.SliceTo()
	if ctx.sS == "" {
		return false
	}
	if !env.SliceFrom("ετ") {
		return false
	}
	c := env.Cursor
	env.Insert(env.Cursor, env.Cursor, ctx.sS)
	env.Cursor = c
	return true
}

// step5d deletes an -οντας/-ωντας participle ending (a48) and restores -οντ
// after the literal αρχ stem or -ωντ after κρε.
func step5d(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a48, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !env.SliceDel() {
		return false
	}
	ctx.bTest1 = false
lab0:
	for {
		v1 := env.Limit - env.Cursor
	lab1:
		for {
			env.Ket = env.Cursor
			if !env.EqSB("αρχ") {
				break lab1
			}
			env.Bra = env.Cursor
			if env.Cursor > env.LimitBackward {
				break lab1
			}
			ctx.sS = env.SliceTo()
			if ctx.sS == "" {
				return false
			}
			if !env.SliceFrom("οντ") {
				return false
			}
			c := env.Cursor
			env.Insert(env.Cursor, env.Cursor, ctx.sS)
			env.Cursor = c
			break lab0
		}
		env.Cursor = env.Limit - v1
		env.Ket = env.Cursor
		if !env.EqSB("κρε") {
			return false
		}
		env.Bra = env.Cursor
		ctx.sS = env.SliceTo()
		if ctx.sS == "" {
			return false
		}
		if !env.SliceFrom("ωντ") {
			return false
		}
		c := env.Cursor
		env.Insert(env.Cursor, env.Cursor, ctx.sS)
		env.Cursor = c
		break lab0
	}
	return true
}

// step5e deletes an -ομαστε/-ιομαστε ending (a49) and restores -ομαστ after the
// literal ον stem.
func step5e(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a49, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !env.SliceDel() {
		return false
	}
	ctx.bTest1 = false
	env.Ket = env.Cursor
	if !env.EqSB("ον") {
		return false
	}
	env.Bra = env.Cursor
	if env.Cursor > env.LimitBackward {
		return false
	}
	ctx.sS = env.SliceTo()
	if ctx.sS == "" {
		return false
	}
	if !env.SliceFrom("ομαστ") {
		return false
	}
	c := env.Cursor
	env.Insert(env.Cursor, env.Cursor, ctx.sS)
	env.Cursor = c
	return true
}

// step5f handles the -ιεστε/-εστε endings, restoring -ιεστ for the captured
// prefix stems in a50/a51.
func step5f(env *snowball.Env, ctx *context) bool {
	v1 := env.Limit - env.Cursor
lab0:
	for {
		env.Ket = env.Cursor
		if !env.EqSB("ιεστε") {
			break lab0
		}
		env.Bra = env.Cursor
		if !env.SliceDel() {
			return false
		}
		ctx.bTest1 = false
		env.Ket = env.Cursor
		if snowball.FindAmongB(env, a50, ctx) == 0 {
			break lab0
		}
		env.Bra = env.Cursor
		if env.Cursor > env.LimitBackward {
			break lab0
		}
		ctx.sS = env.SliceTo()
		if ctx.sS == "" {
			return false
		}
		if !env.SliceFrom("ιεστ") {
			return false
		}
		c := env.Cursor
		env.Insert(env.Cursor, env.Cursor, ctx.sS)
		env.Cursor = c
		break lab0
	}
	env.Cursor = env.Limit - v1
	env.Ket = env.Cursor
	if !env.EqSB("εστε") {
		return false
	}
	env.Bra = env.Cursor
	if !env.SliceDel() {
		return false
	}
	ctx.bTest1 = false
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a51, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if env.Cursor > env.LimitBackward {
		return false
	}
	ctx.sS = env.SliceTo()
	if ctx.sS == "" {
		return false
	}
	if !env.SliceFrom("ιεστ") {
		return false
	}
	c := env.Cursor
	env.Insert(env.Cursor, env.Cursor, ctx.sS)
	env.Cursor = c
	return true
}

// step5g handles the -ηκες/-ηκα/-ηκε endings, deleting the optional -ηθηκ
// extension in a52 and restoring -ηκ for the captured prefix stems in a53/a54.
func step5g(env *snowball.Env, ctx *context) bool {
	v1 := env.Limit - env.Cursor
lab0:
	for {
		env.Ket = env.Cursor
		if snowball.FindAmongB(env, a52, ctx) == 0 {
			break lab0
		}
		env.Bra = env.Cursor
		if !env.SliceDel() {
			return false
		}
		ctx.bTest1 = false
		break lab0
	}
	env.Cursor = env.Limit - v1
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a55, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !env.SliceDel() {
		return false
	}
	ctx.bTest1 = false
lab1:
	for {
		v2 := env.Limit - env.Cursor
	lab2:
		for {
			env.Ket = env.Cursor
			if snowball.FindAmongB(env, a53, ctx) == 0 {
				break lab2
			}
			env.Bra = env.Cursor
			ctx.sS = env.SliceTo()
			if ctx.sS == "" {
				return false
			}
			if !env.SliceFrom("ηκ") {
				return false
			}
			c := env.Cursor
			env.Insert(env.Cursor, env.Cursor, ctx.sS)
			env.Cursor = c
			break lab1
		}
		env.Cursor = env.Limit - v2
		env.Ket = env.Cursor
		if snowball.FindAmongB(env, a54, ctx) == 0 {
			return false
		}
		env.Bra = env.Cursor
		if env.Cursor > env.LimitBackward {
			return false
		}
		ctx.sS = env.SliceTo()
		if ctx.sS == "" {
			return false
		}
		if !env.SliceFrom("ηκ") {
			return false
		}
		c := env.Cursor
		env.Insert(env.Cursor, env.Cursor, ctx.sS)
		env.Cursor = c
		break lab1
	}
	return true
}

// step5h deletes an -ουσες/-ουσα/-ουσε ending (a58) and restores -ους after the
// captured prefix stems in a56/a57.
func step5h(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a58, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !env.SliceDel() {
		return false
	}
	ctx.bTest1 = false
lab0:
	for {
		v1 := env.Limit - env.Cursor
	lab1:
		for {
			env.Ket = env.Cursor
			if snowball.FindAmongB(env, a56, ctx) == 0 {
				break lab1
			}
			env.Bra = env.Cursor
			ctx.sS = env.SliceTo()
			if ctx.sS == "" {
				return false
			}
			if !env.SliceFrom("ουσ") {
				return false
			}
			c := env.Cursor
			env.Insert(env.Cursor, env.Cursor, ctx.sS)
			env.Cursor = c
			break lab0
		}
		env.Cursor = env.Limit - v1
		env.Ket = env.Cursor
		if snowball.FindAmongB(env, a57, ctx) == 0 {
			return false
		}
		env.Bra = env.Cursor
		if env.Cursor > env.LimitBackward {
			return false
		}
		ctx.sS = env.SliceTo()
		if ctx.sS == "" {
			return false
		}
		if !env.SliceFrom("ουσ") {
			return false
		}
		c := env.Cursor
		env.Insert(env.Cursor, env.Cursor, ctx.sS)
		env.Cursor = c
		break lab0
	}
	return true
}

// step5i deletes an -αγες/-αγα/-αγε ending (a62) and restores -αγ after the
// literal κολλ stem or, unless blocked by a59, for the captured prefix stems in
// a60/a61.
func step5i(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a62, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !env.SliceDel() {
		return false
	}
	ctx.bTest1 = false
lab0:
	for {
		v1 := env.Limit - env.Cursor
	lab1:
		for {
			env.Ket = env.Cursor
			if !env.EqSB("κολλ") {
				break lab1
			}
			env.Bra = env.Cursor
			ctx.sS = env.SliceTo()
			if ctx.sS == "" {
				return false
			}
			if !env.SliceFrom("αγ") {
				return false
			}
			c := env.Cursor
			env.Insert(env.Cursor, env.Cursor, ctx.sS)
			env.Cursor = c
			break lab0
		}
		env.Cursor = env.Limit - v1
		v2 := env.Limit - env.Cursor
	lab2:
		for {
			env.Ket = env.Cursor
			if snowball.FindAmongB(env, a59, ctx) == 0 {
				break lab2
			}
			env.Bra = env.Cursor
			return false
		}
		env.Cursor = env.Limit - v2
	lab3:
		for {
			v3 := env.Limit - env.Cursor
		lab4:
			for {
				env.Ket = env.Cursor
				if snowball.FindAmongB(env, a60, ctx) == 0 {
					break lab4
				}
				env.Bra = env.Cursor
				ctx.sS = env.SliceTo()
				if ctx.sS == "" {
					return false
				}
				if !env.SliceFrom("αγ") {
					return false
				}
				c := env.Cursor
				env.Insert(env.Cursor, env.Cursor, ctx.sS)
				env.Cursor = c
				break lab3
			}
			env.Cursor = env.Limit - v3
			env.Ket = env.Cursor
			if snowball.FindAmongB(env, a61, ctx) == 0 {
				return false
			}
			env.Bra = env.Cursor
			if env.Cursor > env.LimitBackward {
				return false
			}
			ctx.sS = env.SliceTo()
			if ctx.sS == "" {
				return false
			}
			if !env.SliceFrom("αγ") {
				return false
			}
			c := env.Cursor
			env.Insert(env.Cursor, env.Cursor, ctx.sS)
			env.Cursor = c
			break lab3
		}
		break lab0
	}
	return true
}

// step5j deletes an -ησου/-ησα/-ησε ending (a63) and restores -ησ for the
// captured prefix stems in a64.
func step5j(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a63, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !env.SliceDel() {
		return false
	}
	ctx.bTest1 = false
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a64, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if env.Cursor > env.LimitBackward {
		return false
	}
	ctx.sS = env.SliceTo()
	if ctx.sS == "" {
		return false
	}
	if !env.SliceFrom("ησ") {
		return false
	}
	c := env.Cursor
	env.Insert(env.Cursor, env.Cursor, ctx.sS)
	env.Cursor = c
	return true
}

// step5k deletes an -ηστε ending (a65) and restores -ηστ for the captured prefix
// stems in a66.
func step5k(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a65, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !env.SliceDel() {
		return false
	}
	ctx.bTest1 = false
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a66, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if env.Cursor > env.LimitBackward {
		return false
	}
	ctx.sS = env.SliceTo()
	if ctx.sS == "" {
		return false
	}
	if !env.SliceFrom("ηστ") {
		return false
	}
	c := env.Cursor
	env.Insert(env.Cursor, env.Cursor, ctx.sS)
	env.Cursor = c
	return true
}

// step5l deletes an -ουνε/-ησουνε/-ηθουνε ending (a67) and restores -ουν for the
// captured prefix stems in a68.
func step5l(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a67, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !env.SliceDel() {
		return false
	}
	ctx.bTest1 = false
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a68, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if env.Cursor > env.LimitBackward {
		return false
	}
	ctx.sS = env.SliceTo()
	if ctx.sS == "" {
		return false
	}
	if !env.SliceFrom("ουν") {
		return false
	}
	c := env.Cursor
	env.Insert(env.Cursor, env.Cursor, ctx.sS)
	env.Cursor = c
	return true
}

// step5m deletes an -ουμε/-ησουμε/-ηθουμε ending (a69) and restores -ουμ for the
// captured prefix stems in a70.
func step5m(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a69, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !env.SliceDel() {
		return false
	}
	ctx.bTest1 = false
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a70, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if env.Cursor > env.LimitBackward {
		return false
	}
	ctx.sS = env.SliceTo()
	if ctx.sS == "" {
		return false
	}
	if !env.SliceFrom("ουμ") {
		return false
	}
	c := env.Cursor
	env.Insert(env.Cursor, env.Cursor, ctx.sS)
	env.Cursor = c
	return true
}

// step6 restores -μα for the -ματα/-ματων neuter stems (a71), then — only if
// test1 is still set — deletes the broad set of regular inflectional endings in
// a72.
func step6(env *snowball.Env, ctx *context) bool {
	v1 := env.Limit - env.Cursor
lab0:
	for {
		env.Ket = env.Cursor
		if snowball.FindAmongB(env, a71, ctx) == 0 {
			break lab0
		}
		env.Bra = env.Cursor
		if !env.SliceFrom("μα") {
			return false
		}
		break lab0
	}
	env.Cursor = env.Limit - v1
	if !ctx.bTest1 {
		return false
	}
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a72, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !env.SliceDel() {
		return false
	}
	return true
}

// step7 deletes the residual comparative/superlative endings in a73.
func step7(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a73, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !env.SliceDel() {
		return false
	}
	return true
}

// Stem runs the Snowball greek algorithm over env, mirroring the generated
// `stem` entry point: it lower-cases the word, then applies the noun/verb suffix
// steps backwards. It always returns true; the result is the mutated env.
func Stem(env *snowball.Env) bool {
	ctx := &context{}
	env.LimitBackward = env.Cursor
	env.Cursor = env.Limit
	v1 := env.Limit - env.Cursor
lab0:
	for {
		if !tolower(env, ctx) {
			break lab0
		}
		break lab0
	}
	env.Cursor = env.Limit - v1
	if !hasMinLength(env, ctx) {
		return false
	}
	ctx.bTest1 = true
	v2 := env.Limit - env.Cursor
lab1:
	for {
		if !step1(env, ctx) {
			break lab1
		}
		break lab1
	}
	env.Cursor = env.Limit - v2
	v3 := env.Limit - env.Cursor
lab2:
	for {
		if !steps1(env, ctx) {
			break lab2
		}
		break lab2
	}
	env.Cursor = env.Limit - v3
	v4 := env.Limit - env.Cursor
lab3:
	for {
		if !steps2(env, ctx) {
			break lab3
		}
		break lab3
	}
	env.Cursor = env.Limit - v4
	v5 := env.Limit - env.Cursor
lab4:
	for {
		if !steps3(env, ctx) {
			break lab4
		}
		break lab4
	}
	env.Cursor = env.Limit - v5
	v6 := env.Limit - env.Cursor
lab5:
	for {
		if !steps4(env, ctx) {
			break lab5
		}
		break lab5
	}
	env.Cursor = env.Limit - v6
	v7 := env.Limit - env.Cursor
lab6:
	for {
		if !steps5(env, ctx) {
			break lab6
		}
		break lab6
	}
	env.Cursor = env.Limit - v7
	v8 := env.Limit - env.Cursor
lab7:
	for {
		if !steps6(env, ctx) {
			break lab7
		}
		break lab7
	}
	env.Cursor = env.Limit - v8
	v9 := env.Limit - env.Cursor
lab8:
	for {
		if !steps7(env, ctx) {
			break lab8
		}
		break lab8
	}
	env.Cursor = env.Limit - v9
	v10 := env.Limit - env.Cursor
lab9:
	for {
		if !steps8(env, ctx) {
			break lab9
		}
		break lab9
	}
	env.Cursor = env.Limit - v10
	v11 := env.Limit - env.Cursor
lab10:
	for {
		if !steps9(env, ctx) {
			break lab10
		}
		break lab10
	}
	env.Cursor = env.Limit - v11
	v12 := env.Limit - env.Cursor
lab11:
	for {
		if !steps10(env, ctx) {
			break lab11
		}
		break lab11
	}
	env.Cursor = env.Limit - v12
	v13 := env.Limit - env.Cursor
lab12:
	for {
		if !step2a(env, ctx) {
			break lab12
		}
		break lab12
	}
	env.Cursor = env.Limit - v13
	v14 := env.Limit - env.Cursor
lab13:
	for {
		if !step2b(env, ctx) {
			break lab13
		}
		break lab13
	}
	env.Cursor = env.Limit - v14
	v15 := env.Limit - env.Cursor
lab14:
	for {
		if !step2c(env, ctx) {
			break lab14
		}
		break lab14
	}
	env.Cursor = env.Limit - v15
	v16 := env.Limit - env.Cursor
lab15:
	for {
		if !step2d(env, ctx) {
			break lab15
		}
		break lab15
	}
	env.Cursor = env.Limit - v16
	v17 := env.Limit - env.Cursor
lab16:
	for {
		if !step3(env, ctx) {
			break lab16
		}
		break lab16
	}
	env.Cursor = env.Limit - v17
	v18 := env.Limit - env.Cursor
lab17:
	for {
		if !step4(env, ctx) {
			break lab17
		}
		break lab17
	}
	env.Cursor = env.Limit - v18
	v19 := env.Limit - env.Cursor
lab18:
	for {
		if !step5a(env, ctx) {
			break lab18
		}
		break lab18
	}
	env.Cursor = env.Limit - v19
	v20 := env.Limit - env.Cursor
lab19:
	for {
		if !step5b(env, ctx) {
			break lab19
		}
		break lab19
	}
	env.Cursor = env.Limit - v20
	v21 := env.Limit - env.Cursor
lab20:
	for {
		if !step5c(env, ctx) {
			break lab20
		}
		break lab20
	}
	env.Cursor = env.Limit - v21
	v22 := env.Limit - env.Cursor
lab21:
	for {
		if !step5d(env, ctx) {
			break lab21
		}
		break lab21
	}
	env.Cursor = env.Limit - v22
	v23 := env.Limit - env.Cursor
lab22:
	for {
		if !step5e(env, ctx) {
			break lab22
		}
		break lab22
	}
	env.Cursor = env.Limit - v23
	v24 := env.Limit - env.Cursor
lab23:
	for {
		if !step5f(env, ctx) {
			break lab23
		}
		break lab23
	}
	env.Cursor = env.Limit - v24
	v25 := env.Limit - env.Cursor
lab24:
	for {
		if !step5g(env, ctx) {
			break lab24
		}
		break lab24
	}
	env.Cursor = env.Limit - v25
	v26 := env.Limit - env.Cursor
lab25:
	for {
		if !step5h(env, ctx) {
			break lab25
		}
		break lab25
	}
	env.Cursor = env.Limit - v26
	v27 := env.Limit - env.Cursor
lab26:
	for {
		if !step5j(env, ctx) {
			break lab26
		}
		break lab26
	}
	env.Cursor = env.Limit - v27
	v28 := env.Limit - env.Cursor
lab27:
	for {
		if !step5i(env, ctx) {
			break lab27
		}
		break lab27
	}
	env.Cursor = env.Limit - v28
	v29 := env.Limit - env.Cursor
lab28:
	for {
		if !step5k(env, ctx) {
			break lab28
		}
		break lab28
	}
	env.Cursor = env.Limit - v29
	v30 := env.Limit - env.Cursor
lab29:
	for {
		if !step5l(env, ctx) {
			break lab29
		}
		break lab29
	}
	env.Cursor = env.Limit - v30
	v31 := env.Limit - env.Cursor
lab30:
	for {
		if !step5m(env, ctx) {
			break lab30
		}
		break lab30
	}
	env.Cursor = env.Limit - v31
	v32 := env.Limit - env.Cursor
lab31:
	for {
		if !step6(env, ctx) {
			break lab31
		}
		break lab31
	}
	env.Cursor = env.Limit - v32
	v33 := env.Limit - env.Cursor
lab32:
	for {
		if !step7(env, ctx) {
			break lab32
		}
		break lab32
	}
	env.Cursor = env.Limit - v33
	env.Cursor = env.LimitBackward
	return true
}
