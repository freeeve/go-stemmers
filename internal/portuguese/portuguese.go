// Package portuguese is a byte-faithful Go port of rust-stemmers' generated
// Snowball "portuguese" stemmer. It produces output identical to
// rust-stemmers 1.2.0's Portuguese algorithm; the canonical Snowball
// portuguese vocabulary is the conformance oracle.
//
// The control flow mirrors the generated Rust: each Snowball labelled block
// (`'lab: loop { … break 'lab }`) becomes a Go labelled `for { … break LABEL }`,
// and the result-code dispatch (`else if among_var == N`) becomes a `switch`.
package portuguese

import "github.com/freeeve/go-stemmers/internal/snowball"

// context holds the per-run state the generated algorithm threads through its
// routines: the RV/R1/R2 region marks.
type context struct {
	iP2 int
	iP1 int
	iPV int
}

var a0 = []snowball.Among[context]{
	{Str: "", SubstringI: -1, Result: 3},
	{Str: "ã", SubstringI: 0, Result: 1},
	{Str: "õ", SubstringI: 0, Result: 2},
}

var a1 = []snowball.Among[context]{
	{Str: "", SubstringI: -1, Result: 3},
	{Str: "a~", SubstringI: 0, Result: 1},
	{Str: "o~", SubstringI: 0, Result: 2},
}

var a2 = []snowball.Among[context]{
	{Str: "ic", SubstringI: -1, Result: -1},
	{Str: "ad", SubstringI: -1, Result: -1},
	{Str: "os", SubstringI: -1, Result: -1},
	{Str: "iv", SubstringI: -1, Result: 1},
}

var a3 = []snowball.Among[context]{
	{Str: "ante", SubstringI: -1, Result: 1},
	{Str: "avel", SubstringI: -1, Result: 1},
	{Str: "ível", SubstringI: -1, Result: 1},
}

var a4 = []snowball.Among[context]{
	{Str: "ic", SubstringI: -1, Result: 1},
	{Str: "abil", SubstringI: -1, Result: 1},
	{Str: "iv", SubstringI: -1, Result: 1},
}

var a5 = []snowball.Among[context]{
	{Str: "ica", SubstringI: -1, Result: 1},
	{Str: "ância", SubstringI: -1, Result: 1},
	{Str: "ência", SubstringI: -1, Result: 4},
	{Str: "logia", SubstringI: -1, Result: 2},
	{Str: "ira", SubstringI: -1, Result: 9},
	{Str: "adora", SubstringI: -1, Result: 1},
	{Str: "osa", SubstringI: -1, Result: 1},
	{Str: "ista", SubstringI: -1, Result: 1},
	{Str: "iva", SubstringI: -1, Result: 8},
	{Str: "eza", SubstringI: -1, Result: 1},
	{Str: "idade", SubstringI: -1, Result: 7},
	{Str: "ante", SubstringI: -1, Result: 1},
	{Str: "mente", SubstringI: -1, Result: 6},
	{Str: "amente", SubstringI: 12, Result: 5},
	{Str: "ável", SubstringI: -1, Result: 1},
	{Str: "ível", SubstringI: -1, Result: 1},
	{Str: "ico", SubstringI: -1, Result: 1},
	{Str: "ismo", SubstringI: -1, Result: 1},
	{Str: "oso", SubstringI: -1, Result: 1},
	{Str: "amento", SubstringI: -1, Result: 1},
	{Str: "imento", SubstringI: -1, Result: 1},
	{Str: "ivo", SubstringI: -1, Result: 8},
	{Str: "aça~o", SubstringI: -1, Result: 1},
	{Str: "uça~o", SubstringI: -1, Result: 3},
	{Str: "ador", SubstringI: -1, Result: 1},
	{Str: "icas", SubstringI: -1, Result: 1},
	{Str: "ências", SubstringI: -1, Result: 4},
	{Str: "logias", SubstringI: -1, Result: 2},
	{Str: "iras", SubstringI: -1, Result: 9},
	{Str: "adoras", SubstringI: -1, Result: 1},
	{Str: "osas", SubstringI: -1, Result: 1},
	{Str: "istas", SubstringI: -1, Result: 1},
	{Str: "ivas", SubstringI: -1, Result: 8},
	{Str: "ezas", SubstringI: -1, Result: 1},
	{Str: "idades", SubstringI: -1, Result: 7},
	{Str: "adores", SubstringI: -1, Result: 1},
	{Str: "antes", SubstringI: -1, Result: 1},
	{Str: "aço~es", SubstringI: -1, Result: 1},
	{Str: "uço~es", SubstringI: -1, Result: 3},
	{Str: "icos", SubstringI: -1, Result: 1},
	{Str: "ismos", SubstringI: -1, Result: 1},
	{Str: "osos", SubstringI: -1, Result: 1},
	{Str: "amentos", SubstringI: -1, Result: 1},
	{Str: "imentos", SubstringI: -1, Result: 1},
	{Str: "ivos", SubstringI: -1, Result: 8},
}

var a6 = []snowball.Among[context]{
	{Str: "ada", SubstringI: -1, Result: 1},
	{Str: "ida", SubstringI: -1, Result: 1},
	{Str: "ia", SubstringI: -1, Result: 1},
	{Str: "aria", SubstringI: 2, Result: 1},
	{Str: "eria", SubstringI: 2, Result: 1},
	{Str: "iria", SubstringI: 2, Result: 1},
	{Str: "ara", SubstringI: -1, Result: 1},
	{Str: "era", SubstringI: -1, Result: 1},
	{Str: "ira", SubstringI: -1, Result: 1},
	{Str: "ava", SubstringI: -1, Result: 1},
	{Str: "asse", SubstringI: -1, Result: 1},
	{Str: "esse", SubstringI: -1, Result: 1},
	{Str: "isse", SubstringI: -1, Result: 1},
	{Str: "aste", SubstringI: -1, Result: 1},
	{Str: "este", SubstringI: -1, Result: 1},
	{Str: "iste", SubstringI: -1, Result: 1},
	{Str: "ei", SubstringI: -1, Result: 1},
	{Str: "arei", SubstringI: 16, Result: 1},
	{Str: "erei", SubstringI: 16, Result: 1},
	{Str: "irei", SubstringI: 16, Result: 1},
	{Str: "am", SubstringI: -1, Result: 1},
	{Str: "iam", SubstringI: 20, Result: 1},
	{Str: "ariam", SubstringI: 21, Result: 1},
	{Str: "eriam", SubstringI: 21, Result: 1},
	{Str: "iriam", SubstringI: 21, Result: 1},
	{Str: "aram", SubstringI: 20, Result: 1},
	{Str: "eram", SubstringI: 20, Result: 1},
	{Str: "iram", SubstringI: 20, Result: 1},
	{Str: "avam", SubstringI: 20, Result: 1},
	{Str: "em", SubstringI: -1, Result: 1},
	{Str: "arem", SubstringI: 29, Result: 1},
	{Str: "erem", SubstringI: 29, Result: 1},
	{Str: "irem", SubstringI: 29, Result: 1},
	{Str: "assem", SubstringI: 29, Result: 1},
	{Str: "essem", SubstringI: 29, Result: 1},
	{Str: "issem", SubstringI: 29, Result: 1},
	{Str: "ado", SubstringI: -1, Result: 1},
	{Str: "ido", SubstringI: -1, Result: 1},
	{Str: "ando", SubstringI: -1, Result: 1},
	{Str: "endo", SubstringI: -1, Result: 1},
	{Str: "indo", SubstringI: -1, Result: 1},
	{Str: "ara~o", SubstringI: -1, Result: 1},
	{Str: "era~o", SubstringI: -1, Result: 1},
	{Str: "ira~o", SubstringI: -1, Result: 1},
	{Str: "ar", SubstringI: -1, Result: 1},
	{Str: "er", SubstringI: -1, Result: 1},
	{Str: "ir", SubstringI: -1, Result: 1},
	{Str: "as", SubstringI: -1, Result: 1},
	{Str: "adas", SubstringI: 47, Result: 1},
	{Str: "idas", SubstringI: 47, Result: 1},
	{Str: "ias", SubstringI: 47, Result: 1},
	{Str: "arias", SubstringI: 50, Result: 1},
	{Str: "erias", SubstringI: 50, Result: 1},
	{Str: "irias", SubstringI: 50, Result: 1},
	{Str: "aras", SubstringI: 47, Result: 1},
	{Str: "eras", SubstringI: 47, Result: 1},
	{Str: "iras", SubstringI: 47, Result: 1},
	{Str: "avas", SubstringI: 47, Result: 1},
	{Str: "es", SubstringI: -1, Result: 1},
	{Str: "ardes", SubstringI: 58, Result: 1},
	{Str: "erdes", SubstringI: 58, Result: 1},
	{Str: "irdes", SubstringI: 58, Result: 1},
	{Str: "ares", SubstringI: 58, Result: 1},
	{Str: "eres", SubstringI: 58, Result: 1},
	{Str: "ires", SubstringI: 58, Result: 1},
	{Str: "asses", SubstringI: 58, Result: 1},
	{Str: "esses", SubstringI: 58, Result: 1},
	{Str: "isses", SubstringI: 58, Result: 1},
	{Str: "astes", SubstringI: 58, Result: 1},
	{Str: "estes", SubstringI: 58, Result: 1},
	{Str: "istes", SubstringI: 58, Result: 1},
	{Str: "is", SubstringI: -1, Result: 1},
	{Str: "ais", SubstringI: 71, Result: 1},
	{Str: "eis", SubstringI: 71, Result: 1},
	{Str: "areis", SubstringI: 73, Result: 1},
	{Str: "ereis", SubstringI: 73, Result: 1},
	{Str: "ireis", SubstringI: 73, Result: 1},
	{Str: "áreis", SubstringI: 73, Result: 1},
	{Str: "éreis", SubstringI: 73, Result: 1},
	{Str: "íreis", SubstringI: 73, Result: 1},
	{Str: "ásseis", SubstringI: 73, Result: 1},
	{Str: "ésseis", SubstringI: 73, Result: 1},
	{Str: "ísseis", SubstringI: 73, Result: 1},
	{Str: "áveis", SubstringI: 73, Result: 1},
	{Str: "íeis", SubstringI: 73, Result: 1},
	{Str: "aríeis", SubstringI: 84, Result: 1},
	{Str: "eríeis", SubstringI: 84, Result: 1},
	{Str: "iríeis", SubstringI: 84, Result: 1},
	{Str: "ados", SubstringI: -1, Result: 1},
	{Str: "idos", SubstringI: -1, Result: 1},
	{Str: "amos", SubstringI: -1, Result: 1},
	{Str: "áramos", SubstringI: 90, Result: 1},
	{Str: "éramos", SubstringI: 90, Result: 1},
	{Str: "íramos", SubstringI: 90, Result: 1},
	{Str: "ávamos", SubstringI: 90, Result: 1},
	{Str: "íamos", SubstringI: 90, Result: 1},
	{Str: "aríamos", SubstringI: 95, Result: 1},
	{Str: "eríamos", SubstringI: 95, Result: 1},
	{Str: "iríamos", SubstringI: 95, Result: 1},
	{Str: "emos", SubstringI: -1, Result: 1},
	{Str: "aremos", SubstringI: 99, Result: 1},
	{Str: "eremos", SubstringI: 99, Result: 1},
	{Str: "iremos", SubstringI: 99, Result: 1},
	{Str: "ássemos", SubstringI: 99, Result: 1},
	{Str: "êssemos", SubstringI: 99, Result: 1},
	{Str: "íssemos", SubstringI: 99, Result: 1},
	{Str: "imos", SubstringI: -1, Result: 1},
	{Str: "armos", SubstringI: -1, Result: 1},
	{Str: "ermos", SubstringI: -1, Result: 1},
	{Str: "irmos", SubstringI: -1, Result: 1},
	{Str: "ámos", SubstringI: -1, Result: 1},
	{Str: "arás", SubstringI: -1, Result: 1},
	{Str: "erás", SubstringI: -1, Result: 1},
	{Str: "irás", SubstringI: -1, Result: 1},
	{Str: "eu", SubstringI: -1, Result: 1},
	{Str: "iu", SubstringI: -1, Result: 1},
	{Str: "ou", SubstringI: -1, Result: 1},
	{Str: "ará", SubstringI: -1, Result: 1},
	{Str: "erá", SubstringI: -1, Result: 1},
	{Str: "irá", SubstringI: -1, Result: 1},
}

var a7 = []snowball.Among[context]{
	{Str: "a", SubstringI: -1, Result: 1},
	{Str: "i", SubstringI: -1, Result: 1},
	{Str: "o", SubstringI: -1, Result: 1},
	{Str: "os", SubstringI: -1, Result: 1},
	{Str: "á", SubstringI: -1, Result: 1},
	{Str: "í", SubstringI: -1, Result: 1},
	{Str: "ó", SubstringI: -1, Result: 1},
}

var a8 = []snowball.Among[context]{
	{Str: "e", SubstringI: -1, Result: 1},
	{Str: "ç", SubstringI: -1, Result: 2},
	{Str: "é", SubstringI: -1, Result: 1},
	{Str: "ê", SubstringI: -1, Result: 1},
}

var gV = []byte{17, 65, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 19, 12, 2}

// prelude decomposes the nasalised vowels ã/õ into a~/o~ so the suffix tables
// can match the tilde as a separate character.
func prelude(env *snowball.Env, ctx *context) bool {
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
				if !env.SliceFrom("a~") {
					return false
				}
			case 2:
				if !env.SliceFrom("o~") {
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
				if !env.InGrouping(gV, 97, 250) {
					break lab2
				}
			lab3:
				for {
					v3 := env.Cursor
				lab4:
					for {
						if !env.OutGrouping(gV, 97, 250) {
							break lab4
						}
					golab5:
						for {
							for {
								if !env.InGrouping(gV, 97, 250) {
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
					if !env.InGrouping(gV, 97, 250) {
						break lab2
					}
				golab7:
					for {
						for {
							if !env.OutGrouping(gV, 97, 250) {
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
			if !env.OutGrouping(gV, 97, 250) {
				break lab0
			}
		lab9:
			for {
				v6 := env.Cursor
			lab10:
				for {
					if !env.OutGrouping(gV, 97, 250) {
						break lab10
					}
				golab11:
					for {
						for {
							if !env.InGrouping(gV, 97, 250) {
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
				if !env.InGrouping(gV, 97, 250) {
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
				if !env.InGrouping(gV, 97, 250) {
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
				if !env.OutGrouping(gV, 97, 250) {
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
				if !env.InGrouping(gV, 97, 250) {
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
				if !env.OutGrouping(gV, 97, 250) {
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

// postlude recomposes the a~/o~ sequences introduced by prelude back into ã/õ.
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
				if !env.SliceFrom("ã") {
					return false
				}
			case 2:
				if !env.SliceFrom("õ") {
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

// standardSuffix rewrites or deletes the derivational suffixes in a5, with the
// post-deletion fix-ups in a2/a3/a4.
func standardSuffix(env *snowball.Env, ctx *context) bool {
	var amongVar int32
	env.Ket = env.Cursor
	amongVar = snowball.FindAmongB(env, a5, ctx)
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
		if !env.SliceFrom("log") {
			return false
		}
	case 3:
		if !r2(env, ctx) {
			return false
		}
		if !env.SliceFrom("u") {
			return false
		}
	case 4:
		if !r2(env, ctx) {
			return false
		}
		if !env.SliceFrom("ente") {
			return false
		}
	case 5:
		if !r1(env, ctx) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
		v1 := env.Limit - env.Cursor
	lab0:
		for {
			env.Ket = env.Cursor
			amongVar = snowball.FindAmongB(env, a2, ctx)
			if amongVar == 0 {
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
			switch amongVar {
			case 1:
				env.Ket = env.Cursor
				if !env.EqSB("at") {
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
			}
			break lab0
		}
	case 6:
		if !r2(env, ctx) {
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
			switch amongVar {
			case 1:
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
			break lab3
		}
	case 9:
		if !rV(env, ctx) {
			return false
		}
		if !env.EqSB("e") {
			return false
		}
		if !env.SliceFrom("ir") {
			return false
		}
	}
	return true
}

// verbSuffix deletes the verb endings in a6 within region RV.
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
	amongVar := snowball.FindAmongB(env, a6, ctx)
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

// residualSuffix deletes a residual vowel suffix in a7 within region RV.
func residualSuffix(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a7, ctx)
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
	}
	return true
}

// residualForm cleans up the final e/é/ê (deleting it, and a preceding gu/ci)
// and turns a final ç into c.
func residualForm(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a8, ctx)
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
		env.Ket = env.Cursor
	lab0:
		for {
			v1 := env.Limit - env.Cursor
		lab1:
			for {
				if !env.EqSB("u") {
					break lab1
				}
				env.Bra = env.Cursor
				v2 := env.Limit - env.Cursor
				if !env.EqSB("g") {
					break lab1
				}
				env.Cursor = env.Limit - v2
				break lab0
			}
			env.Cursor = env.Limit - v1
			if !env.EqSB("i") {
				return false
			}
			env.Bra = env.Cursor
			v3 := env.Limit - env.Cursor
			if !env.EqSB("c") {
				return false
			}
			env.Cursor = env.Limit - v3
			break lab0
		}
		if !rV(env, ctx) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	case 2:
		if !env.SliceFrom("c") {
			return false
		}
	}
	return true
}

// Stem runs the Snowball portuguese algorithm over env, mirroring the generated
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
					if !verbSuffix(env, ctx) {
						break lab4
					}
					break lab5
				}
				env.Cursor = env.Limit - v5
				v7 := env.Limit - env.Cursor
			lab7:
				for {
					env.Ket = env.Cursor
					if !env.EqSB("i") {
						break lab7
					}
					env.Bra = env.Cursor
					v8 := env.Limit - env.Cursor
					if !env.EqSB("c") {
						break lab7
					}
					env.Cursor = env.Limit - v8
					if !rV(env, ctx) {
						break lab7
					}
					if !env.SliceDel() {
						return false
					}
					break lab7
				}
				env.Cursor = env.Limit - v7
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
lab8:
	for {
		if !residualForm(env, ctx) {
			break lab8
		}
		break lab8
	}
	env.Cursor = env.Limit - v9
	env.Cursor = env.LimitBackward
	v10 := env.Cursor
lab9:
	for {
		if !postlude(env, ctx) {
			break lab9
		}
		break lab9
	}
	env.Cursor = v10
	return true
}
