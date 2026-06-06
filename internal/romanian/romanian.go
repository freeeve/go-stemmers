// Package romanian is a byte-faithful Go port of rust-stemmers' generated
// Snowball "romanian" stemmer. It produces output identical to rust-stemmers
// 1.2.0's Romanian algorithm; the canonical Snowball romanian vocabulary is the
// conformance oracle.
//
// The control flow mirrors the generated Rust: each Snowball labelled block
// (`'lab: loop { … break 'lab }`) becomes a Go labelled `for { … break LABEL }`,
// and the result-code dispatch (`else if among_var == N`) becomes a `switch`.
package romanian

import "github.com/freeeve/go-stemmers/internal/snowball"

// context holds the per-run state the generated algorithm threads through its
// routines: whether a standard suffix was removed and the RV/R1/R2 region marks.
type context struct {
	bStandardSuffixRemoved bool
	iP2                    int
	iP1                    int
	iPV                    int
}

var a0 = []snowball.Among[context]{
	{Str: "", SubstringI: -1, Result: 3},
	{Str: "I", SubstringI: 0, Result: 1},
	{Str: "U", SubstringI: 0, Result: 2},
}

var a1 = []snowball.Among[context]{
	{Str: "ea", SubstringI: -1, Result: 3},
	{Str: "aţia", SubstringI: -1, Result: 7},
	{Str: "aua", SubstringI: -1, Result: 2},
	{Str: "iua", SubstringI: -1, Result: 4},
	{Str: "aţie", SubstringI: -1, Result: 7},
	{Str: "ele", SubstringI: -1, Result: 3},
	{Str: "ile", SubstringI: -1, Result: 5},
	{Str: "iile", SubstringI: 6, Result: 4},
	{Str: "iei", SubstringI: -1, Result: 4},
	{Str: "atei", SubstringI: -1, Result: 6},
	{Str: "ii", SubstringI: -1, Result: 4},
	{Str: "ului", SubstringI: -1, Result: 1},
	{Str: "ul", SubstringI: -1, Result: 1},
	{Str: "elor", SubstringI: -1, Result: 3},
	{Str: "ilor", SubstringI: -1, Result: 4},
	{Str: "iilor", SubstringI: 14, Result: 4},
}

var a2 = []snowball.Among[context]{
	{Str: "icala", SubstringI: -1, Result: 4},
	{Str: "iciva", SubstringI: -1, Result: 4},
	{Str: "ativa", SubstringI: -1, Result: 5},
	{Str: "itiva", SubstringI: -1, Result: 6},
	{Str: "icale", SubstringI: -1, Result: 4},
	{Str: "aţiune", SubstringI: -1, Result: 5},
	{Str: "iţiune", SubstringI: -1, Result: 6},
	{Str: "atoare", SubstringI: -1, Result: 5},
	{Str: "itoare", SubstringI: -1, Result: 6},
	{Str: "ătoare", SubstringI: -1, Result: 5},
	{Str: "icitate", SubstringI: -1, Result: 4},
	{Str: "abilitate", SubstringI: -1, Result: 1},
	{Str: "ibilitate", SubstringI: -1, Result: 2},
	{Str: "ivitate", SubstringI: -1, Result: 3},
	{Str: "icive", SubstringI: -1, Result: 4},
	{Str: "ative", SubstringI: -1, Result: 5},
	{Str: "itive", SubstringI: -1, Result: 6},
	{Str: "icali", SubstringI: -1, Result: 4},
	{Str: "atori", SubstringI: -1, Result: 5},
	{Str: "icatori", SubstringI: 18, Result: 4},
	{Str: "itori", SubstringI: -1, Result: 6},
	{Str: "ători", SubstringI: -1, Result: 5},
	{Str: "icitati", SubstringI: -1, Result: 4},
	{Str: "abilitati", SubstringI: -1, Result: 1},
	{Str: "ivitati", SubstringI: -1, Result: 3},
	{Str: "icivi", SubstringI: -1, Result: 4},
	{Str: "ativi", SubstringI: -1, Result: 5},
	{Str: "itivi", SubstringI: -1, Result: 6},
	{Str: "icităi", SubstringI: -1, Result: 4},
	{Str: "abilităi", SubstringI: -1, Result: 1},
	{Str: "ivităi", SubstringI: -1, Result: 3},
	{Str: "icităţi", SubstringI: -1, Result: 4},
	{Str: "abilităţi", SubstringI: -1, Result: 1},
	{Str: "ivităţi", SubstringI: -1, Result: 3},
	{Str: "ical", SubstringI: -1, Result: 4},
	{Str: "ator", SubstringI: -1, Result: 5},
	{Str: "icator", SubstringI: 35, Result: 4},
	{Str: "itor", SubstringI: -1, Result: 6},
	{Str: "ător", SubstringI: -1, Result: 5},
	{Str: "iciv", SubstringI: -1, Result: 4},
	{Str: "ativ", SubstringI: -1, Result: 5},
	{Str: "itiv", SubstringI: -1, Result: 6},
	{Str: "icală", SubstringI: -1, Result: 4},
	{Str: "icivă", SubstringI: -1, Result: 4},
	{Str: "ativă", SubstringI: -1, Result: 5},
	{Str: "itivă", SubstringI: -1, Result: 6},
}

var a3 = []snowball.Among[context]{
	{Str: "ica", SubstringI: -1, Result: 1},
	{Str: "abila", SubstringI: -1, Result: 1},
	{Str: "ibila", SubstringI: -1, Result: 1},
	{Str: "oasa", SubstringI: -1, Result: 1},
	{Str: "ata", SubstringI: -1, Result: 1},
	{Str: "ita", SubstringI: -1, Result: 1},
	{Str: "anta", SubstringI: -1, Result: 1},
	{Str: "ista", SubstringI: -1, Result: 3},
	{Str: "uta", SubstringI: -1, Result: 1},
	{Str: "iva", SubstringI: -1, Result: 1},
	{Str: "ic", SubstringI: -1, Result: 1},
	{Str: "ice", SubstringI: -1, Result: 1},
	{Str: "abile", SubstringI: -1, Result: 1},
	{Str: "ibile", SubstringI: -1, Result: 1},
	{Str: "isme", SubstringI: -1, Result: 3},
	{Str: "iune", SubstringI: -1, Result: 2},
	{Str: "oase", SubstringI: -1, Result: 1},
	{Str: "ate", SubstringI: -1, Result: 1},
	{Str: "itate", SubstringI: 17, Result: 1},
	{Str: "ite", SubstringI: -1, Result: 1},
	{Str: "ante", SubstringI: -1, Result: 1},
	{Str: "iste", SubstringI: -1, Result: 3},
	{Str: "ute", SubstringI: -1, Result: 1},
	{Str: "ive", SubstringI: -1, Result: 1},
	{Str: "ici", SubstringI: -1, Result: 1},
	{Str: "abili", SubstringI: -1, Result: 1},
	{Str: "ibili", SubstringI: -1, Result: 1},
	{Str: "iuni", SubstringI: -1, Result: 2},
	{Str: "atori", SubstringI: -1, Result: 1},
	{Str: "osi", SubstringI: -1, Result: 1},
	{Str: "ati", SubstringI: -1, Result: 1},
	{Str: "itati", SubstringI: 30, Result: 1},
	{Str: "iti", SubstringI: -1, Result: 1},
	{Str: "anti", SubstringI: -1, Result: 1},
	{Str: "isti", SubstringI: -1, Result: 3},
	{Str: "uti", SubstringI: -1, Result: 1},
	{Str: "işti", SubstringI: -1, Result: 3},
	{Str: "ivi", SubstringI: -1, Result: 1},
	{Str: "ităi", SubstringI: -1, Result: 1},
	{Str: "oşi", SubstringI: -1, Result: 1},
	{Str: "ităţi", SubstringI: -1, Result: 1},
	{Str: "abil", SubstringI: -1, Result: 1},
	{Str: "ibil", SubstringI: -1, Result: 1},
	{Str: "ism", SubstringI: -1, Result: 3},
	{Str: "ator", SubstringI: -1, Result: 1},
	{Str: "os", SubstringI: -1, Result: 1},
	{Str: "at", SubstringI: -1, Result: 1},
	{Str: "it", SubstringI: -1, Result: 1},
	{Str: "ant", SubstringI: -1, Result: 1},
	{Str: "ist", SubstringI: -1, Result: 3},
	{Str: "ut", SubstringI: -1, Result: 1},
	{Str: "iv", SubstringI: -1, Result: 1},
	{Str: "ică", SubstringI: -1, Result: 1},
	{Str: "abilă", SubstringI: -1, Result: 1},
	{Str: "ibilă", SubstringI: -1, Result: 1},
	{Str: "oasă", SubstringI: -1, Result: 1},
	{Str: "ată", SubstringI: -1, Result: 1},
	{Str: "ită", SubstringI: -1, Result: 1},
	{Str: "antă", SubstringI: -1, Result: 1},
	{Str: "istă", SubstringI: -1, Result: 3},
	{Str: "ută", SubstringI: -1, Result: 1},
	{Str: "ivă", SubstringI: -1, Result: 1},
}

var a4 = []snowball.Among[context]{
	{Str: "ea", SubstringI: -1, Result: 1},
	{Str: "ia", SubstringI: -1, Result: 1},
	{Str: "esc", SubstringI: -1, Result: 1},
	{Str: "ăsc", SubstringI: -1, Result: 1},
	{Str: "ind", SubstringI: -1, Result: 1},
	{Str: "ând", SubstringI: -1, Result: 1},
	{Str: "are", SubstringI: -1, Result: 1},
	{Str: "ere", SubstringI: -1, Result: 1},
	{Str: "ire", SubstringI: -1, Result: 1},
	{Str: "âre", SubstringI: -1, Result: 1},
	{Str: "se", SubstringI: -1, Result: 2},
	{Str: "ase", SubstringI: 10, Result: 1},
	{Str: "sese", SubstringI: 10, Result: 2},
	{Str: "ise", SubstringI: 10, Result: 1},
	{Str: "use", SubstringI: 10, Result: 1},
	{Str: "âse", SubstringI: 10, Result: 1},
	{Str: "eşte", SubstringI: -1, Result: 1},
	{Str: "ăşte", SubstringI: -1, Result: 1},
	{Str: "eze", SubstringI: -1, Result: 1},
	{Str: "ai", SubstringI: -1, Result: 1},
	{Str: "eai", SubstringI: 19, Result: 1},
	{Str: "iai", SubstringI: 19, Result: 1},
	{Str: "sei", SubstringI: -1, Result: 2},
	{Str: "eşti", SubstringI: -1, Result: 1},
	{Str: "ăşti", SubstringI: -1, Result: 1},
	{Str: "ui", SubstringI: -1, Result: 1},
	{Str: "ezi", SubstringI: -1, Result: 1},
	{Str: "aşi", SubstringI: -1, Result: 1},
	{Str: "seşi", SubstringI: -1, Result: 2},
	{Str: "aseşi", SubstringI: 28, Result: 1},
	{Str: "seseşi", SubstringI: 28, Result: 2},
	{Str: "iseşi", SubstringI: 28, Result: 1},
	{Str: "useşi", SubstringI: 28, Result: 1},
	{Str: "âseşi", SubstringI: 28, Result: 1},
	{Str: "işi", SubstringI: -1, Result: 1},
	{Str: "uşi", SubstringI: -1, Result: 1},
	{Str: "âşi", SubstringI: -1, Result: 1},
	{Str: "âi", SubstringI: -1, Result: 1},
	{Str: "aţi", SubstringI: -1, Result: 2},
	{Str: "eaţi", SubstringI: 38, Result: 1},
	{Str: "iaţi", SubstringI: 38, Result: 1},
	{Str: "eţi", SubstringI: -1, Result: 2},
	{Str: "iţi", SubstringI: -1, Result: 2},
	{Str: "arăţi", SubstringI: -1, Result: 1},
	{Str: "serăţi", SubstringI: -1, Result: 2},
	{Str: "aserăţi", SubstringI: 44, Result: 1},
	{Str: "seserăţi", SubstringI: 44, Result: 2},
	{Str: "iserăţi", SubstringI: 44, Result: 1},
	{Str: "userăţi", SubstringI: 44, Result: 1},
	{Str: "âserăţi", SubstringI: 44, Result: 1},
	{Str: "irăţi", SubstringI: -1, Result: 1},
	{Str: "urăţi", SubstringI: -1, Result: 1},
	{Str: "ârăţi", SubstringI: -1, Result: 1},
	{Str: "âţi", SubstringI: -1, Result: 2},
	{Str: "am", SubstringI: -1, Result: 1},
	{Str: "eam", SubstringI: 54, Result: 1},
	{Str: "iam", SubstringI: 54, Result: 1},
	{Str: "em", SubstringI: -1, Result: 2},
	{Str: "asem", SubstringI: 57, Result: 1},
	{Str: "sesem", SubstringI: 57, Result: 2},
	{Str: "isem", SubstringI: 57, Result: 1},
	{Str: "usem", SubstringI: 57, Result: 1},
	{Str: "âsem", SubstringI: 57, Result: 1},
	{Str: "im", SubstringI: -1, Result: 2},
	{Str: "ăm", SubstringI: -1, Result: 2},
	{Str: "arăm", SubstringI: 64, Result: 1},
	{Str: "serăm", SubstringI: 64, Result: 2},
	{Str: "aserăm", SubstringI: 66, Result: 1},
	{Str: "seserăm", SubstringI: 66, Result: 2},
	{Str: "iserăm", SubstringI: 66, Result: 1},
	{Str: "userăm", SubstringI: 66, Result: 1},
	{Str: "âserăm", SubstringI: 66, Result: 1},
	{Str: "irăm", SubstringI: 64, Result: 1},
	{Str: "urăm", SubstringI: 64, Result: 1},
	{Str: "ârăm", SubstringI: 64, Result: 1},
	{Str: "âm", SubstringI: -1, Result: 2},
	{Str: "au", SubstringI: -1, Result: 1},
	{Str: "eau", SubstringI: 76, Result: 1},
	{Str: "iau", SubstringI: 76, Result: 1},
	{Str: "indu", SubstringI: -1, Result: 1},
	{Str: "ându", SubstringI: -1, Result: 1},
	{Str: "ez", SubstringI: -1, Result: 1},
	{Str: "ească", SubstringI: -1, Result: 1},
	{Str: "ară", SubstringI: -1, Result: 1},
	{Str: "seră", SubstringI: -1, Result: 2},
	{Str: "aseră", SubstringI: 84, Result: 1},
	{Str: "seseră", SubstringI: 84, Result: 2},
	{Str: "iseră", SubstringI: 84, Result: 1},
	{Str: "useră", SubstringI: 84, Result: 1},
	{Str: "âseră", SubstringI: 84, Result: 1},
	{Str: "iră", SubstringI: -1, Result: 1},
	{Str: "ură", SubstringI: -1, Result: 1},
	{Str: "âră", SubstringI: -1, Result: 1},
	{Str: "ează", SubstringI: -1, Result: 1},
}

var a5 = []snowball.Among[context]{
	{Str: "a", SubstringI: -1, Result: 1},
	{Str: "e", SubstringI: -1, Result: 1},
	{Str: "ie", SubstringI: 1, Result: 1},
	{Str: "i", SubstringI: -1, Result: 1},
	{Str: "ă", SubstringI: -1, Result: 1},
}

var gV = []byte{17, 65, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 32, 0, 0, 4}

// prelude upper-cases the semivowel u/i between two vowels to U/I so the
// suffix-stripping steps treat them as consonants.
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
					if !env.InGrouping(gV, 97, 259) {
						break lab3
					}
					env.Bra = env.Cursor
				lab4:
					for {
						v3 := env.Cursor
					lab5:
						for {
							if !env.EqS("u") {
								break lab5
							}
							env.Ket = env.Cursor
							if !env.InGrouping(gV, 97, 259) {
								break lab5
							}
							if !env.SliceFrom("U") {
								return false
							}
							break lab4
						}
						env.Cursor = v3
						if !env.EqS("i") {
							break lab3
						}
						env.Ket = env.Cursor
						if !env.InGrouping(gV, 97, 259) {
							break lab3
						}
						if !env.SliceFrom("I") {
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

// markRegions sets the RV (iPV), R1 (iP1) and R2 (iP2) region boundaries.
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
				if !env.InGrouping(gV, 97, 259) {
					break lab2
				}
			lab3:
				for {
					v3 := env.Cursor
				lab4:
					for {
						if !env.OutGrouping(gV, 97, 259) {
							break lab4
						}
					golab5:
						for {
							for {
								if !env.InGrouping(gV, 97, 259) {
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
					if !env.InGrouping(gV, 97, 259) {
						break lab2
					}
				golab7:
					for {
						for {
							if !env.OutGrouping(gV, 97, 259) {
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
			if !env.OutGrouping(gV, 97, 259) {
				break lab0
			}
		lab9:
			for {
				v6 := env.Cursor
			lab10:
				for {
					if !env.OutGrouping(gV, 97, 259) {
						break lab10
					}
				golab11:
					for {
						for {
							if !env.InGrouping(gV, 97, 259) {
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
				if !env.InGrouping(gV, 97, 259) {
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
				if !env.InGrouping(gV, 97, 259) {
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
				if !env.OutGrouping(gV, 97, 259) {
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
				if !env.InGrouping(gV, 97, 259) {
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
				if !env.OutGrouping(gV, 97, 259) {
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

// postlude restores the U/I introduced by the prelude back to u/i.
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

// step0 strips or rewrites the noun/adjective plural and article endings in a1
// when they fall in R1.
func step0(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a1, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !r1(env, ctx) {
		return false
	}
	switch amongVar {
	case 1:
		if !env.SliceDel() {
			return false
		}
	case 2:
		if !env.SliceFrom("a") {
			return false
		}
	case 3:
		if !env.SliceFrom("e") {
			return false
		}
	case 4:
		if !env.SliceFrom("i") {
			return false
		}
	case 5:
		v1 := env.Limit - env.Cursor
	lab0:
		for {
			if !env.EqSB("ab") {
				break lab0
			}
			return false
		}
		env.Cursor = env.Limit - v1
		if !env.SliceFrom("i") {
			return false
		}
	case 6:
		if !env.SliceFrom("at") {
			return false
		}
	case 7:
		if !env.SliceFrom("aţi") {
			return false
		}
	}
	return true
}

// comboSuffix rewrites the combining derivational suffixes in a2 (in R1) to their
// shorter forms and records that a standard suffix was removed.
func comboSuffix(env *snowball.Env, ctx *context) bool {
	v1 := env.Limit - env.Cursor
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a2, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !r1(env, ctx) {
		return false
	}
	switch amongVar {
	case 1:
		if !env.SliceFrom("abil") {
			return false
		}
	case 2:
		if !env.SliceFrom("ibil") {
			return false
		}
	case 3:
		if !env.SliceFrom("iv") {
			return false
		}
	case 4:
		if !env.SliceFrom("ic") {
			return false
		}
	case 5:
		if !env.SliceFrom("at") {
			return false
		}
	case 6:
		if !env.SliceFrom("it") {
			return false
		}
	}
	ctx.bStandardSuffixRemoved = true
	env.Cursor = env.Limit - v1
	return true
}

// standardSuffix repeatedly applies comboSuffix then removes or rewrites the
// derivational suffixes in a3 when they fall in R2.
func standardSuffix(env *snowball.Env, ctx *context) bool {
	ctx.bStandardSuffixRemoved = false
replab0:
	for {
		v1 := env.Limit - env.Cursor
	lab1:
		for once := 0; once < 1; once++ {
			if !comboSuffix(env, ctx) {
				break lab1
			}
			continue replab0
		}
		env.Cursor = env.Limit - v1
		break replab0
	}
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a3, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !r2(env, ctx) {
		return false
	}
	switch amongVar {
	case 1:
		if !env.SliceDel() {
			return false
		}
	case 2:
		if !env.EqSB("ţ") {
			return false
		}
		env.Bra = env.Cursor
		if !env.SliceFrom("t") {
			return false
		}
	case 3:
		if !env.SliceFrom("ist") {
			return false
		}
	}
	ctx.bStandardSuffixRemoved = true
	return true
}

// verbSuffix removes the verb conjugation endings in a4 within region RV.
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
	amongVar := snowball.FindAmongB(env, a4, ctx)
	if amongVar == 0 {
		env.LimitBackward = v2
		return false
	}
	env.Bra = env.Cursor
	switch amongVar {
	case 1:
	lab0:
		for {
			v3 := env.Limit - env.Cursor
		lab1:
			for {
				if !env.OutGroupingB(gV, 97, 259) {
					break lab1
				}
				break lab0
			}
			env.Cursor = env.Limit - v3
			if !env.EqSB("u") {
				env.LimitBackward = v2
				return false
			}
			break lab0
		}
		if !env.SliceDel() {
			return false
		}
	case 2:
		if !env.SliceDel() {
			return false
		}
	}
	env.LimitBackward = v2
	return true
}

// vowelSuffix removes a final vowel ending in a5 when it falls in region RV.
func vowelSuffix(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a5, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !rv(env, ctx) {
		return false
	}
	switch amongVar {
	case 1:
		if !env.SliceDel() {
			return false
		}
	}
	return true
}

// Stem runs the Snowball romanian algorithm over env, mirroring the generated
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
	step0(env, ctx)
	env.Cursor = env.Limit - v3

	v4 := env.Limit - env.Cursor
	standardSuffix(env, ctx)
	env.Cursor = env.Limit - v4

	v5 := env.Limit - env.Cursor
lab4:
	for {
	lab5:
		for {
			v6 := env.Limit - env.Cursor
		lab6:
			for {
				if !ctx.bStandardSuffixRemoved {
					break lab6
				}
				break lab5
			}
			env.Cursor = env.Limit - v6
			if !verbSuffix(env, ctx) {
				break lab4
			}
			break lab5
		}
		break lab4
	}
	env.Cursor = env.Limit - v5

	v7 := env.Limit - env.Cursor
	vowelSuffix(env, ctx)
	env.Cursor = env.Limit - v7

	env.Cursor = env.LimitBackward

	v8 := env.Cursor
	postlude(env, ctx)
	env.Cursor = v8

	return true
}
