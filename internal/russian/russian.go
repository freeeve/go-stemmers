// Package russian is a byte-faithful Go port of rust-stemmers' generated
// Snowball "russian" stemmer. It produces output identical to
// rust-stemmers 1.2.0's Russian algorithm.
//
// The control flow mirrors the generated Rust: each Snowball labelled block
// (`'lab: loop { … break 'lab }`) becomes a Go labelled `for { … break LABEL }`,
// and the result-code dispatch (`else if among_var == N`) becomes a `switch`.
// The among-table strings are Cyrillic (multi-byte UTF-8) and are carried
// verbatim; the grouping bytes and the 1072..1103 codepoint range match the
// generated source exactly.
package russian

import "github.com/freeeve/go-stemmers/internal/snowball"

// context holds the per-run state the generated algorithm threads through its
// routines: the R2 (iP2) and RV (iPV) region boundaries.
type context struct {
	iP2 int
	iPV int
}

var a0 = []snowball.Among[context]{
	{Str: "вшись", SubstringI: -1, Result: 1},
	{Str: "ывшись", SubstringI: 0, Result: 2},
	{Str: "ившись", SubstringI: 0, Result: 2},
	{Str: "в", SubstringI: -1, Result: 1},
	{Str: "ыв", SubstringI: 3, Result: 2},
	{Str: "ив", SubstringI: 3, Result: 2},
	{Str: "вши", SubstringI: -1, Result: 1},
	{Str: "ывши", SubstringI: 6, Result: 2},
	{Str: "ивши", SubstringI: 6, Result: 2},
}

var a1 = []snowball.Among[context]{
	{Str: "ему", SubstringI: -1, Result: 1},
	{Str: "ому", SubstringI: -1, Result: 1},
	{Str: "ых", SubstringI: -1, Result: 1},
	{Str: "их", SubstringI: -1, Result: 1},
	{Str: "ую", SubstringI: -1, Result: 1},
	{Str: "юю", SubstringI: -1, Result: 1},
	{Str: "ею", SubstringI: -1, Result: 1},
	{Str: "ою", SubstringI: -1, Result: 1},
	{Str: "яя", SubstringI: -1, Result: 1},
	{Str: "ая", SubstringI: -1, Result: 1},
	{Str: "ые", SubstringI: -1, Result: 1},
	{Str: "ее", SubstringI: -1, Result: 1},
	{Str: "ие", SubstringI: -1, Result: 1},
	{Str: "ое", SubstringI: -1, Result: 1},
	{Str: "ыми", SubstringI: -1, Result: 1},
	{Str: "ими", SubstringI: -1, Result: 1},
	{Str: "ый", SubstringI: -1, Result: 1},
	{Str: "ей", SubstringI: -1, Result: 1},
	{Str: "ий", SubstringI: -1, Result: 1},
	{Str: "ой", SubstringI: -1, Result: 1},
	{Str: "ым", SubstringI: -1, Result: 1},
	{Str: "ем", SubstringI: -1, Result: 1},
	{Str: "им", SubstringI: -1, Result: 1},
	{Str: "ом", SubstringI: -1, Result: 1},
	{Str: "его", SubstringI: -1, Result: 1},
	{Str: "ого", SubstringI: -1, Result: 1},
}

var a2 = []snowball.Among[context]{
	{Str: "вш", SubstringI: -1, Result: 1},
	{Str: "ывш", SubstringI: 0, Result: 2},
	{Str: "ивш", SubstringI: 0, Result: 2},
	{Str: "щ", SubstringI: -1, Result: 1},
	{Str: "ющ", SubstringI: 3, Result: 1},
	{Str: "ующ", SubstringI: 4, Result: 2},
	{Str: "ем", SubstringI: -1, Result: 1},
	{Str: "нн", SubstringI: -1, Result: 1},
}

var a3 = []snowball.Among[context]{
	{Str: "сь", SubstringI: -1, Result: 1},
	{Str: "ся", SubstringI: -1, Result: 1},
}

var a4 = []snowball.Among[context]{
	{Str: "ыт", SubstringI: -1, Result: 2},
	{Str: "ют", SubstringI: -1, Result: 1},
	{Str: "уют", SubstringI: 1, Result: 2},
	{Str: "ят", SubstringI: -1, Result: 2},
	{Str: "ет", SubstringI: -1, Result: 1},
	{Str: "ует", SubstringI: 4, Result: 2},
	{Str: "ит", SubstringI: -1, Result: 2},
	{Str: "ны", SubstringI: -1, Result: 1},
	{Str: "ены", SubstringI: 7, Result: 2},
	{Str: "ть", SubstringI: -1, Result: 1},
	{Str: "ыть", SubstringI: 9, Result: 2},
	{Str: "ить", SubstringI: 9, Result: 2},
	{Str: "ешь", SubstringI: -1, Result: 1},
	{Str: "ишь", SubstringI: -1, Result: 2},
	{Str: "ю", SubstringI: -1, Result: 2},
	{Str: "ую", SubstringI: 14, Result: 2},
	{Str: "ла", SubstringI: -1, Result: 1},
	{Str: "ыла", SubstringI: 16, Result: 2},
	{Str: "ила", SubstringI: 16, Result: 2},
	{Str: "на", SubstringI: -1, Result: 1},
	{Str: "ена", SubstringI: 19, Result: 2},
	{Str: "ете", SubstringI: -1, Result: 1},
	{Str: "ите", SubstringI: -1, Result: 2},
	{Str: "йте", SubstringI: -1, Result: 1},
	{Str: "уйте", SubstringI: 23, Result: 2},
	{Str: "ейте", SubstringI: 23, Result: 2},
	{Str: "ли", SubstringI: -1, Result: 1},
	{Str: "ыли", SubstringI: 26, Result: 2},
	{Str: "или", SubstringI: 26, Result: 2},
	{Str: "й", SubstringI: -1, Result: 1},
	{Str: "уй", SubstringI: 29, Result: 2},
	{Str: "ей", SubstringI: 29, Result: 2},
	{Str: "л", SubstringI: -1, Result: 1},
	{Str: "ыл", SubstringI: 32, Result: 2},
	{Str: "ил", SubstringI: 32, Result: 2},
	{Str: "ым", SubstringI: -1, Result: 2},
	{Str: "ем", SubstringI: -1, Result: 1},
	{Str: "им", SubstringI: -1, Result: 2},
	{Str: "н", SubstringI: -1, Result: 1},
	{Str: "ен", SubstringI: 38, Result: 2},
	{Str: "ло", SubstringI: -1, Result: 1},
	{Str: "ыло", SubstringI: 40, Result: 2},
	{Str: "ило", SubstringI: 40, Result: 2},
	{Str: "но", SubstringI: -1, Result: 1},
	{Str: "ено", SubstringI: 43, Result: 2},
	{Str: "нно", SubstringI: 43, Result: 1},
}

var a5 = []snowball.Among[context]{
	{Str: "у", SubstringI: -1, Result: 1},
	{Str: "ях", SubstringI: -1, Result: 1},
	{Str: "иях", SubstringI: 1, Result: 1},
	{Str: "ах", SubstringI: -1, Result: 1},
	{Str: "ы", SubstringI: -1, Result: 1},
	{Str: "ь", SubstringI: -1, Result: 1},
	{Str: "ю", SubstringI: -1, Result: 1},
	{Str: "ью", SubstringI: 6, Result: 1},
	{Str: "ию", SubstringI: 6, Result: 1},
	{Str: "я", SubstringI: -1, Result: 1},
	{Str: "ья", SubstringI: 9, Result: 1},
	{Str: "ия", SubstringI: 9, Result: 1},
	{Str: "а", SubstringI: -1, Result: 1},
	{Str: "ев", SubstringI: -1, Result: 1},
	{Str: "ов", SubstringI: -1, Result: 1},
	{Str: "е", SubstringI: -1, Result: 1},
	{Str: "ье", SubstringI: 15, Result: 1},
	{Str: "ие", SubstringI: 15, Result: 1},
	{Str: "и", SubstringI: -1, Result: 1},
	{Str: "еи", SubstringI: 18, Result: 1},
	{Str: "ии", SubstringI: 18, Result: 1},
	{Str: "ями", SubstringI: 18, Result: 1},
	{Str: "иями", SubstringI: 21, Result: 1},
	{Str: "ами", SubstringI: 18, Result: 1},
	{Str: "й", SubstringI: -1, Result: 1},
	{Str: "ей", SubstringI: 24, Result: 1},
	{Str: "ией", SubstringI: 25, Result: 1},
	{Str: "ий", SubstringI: 24, Result: 1},
	{Str: "ой", SubstringI: 24, Result: 1},
	{Str: "ям", SubstringI: -1, Result: 1},
	{Str: "иям", SubstringI: 29, Result: 1},
	{Str: "ам", SubstringI: -1, Result: 1},
	{Str: "ем", SubstringI: -1, Result: 1},
	{Str: "ием", SubstringI: 32, Result: 1},
	{Str: "ом", SubstringI: -1, Result: 1},
	{Str: "о", SubstringI: -1, Result: 1},
}

var a6 = []snowball.Among[context]{
	{Str: "ост", SubstringI: -1, Result: 1},
	{Str: "ость", SubstringI: -1, Result: 1},
}

var a7 = []snowball.Among[context]{
	{Str: "ейш", SubstringI: -1, Result: 1},
	{Str: "ь", SubstringI: -1, Result: 3},
	{Str: "ейше", SubstringI: -1, Result: 1},
	{Str: "н", SubstringI: -1, Result: 2},
}

var gV = []byte{33, 65, 8, 232}

// markRegions sets the RV (iPV) and R2 (iP2) region boundaries used to gate the
// suffix-stripping routines.
func markRegions(env *snowball.Env, ctx *context) bool {
	ctx.iPV = env.Limit
	ctx.iP2 = env.Limit
	v1 := env.Cursor
lab0:
	for {
	golab1:
		for {
			for {
				if !env.InGrouping(gV, 1072, 1103) {
					break
				}
				break golab1
			}
			if env.Cursor >= env.Limit {
				break lab0
			}
			env.NextChar()
		}
		ctx.iPV = env.Cursor
	golab3:
		for {
			for {
				if !env.OutGrouping(gV, 1072, 1103) {
					break
				}
				break golab3
			}
			if env.Cursor >= env.Limit {
				break lab0
			}
			env.NextChar()
		}
	golab5:
		for {
			for {
				if !env.InGrouping(gV, 1072, 1103) {
					break
				}
				break golab5
			}
			if env.Cursor >= env.Limit {
				break lab0
			}
			env.NextChar()
		}
	golab7:
		for {
			for {
				if !env.OutGrouping(gV, 1072, 1103) {
					break
				}
				break golab7
			}
			if env.Cursor >= env.Limit {
				break lab0
			}
			env.NextChar()
		}
		ctx.iP2 = env.Cursor
		break lab0
	}
	env.Cursor = v1
	return true
}

// r2 reports whether the cursor is within region R2.
func r2(env *snowball.Env, ctx *context) bool {
	return ctx.iP2 <= env.Cursor
}

// perfectiveGerund strips perfective-gerund endings (a0), with the а/я vowel
// guard required before the shorter group-1 endings.
func perfectiveGerund(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a0, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	switch amongVar {
	case 1:
	lab0:
		for {
			v1 := env.Limit - env.Cursor
		lab1:
			for {
				if !env.EqSB("а") {
					break lab1
				}
				break lab0
			}
			env.Cursor = env.Limit - v1
			if !env.EqSB("я") {
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
	return true
}

// adjective strips adjectival agreement endings (a1).
func adjective(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a1, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	switch amongVar {
	case 1:
		if !env.SliceDel() {
			return false
		}
	}
	return true
}

// adjectival strips an adjective ending and then any participle suffix (a2),
// with the а/я vowel guard before the shorter group-1 suffix.
func adjectival(env *snowball.Env, ctx *context) bool {
	if !adjective(env, ctx) {
		return false
	}
	v1 := env.Limit - env.Cursor
lab0:
	for {
		env.Ket = env.Cursor
		amongVar := snowball.FindAmongB(env, a2, ctx)
		if amongVar == 0 {
			env.Cursor = env.Limit - v1
			break lab0
		}
		env.Bra = env.Cursor
		switch amongVar {
		case 1:
		lab1:
			for {
				v2 := env.Limit - env.Cursor
			lab2:
				for {
					if !env.EqSB("а") {
						break lab2
					}
					break lab1
				}
				env.Cursor = env.Limit - v2
				if !env.EqSB("я") {
					env.Cursor = env.Limit - v1
					break lab0
				}
				break lab1
			}
			if !env.SliceDel() {
				return false
			}
		case 2:
			if !env.SliceDel() {
				return false
			}
		}
		break lab0
	}
	return true
}

// reflexive strips the reflexive endings ся/сь (a3).
func reflexive(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a3, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	switch amongVar {
	case 1:
		if !env.SliceDel() {
			return false
		}
	}
	return true
}

// verb strips verb conjugation endings (a4), with the а/я vowel guard before the
// shorter group-1 endings.
func verb(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a4, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	switch amongVar {
	case 1:
	lab0:
		for {
			v1 := env.Limit - env.Cursor
		lab1:
			for {
				if !env.EqSB("а") {
					break lab1
				}
				break lab0
			}
			env.Cursor = env.Limit - v1
			if !env.EqSB("я") {
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
	return true
}

// noun strips noun case endings (a5).
func noun(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a5, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	switch amongVar {
	case 1:
		if !env.SliceDel() {
			return false
		}
	}
	return true
}

// derivational removes the derivational suffix ост/ость (a6) when it lies within
// region R2.
func derivational(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a6, ctx)
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
	}
	return true
}

// tidyUp performs the final clean-up (a7): undoubling нн, dropping a superlative
// ейш(е), and removing a trailing soft sign ь.
func tidyUp(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a7, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	switch amongVar {
	case 1:
		if !env.SliceDel() {
			return false
		}
		env.Ket = env.Cursor
		if !env.EqSB("н") {
			return false
		}
		env.Bra = env.Cursor
		if !env.EqSB("н") {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	case 2:
		if !env.EqSB("н") {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	case 3:
		if !env.SliceDel() {
			return false
		}
	}
	return true
}

// Stem runs the Snowball russian algorithm over env, mirroring the generated
// `stem` entry point. It always returns true; the result is the mutated env.
func Stem(env *snowball.Env) bool {
	ctx, _ := env.Scratch.(*context)
	if ctx == nil {
		ctx = &context{}
		env.Scratch = ctx
	}
	*ctx = context{}
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
	if env.Cursor < ctx.iPV {
		return false
	}
	env.Cursor = ctx.iPV
	v3 := env.LimitBackward
	env.LimitBackward = env.Cursor
	env.Cursor = env.Limit - v2
	v4 := env.Limit - env.Cursor
lab1:
	for {
	lab2:
		for {
			v5 := env.Limit - env.Cursor
		lab3:
			for {
				if !perfectiveGerund(env, ctx) {
					break lab3
				}
				break lab2
			}
			env.Cursor = env.Limit - v5
			v6 := env.Limit - env.Cursor
		lab4:
			for {
				if !reflexive(env, ctx) {
					env.Cursor = env.Limit - v6
					break lab4
				}
				break lab4
			}
		lab5:
			for {
				v7 := env.Limit - env.Cursor
			lab6:
				for {
					if !adjectival(env, ctx) {
						break lab6
					}
					break lab5
				}
				env.Cursor = env.Limit - v7
			lab7:
				for {
					if !verb(env, ctx) {
						break lab7
					}
					break lab5
				}
				env.Cursor = env.Limit - v7
				if !noun(env, ctx) {
					break lab1
				}
				break lab5
			}
			break lab2
		}
		break lab1
	}
	env.Cursor = env.Limit - v4
	v8 := env.Limit - env.Cursor
lab8:
	for {
		env.Ket = env.Cursor
		if !env.EqSB("и") {
			env.Cursor = env.Limit - v8
			break lab8
		}
		env.Bra = env.Cursor
		if !env.SliceDel() {
			return false
		}
		break lab8
	}
	v9 := env.Limit - env.Cursor
lab9:
	for {
		if !derivational(env, ctx) {
			break lab9
		}
		break lab9
	}
	env.Cursor = env.Limit - v9
	v10 := env.Limit - env.Cursor
lab10:
	for {
		if !tidyUp(env, ctx) {
			break lab10
		}
		break lab10
	}
	env.Cursor = env.Limit - v10
	env.LimitBackward = v3
	env.Cursor = env.LimitBackward
	return true
}
