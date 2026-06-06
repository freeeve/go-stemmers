// Package french is a byte-faithful Go port of rust-stemmers' generated
// Snowball "french" stemmer. It produces output identical to rust-stemmers
// 1.2.0's French algorithm; the canonical Snowball french vocabulary is the
// conformance oracle.
//
// The control flow mirrors the generated Rust: each Snowball labelled block
// (`'lab: loop { … break 'lab }`) becomes a Go labelled `for { … break LABEL }`,
// and the result-code dispatch (`else if among_var == N`) becomes a `switch`.
// Accented suffix literals (é, è, ç, î, â, ë …) are pasted verbatim.
package french

import "github.com/freeeve/go-stemmers/internal/snowball"

// context holds the per-run region marks the generated algorithm threads through
// its routines: RV (iPV), R1 (iP1) and R2 (iP2).
type context struct {
	iP2 int
	iP1 int
	iPV int
}

var a0 = []snowball.Among[context]{
	{Str: "col", SubstringI: -1, Result: -1},
	{Str: "par", SubstringI: -1, Result: -1},
	{Str: "tap", SubstringI: -1, Result: -1},
}

var a1 = []snowball.Among[context]{
	{Str: "", SubstringI: -1, Result: 4},
	{Str: "I", SubstringI: 0, Result: 1},
	{Str: "U", SubstringI: 0, Result: 2},
	{Str: "Y", SubstringI: 0, Result: 3},
}

var a2 = []snowball.Among[context]{
	{Str: "iqU", SubstringI: -1, Result: 3},
	{Str: "abl", SubstringI: -1, Result: 3},
	{Str: "Ièr", SubstringI: -1, Result: 4},
	{Str: "ièr", SubstringI: -1, Result: 4},
	{Str: "eus", SubstringI: -1, Result: 2},
	{Str: "iv", SubstringI: -1, Result: 1},
}

var a3 = []snowball.Among[context]{
	{Str: "ic", SubstringI: -1, Result: 2},
	{Str: "abil", SubstringI: -1, Result: 1},
	{Str: "iv", SubstringI: -1, Result: 3},
}

var a4 = []snowball.Among[context]{
	{Str: "iqUe", SubstringI: -1, Result: 1},
	{Str: "atrice", SubstringI: -1, Result: 2},
	{Str: "ance", SubstringI: -1, Result: 1},
	{Str: "ence", SubstringI: -1, Result: 5},
	{Str: "logie", SubstringI: -1, Result: 3},
	{Str: "able", SubstringI: -1, Result: 1},
	{Str: "isme", SubstringI: -1, Result: 1},
	{Str: "euse", SubstringI: -1, Result: 11},
	{Str: "iste", SubstringI: -1, Result: 1},
	{Str: "ive", SubstringI: -1, Result: 8},
	{Str: "if", SubstringI: -1, Result: 8},
	{Str: "usion", SubstringI: -1, Result: 4},
	{Str: "ation", SubstringI: -1, Result: 2},
	{Str: "ution", SubstringI: -1, Result: 4},
	{Str: "ateur", SubstringI: -1, Result: 2},
	{Str: "iqUes", SubstringI: -1, Result: 1},
	{Str: "atrices", SubstringI: -1, Result: 2},
	{Str: "ances", SubstringI: -1, Result: 1},
	{Str: "ences", SubstringI: -1, Result: 5},
	{Str: "logies", SubstringI: -1, Result: 3},
	{Str: "ables", SubstringI: -1, Result: 1},
	{Str: "ismes", SubstringI: -1, Result: 1},
	{Str: "euses", SubstringI: -1, Result: 11},
	{Str: "istes", SubstringI: -1, Result: 1},
	{Str: "ives", SubstringI: -1, Result: 8},
	{Str: "ifs", SubstringI: -1, Result: 8},
	{Str: "usions", SubstringI: -1, Result: 4},
	{Str: "ations", SubstringI: -1, Result: 2},
	{Str: "utions", SubstringI: -1, Result: 4},
	{Str: "ateurs", SubstringI: -1, Result: 2},
	{Str: "ments", SubstringI: -1, Result: 15},
	{Str: "ements", SubstringI: 30, Result: 6},
	{Str: "issements", SubstringI: 31, Result: 12},
	{Str: "ités", SubstringI: -1, Result: 7},
	{Str: "ment", SubstringI: -1, Result: 15},
	{Str: "ement", SubstringI: 34, Result: 6},
	{Str: "issement", SubstringI: 35, Result: 12},
	{Str: "amment", SubstringI: 34, Result: 13},
	{Str: "emment", SubstringI: 34, Result: 14},
	{Str: "aux", SubstringI: -1, Result: 10},
	{Str: "eaux", SubstringI: 39, Result: 9},
	{Str: "eux", SubstringI: -1, Result: 1},
	{Str: "ité", SubstringI: -1, Result: 7},
}

var a5 = []snowball.Among[context]{
	{Str: "ira", SubstringI: -1, Result: 1},
	{Str: "ie", SubstringI: -1, Result: 1},
	{Str: "isse", SubstringI: -1, Result: 1},
	{Str: "issante", SubstringI: -1, Result: 1},
	{Str: "i", SubstringI: -1, Result: 1},
	{Str: "irai", SubstringI: 4, Result: 1},
	{Str: "ir", SubstringI: -1, Result: 1},
	{Str: "iras", SubstringI: -1, Result: 1},
	{Str: "ies", SubstringI: -1, Result: 1},
	{Str: "îmes", SubstringI: -1, Result: 1},
	{Str: "isses", SubstringI: -1, Result: 1},
	{Str: "issantes", SubstringI: -1, Result: 1},
	{Str: "îtes", SubstringI: -1, Result: 1},
	{Str: "is", SubstringI: -1, Result: 1},
	{Str: "irais", SubstringI: 13, Result: 1},
	{Str: "issais", SubstringI: 13, Result: 1},
	{Str: "irions", SubstringI: -1, Result: 1},
	{Str: "issions", SubstringI: -1, Result: 1},
	{Str: "irons", SubstringI: -1, Result: 1},
	{Str: "issons", SubstringI: -1, Result: 1},
	{Str: "issants", SubstringI: -1, Result: 1},
	{Str: "it", SubstringI: -1, Result: 1},
	{Str: "irait", SubstringI: 21, Result: 1},
	{Str: "issait", SubstringI: 21, Result: 1},
	{Str: "issant", SubstringI: -1, Result: 1},
	{Str: "iraIent", SubstringI: -1, Result: 1},
	{Str: "issaIent", SubstringI: -1, Result: 1},
	{Str: "irent", SubstringI: -1, Result: 1},
	{Str: "issent", SubstringI: -1, Result: 1},
	{Str: "iront", SubstringI: -1, Result: 1},
	{Str: "ît", SubstringI: -1, Result: 1},
	{Str: "iriez", SubstringI: -1, Result: 1},
	{Str: "issiez", SubstringI: -1, Result: 1},
	{Str: "irez", SubstringI: -1, Result: 1},
	{Str: "issez", SubstringI: -1, Result: 1},
}

var a6 = []snowball.Among[context]{
	{Str: "a", SubstringI: -1, Result: 3},
	{Str: "era", SubstringI: 0, Result: 2},
	{Str: "asse", SubstringI: -1, Result: 3},
	{Str: "ante", SubstringI: -1, Result: 3},
	{Str: "ée", SubstringI: -1, Result: 2},
	{Str: "ai", SubstringI: -1, Result: 3},
	{Str: "erai", SubstringI: 5, Result: 2},
	{Str: "er", SubstringI: -1, Result: 2},
	{Str: "as", SubstringI: -1, Result: 3},
	{Str: "eras", SubstringI: 8, Result: 2},
	{Str: "âmes", SubstringI: -1, Result: 3},
	{Str: "asses", SubstringI: -1, Result: 3},
	{Str: "antes", SubstringI: -1, Result: 3},
	{Str: "âtes", SubstringI: -1, Result: 3},
	{Str: "ées", SubstringI: -1, Result: 2},
	{Str: "ais", SubstringI: -1, Result: 3},
	{Str: "erais", SubstringI: 15, Result: 2},
	{Str: "ions", SubstringI: -1, Result: 1},
	{Str: "erions", SubstringI: 17, Result: 2},
	{Str: "assions", SubstringI: 17, Result: 3},
	{Str: "erons", SubstringI: -1, Result: 2},
	{Str: "ants", SubstringI: -1, Result: 3},
	{Str: "és", SubstringI: -1, Result: 2},
	{Str: "ait", SubstringI: -1, Result: 3},
	{Str: "erait", SubstringI: 23, Result: 2},
	{Str: "ant", SubstringI: -1, Result: 3},
	{Str: "aIent", SubstringI: -1, Result: 3},
	{Str: "eraIent", SubstringI: 26, Result: 2},
	{Str: "èrent", SubstringI: -1, Result: 2},
	{Str: "assent", SubstringI: -1, Result: 3},
	{Str: "eront", SubstringI: -1, Result: 2},
	{Str: "ât", SubstringI: -1, Result: 3},
	{Str: "ez", SubstringI: -1, Result: 2},
	{Str: "iez", SubstringI: 32, Result: 2},
	{Str: "eriez", SubstringI: 33, Result: 2},
	{Str: "assiez", SubstringI: 33, Result: 3},
	{Str: "erez", SubstringI: 32, Result: 2},
	{Str: "é", SubstringI: -1, Result: 2},
}

var a7 = []snowball.Among[context]{
	{Str: "e", SubstringI: -1, Result: 3},
	{Str: "Ière", SubstringI: 0, Result: 2},
	{Str: "ière", SubstringI: 0, Result: 2},
	{Str: "ion", SubstringI: -1, Result: 1},
	{Str: "Ier", SubstringI: -1, Result: 2},
	{Str: "ier", SubstringI: -1, Result: 2},
	{Str: "ë", SubstringI: -1, Result: 4},
}

var a8 = []snowball.Among[context]{
	{Str: "ell", SubstringI: -1, Result: -1},
	{Str: "eill", SubstringI: -1, Result: -1},
	{Str: "enn", SubstringI: -1, Result: -1},
	{Str: "onn", SubstringI: -1, Result: -1},
	{Str: "ett", SubstringI: -1, Result: -1},
}

var gV = []byte{17, 65, 16, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 130, 103, 8, 5}

var gKeepWithS = []byte{1, 65, 20, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128}

// prelude upper-cases the semivowel forms of u/i (when flanked by vowels), y
// (when adjacent to a vowel) and the u of "qu" to U/I/Y, so later steps treat
// them as consonants.
func prelude(env *snowball.Env, ctx *context) bool {
replab0:
	for {
		v1 := env.Cursor
	lab1:
		for once := 0; once < 1; once++ {
		golab2:
			for {
				v2 := env.Cursor
			lab3:
				for {
				lab4:
					for {
						v3 := env.Cursor
					lab5:
						for {
							if !env.InGrouping(gV, 97, 251) {
								break lab5
							}
							env.Bra = env.Cursor
						lab6:
							for {
								v4 := env.Cursor
							lab7:
								for {
									if !env.EqS("u") {
										break lab7
									}
									env.Ket = env.Cursor
									if !env.InGrouping(gV, 97, 251) {
										break lab7
									}
									if !env.SliceFrom("U") {
										return false
									}
									break lab6
								}
								env.Cursor = v4
							lab8:
								for {
									if !env.EqS("i") {
										break lab8
									}
									env.Ket = env.Cursor
									if !env.InGrouping(gV, 97, 251) {
										break lab8
									}
									if !env.SliceFrom("I") {
										return false
									}
									break lab6
								}
								env.Cursor = v4
								if !env.EqS("y") {
									break lab5
								}
								env.Ket = env.Cursor
								if !env.SliceFrom("Y") {
									return false
								}
								break lab6
							}
							break lab4
						}
						env.Cursor = v3
					lab9:
						for {
							env.Bra = env.Cursor
							if !env.EqS("y") {
								break lab9
							}
							env.Ket = env.Cursor
							if !env.InGrouping(gV, 97, 251) {
								break lab9
							}
							if !env.SliceFrom("Y") {
								return false
							}
							break lab4
						}
						env.Cursor = v3
						if !env.EqS("q") {
							break lab3
						}
						env.Bra = env.Cursor
						if !env.EqS("u") {
							break lab3
						}
						env.Ket = env.Cursor
						if !env.SliceFrom("U") {
							return false
						}
						break lab4
					}
					env.Cursor = v2
					break golab2
				}
				env.Cursor = v2
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

// markRegions computes the RV (iPV), R1 (iP1) and R2 (iP2) region boundaries
// used by the suffix-removal steps.
func markRegions(env *snowball.Env, ctx *context) bool {
	ctx.iPV = env.Limit
	ctx.iP1 = env.Limit
	ctx.iP2 = env.Limit
	v1 := env.Cursor
lab0:
	for {
	lab1:
		for {
			v2 := env.Cursor
		lab2:
			for {
				if !env.InGrouping(gV, 97, 251) {
					break lab2
				}
				if !env.InGrouping(gV, 97, 251) {
					break lab2
				}
				if env.Cursor >= env.Limit {
					break lab2
				}
				env.NextChar()
				break lab1
			}
			env.Cursor = v2
		lab3:
			for {
				if snowball.FindAmong(env, a0, ctx) == 0 {
					break lab3
				}
				break lab1
			}
			env.Cursor = v2
			if env.Cursor >= env.Limit {
				break lab0
			}
			env.NextChar()
		golab4:
			for {
				for {
					if !env.InGrouping(gV, 97, 251) {
						break
					}
					break golab4
				}
				if env.Cursor >= env.Limit {
					break lab0
				}
				env.NextChar()
			}
			break lab1
		}
		ctx.iPV = env.Cursor
		break lab0
	}
	env.Cursor = v1
	v4 := env.Cursor
lab6:
	for {
	golab7:
		for {
			for {
				if !env.InGrouping(gV, 97, 251) {
					break
				}
				break golab7
			}
			if env.Cursor >= env.Limit {
				break lab6
			}
			env.NextChar()
		}
	golab9:
		for {
			for {
				if !env.OutGrouping(gV, 97, 251) {
					break
				}
				break golab9
			}
			if env.Cursor >= env.Limit {
				break lab6
			}
			env.NextChar()
		}
		ctx.iP1 = env.Cursor
	golab11:
		for {
			for {
				if !env.InGrouping(gV, 97, 251) {
					break
				}
				break golab11
			}
			if env.Cursor >= env.Limit {
				break lab6
			}
			env.NextChar()
		}
	golab13:
		for {
			for {
				if !env.OutGrouping(gV, 97, 251) {
					break
				}
				break golab13
			}
			if env.Cursor >= env.Limit {
				break lab6
			}
			env.NextChar()
		}
		ctx.iP2 = env.Cursor
		break lab6
	}
	env.Cursor = v4
	return true
}

// postlude restores the I/U/Y introduced by the prelude back to lowercase
// i/u/y.
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
				if !env.SliceFrom("i") {
					return false
				}
			case 2:
				if !env.SliceFrom("u") {
					return false
				}
			case 3:
				if !env.SliceFrom("y") {
					return false
				}
			case 4:
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

// rV reports whether the cursor is within region RV.
func rV(env *snowball.Env, ctx *context) bool {
	return ctx.iPV <= env.Cursor
}

// r1 reports whether the cursor is within region R1.
func r1(env *snowball.Env, ctx *context) bool {
	return ctx.iP1 <= env.Cursor
}

// r2 reports whether the cursor is within region R2.
func r2(env *snowball.Env, ctx *context) bool {
	return ctx.iP2 <= env.Cursor
}

// standardSuffix removes the standard derivational suffixes in a4 (with the
// follow-up clean-ups over a2/a3) subject to the region constraints.
func standardSuffix(env *snowball.Env, ctx *context) bool {
	var amongVar int32
	env.Ket = env.Cursor
	amongVar = snowball.FindAmongB(env, a4, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	switch amongVar {
	case 1:
		if !r2(env, ctx) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	case 2:
		if !r2(env, ctx) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
		v1 := env.Limit - env.Cursor
	lab0:
		for {
			env.Ket = env.Cursor
			if !env.EqSB("ic") {
				env.Cursor = env.Limit - v1
				break lab0
			}
			env.Bra = env.Cursor
		lab1:
			for {
				v2 := env.Limit - env.Cursor
			lab2:
				for {
					if !r2(env, ctx) {
						break lab2
					}
					if !env.SliceDel() {
						return false
					}
					break lab1
				}
				env.Cursor = env.Limit - v2
				if !env.SliceFrom("iqU") {
					return false
				}
				break lab1
			}
			break lab0
		}
	case 3:
		if !r2(env, ctx) {
			return false
		}
		if !env.SliceFrom("log") {
			return false
		}
	case 4:
		if !r2(env, ctx) {
			return false
		}
		if !env.SliceFrom("u") {
			return false
		}
	case 5:
		if !r2(env, ctx) {
			return false
		}
		if !env.SliceFrom("ent") {
			return false
		}
	case 6:
		if !rV(env, ctx) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
		v3 := env.Limit - env.Cursor
	lab3:
		for {
			env.Ket = env.Cursor
			amongVar = snowball.FindAmongB(env, a2, ctx)
			if amongVar == 0 {
				env.Cursor = env.Limit - v3
				break lab3
			}
			env.Bra = env.Cursor
			switch amongVar {
			case 1:
				if !r2(env, ctx) {
					env.Cursor = env.Limit - v3
					break lab3
				}
				if !env.SliceDel() {
					return false
				}
				env.Ket = env.Cursor
				if !env.EqSB("at") {
					env.Cursor = env.Limit - v3
					break lab3
				}
				env.Bra = env.Cursor
				if !r2(env, ctx) {
					env.Cursor = env.Limit - v3
					break lab3
				}
				if !env.SliceDel() {
					return false
				}
			case 2:
			lab4:
				for {
					v4 := env.Limit - env.Cursor
				lab5:
					for {
						if !r2(env, ctx) {
							break lab5
						}
						if !env.SliceDel() {
							return false
						}
						break lab4
					}
					env.Cursor = env.Limit - v4
					if !r1(env, ctx) {
						env.Cursor = env.Limit - v3
						break lab3
					}
					if !env.SliceFrom("eux") {
						return false
					}
					break lab4
				}
			case 3:
				if !r2(env, ctx) {
					env.Cursor = env.Limit - v3
					break lab3
				}
				if !env.SliceDel() {
					return false
				}
			case 4:
				if !rV(env, ctx) {
					env.Cursor = env.Limit - v3
					break lab3
				}
				if !env.SliceFrom("i") {
					return false
				}
			}
			break lab3
		}
	case 7:
		if !r2(env, ctx) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
		v5 := env.Limit - env.Cursor
	lab6:
		for {
			env.Ket = env.Cursor
			amongVar = snowball.FindAmongB(env, a3, ctx)
			if amongVar == 0 {
				env.Cursor = env.Limit - v5
				break lab6
			}
			env.Bra = env.Cursor
			switch amongVar {
			case 1:
			lab7:
				for {
					v6 := env.Limit - env.Cursor
				lab8:
					for {
						if !r2(env, ctx) {
							break lab8
						}
						if !env.SliceDel() {
							return false
						}
						break lab7
					}
					env.Cursor = env.Limit - v6
					if !env.SliceFrom("abl") {
						return false
					}
					break lab7
				}
			case 2:
			lab9:
				for {
					v7 := env.Limit - env.Cursor
				lab10:
					for {
						if !r2(env, ctx) {
							break lab10
						}
						if !env.SliceDel() {
							return false
						}
						break lab9
					}
					env.Cursor = env.Limit - v7
					if !env.SliceFrom("iqU") {
						return false
					}
					break lab9
				}
			case 3:
				if !r2(env, ctx) {
					env.Cursor = env.Limit - v5
					break lab6
				}
				if !env.SliceDel() {
					return false
				}
			}
			break lab6
		}
	case 8:
		if !r2(env, ctx) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
		v8 := env.Limit - env.Cursor
	lab11:
		for {
			env.Ket = env.Cursor
			if !env.EqSB("at") {
				env.Cursor = env.Limit - v8
				break lab11
			}
			env.Bra = env.Cursor
			if !r2(env, ctx) {
				env.Cursor = env.Limit - v8
				break lab11
			}
			if !env.SliceDel() {
				return false
			}
			env.Ket = env.Cursor
			if !env.EqSB("ic") {
				env.Cursor = env.Limit - v8
				break lab11
			}
			env.Bra = env.Cursor
		lab12:
			for {
				v9 := env.Limit - env.Cursor
			lab13:
				for {
					if !r2(env, ctx) {
						break lab13
					}
					if !env.SliceDel() {
						return false
					}
					break lab12
				}
				env.Cursor = env.Limit - v9
				if !env.SliceFrom("iqU") {
					return false
				}
				break lab12
			}
			break lab11
		}
	case 9:
		if !env.SliceFrom("eau") {
			return false
		}
	case 10:
		if !r1(env, ctx) {
			return false
		}
		if !env.SliceFrom("al") {
			return false
		}
	case 11:
	lab14:
		for {
			v10 := env.Limit - env.Cursor
		lab15:
			for {
				if !r2(env, ctx) {
					break lab15
				}
				if !env.SliceDel() {
					return false
				}
				break lab14
			}
			env.Cursor = env.Limit - v10
			if !r1(env, ctx) {
				return false
			}
			if !env.SliceFrom("eux") {
				return false
			}
			break lab14
		}
	case 12:
		if !r1(env, ctx) {
			return false
		}
		if !env.OutGroupingB(gV, 97, 251) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	case 13:
		if !rV(env, ctx) {
			return false
		}
		if !env.SliceFrom("ant") {
			return false
		}
		return false
	case 14:
		if !rV(env, ctx) {
			return false
		}
		if !env.SliceFrom("ent") {
			return false
		}
		return false
	case 15:
		v11 := env.Limit - env.Cursor
		if !env.InGroupingB(gV, 97, 251) {
			return false
		}
		if !rV(env, ctx) {
			return false
		}
		env.Cursor = env.Limit - v11
		if !env.SliceDel() {
			return false
		}
		return false
	}
	return true
}

// iVerbSuffix removes i-verb suffixes (a5) within RV when preceded by a
// non-vowel.
func iVerbSuffix(env *snowball.Env, ctx *context) bool {
	var amongVar int32
	v1 := env.Limit - env.Cursor
	if env.Cursor < ctx.iPV {
		return false
	}
	env.Cursor = ctx.iPV
	v2 := env.LimitBackward
	env.LimitBackward = env.Cursor
	env.Cursor = env.Limit - v1
	env.Ket = env.Cursor
	amongVar = snowball.FindAmongB(env, a5, ctx)
	if amongVar == 0 {
		env.LimitBackward = v2
		return false
	}
	env.Bra = env.Cursor
	switch amongVar {
	case 1:
		if !env.OutGroupingB(gV, 97, 251) {
			env.LimitBackward = v2
			return false
		}
		if !env.SliceDel() {
			return false
		}
	}
	env.LimitBackward = v2
	return true
}

// verbSuffix removes the remaining verb suffixes (a6) within RV.
func verbSuffix(env *snowball.Env, ctx *context) bool {
	var amongVar int32
	v1 := env.Limit - env.Cursor
	if env.Cursor < ctx.iPV {
		return false
	}
	env.Cursor = ctx.iPV
	v2 := env.LimitBackward
	env.LimitBackward = env.Cursor
	env.Cursor = env.Limit - v1
	env.Ket = env.Cursor
	amongVar = snowball.FindAmongB(env, a6, ctx)
	if amongVar == 0 {
		env.LimitBackward = v2
		return false
	}
	env.Bra = env.Cursor
	switch amongVar {
	case 1:
		if !r2(env, ctx) {
			env.LimitBackward = v2
			return false
		}
		if !env.SliceDel() {
			return false
		}
	case 2:
		if !env.SliceDel() {
			return false
		}
	case 3:
		if !env.SliceDel() {
			return false
		}
		v3 := env.Limit - env.Cursor
	lab0:
		for {
			env.Ket = env.Cursor
			if !env.EqSB("e") {
				env.Cursor = env.Limit - v3
				break lab0
			}
			env.Bra = env.Cursor
			if !env.SliceDel() {
				return false
			}
			break lab0
		}
	}
	env.LimitBackward = v2
	return true
}

// residualSuffix removes residual suffixes: a trailing s before a non-keep
// character, then the endings in a7 within RV.
func residualSuffix(env *snowball.Env, ctx *context) bool {
	var amongVar int32
	v1 := env.Limit - env.Cursor
lab0:
	for {
		env.Ket = env.Cursor
		if !env.EqSB("s") {
			env.Cursor = env.Limit - v1
			break lab0
		}
		env.Bra = env.Cursor
		v2 := env.Limit - env.Cursor
		if !env.OutGroupingB(gKeepWithS, 97, 232) {
			env.Cursor = env.Limit - v1
			break lab0
		}
		env.Cursor = env.Limit - v2
		if !env.SliceDel() {
			return false
		}
		break lab0
	}
	v3 := env.Limit - env.Cursor
	if env.Cursor < ctx.iPV {
		return false
	}
	env.Cursor = ctx.iPV
	v4 := env.LimitBackward
	env.LimitBackward = env.Cursor
	env.Cursor = env.Limit - v3
	env.Ket = env.Cursor
	amongVar = snowball.FindAmongB(env, a7, ctx)
	if amongVar == 0 {
		env.LimitBackward = v4
		return false
	}
	env.Bra = env.Cursor
	switch amongVar {
	case 1:
		if !r2(env, ctx) {
			env.LimitBackward = v4
			return false
		}
	lab1:
		for {
			v5 := env.Limit - env.Cursor
		lab2:
			for {
				if !env.EqSB("s") {
					break lab2
				}
				break lab1
			}
			env.Cursor = env.Limit - v5
			if !env.EqSB("t") {
				env.LimitBackward = v4
				return false
			}
			break lab1
		}
		if !env.SliceDel() {
			return false
		}
	case 2:
		if !env.SliceFrom("i") {
			return false
		}
	case 3:
		if !env.SliceDel() {
			return false
		}
	case 4:
		if !env.EqSB("gu") {
			env.LimitBackward = v4
			return false
		}
		if !env.SliceDel() {
			return false
		}
	}
	env.LimitBackward = v4
	return true
}

// unDouble undoubles a final doubled consonant (ell, enn, ett …).
func unDouble(env *snowball.Env, ctx *context) bool {
	v1 := env.Limit - env.Cursor
	if snowball.FindAmongB(env, a8, ctx) == 0 {
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

// unAccent removes the accent from a final é/è preceded by at least one
// consonant.
func unAccent(env *snowball.Env, ctx *context) bool {
	v1 := 1
replab0:
	for {
	lab1:
		for once := 0; once < 1; once++ {
			if !env.OutGroupingB(gV, 97, 251) {
				break lab1
			}
			v1--
			continue replab0
		}
		break replab0
	}
	if v1 > 0 {
		return false
	}
	env.Ket = env.Cursor
lab2:
	for {
		v3 := env.Limit - env.Cursor
	lab3:
		for {
			if !env.EqSB("é") {
				break lab3
			}
			break lab2
		}
		env.Cursor = env.Limit - v3
		if !env.EqSB("è") {
			return false
		}
		break lab2
	}
	env.Bra = env.Cursor
	if !env.SliceFrom("e") {
		return false
	}
	return true
}

// Stem runs the Snowball french algorithm over env, mirroring the generated
// `stem` entry point. It always returns true; the result is the mutated env.
func Stem(env *snowball.Env) bool {
	ctx, _ := env.Scratch.(*context)
	if ctx == nil {
		ctx = &context{}
		env.Scratch = ctx
	}
	*ctx = context{}

	v1 := env.Cursor
	prelude(env, ctx)
	env.Cursor = v1

	v2 := env.Cursor
	markRegions(env, ctx)
	env.Cursor = v2

	env.LimitBackward = env.Cursor
	env.Cursor = env.Limit

	v3 := env.Limit - env.Cursor
lab2:
	for {
	lab3:
		for {
			v4 := env.Limit - env.Cursor
		lab4:
			for {
				v5 := env.Limit - env.Cursor
			lab5:
				for {
					v6 := env.Limit - env.Cursor
				lab6:
					for {
						if !standardSuffix(env, ctx) {
							break lab6
						}
						break lab5
					}
					env.Cursor = env.Limit - v6
				lab7:
					for {
						if !iVerbSuffix(env, ctx) {
							break lab7
						}
						break lab5
					}
					env.Cursor = env.Limit - v6
					if !verbSuffix(env, ctx) {
						break lab4
					}
					break lab5
				}
				env.Cursor = env.Limit - v5
				v7 := env.Limit - env.Cursor
			lab8:
				for {
					env.Ket = env.Cursor
				lab9:
					for {
						v8 := env.Limit - env.Cursor
					lab10:
						for {
							if !env.EqSB("Y") {
								break lab10
							}
							env.Bra = env.Cursor
							if !env.SliceFrom("i") {
								return false
							}
							break lab9
						}
						env.Cursor = env.Limit - v8
						if !env.EqSB("ç") {
							env.Cursor = env.Limit - v7
							break lab8
						}
						env.Bra = env.Cursor
						if !env.SliceFrom("c") {
							return false
						}
						break lab9
					}
					break lab8
				}
				break lab3
			}
			env.Cursor = env.Limit - v4
			if !residualSuffix(env, ctx) {
				break lab2
			}
			break lab3
		}
		break lab2
	}
	env.Cursor = env.Limit - v3

	v9 := env.Limit - env.Cursor
	unDouble(env, ctx)
	env.Cursor = env.Limit - v9

	v10 := env.Limit - env.Cursor
	unAccent(env, ctx)
	env.Cursor = env.Limit - v10

	env.Cursor = env.LimitBackward

	v11 := env.Cursor
	postlude(env, ctx)
	env.Cursor = v11

	return true
}
