// Package spanish is a byte-faithful Go port of rust-stemmers' generated
// Snowball "spanish" stemmer. It produces output identical to rust-stemmers
// 1.2.0's Spanish algorithm; the canonical Snowball spanish vocabulary is the
// conformance oracle.
//
// The control flow mirrors the generated Rust: each Snowball labelled block
// (`'lab: loop { … break 'lab }`) becomes a Go labelled `for { … break LABEL }`,
// and the result-code dispatch (`else if among_var == N`) becomes a `switch`.
package spanish

import "github.com/freeeve/go-stemmers/internal/snowball"

// context holds the per-run region marks the generated algorithm threads through
// its routines: the RV (iPV), R1 (iP1), and R2 (iP2) boundaries.
type context struct {
	iP2 int
	iP1 int
	iPV int
}

var a0 = []snowball.Among[context]{
	{Str: "", SubstringI: -1, Result: 6},
	{Str: "á", SubstringI: 0, Result: 1},
	{Str: "é", SubstringI: 0, Result: 2},
	{Str: "í", SubstringI: 0, Result: 3},
	{Str: "ó", SubstringI: 0, Result: 4},
	{Str: "ú", SubstringI: 0, Result: 5},
}

var a1 = []snowball.Among[context]{
	{Str: "la", SubstringI: -1, Result: -1},
	{Str: "sela", SubstringI: 0, Result: -1},
	{Str: "le", SubstringI: -1, Result: -1},
	{Str: "me", SubstringI: -1, Result: -1},
	{Str: "se", SubstringI: -1, Result: -1},
	{Str: "lo", SubstringI: -1, Result: -1},
	{Str: "selo", SubstringI: 5, Result: -1},
	{Str: "las", SubstringI: -1, Result: -1},
	{Str: "selas", SubstringI: 7, Result: -1},
	{Str: "les", SubstringI: -1, Result: -1},
	{Str: "los", SubstringI: -1, Result: -1},
	{Str: "selos", SubstringI: 10, Result: -1},
	{Str: "nos", SubstringI: -1, Result: -1},
}

var a2 = []snowball.Among[context]{
	{Str: "ando", SubstringI: -1, Result: 6},
	{Str: "iendo", SubstringI: -1, Result: 6},
	{Str: "yendo", SubstringI: -1, Result: 7},
	{Str: "ándo", SubstringI: -1, Result: 2},
	{Str: "iéndo", SubstringI: -1, Result: 1},
	{Str: "ar", SubstringI: -1, Result: 6},
	{Str: "er", SubstringI: -1, Result: 6},
	{Str: "ir", SubstringI: -1, Result: 6},
	{Str: "ár", SubstringI: -1, Result: 3},
	{Str: "ér", SubstringI: -1, Result: 4},
	{Str: "ír", SubstringI: -1, Result: 5},
}

var a3 = []snowball.Among[context]{
	{Str: "ic", SubstringI: -1, Result: -1},
	{Str: "ad", SubstringI: -1, Result: -1},
	{Str: "os", SubstringI: -1, Result: -1},
	{Str: "iv", SubstringI: -1, Result: 1},
}

var a4 = []snowball.Among[context]{
	{Str: "able", SubstringI: -1, Result: 1},
	{Str: "ible", SubstringI: -1, Result: 1},
	{Str: "ante", SubstringI: -1, Result: 1},
}

var a5 = []snowball.Among[context]{
	{Str: "ic", SubstringI: -1, Result: 1},
	{Str: "abil", SubstringI: -1, Result: 1},
	{Str: "iv", SubstringI: -1, Result: 1},
}

var a6 = []snowball.Among[context]{
	{Str: "ica", SubstringI: -1, Result: 1},
	{Str: "ancia", SubstringI: -1, Result: 2},
	{Str: "encia", SubstringI: -1, Result: 5},
	{Str: "adora", SubstringI: -1, Result: 2},
	{Str: "osa", SubstringI: -1, Result: 1},
	{Str: "ista", SubstringI: -1, Result: 1},
	{Str: "iva", SubstringI: -1, Result: 9},
	{Str: "anza", SubstringI: -1, Result: 1},
	{Str: "logía", SubstringI: -1, Result: 3},
	{Str: "idad", SubstringI: -1, Result: 8},
	{Str: "able", SubstringI: -1, Result: 1},
	{Str: "ible", SubstringI: -1, Result: 1},
	{Str: "ante", SubstringI: -1, Result: 2},
	{Str: "mente", SubstringI: -1, Result: 7},
	{Str: "amente", SubstringI: 13, Result: 6},
	{Str: "ación", SubstringI: -1, Result: 2},
	{Str: "ución", SubstringI: -1, Result: 4},
	{Str: "ico", SubstringI: -1, Result: 1},
	{Str: "ismo", SubstringI: -1, Result: 1},
	{Str: "oso", SubstringI: -1, Result: 1},
	{Str: "amiento", SubstringI: -1, Result: 1},
	{Str: "imiento", SubstringI: -1, Result: 1},
	{Str: "ivo", SubstringI: -1, Result: 9},
	{Str: "ador", SubstringI: -1, Result: 2},
	{Str: "icas", SubstringI: -1, Result: 1},
	{Str: "ancias", SubstringI: -1, Result: 2},
	{Str: "encias", SubstringI: -1, Result: 5},
	{Str: "adoras", SubstringI: -1, Result: 2},
	{Str: "osas", SubstringI: -1, Result: 1},
	{Str: "istas", SubstringI: -1, Result: 1},
	{Str: "ivas", SubstringI: -1, Result: 9},
	{Str: "anzas", SubstringI: -1, Result: 1},
	{Str: "logías", SubstringI: -1, Result: 3},
	{Str: "idades", SubstringI: -1, Result: 8},
	{Str: "ables", SubstringI: -1, Result: 1},
	{Str: "ibles", SubstringI: -1, Result: 1},
	{Str: "aciones", SubstringI: -1, Result: 2},
	{Str: "uciones", SubstringI: -1, Result: 4},
	{Str: "adores", SubstringI: -1, Result: 2},
	{Str: "antes", SubstringI: -1, Result: 2},
	{Str: "icos", SubstringI: -1, Result: 1},
	{Str: "ismos", SubstringI: -1, Result: 1},
	{Str: "osos", SubstringI: -1, Result: 1},
	{Str: "amientos", SubstringI: -1, Result: 1},
	{Str: "imientos", SubstringI: -1, Result: 1},
	{Str: "ivos", SubstringI: -1, Result: 9},
}

var a7 = []snowball.Among[context]{
	{Str: "ya", SubstringI: -1, Result: 1},
	{Str: "ye", SubstringI: -1, Result: 1},
	{Str: "yan", SubstringI: -1, Result: 1},
	{Str: "yen", SubstringI: -1, Result: 1},
	{Str: "yeron", SubstringI: -1, Result: 1},
	{Str: "yendo", SubstringI: -1, Result: 1},
	{Str: "yo", SubstringI: -1, Result: 1},
	{Str: "yas", SubstringI: -1, Result: 1},
	{Str: "yes", SubstringI: -1, Result: 1},
	{Str: "yais", SubstringI: -1, Result: 1},
	{Str: "yamos", SubstringI: -1, Result: 1},
	{Str: "yó", SubstringI: -1, Result: 1},
}

var a8 = []snowball.Among[context]{
	{Str: "aba", SubstringI: -1, Result: 2},
	{Str: "ada", SubstringI: -1, Result: 2},
	{Str: "ida", SubstringI: -1, Result: 2},
	{Str: "ara", SubstringI: -1, Result: 2},
	{Str: "iera", SubstringI: -1, Result: 2},
	{Str: "ía", SubstringI: -1, Result: 2},
	{Str: "aría", SubstringI: 5, Result: 2},
	{Str: "ería", SubstringI: 5, Result: 2},
	{Str: "iría", SubstringI: 5, Result: 2},
	{Str: "ad", SubstringI: -1, Result: 2},
	{Str: "ed", SubstringI: -1, Result: 2},
	{Str: "id", SubstringI: -1, Result: 2},
	{Str: "ase", SubstringI: -1, Result: 2},
	{Str: "iese", SubstringI: -1, Result: 2},
	{Str: "aste", SubstringI: -1, Result: 2},
	{Str: "iste", SubstringI: -1, Result: 2},
	{Str: "an", SubstringI: -1, Result: 2},
	{Str: "aban", SubstringI: 16, Result: 2},
	{Str: "aran", SubstringI: 16, Result: 2},
	{Str: "ieran", SubstringI: 16, Result: 2},
	{Str: "ían", SubstringI: 16, Result: 2},
	{Str: "arían", SubstringI: 20, Result: 2},
	{Str: "erían", SubstringI: 20, Result: 2},
	{Str: "irían", SubstringI: 20, Result: 2},
	{Str: "en", SubstringI: -1, Result: 1},
	{Str: "asen", SubstringI: 24, Result: 2},
	{Str: "iesen", SubstringI: 24, Result: 2},
	{Str: "aron", SubstringI: -1, Result: 2},
	{Str: "ieron", SubstringI: -1, Result: 2},
	{Str: "arán", SubstringI: -1, Result: 2},
	{Str: "erán", SubstringI: -1, Result: 2},
	{Str: "irán", SubstringI: -1, Result: 2},
	{Str: "ado", SubstringI: -1, Result: 2},
	{Str: "ido", SubstringI: -1, Result: 2},
	{Str: "ando", SubstringI: -1, Result: 2},
	{Str: "iendo", SubstringI: -1, Result: 2},
	{Str: "ar", SubstringI: -1, Result: 2},
	{Str: "er", SubstringI: -1, Result: 2},
	{Str: "ir", SubstringI: -1, Result: 2},
	{Str: "as", SubstringI: -1, Result: 2},
	{Str: "abas", SubstringI: 39, Result: 2},
	{Str: "adas", SubstringI: 39, Result: 2},
	{Str: "idas", SubstringI: 39, Result: 2},
	{Str: "aras", SubstringI: 39, Result: 2},
	{Str: "ieras", SubstringI: 39, Result: 2},
	{Str: "ías", SubstringI: 39, Result: 2},
	{Str: "arías", SubstringI: 45, Result: 2},
	{Str: "erías", SubstringI: 45, Result: 2},
	{Str: "irías", SubstringI: 45, Result: 2},
	{Str: "es", SubstringI: -1, Result: 1},
	{Str: "ases", SubstringI: 49, Result: 2},
	{Str: "ieses", SubstringI: 49, Result: 2},
	{Str: "abais", SubstringI: -1, Result: 2},
	{Str: "arais", SubstringI: -1, Result: 2},
	{Str: "ierais", SubstringI: -1, Result: 2},
	{Str: "íais", SubstringI: -1, Result: 2},
	{Str: "aríais", SubstringI: 55, Result: 2},
	{Str: "eríais", SubstringI: 55, Result: 2},
	{Str: "iríais", SubstringI: 55, Result: 2},
	{Str: "aseis", SubstringI: -1, Result: 2},
	{Str: "ieseis", SubstringI: -1, Result: 2},
	{Str: "asteis", SubstringI: -1, Result: 2},
	{Str: "isteis", SubstringI: -1, Result: 2},
	{Str: "áis", SubstringI: -1, Result: 2},
	{Str: "éis", SubstringI: -1, Result: 1},
	{Str: "aréis", SubstringI: 64, Result: 2},
	{Str: "eréis", SubstringI: 64, Result: 2},
	{Str: "iréis", SubstringI: 64, Result: 2},
	{Str: "ados", SubstringI: -1, Result: 2},
	{Str: "idos", SubstringI: -1, Result: 2},
	{Str: "amos", SubstringI: -1, Result: 2},
	{Str: "ábamos", SubstringI: 70, Result: 2},
	{Str: "áramos", SubstringI: 70, Result: 2},
	{Str: "iéramos", SubstringI: 70, Result: 2},
	{Str: "íamos", SubstringI: 70, Result: 2},
	{Str: "aríamos", SubstringI: 74, Result: 2},
	{Str: "eríamos", SubstringI: 74, Result: 2},
	{Str: "iríamos", SubstringI: 74, Result: 2},
	{Str: "emos", SubstringI: -1, Result: 1},
	{Str: "aremos", SubstringI: 78, Result: 2},
	{Str: "eremos", SubstringI: 78, Result: 2},
	{Str: "iremos", SubstringI: 78, Result: 2},
	{Str: "ásemos", SubstringI: 78, Result: 2},
	{Str: "iésemos", SubstringI: 78, Result: 2},
	{Str: "imos", SubstringI: -1, Result: 2},
	{Str: "arás", SubstringI: -1, Result: 2},
	{Str: "erás", SubstringI: -1, Result: 2},
	{Str: "irás", SubstringI: -1, Result: 2},
	{Str: "ís", SubstringI: -1, Result: 2},
	{Str: "ará", SubstringI: -1, Result: 2},
	{Str: "erá", SubstringI: -1, Result: 2},
	{Str: "irá", SubstringI: -1, Result: 2},
	{Str: "aré", SubstringI: -1, Result: 2},
	{Str: "eré", SubstringI: -1, Result: 2},
	{Str: "iré", SubstringI: -1, Result: 2},
	{Str: "ió", SubstringI: -1, Result: 2},
}

var a9 = []snowball.Among[context]{
	{Str: "a", SubstringI: -1, Result: 1},
	{Str: "e", SubstringI: -1, Result: 2},
	{Str: "o", SubstringI: -1, Result: 1},
	{Str: "os", SubstringI: -1, Result: 1},
	{Str: "á", SubstringI: -1, Result: 1},
	{Str: "é", SubstringI: -1, Result: 2},
	{Str: "í", SubstringI: -1, Result: 1},
	{Str: "ó", SubstringI: -1, Result: 1},
}

var gV = []byte{17, 65, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 17, 4, 10}

// markRegions sets the RV (iPV), R1 (iP1) and R2 (iP2) region boundaries used by
// the suffix-stripping routines.
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
				if !env.InGrouping(gV, 97, 252) {
					break lab2
				}
			lab3:
				for {
					v3 := env.Cursor
				lab4:
					for {
						if !env.OutGrouping(gV, 97, 252) {
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
								break lab4
							}
							env.NextChar()
						}
						break lab3
					}
					env.Cursor = v3
					if !env.InGrouping(gV, 97, 252) {
						break lab2
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
							break lab2
						}
						env.NextChar()
					}
					break lab3
				}
				break lab1
			}
			env.Cursor = v2
			if !env.OutGrouping(gV, 97, 252) {
				break lab0
			}
		lab9:
			for {
				v6 := env.Cursor
			lab10:
				for {
					if !env.OutGrouping(gV, 97, 252) {
						break lab10
					}
				golab11:
					for {
						for {
							if !env.InGrouping(gV, 97, 252) {
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
				if !env.InGrouping(gV, 97, 252) {
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
				if !env.InGrouping(gV, 97, 252) {
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
				if !env.OutGrouping(gV, 97, 252) {
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
				if !env.InGrouping(gV, 97, 252) {
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
				if !env.OutGrouping(gV, 97, 252) {
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

// postlude restores the acute-accented vowels (á→a, é→e, í→i, ó→o, ú→u) that
// survived to the end of the word back to their plain forms.
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

// attachedPronoun strips an enclitic object pronoun (la, le, me, se, lo, nos …)
// from a gerund or infinitive form within RV, restoring the verb ending.
func attachedPronoun(env *snowball.Env, ctx *context) bool {
	var amongVar int32
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a1, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	amongVar = snowball.FindAmongB(env, a2, ctx)
	if amongVar == 0 {
		return false
	}
	if !rV(env, ctx) {
		return false
	}
	switch amongVar {
	case 1:
		env.Bra = env.Cursor
		if !env.SliceFrom("iendo") {
			return false
		}
	case 2:
		env.Bra = env.Cursor
		if !env.SliceFrom("ando") {
			return false
		}
	case 3:
		env.Bra = env.Cursor
		if !env.SliceFrom("ar") {
			return false
		}
	case 4:
		env.Bra = env.Cursor
		if !env.SliceFrom("er") {
			return false
		}
	case 5:
		env.Bra = env.Cursor
		if !env.SliceFrom("ir") {
			return false
		}
	case 6:
		if !env.SliceDel() {
			return false
		}
	case 7:
		if !env.EqSB("u") {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	}
	return true
}

// standardSuffix rewrites or deletes the standard derivational suffixes in a6
// (within R1/R2), with the chained clean-up passes over a3/a4/a5.
func standardSuffix(env *snowball.Env, ctx *context) bool {
	var amongVar int32
	env.Ket = env.Cursor
	amongVar = snowball.FindAmongB(env, a6, ctx)
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
			amongVar = snowball.FindAmongB(env, a3, ctx)
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
	case 7:
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
			amongVar = snowball.FindAmongB(env, a4, ctx)
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
	case 8:
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
			amongVar = snowball.FindAmongB(env, a5, ctx)
			if amongVar == 0 {
				env.Cursor = env.Limit - v4
				break lab3
			}
			env.Bra = env.Cursor
			switch amongVar {
			case 1:
				if !r2(env, ctx) {
					env.Cursor = env.Limit - v4
					break lab3
				}
				if !env.SliceDel() {
					return false
				}
			}
			break lab3
		}
	case 9:
		if !r2(env, ctx) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
		v5 := env.Limit - env.Cursor
	lab4:
		for {
			env.Ket = env.Cursor
			if !env.EqSB("at") {
				env.Cursor = env.Limit - v5
				break lab4
			}
			env.Bra = env.Cursor
			if !r2(env, ctx) {
				env.Cursor = env.Limit - v5
				break lab4
			}
			if !env.SliceDel() {
				return false
			}
			break lab4
		}
	}
	return true
}

// yVerbSuffix removes a y-initial verb ending (after a u) within RV.
func yVerbSuffix(env *snowball.Env, ctx *context) bool {
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
	amongVar = snowball.FindAmongB(env, a7, ctx)
	if amongVar == 0 {
		env.LimitBackward = v2
		return false
	}
	env.Bra = env.Cursor
	env.LimitBackward = v2
	switch amongVar {
	case 1:
		if !env.EqSB("u") {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	}
	return true
}

// verbSuffix removes the remaining verb conjugation endings in a8 within RV,
// guarding the -gu→-g case.
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
	amongVar = snowball.FindAmongB(env, a8, ctx)
	if amongVar == 0 {
		env.LimitBackward = v2
		return false
	}
	env.Bra = env.Cursor
	env.LimitBackward = v2
	switch amongVar {
	case 1:
		v3 := env.Limit - env.Cursor
	lab0:
		for {
			if !env.EqSB("u") {
				env.Cursor = env.Limit - v3
				break lab0
			}
			v4 := env.Limit - env.Cursor
			if !env.EqSB("g") {
				env.Cursor = env.Limit - v3
				break lab0
			}
			env.Cursor = env.Limit - v4
			break lab0
		}
		env.Bra = env.Cursor
		if !env.SliceDel() {
			return false
		}
	case 2:
		if !env.SliceDel() {
			return false
		}
	}
	return true
}

// residualSuffix removes a residual final vowel within RV, with the -gu→-g
// guard for the -e case.
func residualSuffix(env *snowball.Env, ctx *context) bool {
	var amongVar int32
	env.Ket = env.Cursor
	amongVar = snowball.FindAmongB(env, a9, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	switch amongVar {
	case 1:
		if !rV(env, ctx) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	case 2:
		if !rV(env, ctx) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
		v1 := env.Limit - env.Cursor
	lab0:
		for {
			env.Ket = env.Cursor
			if !env.EqSB("u") {
				env.Cursor = env.Limit - v1
				break lab0
			}
			env.Bra = env.Cursor
			v2 := env.Limit - env.Cursor
			if !env.EqSB("g") {
				env.Cursor = env.Limit - v1
				break lab0
			}
			env.Cursor = env.Limit - v2
			if !rV(env, ctx) {
				env.Cursor = env.Limit - v1
				break lab0
			}
			if !env.SliceDel() {
				return false
			}
			break lab0
		}
	}
	return true
}

// Stem runs the Snowball spanish algorithm over env, mirroring the generated
// `stem` entry point. It always returns true; the result is the mutated env.
func Stem(env *snowball.Env) bool {
	ctx := &context{}
	v1 := env.Cursor
lab0:
	for {
		if !markRegions(env, ctx) {
			break lab0
		}
		break lab0
	}
	env.Cursor = v1
	env.LimitBackward = env.Cursor
	env.Cursor = env.Limit
	v2 := env.Limit - env.Cursor
lab1:
	for {
		if !attachedPronoun(env, ctx) {
			break lab1
		}
		break lab1
	}
	env.Cursor = env.Limit - v2
	v3 := env.Limit - env.Cursor
lab2:
	for {
	lab3:
		for {
			v4 := env.Limit - env.Cursor
		lab4:
			for {
				if !standardSuffix(env, ctx) {
					break lab4
				}
				break lab3
			}
			env.Cursor = env.Limit - v4
		lab5:
			for {
				if !yVerbSuffix(env, ctx) {
					break lab5
				}
				break lab3
			}
			env.Cursor = env.Limit - v4
			if !verbSuffix(env, ctx) {
				break lab2
			}
			break lab3
		}
		break lab2
	}
	env.Cursor = env.Limit - v3
	v5 := env.Limit - env.Cursor
lab6:
	for {
		if !residualSuffix(env, ctx) {
			break lab6
		}
		break lab6
	}
	env.Cursor = env.Limit - v5
	env.Cursor = env.LimitBackward
	v6 := env.Cursor
lab7:
	for {
		if !postlude(env, ctx) {
			break lab7
		}
		break lab7
	}
	env.Cursor = v6
	return true
}
