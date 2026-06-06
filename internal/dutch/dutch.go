// Package dutch is a byte-faithful Go port of rust-stemmers' generated Snowball
// "dutch" stemmer. It produces output identical to rust-stemmers 1.2.0's Dutch
// algorithm; the canonical Snowball dutch vocabulary is the conformance oracle.
//
// The control flow mirrors the generated Rust: each Snowball labelled block
// (`'lab: loop { … break 'lab }`) becomes a Go labelled `for { … break LABEL }`,
// and the result-code dispatch (`else if among_var == N`) becomes a `switch`.
package dutch

import "github.com/freeeve/go-stemmers/internal/snowball"

// context holds the per-run state the generated algorithm threads through its
// routines: the R1/R2 region marks and whether an e was removed (e_found).
type context struct {
	iP2     int
	iP1     int
	bEFound bool
}

var a0 = []snowball.Among[context]{
	{Str: "", SubstringI: -1, Result: 6},
	{Str: "á", SubstringI: 0, Result: 1},
	{Str: "ä", SubstringI: 0, Result: 1},
	{Str: "é", SubstringI: 0, Result: 2},
	{Str: "ë", SubstringI: 0, Result: 2},
	{Str: "í", SubstringI: 0, Result: 3},
	{Str: "ï", SubstringI: 0, Result: 3},
	{Str: "ó", SubstringI: 0, Result: 4},
	{Str: "ö", SubstringI: 0, Result: 4},
	{Str: "ú", SubstringI: 0, Result: 5},
	{Str: "ü", SubstringI: 0, Result: 5},
}

var a1 = []snowball.Among[context]{
	{Str: "", SubstringI: -1, Result: 3},
	{Str: "I", SubstringI: 0, Result: 2},
	{Str: "Y", SubstringI: 0, Result: 1},
}

var a2 = []snowball.Among[context]{
	{Str: "dd", SubstringI: -1, Result: -1},
	{Str: "kk", SubstringI: -1, Result: -1},
	{Str: "tt", SubstringI: -1, Result: -1},
}

var a3 = []snowball.Among[context]{
	{Str: "ene", SubstringI: -1, Result: 2},
	{Str: "se", SubstringI: -1, Result: 3},
	{Str: "en", SubstringI: -1, Result: 2},
	{Str: "heden", SubstringI: 2, Result: 1},
	{Str: "s", SubstringI: -1, Result: 3},
}

var a4 = []snowball.Among[context]{
	{Str: "end", SubstringI: -1, Result: 1},
	{Str: "ig", SubstringI: -1, Result: 2},
	{Str: "ing", SubstringI: -1, Result: 1},
	{Str: "lijk", SubstringI: -1, Result: 3},
	{Str: "baar", SubstringI: -1, Result: 4},
	{Str: "bar", SubstringI: -1, Result: 5},
}

var a5 = []snowball.Among[context]{
	{Str: "aa", SubstringI: -1, Result: -1},
	{Str: "ee", SubstringI: -1, Result: -1},
	{Str: "oo", SubstringI: -1, Result: -1},
	{Str: "uu", SubstringI: -1, Result: -1},
}

var gV = []byte{17, 65, 16, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128}
var gVI = []byte{1, 0, 0, 17, 65, 16, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128}
var gVJ = []byte{17, 67, 16, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128}

// prelude folds accented vowels (á→a, é→e, …) to their plain forms and marks a
// vowel-flanked y/i as the consonant Y/I so later steps treat them correctly.
func prelude(env *snowball.Env, ctx *context) bool {
	var amongVar int32
	v1 := env.Cursor
replab0:
	for {
		v2 := env.Cursor
	lab1:
		for once := 0; once < 1; once++ {
			env.Bra = env.Cursor
			amongVar = snowball.FindAmong(env, a0, ctx)
			if amongVar == 0 {
				break lab1
			}
			env.Ket = env.Cursor
			switch amongVar {
			case 1:
				if !env.SliceFrom("a") {
					return false
				}
			case 2:
				if !env.SliceFrom("e") {
					return false
				}
			case 3:
				if !env.SliceFrom("i") {
					return false
				}
			case 4:
				if !env.SliceFrom("o") {
					return false
				}
			case 5:
				if !env.SliceFrom("u") {
					return false
				}
			case 6:
				if env.Cursor >= env.Limit {
					break lab1
				}
				env.NextChar()
			}
			continue replab0
		}
		env.Cursor = v2
		break replab0
	}
	env.Cursor = v1
	v3 := env.Cursor
lab2:
	for {
		env.Bra = env.Cursor
		if !env.EqS("y") {
			env.Cursor = v3
			break lab2
		}
		env.Ket = env.Cursor
		if !env.SliceFrom("Y") {
			return false
		}
		break lab2
	}
replab3:
	for {
		v4 := env.Cursor
	lab4:
		for once := 0; once < 1; once++ {
		golab5:
			for {
				v5 := env.Cursor
			lab6:
				for {
					if !env.InGrouping(gV, 97, 232) {
						break lab6
					}
					env.Bra = env.Cursor
				lab7:
					for {
						v6 := env.Cursor
					lab8:
						for {
							if !env.EqS("i") {
								break lab8
							}
							env.Ket = env.Cursor
							if !env.InGrouping(gV, 97, 232) {
								break lab8
							}
							if !env.SliceFrom("I") {
								return false
							}
							break lab7
						}
						env.Cursor = v6
						if !env.EqS("y") {
							break lab6
						}
						env.Ket = env.Cursor
						if !env.SliceFrom("Y") {
							return false
						}
						break lab7
					}
					env.Cursor = v5
					break golab5
				}
				env.Cursor = v5
				if env.Cursor >= env.Limit {
					break lab4
				}
				env.NextChar()
			}
			continue replab3
		}
		env.Cursor = v4
		break replab3
	}
	return true
}

// markRegions sets the R1 (iP1) and R2 (iP2) region boundaries, forcing R1 to be
// at least position 3 from the start of the word.
func markRegions(env *snowball.Env, ctx *context) bool {
	ctx.iP1 = env.Limit
	ctx.iP2 = env.Limit
golab0:
	for {
		for {
			if !env.InGrouping(gV, 97, 232) {
				break
			}
			break golab0
		}
		if env.Cursor >= env.Limit {
			return false
		}
		env.NextChar()
	}
golab2:
	for {
		for {
			if !env.OutGrouping(gV, 97, 232) {
				break
			}
			break golab2
		}
		if env.Cursor >= env.Limit {
			return false
		}
		env.NextChar()
	}
	ctx.iP1 = env.Cursor
lab4:
	for {
		if !(ctx.iP1 < 3) {
			break lab4
		}
		ctx.iP1 = 3
		break lab4
	}
golab5:
	for {
		for {
			if !env.InGrouping(gV, 97, 232) {
				break
			}
			break golab5
		}
		if env.Cursor >= env.Limit {
			return false
		}
		env.NextChar()
	}
golab7:
	for {
		for {
			if !env.OutGrouping(gV, 97, 232) {
				break
			}
			break golab7
		}
		if env.Cursor >= env.Limit {
			return false
		}
		env.NextChar()
	}
	ctx.iP2 = env.Cursor
	return true
}

// postlude restores the consonant marks introduced by the prelude, turning Y
// back into y and I back into i.
func postlude(env *snowball.Env, ctx *context) bool {
	var amongVar int32
replab0:
	for {
		v1 := env.Cursor
	lab1:
		for once := 0; once < 1; once++ {
			env.Bra = env.Cursor
			amongVar = snowball.FindAmong(env, a1, ctx)
			if amongVar == 0 {
				break lab1
			}
			env.Ket = env.Cursor
			switch amongVar {
			case 1:
				if !env.SliceFrom("y") {
					return false
				}
			case 2:
				if !env.SliceFrom("i") {
					return false
				}
			case 3:
				if env.Cursor >= env.Limit {
					break lab1
				}
				env.NextChar()
			}
			continue replab0
		}
		env.Cursor = v1
		break replab0
	}
	return true
}

// r1 reports whether the cursor is within region R1.
func r1(env *snowball.Env, ctx *context) bool {
	return ctx.iP1 <= env.Cursor
}

// r2 reports whether the cursor is within region R2.
func r2(env *snowball.Env, ctx *context) bool {
	return ctx.iP2 <= env.Cursor
}

// undouble removes a doubled final consonant (dd, kk, tt) left behind after a
// suffix deletion.
func undouble(env *snowball.Env, ctx *context) bool {
	v1 := env.Limit - env.Cursor
	if snowball.FindAmongB(env, a2, ctx) == 0 {
		return false
	}
	env.Cursor = env.Limit - v1
	env.Ket = env.Cursor
	if env.Cursor <= env.LimitBackward {
		return false
	}
	env.PreviousChar()
	env.Bra = env.Cursor
	if !env.SliceDel() {
		return false
	}
	return true
}

// eEnding removes a final e in R1 when preceded by a non-vowel, sets e_found, and
// undoubles the resulting consonant.
func eEnding(env *snowball.Env, ctx *context) bool {
	ctx.bEFound = false
	env.Ket = env.Cursor
	if !env.EqSB("e") {
		return false
	}
	env.Bra = env.Cursor
	if !r1(env, ctx) {
		return false
	}
	v1 := env.Limit - env.Cursor
	if !env.OutGroupingB(gV, 97, 232) {
		return false
	}
	env.Cursor = env.Limit - v1
	if !env.SliceDel() {
		return false
	}
	ctx.bEFound = true
	if !undouble(env, ctx) {
		return false
	}
	return true
}

// enEnding removes a final en in R1 when preceded by a non-vowel (but not gem),
// then undoubles.
func enEnding(env *snowball.Env, ctx *context) bool {
	if !r1(env, ctx) {
		return false
	}
	v1 := env.Limit - env.Cursor
	if !env.OutGroupingB(gV, 97, 232) {
		return false
	}
	env.Cursor = env.Limit - v1
	v2 := env.Limit - env.Cursor
lab0:
	for {
		if !env.EqSB("gem") {
			break lab0
		}
		return false
	}
	env.Cursor = env.Limit - v2
	if !env.SliceDel() {
		return false
	}
	if !undouble(env, ctx) {
		return false
	}
	return true
}

// standardSuffix is the main backwards suffix-stripping routine: it handles the
// heden/ene/se group, e/en endings, the heid suffix, the derivational endings in
// a4, and a final undouble of long vowels.
func standardSuffix(env *snowball.Env, ctx *context) bool {
	var amongVar int32
	v1 := env.Limit - env.Cursor
lab0:
	for {
		env.Ket = env.Cursor
		amongVar = snowball.FindAmongB(env, a3, ctx)
		if amongVar == 0 {
			break lab0
		}
		env.Bra = env.Cursor
		switch amongVar {
		case 1:
			if !r1(env, ctx) {
				break lab0
			}
			if !env.SliceFrom("heid") {
				return false
			}
		case 2:
			if !enEnding(env, ctx) {
				break lab0
			}
		case 3:
			if !r1(env, ctx) {
				break lab0
			}
			if !env.OutGroupingB(gVJ, 97, 232) {
				break lab0
			}
			if !env.SliceDel() {
				return false
			}
		}
		break lab0
	}
	env.Cursor = env.Limit - v1
	v2 := env.Limit - env.Cursor
lab1:
	for {
		if !eEnding(env, ctx) {
			break lab1
		}
		break lab1
	}
	env.Cursor = env.Limit - v2
	v3 := env.Limit - env.Cursor
lab2:
	for {
		env.Ket = env.Cursor
		if !env.EqSB("heid") {
			break lab2
		}
		env.Bra = env.Cursor
		if !r2(env, ctx) {
			break lab2
		}
		v4 := env.Limit - env.Cursor
	lab3:
		for {
			if !env.EqSB("c") {
				break lab3
			}
			break lab2
		}
		env.Cursor = env.Limit - v4
		if !env.SliceDel() {
			return false
		}
		env.Ket = env.Cursor
		if !env.EqSB("en") {
			break lab2
		}
		env.Bra = env.Cursor
		if !enEnding(env, ctx) {
			break lab2
		}
		break lab2
	}
	env.Cursor = env.Limit - v3
	v5 := env.Limit - env.Cursor
lab4:
	for {
		env.Ket = env.Cursor
		amongVar = snowball.FindAmongB(env, a4, ctx)
		if amongVar == 0 {
			break lab4
		}
		env.Bra = env.Cursor
		switch amongVar {
		case 1:
			if !r2(env, ctx) {
				break lab4
			}
			if !env.SliceDel() {
				return false
			}
		lab5:
			for {
				v6 := env.Limit - env.Cursor
			lab6:
				for {
					env.Ket = env.Cursor
					if !env.EqSB("ig") {
						break lab6
					}
					env.Bra = env.Cursor
					if !r2(env, ctx) {
						break lab6
					}
					v7 := env.Limit - env.Cursor
				lab7:
					for {
						if !env.EqSB("e") {
							break lab7
						}
						break lab6
					}
					env.Cursor = env.Limit - v7
					if !env.SliceDel() {
						return false
					}
					break lab5
				}
				env.Cursor = env.Limit - v6
				if !undouble(env, ctx) {
					break lab4
				}
				break lab5
			}
		case 2:
			if !r2(env, ctx) {
				break lab4
			}
			v8 := env.Limit - env.Cursor
		lab8:
			for {
				if !env.EqSB("e") {
					break lab8
				}
				break lab4
			}
			env.Cursor = env.Limit - v8
			if !env.SliceDel() {
				return false
			}
		case 3:
			if !r2(env, ctx) {
				break lab4
			}
			if !env.SliceDel() {
				return false
			}
			if !eEnding(env, ctx) {
				break lab4
			}
		case 4:
			if !r2(env, ctx) {
				break lab4
			}
			if !env.SliceDel() {
				return false
			}
		case 5:
			if !r2(env, ctx) {
				break lab4
			}
			if !ctx.bEFound {
				break lab4
			}
			if !env.SliceDel() {
				return false
			}
		}
		break lab4
	}
	env.Cursor = env.Limit - v5
	v9 := env.Limit - env.Cursor
lab9:
	for {
		if !env.OutGroupingB(gVI, 73, 232) {
			break lab9
		}
		v10 := env.Limit - env.Cursor
		if snowball.FindAmongB(env, a5, ctx) == 0 {
			break lab9
		}
		if !env.OutGroupingB(gV, 97, 232) {
			break lab9
		}
		env.Cursor = env.Limit - v10
		env.Ket = env.Cursor
		if env.Cursor <= env.LimitBackward {
			break lab9
		}
		env.PreviousChar()
		env.Bra = env.Cursor
		if !env.SliceDel() {
			return false
		}
		break lab9
	}
	env.Cursor = env.Limit - v9
	return true
}

// Stem runs the Snowball dutch algorithm over env, mirroring the generated `stem`
// entry point. It always returns true; the result is the mutated env.
func Stem(env *snowball.Env) bool {
	ctx := &context{}
	v1 := env.Cursor
lab0:
	for {
		if !prelude(env, ctx) {
			break lab0
		}
		break lab0
	}
	env.Cursor = v1
	v2 := env.Cursor
lab1:
	for {
		if !markRegions(env, ctx) {
			break lab1
		}
		break lab1
	}
	env.Cursor = v2
	env.LimitBackward = env.Cursor
	env.Cursor = env.Limit
	v3 := env.Limit - env.Cursor
lab2:
	for {
		if !standardSuffix(env, ctx) {
			break lab2
		}
		break lab2
	}
	env.Cursor = env.Limit - v3
	env.Cursor = env.LimitBackward
	v4 := env.Cursor
lab3:
	for {
		if !postlude(env, ctx) {
			break lab3
		}
		break lab3
	}
	env.Cursor = v4
	return true
}
