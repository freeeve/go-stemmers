// Package italian is a byte-faithful Go port of rust-stemmers' generated
// Snowball "italian" stemmer. It produces output identical to rust-stemmers
// 1.2.0's Italian algorithm; the canonical Snowball italian vocabulary is the
// conformance oracle.
//
// The control flow mirrors the generated Rust: each Snowball labelled block
// (`'lab: loop { … break 'lab }`) becomes a Go labelled `for { … break LABEL }`,
// and the result-code dispatch (`else if among_var == N`) becomes a `switch`.
package italian

import "github.com/freeeve/go-stemmers/internal/snowball"

// context holds the per-run state the generated algorithm threads through its
// routines: the RV/R1/R2 region marks.
type context struct {
	iP2 int
	iP1 int
	iPV int
}

var a0 = []snowball.Among[context]{
	{Str: "", SubstringI: -1, Result: 7},
	{Str: "qu", SubstringI: 0, Result: 6},
	{Str: "á", SubstringI: 0, Result: 1},
	{Str: "é", SubstringI: 0, Result: 2},
	{Str: "í", SubstringI: 0, Result: 3},
	{Str: "ó", SubstringI: 0, Result: 4},
	{Str: "ú", SubstringI: 0, Result: 5},
}

var a1 = []snowball.Among[context]{
	{Str: "", SubstringI: -1, Result: 3},
	{Str: "I", SubstringI: 0, Result: 1},
	{Str: "U", SubstringI: 0, Result: 2},
}

var a2 = []snowball.Among[context]{
	{Str: "la", SubstringI: -1, Result: -1},
	{Str: "cela", SubstringI: 0, Result: -1},
	{Str: "gliela", SubstringI: 0, Result: -1},
	{Str: "mela", SubstringI: 0, Result: -1},
	{Str: "tela", SubstringI: 0, Result: -1},
	{Str: "vela", SubstringI: 0, Result: -1},
	{Str: "le", SubstringI: -1, Result: -1},
	{Str: "cele", SubstringI: 6, Result: -1},
	{Str: "gliele", SubstringI: 6, Result: -1},
	{Str: "mele", SubstringI: 6, Result: -1},
	{Str: "tele", SubstringI: 6, Result: -1},
	{Str: "vele", SubstringI: 6, Result: -1},
	{Str: "ne", SubstringI: -1, Result: -1},
	{Str: "cene", SubstringI: 12, Result: -1},
	{Str: "gliene", SubstringI: 12, Result: -1},
	{Str: "mene", SubstringI: 12, Result: -1},
	{Str: "sene", SubstringI: 12, Result: -1},
	{Str: "tene", SubstringI: 12, Result: -1},
	{Str: "vene", SubstringI: 12, Result: -1},
	{Str: "ci", SubstringI: -1, Result: -1},
	{Str: "li", SubstringI: -1, Result: -1},
	{Str: "celi", SubstringI: 20, Result: -1},
	{Str: "glieli", SubstringI: 20, Result: -1},
	{Str: "meli", SubstringI: 20, Result: -1},
	{Str: "teli", SubstringI: 20, Result: -1},
	{Str: "veli", SubstringI: 20, Result: -1},
	{Str: "gli", SubstringI: 20, Result: -1},
	{Str: "mi", SubstringI: -1, Result: -1},
	{Str: "si", SubstringI: -1, Result: -1},
	{Str: "ti", SubstringI: -1, Result: -1},
	{Str: "vi", SubstringI: -1, Result: -1},
	{Str: "lo", SubstringI: -1, Result: -1},
	{Str: "celo", SubstringI: 31, Result: -1},
	{Str: "glielo", SubstringI: 31, Result: -1},
	{Str: "melo", SubstringI: 31, Result: -1},
	{Str: "telo", SubstringI: 31, Result: -1},
	{Str: "velo", SubstringI: 31, Result: -1},
}

var a3 = []snowball.Among[context]{
	{Str: "ando", SubstringI: -1, Result: 1},
	{Str: "endo", SubstringI: -1, Result: 1},
	{Str: "ar", SubstringI: -1, Result: 2},
	{Str: "er", SubstringI: -1, Result: 2},
	{Str: "ir", SubstringI: -1, Result: 2},
}

var a4 = []snowball.Among[context]{
	{Str: "ic", SubstringI: -1, Result: -1},
	{Str: "abil", SubstringI: -1, Result: -1},
	{Str: "os", SubstringI: -1, Result: -1},
	{Str: "iv", SubstringI: -1, Result: 1},
}

var a5 = []snowball.Among[context]{
	{Str: "ic", SubstringI: -1, Result: 1},
	{Str: "abil", SubstringI: -1, Result: 1},
	{Str: "iv", SubstringI: -1, Result: 1},
}

var a6 = []snowball.Among[context]{
	{Str: "ica", SubstringI: -1, Result: 1},
	{Str: "logia", SubstringI: -1, Result: 3},
	{Str: "osa", SubstringI: -1, Result: 1},
	{Str: "ista", SubstringI: -1, Result: 1},
	{Str: "iva", SubstringI: -1, Result: 9},
	{Str: "anza", SubstringI: -1, Result: 1},
	{Str: "enza", SubstringI: -1, Result: 5},
	{Str: "ice", SubstringI: -1, Result: 1},
	{Str: "atrice", SubstringI: 7, Result: 1},
	{Str: "iche", SubstringI: -1, Result: 1},
	{Str: "logie", SubstringI: -1, Result: 3},
	{Str: "abile", SubstringI: -1, Result: 1},
	{Str: "ibile", SubstringI: -1, Result: 1},
	{Str: "usione", SubstringI: -1, Result: 4},
	{Str: "azione", SubstringI: -1, Result: 2},
	{Str: "uzione", SubstringI: -1, Result: 4},
	{Str: "atore", SubstringI: -1, Result: 2},
	{Str: "ose", SubstringI: -1, Result: 1},
	{Str: "ante", SubstringI: -1, Result: 1},
	{Str: "mente", SubstringI: -1, Result: 1},
	{Str: "amente", SubstringI: 19, Result: 7},
	{Str: "iste", SubstringI: -1, Result: 1},
	{Str: "ive", SubstringI: -1, Result: 9},
	{Str: "anze", SubstringI: -1, Result: 1},
	{Str: "enze", SubstringI: -1, Result: 5},
	{Str: "ici", SubstringI: -1, Result: 1},
	{Str: "atrici", SubstringI: 25, Result: 1},
	{Str: "ichi", SubstringI: -1, Result: 1},
	{Str: "abili", SubstringI: -1, Result: 1},
	{Str: "ibili", SubstringI: -1, Result: 1},
	{Str: "ismi", SubstringI: -1, Result: 1},
	{Str: "usioni", SubstringI: -1, Result: 4},
	{Str: "azioni", SubstringI: -1, Result: 2},
	{Str: "uzioni", SubstringI: -1, Result: 4},
	{Str: "atori", SubstringI: -1, Result: 2},
	{Str: "osi", SubstringI: -1, Result: 1},
	{Str: "anti", SubstringI: -1, Result: 1},
	{Str: "amenti", SubstringI: -1, Result: 6},
	{Str: "imenti", SubstringI: -1, Result: 6},
	{Str: "isti", SubstringI: -1, Result: 1},
	{Str: "ivi", SubstringI: -1, Result: 9},
	{Str: "ico", SubstringI: -1, Result: 1},
	{Str: "ismo", SubstringI: -1, Result: 1},
	{Str: "oso", SubstringI: -1, Result: 1},
	{Str: "amento", SubstringI: -1, Result: 6},
	{Str: "imento", SubstringI: -1, Result: 6},
	{Str: "ivo", SubstringI: -1, Result: 9},
	{Str: "ità", SubstringI: -1, Result: 8},
	{Str: "istà", SubstringI: -1, Result: 1},
	{Str: "istè", SubstringI: -1, Result: 1},
	{Str: "istì", SubstringI: -1, Result: 1},
}

var a7 = []snowball.Among[context]{
	{Str: "isca", SubstringI: -1, Result: 1},
	{Str: "enda", SubstringI: -1, Result: 1},
	{Str: "ata", SubstringI: -1, Result: 1},
	{Str: "ita", SubstringI: -1, Result: 1},
	{Str: "uta", SubstringI: -1, Result: 1},
	{Str: "ava", SubstringI: -1, Result: 1},
	{Str: "eva", SubstringI: -1, Result: 1},
	{Str: "iva", SubstringI: -1, Result: 1},
	{Str: "erebbe", SubstringI: -1, Result: 1},
	{Str: "irebbe", SubstringI: -1, Result: 1},
	{Str: "isce", SubstringI: -1, Result: 1},
	{Str: "ende", SubstringI: -1, Result: 1},
	{Str: "are", SubstringI: -1, Result: 1},
	{Str: "ere", SubstringI: -1, Result: 1},
	{Str: "ire", SubstringI: -1, Result: 1},
	{Str: "asse", SubstringI: -1, Result: 1},
	{Str: "ate", SubstringI: -1, Result: 1},
	{Str: "avate", SubstringI: 16, Result: 1},
	{Str: "evate", SubstringI: 16, Result: 1},
	{Str: "ivate", SubstringI: 16, Result: 1},
	{Str: "ete", SubstringI: -1, Result: 1},
	{Str: "erete", SubstringI: 20, Result: 1},
	{Str: "irete", SubstringI: 20, Result: 1},
	{Str: "ite", SubstringI: -1, Result: 1},
	{Str: "ereste", SubstringI: -1, Result: 1},
	{Str: "ireste", SubstringI: -1, Result: 1},
	{Str: "ute", SubstringI: -1, Result: 1},
	{Str: "erai", SubstringI: -1, Result: 1},
	{Str: "irai", SubstringI: -1, Result: 1},
	{Str: "isci", SubstringI: -1, Result: 1},
	{Str: "endi", SubstringI: -1, Result: 1},
	{Str: "erei", SubstringI: -1, Result: 1},
	{Str: "irei", SubstringI: -1, Result: 1},
	{Str: "assi", SubstringI: -1, Result: 1},
	{Str: "ati", SubstringI: -1, Result: 1},
	{Str: "iti", SubstringI: -1, Result: 1},
	{Str: "eresti", SubstringI: -1, Result: 1},
	{Str: "iresti", SubstringI: -1, Result: 1},
	{Str: "uti", SubstringI: -1, Result: 1},
	{Str: "avi", SubstringI: -1, Result: 1},
	{Str: "evi", SubstringI: -1, Result: 1},
	{Str: "ivi", SubstringI: -1, Result: 1},
	{Str: "isco", SubstringI: -1, Result: 1},
	{Str: "ando", SubstringI: -1, Result: 1},
	{Str: "endo", SubstringI: -1, Result: 1},
	{Str: "Yamo", SubstringI: -1, Result: 1},
	{Str: "iamo", SubstringI: -1, Result: 1},
	{Str: "avamo", SubstringI: -1, Result: 1},
	{Str: "evamo", SubstringI: -1, Result: 1},
	{Str: "ivamo", SubstringI: -1, Result: 1},
	{Str: "eremo", SubstringI: -1, Result: 1},
	{Str: "iremo", SubstringI: -1, Result: 1},
	{Str: "assimo", SubstringI: -1, Result: 1},
	{Str: "ammo", SubstringI: -1, Result: 1},
	{Str: "emmo", SubstringI: -1, Result: 1},
	{Str: "eremmo", SubstringI: 54, Result: 1},
	{Str: "iremmo", SubstringI: 54, Result: 1},
	{Str: "immo", SubstringI: -1, Result: 1},
	{Str: "ano", SubstringI: -1, Result: 1},
	{Str: "iscano", SubstringI: 58, Result: 1},
	{Str: "avano", SubstringI: 58, Result: 1},
	{Str: "evano", SubstringI: 58, Result: 1},
	{Str: "ivano", SubstringI: 58, Result: 1},
	{Str: "eranno", SubstringI: -1, Result: 1},
	{Str: "iranno", SubstringI: -1, Result: 1},
	{Str: "ono", SubstringI: -1, Result: 1},
	{Str: "iscono", SubstringI: 65, Result: 1},
	{Str: "arono", SubstringI: 65, Result: 1},
	{Str: "erono", SubstringI: 65, Result: 1},
	{Str: "irono", SubstringI: 65, Result: 1},
	{Str: "erebbero", SubstringI: -1, Result: 1},
	{Str: "irebbero", SubstringI: -1, Result: 1},
	{Str: "assero", SubstringI: -1, Result: 1},
	{Str: "essero", SubstringI: -1, Result: 1},
	{Str: "issero", SubstringI: -1, Result: 1},
	{Str: "ato", SubstringI: -1, Result: 1},
	{Str: "ito", SubstringI: -1, Result: 1},
	{Str: "uto", SubstringI: -1, Result: 1},
	{Str: "avo", SubstringI: -1, Result: 1},
	{Str: "evo", SubstringI: -1, Result: 1},
	{Str: "ivo", SubstringI: -1, Result: 1},
	{Str: "ar", SubstringI: -1, Result: 1},
	{Str: "ir", SubstringI: -1, Result: 1},
	{Str: "erà", SubstringI: -1, Result: 1},
	{Str: "irà", SubstringI: -1, Result: 1},
	{Str: "erò", SubstringI: -1, Result: 1},
	{Str: "irò", SubstringI: -1, Result: 1},
}

var gV = []byte{17, 65, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 128, 8, 2, 1}
var gAEIO = []byte{17, 65, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 128, 8, 2}
var gCG = []byte{17}

// prelude normalises acute-accented vowels to grave accents and marks vowel
// u/i between vowels as the upper-case U/I markers the later steps treat as
// consonants.
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
				if !env.SliceFrom("à") {
					return false
				}
			case 2:
				if !env.SliceFrom("è") {
					return false
				}
			case 3:
				if !env.SliceFrom("ì") {
					return false
				}
			case 4:
				if !env.SliceFrom("ò") {
					return false
				}
			case 5:
				if !env.SliceFrom("ù") {
					return false
				}
			case 6:
				if !env.SliceFrom("qU") {
					return false
				}
			case 7:
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
replab2:
	for {
		v3 := env.Cursor
	lab3:
		for once := 0; once < 1; once++ {
		golab4:
			for {
				v4 := env.Cursor
			lab5:
				for {
					if !env.InGrouping(gV, 97, 249) {
						break lab5
					}
					env.Bra = env.Cursor
				lab6:
					for {
						v5 := env.Cursor
					lab7:
						for {
							if !env.EqS("u") {
								break lab7
							}
							env.Ket = env.Cursor
							if !env.InGrouping(gV, 97, 249) {
								break lab7
							}
							if !env.SliceFrom("U") {
								return false
							}
							break lab6
						}
						env.Cursor = v5
						if !env.EqS("i") {
							break lab5
						}
						env.Ket = env.Cursor
						if !env.InGrouping(gV, 97, 249) {
							break lab5
						}
						if !env.SliceFrom("I") {
							return false
						}
						break lab6
					}
					env.Cursor = v4
					break golab4
				}
				env.Cursor = v4
				if env.Cursor >= env.Limit {
					break lab3
				}
				env.NextChar()
			}
			continue replab2
		}
		env.Cursor = v3
		break replab2
	}
	return true
}

// markRegions sets the RV (iPV), R1 (iP1) and R2 (iP2) region boundaries used to
// gate the suffix-removal steps.
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
				if !env.InGrouping(gV, 97, 249) {
					break lab2
				}
			lab3:
				for {
					v3 := env.Cursor
				lab4:
					for {
						if !env.OutGrouping(gV, 97, 249) {
							break lab4
						}
					golab5:
						for {
							for {
								if !env.InGrouping(gV, 97, 249) {
									break
								}
								break golab5
							}
							if env.Cursor >= env.Limit {
								break lab4
							}
							env.NextChar()
						}
						break lab3
					}
					env.Cursor = v3
					if !env.InGrouping(gV, 97, 249) {
						break lab2
					}
				golab7:
					for {
						for {
							if !env.OutGrouping(gV, 97, 249) {
								break
							}
							break golab7
						}
						if env.Cursor >= env.Limit {
							break lab2
						}
						env.NextChar()
					}
					break lab3
				}
				break lab1
			}
			env.Cursor = v2
			if !env.OutGrouping(gV, 97, 249) {
				break lab0
			}
		lab9:
			for {
				v6 := env.Cursor
			lab10:
				for {
					if !env.OutGrouping(gV, 97, 249) {
						break lab10
					}
				golab11:
					for {
						for {
							if !env.InGrouping(gV, 97, 249) {
								break
							}
							break golab11
						}
						if env.Cursor >= env.Limit {
							break lab10
						}
						env.NextChar()
					}
					break lab9
				}
				env.Cursor = v6
				if !env.InGrouping(gV, 97, 249) {
					break lab0
				}
				if env.Cursor >= env.Limit {
					break lab0
				}
				env.NextChar()
				break lab9
			}
			break lab1
		}
		ctx.iPV = env.Cursor
		break lab0
	}
	env.Cursor = v1
	v8 := env.Cursor
lab13:
	for {
	golab14:
		for {
			for {
				if !env.InGrouping(gV, 97, 249) {
					break
				}
				break golab14
			}
			if env.Cursor >= env.Limit {
				break lab13
			}
			env.NextChar()
		}
	golab16:
		for {
			for {
				if !env.OutGrouping(gV, 97, 249) {
					break
				}
				break golab16
			}
			if env.Cursor >= env.Limit {
				break lab13
			}
			env.NextChar()
		}
		ctx.iP1 = env.Cursor
	golab18:
		for {
			for {
				if !env.InGrouping(gV, 97, 249) {
					break
				}
				break golab18
			}
			if env.Cursor >= env.Limit {
				break lab13
			}
			env.NextChar()
		}
	golab20:
		for {
			for {
				if !env.OutGrouping(gV, 97, 249) {
					break
				}
				break golab20
			}
			if env.Cursor >= env.Limit {
				break lab13
			}
			env.NextChar()
		}
		ctx.iP2 = env.Cursor
		break lab13
	}
	env.Cursor = v8
	return true
}

// postlude turns the I/U markers left by prelude back into lower-case i/u.
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

// rv reports whether the cursor is within region RV.
func rv(env *snowball.Env, ctx *context) bool {
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

// attachedPronoun removes a clitic pronoun (the a2 endings) following a verb
// form (a3) when in region RV.
func attachedPronoun(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a2, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	amongVar := snowball.FindAmongB(env, a3, ctx)
	if amongVar == 0 {
		return false
	}
	if !rv(env, ctx) {
		return false
	}
	switch amongVar {
	case 1:
		if !env.SliceDel() {
			return false
		}
	case 2:
		if !env.SliceFrom("e") {
			return false
		}
	}
	return true
}

// standardSuffix rewrites or deletes the derivational suffixes in a6, with the
// post-deletion fix-ups in a4/a5.
func standardSuffix(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a6, ctx)
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
			if !r2(env, ctx) {
				env.Cursor = env.Limit - v1
				break lab0
			}
			if !env.SliceDel() {
				return false
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
		if !env.SliceFrom("ente") {
			return false
		}
	case 6:
		if !rv(env, ctx) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	case 7:
		if !r1(env, ctx) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
		v2 := env.Limit - env.Cursor
	lab1:
		for {
			env.Ket = env.Cursor
			amongVar = snowball.FindAmongB(env, a4, ctx)
			if amongVar == 0 {
				env.Cursor = env.Limit - v2
				break lab1
			}
			env.Bra = env.Cursor
			if !r2(env, ctx) {
				env.Cursor = env.Limit - v2
				break lab1
			}
			if !env.SliceDel() {
				return false
			}
			switch amongVar {
			case 1:
				env.Ket = env.Cursor
				if !env.EqSB("at") {
					env.Cursor = env.Limit - v2
					break lab1
				}
				env.Bra = env.Cursor
				if !r2(env, ctx) {
					env.Cursor = env.Limit - v2
					break lab1
				}
				if !env.SliceDel() {
					return false
				}
			}
			break lab1
		}
	case 8:
		if !r2(env, ctx) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
		v3 := env.Limit - env.Cursor
	lab2:
		for {
			env.Ket = env.Cursor
			amongVar = snowball.FindAmongB(env, a5, ctx)
			if amongVar == 0 {
				env.Cursor = env.Limit - v3
				break lab2
			}
			env.Bra = env.Cursor
			switch amongVar {
			case 1:
				if !r2(env, ctx) {
					env.Cursor = env.Limit - v3
					break lab2
				}
				if !env.SliceDel() {
					return false
				}
			}
			break lab2
		}
	case 9:
		if !r2(env, ctx) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
		v4 := env.Limit - env.Cursor
	lab3:
		for {
			env.Ket = env.Cursor
			if !env.EqSB("at") {
				env.Cursor = env.Limit - v4
				break lab3
			}
			env.Bra = env.Cursor
			if !r2(env, ctx) {
				env.Cursor = env.Limit - v4
				break lab3
			}
			if !env.SliceDel() {
				return false
			}
			env.Ket = env.Cursor
			if !env.EqSB("ic") {
				env.Cursor = env.Limit - v4
				break lab3
			}
			env.Bra = env.Cursor
			if !r2(env, ctx) {
				env.Cursor = env.Limit - v4
				break lab3
			}
			if !env.SliceDel() {
				return false
			}
			break lab3
		}
	}
	return true
}

// verbSuffix deletes the verb conjugation endings in a7 within region RV.
func verbSuffix(env *snowball.Env, ctx *context) bool {
	v1 := env.Limit - env.Cursor
	if env.Cursor < ctx.iPV {
		return false
	}
	env.Cursor = ctx.iPV
	v2 := env.LimitBackward
	env.LimitBackward = env.Cursor
	env.Cursor = env.Limit - v1
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a7, ctx)
	if amongVar == 0 {
		env.LimitBackward = v2
		return false
	}
	env.Bra = env.Cursor
	switch amongVar {
	case 1:
		if !env.SliceDel() {
			return false
		}
	}
	env.LimitBackward = v2
	return true
}

// vowelSuffix removes a trailing vowel (and an h after c/g) at the end of the
// word within region RV.
func vowelSuffix(env *snowball.Env, ctx *context) bool {
	v1 := env.Limit - env.Cursor
lab0:
	for {
		env.Ket = env.Cursor
		if !env.InGroupingB(gAEIO, 97, 242) {
			env.Cursor = env.Limit - v1
			break lab0
		}
		env.Bra = env.Cursor
		if !rv(env, ctx) {
			env.Cursor = env.Limit - v1
			break lab0
		}
		if !env.SliceDel() {
			return false
		}
		env.Ket = env.Cursor
		if !env.EqSB("i") {
			env.Cursor = env.Limit - v1
			break lab0
		}
		env.Bra = env.Cursor
		if !rv(env, ctx) {
			env.Cursor = env.Limit - v1
			break lab0
		}
		if !env.SliceDel() {
			return false
		}
		break lab0
	}
	v2 := env.Limit - env.Cursor
lab1:
	for {
		env.Ket = env.Cursor
		if !env.EqSB("h") {
			env.Cursor = env.Limit - v2
			break lab1
		}
		env.Bra = env.Cursor
		if !env.InGroupingB(gCG, 99, 103) {
			env.Cursor = env.Limit - v2
			break lab1
		}
		if !rv(env, ctx) {
			env.Cursor = env.Limit - v2
			break lab1
		}
		if !env.SliceDel() {
			return false
		}
		break lab1
	}
	return true
}

// Stem runs the Snowball italian algorithm over env, mirroring the generated
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
		if !attachedPronoun(env, ctx) {
			break lab2
		}
		break lab2
	}
	env.Cursor = env.Limit - v3
	v4 := env.Limit - env.Cursor
lab3:
	for {
	lab4:
		for {
			v5 := env.Limit - env.Cursor
		lab5:
			for {
				if !standardSuffix(env, ctx) {
					break lab5
				}
				break lab4
			}
			env.Cursor = env.Limit - v5
			if !verbSuffix(env, ctx) {
				break lab3
			}
			break lab4
		}
		break lab3
	}
	env.Cursor = env.Limit - v4
	v6 := env.Limit - env.Cursor
lab6:
	for {
		if !vowelSuffix(env, ctx) {
			break lab6
		}
		break lab6
	}
	env.Cursor = env.Limit - v6
	env.Cursor = env.LimitBackward
	v7 := env.Cursor
lab7:
	for {
		if !postlude(env, ctx) {
			break lab7
		}
		break lab7
	}
	env.Cursor = v7
	return true
}
