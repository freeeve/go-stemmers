// Package german is a byte-faithful Go port of rust-stemmers' generated
// Snowball "german" stemmer. It produces output identical to rust-stemmers
// 1.2.0's German algorithm; the canonical Snowball german vocabulary is the
// conformance oracle.
//
// The control flow mirrors the generated Rust: each Snowball labelled block
// (`'lab: loop { … break 'lab }`) becomes a Go labelled `for { … break LABEL }`,
// and the result-code dispatch (`else if among_var == N`) becomes a `switch`.
package german

import "github.com/freeeve/go-stemmers/internal/snowball"

// context holds the per-run state the generated algorithm threads through its
// routines: the prefix mark i_x and the R1/R2 region marks.
type context struct {
	iX  int
	iP2 int
	iP1 int
}

var a0 = []snowball.Among[context]{
	{Str: "", SubstringI: -1, Result: 6},
	{Str: "U", SubstringI: 0, Result: 2},
	{Str: "Y", SubstringI: 0, Result: 1},
	{Str: "ä", SubstringI: 0, Result: 3},
	{Str: "ö", SubstringI: 0, Result: 4},
	{Str: "ü", SubstringI: 0, Result: 5},
}

var a1 = []snowball.Among[context]{
	{Str: "e", SubstringI: -1, Result: 2},
	{Str: "em", SubstringI: -1, Result: 1},
	{Str: "en", SubstringI: -1, Result: 2},
	{Str: "ern", SubstringI: -1, Result: 1},
	{Str: "er", SubstringI: -1, Result: 1},
	{Str: "s", SubstringI: -1, Result: 3},
	{Str: "es", SubstringI: 5, Result: 2},
}

var a2 = []snowball.Among[context]{
	{Str: "en", SubstringI: -1, Result: 1},
	{Str: "er", SubstringI: -1, Result: 1},
	{Str: "st", SubstringI: -1, Result: 2},
	{Str: "est", SubstringI: 2, Result: 1},
}

var a3 = []snowball.Among[context]{
	{Str: "ig", SubstringI: -1, Result: 1},
	{Str: "lich", SubstringI: -1, Result: 1},
}

var a4 = []snowball.Among[context]{
	{Str: "end", SubstringI: -1, Result: 1},
	{Str: "ig", SubstringI: -1, Result: 2},
	{Str: "ung", SubstringI: -1, Result: 1},
	{Str: "lich", SubstringI: -1, Result: 3},
	{Str: "isch", SubstringI: -1, Result: 2},
	{Str: "ik", SubstringI: -1, Result: 2},
	{Str: "heit", SubstringI: -1, Result: 3},
	{Str: "keit", SubstringI: -1, Result: 4},
}

var gV = []byte{17, 65, 16, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 0, 32, 8}
var gSEnding = []byte{117, 30, 5}
var gStEnding = []byte{117, 30, 4}

// prelude expands ß to ss and upper-cases the vowel-flanked u/y to U/Y so the
// region logic treats them as consonants; postlude later restores them.
func prelude(env *snowball.Env, ctx *context) bool {
	v1 := env.Cursor
replab0:
	for {
		v2 := env.Cursor
	lab1:
		for once := 0; once < 1; once++ {
		lab2:
			for {
				v3 := env.Cursor
			lab3:
				for {
					env.Bra = env.Cursor
					if !env.EqS("ß") {
						break lab3
					}
					env.Ket = env.Cursor
					if !env.SliceFrom("ss") {
						return false
					}
					break lab2
				}
				env.Cursor = v3
				if env.Cursor >= env.Limit {
					break lab1
				}
				env.NextChar()
				break lab2
			}
			continue replab0
		}
		env.Cursor = v2
		break replab0
	}
	env.Cursor = v1
replab4:
	for {
		v4 := env.Cursor
	lab5:
		for once := 0; once < 1; once++ {
		golab6:
			for {
				v5 := env.Cursor
			lab7:
				for {
					if !env.InGrouping(gV, 97, 252) {
						break lab7
					}
					env.Bra = env.Cursor
				lab8:
					for {
						v6 := env.Cursor
					lab9:
						for {
							if !env.EqS("u") {
								break lab9
							}
							env.Ket = env.Cursor
							if !env.InGrouping(gV, 97, 252) {
								break lab9
							}
							if !env.SliceFrom("U") {
								return false
							}
							break lab8
						}
						env.Cursor = v6
						if !env.EqS("y") {
							break lab7
						}
						env.Ket = env.Cursor
						if !env.InGrouping(gV, 97, 252) {
							break lab7
						}
						if !env.SliceFrom("Y") {
							return false
						}
						break lab8
					}
					env.Cursor = v5
					break golab6
				}
				env.Cursor = v5
				if env.Cursor >= env.Limit {
					break lab5
				}
				env.NextChar()
			}
			continue replab4
		}
		env.Cursor = v4
		break replab4
	}
	return true
}

// markRegions sets the R1 (iP1) and R2 (iP2) region boundaries, pushing R1 to at
// least three characters in via the iX prefix mark.
func markRegions(env *snowball.Env, ctx *context) bool {
	ctx.iP1 = env.Limit
	ctx.iP2 = env.Limit
	v1 := env.Cursor
	c := env.ByteIndexForHop(3)
	if 0 > c || c > env.Limit {
		return false
	}
	env.Cursor = c
	ctx.iX = env.Cursor
	env.Cursor = v1
golab0:
	for {
		for {
			if !env.InGrouping(gV, 97, 252) {
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
			if !env.OutGrouping(gV, 97, 252) {
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
		if !(ctx.iP1 < ctx.iX) {
			break lab4
		}
		ctx.iP1 = ctx.iX
		break lab4
	}
golab5:
	for {
		for {
			if !env.InGrouping(gV, 97, 252) {
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
			if !env.OutGrouping(gV, 97, 252) {
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

// postlude restores the marks introduced by prelude: Y→y, U→u, and the umlauts
// ä/ö/ü back to a/o/u.
func postlude(env *snowball.Env, ctx *context) bool {
	var amongVar int32
replab0:
	for {
		v1 := env.Cursor
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
				if !env.SliceFrom("y") {
					return false
				}
			case 2:
				if !env.SliceFrom("u") {
					return false
				}
			case 3:
				if !env.SliceFrom("a") {
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

// standardSuffix strips the German inflectional and derivational suffixes in
// three backward passes (a1 plural/possessive, a2 comparative, a4 derivational),
// honouring the R1/R2 region constraints.
func standardSuffix(env *snowball.Env, ctx *context) bool {
	var amongVar int32
	v1 := env.Limit - env.Cursor
lab0:
	for {
		env.Ket = env.Cursor
		amongVar = snowball.FindAmongB(env, a1, ctx)
		if amongVar == 0 {
			break lab0
		}
		env.Bra = env.Cursor
		if !r1(env, ctx) {
			break lab0
		}
		switch amongVar {
		case 1:
			if !env.SliceDel() {
				return false
			}
		case 2:
			if !env.SliceDel() {
				return false
			}
			v2 := env.Limit - env.Cursor
		lab1:
			for {
				env.Ket = env.Cursor
				if !env.EqSB("s") {
					env.Cursor = env.Limit - v2
					break lab1
				}
				env.Bra = env.Cursor
				if !env.EqSB("nis") {
					env.Cursor = env.Limit - v2
					break lab1
				}
				if !env.SliceDel() {
					return false
				}
				break lab1
			}
		case 3:
			if !env.InGroupingB(gSEnding, 98, 116) {
				break lab0
			}
			if !env.SliceDel() {
				return false
			}
		}
		break lab0
	}
	env.Cursor = env.Limit - v1
	v3 := env.Limit - env.Cursor
lab2:
	for {
		env.Ket = env.Cursor
		amongVar = snowball.FindAmongB(env, a2, ctx)
		if amongVar == 0 {
			break lab2
		}
		env.Bra = env.Cursor
		if !r1(env, ctx) {
			break lab2
		}
		switch amongVar {
		case 1:
			if !env.SliceDel() {
				return false
			}
		case 2:
			if !env.InGroupingB(gStEnding, 98, 116) {
				break lab2
			}
			c := env.ByteIndexForHop(-3)
			if env.LimitBackward > c || c > env.Limit {
				break lab2
			}
			env.Cursor = c
			if !env.SliceDel() {
				return false
			}
		}
		break lab2
	}
	env.Cursor = env.Limit - v3
	v4 := env.Limit - env.Cursor
lab3:
	for {
		env.Ket = env.Cursor
		amongVar = snowball.FindAmongB(env, a4, ctx)
		if amongVar == 0 {
			break lab3
		}
		env.Bra = env.Cursor
		if !r2(env, ctx) {
			break lab3
		}
		switch amongVar {
		case 1:
			if !env.SliceDel() {
				return false
			}
			v5 := env.Limit - env.Cursor
		lab4:
			for {
				env.Ket = env.Cursor
				if !env.EqSB("ig") {
					env.Cursor = env.Limit - v5
					break lab4
				}
				env.Bra = env.Cursor
				v6 := env.Limit - env.Cursor
			lab5:
				for {
					if !env.EqSB("e") {
						break lab5
					}
					env.Cursor = env.Limit - v5
					break lab4
				}
				env.Cursor = env.Limit - v6
				if !r2(env, ctx) {
					env.Cursor = env.Limit - v5
					break lab4
				}
				if !env.SliceDel() {
					return false
				}
				break lab4
			}
		case 2:
			v7 := env.Limit - env.Cursor
		lab6:
			for {
				if !env.EqSB("e") {
					break lab6
				}
				break lab3
			}
			env.Cursor = env.Limit - v7
			if !env.SliceDel() {
				return false
			}
		case 3:
			if !env.SliceDel() {
				return false
			}
			v8 := env.Limit - env.Cursor
		lab7:
			for {
				env.Ket = env.Cursor
			lab8:
				for {
					v9 := env.Limit - env.Cursor
				lab9:
					for {
						if !env.EqSB("er") {
							break lab9
						}
						break lab8
					}
					env.Cursor = env.Limit - v9
					if !env.EqSB("en") {
						env.Cursor = env.Limit - v8
						break lab7
					}
					break lab8
				}
				env.Bra = env.Cursor
				if !r1(env, ctx) {
					env.Cursor = env.Limit - v8
					break lab7
				}
				if !env.SliceDel() {
					return false
				}
				break lab7
			}
		case 4:
			if !env.SliceDel() {
				return false
			}
			v10 := env.Limit - env.Cursor
		lab10:
			for {
				env.Ket = env.Cursor
				amongVar = snowball.FindAmongB(env, a3, ctx)
				if amongVar == 0 {
					env.Cursor = env.Limit - v10
					break lab10
				}
				env.Bra = env.Cursor
				if !r2(env, ctx) {
					env.Cursor = env.Limit - v10
					break lab10
				}
				switch amongVar {
				case 1:
					if !env.SliceDel() {
						return false
					}
				}
				break lab10
			}
		}
		break lab3
	}
	env.Cursor = env.Limit - v4
	return true
}

// Stem runs the Snowball german algorithm over env, mirroring the generated
// `stem` entry point. It always returns true; the result is the mutated env.
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
