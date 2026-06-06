// Package swedish is a byte-faithful Go port of rust-stemmers' generated
// Snowball "swedish" stemmer. It produces output identical to
// rust-stemmers 1.2.0's Swedish algorithm.
//
// The control flow mirrors the generated Rust: each Snowball labelled block
// (`'lab: loop { … break 'lab }`) becomes a Go labelled `for { … break LABEL }`,
// and the result-code dispatch (`else if among_var == N`) becomes a `switch`.
package swedish

import "github.com/freeeve/go-stemmers/internal/snowball"

// context holds the per-run state the generated algorithm threads through its
// routines: the iX hop mark and the R1 (iP1) region boundary.
type context struct {
	iX  int
	iP1 int
}

var a0 = []snowball.Among[context]{
	{Str: "a", SubstringI: -1, Result: 1},
	{Str: "arna", SubstringI: 0, Result: 1},
	{Str: "erna", SubstringI: 0, Result: 1},
	{Str: "heterna", SubstringI: 2, Result: 1},
	{Str: "orna", SubstringI: 0, Result: 1},
	{Str: "ad", SubstringI: -1, Result: 1},
	{Str: "e", SubstringI: -1, Result: 1},
	{Str: "ade", SubstringI: 6, Result: 1},
	{Str: "ande", SubstringI: 6, Result: 1},
	{Str: "arne", SubstringI: 6, Result: 1},
	{Str: "are", SubstringI: 6, Result: 1},
	{Str: "aste", SubstringI: 6, Result: 1},
	{Str: "en", SubstringI: -1, Result: 1},
	{Str: "anden", SubstringI: 12, Result: 1},
	{Str: "aren", SubstringI: 12, Result: 1},
	{Str: "heten", SubstringI: 12, Result: 1},
	{Str: "ern", SubstringI: -1, Result: 1},
	{Str: "ar", SubstringI: -1, Result: 1},
	{Str: "er", SubstringI: -1, Result: 1},
	{Str: "heter", SubstringI: 18, Result: 1},
	{Str: "or", SubstringI: -1, Result: 1},
	{Str: "s", SubstringI: -1, Result: 2},
	{Str: "as", SubstringI: 21, Result: 1},
	{Str: "arnas", SubstringI: 22, Result: 1},
	{Str: "ernas", SubstringI: 22, Result: 1},
	{Str: "ornas", SubstringI: 22, Result: 1},
	{Str: "es", SubstringI: 21, Result: 1},
	{Str: "ades", SubstringI: 26, Result: 1},
	{Str: "andes", SubstringI: 26, Result: 1},
	{Str: "ens", SubstringI: 21, Result: 1},
	{Str: "arens", SubstringI: 29, Result: 1},
	{Str: "hetens", SubstringI: 29, Result: 1},
	{Str: "erns", SubstringI: 21, Result: 1},
	{Str: "at", SubstringI: -1, Result: 1},
	{Str: "andet", SubstringI: -1, Result: 1},
	{Str: "het", SubstringI: -1, Result: 1},
	{Str: "ast", SubstringI: -1, Result: 1},
}

var a1 = []snowball.Among[context]{
	{Str: "dd", SubstringI: -1, Result: -1},
	{Str: "gd", SubstringI: -1, Result: -1},
	{Str: "nn", SubstringI: -1, Result: -1},
	{Str: "dt", SubstringI: -1, Result: -1},
	{Str: "gt", SubstringI: -1, Result: -1},
	{Str: "kt", SubstringI: -1, Result: -1},
	{Str: "tt", SubstringI: -1, Result: -1},
}

var a2 = []snowball.Among[context]{
	{Str: "ig", SubstringI: -1, Result: 1},
	{Str: "lig", SubstringI: 0, Result: 1},
	{Str: "els", SubstringI: -1, Result: 1},
	{Str: "fullt", SubstringI: -1, Result: 3},
	{Str: "löst", SubstringI: -1, Result: 2},
}

var gV = []byte{17, 65, 16, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 24, 0, 32}

var gSEnding = []byte{119, 127, 149}

// markRegions sets the R1 (iP1) region boundary, clamped to be no earlier than
// the iX hop mark three characters in.
func markRegions(env *snowball.Env, ctx *context) bool {
	ctx.iP1 = env.Limit
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
		v2 := env.Cursor
	lab1:
		for {
			if !env.InGrouping(gV, 97, 246) {
				break lab1
			}
			env.Cursor = v2
			break golab0
		}
		env.Cursor = v2
		if env.Cursor >= env.Limit {
			return false
		}
		env.NextChar()
	}
golab2:
	for {
	lab3:
		for {
			if !env.OutGrouping(gV, 97, 246) {
				break lab3
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
	return true
}

// mainSuffix strips the inflectional endings in a0 within R1, with the
// s-ending guard for the -s case.
func mainSuffix(env *snowball.Env, ctx *context) bool {
	var amongVar int32
	v1 := env.Limit - env.Cursor
	if env.Cursor < ctx.iP1 {
		return false
	}
	env.Cursor = ctx.iP1
	v2 := env.LimitBackward
	env.LimitBackward = env.Cursor
	env.Cursor = env.Limit - v1
	env.Ket = env.Cursor
	amongVar = snowball.FindAmongB(env, a0, ctx)
	if amongVar == 0 {
		env.LimitBackward = v2
		return false
	}
	env.Bra = env.Cursor
	env.LimitBackward = v2
	if amongVar == 0 {
		return false
	}
	switch amongVar {
	case 1:
		if !env.SliceDel() {
			return false
		}
	case 2:
		if !env.InGroupingB(gSEnding, 98, 121) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	}
	return true
}

// consonantPair deletes the trailing consonant of a stop-consonant pair in a1
// within R1.
func consonantPair(env *snowball.Env, ctx *context) bool {
	v1 := env.Limit - env.Cursor
	if env.Cursor < ctx.iP1 {
		return false
	}
	env.Cursor = ctx.iP1
	v2 := env.LimitBackward
	env.LimitBackward = env.Cursor
	env.Cursor = env.Limit - v1
	v3 := env.Limit - env.Cursor
	if snowball.FindAmongB(env, a1, ctx) == 0 {
		env.LimitBackward = v2
		return false
	}
	env.Cursor = env.Limit - v3
	env.Ket = env.Cursor
	if env.Cursor <= env.LimitBackward {
		env.LimitBackward = v2
		return false
	}
	env.PreviousChar()
	env.Bra = env.Cursor
	if !env.SliceDel() {
		return false
	}
	env.LimitBackward = v2
	return true
}

// otherSuffix rewrites or deletes the derivational suffixes in a2 within R1.
func otherSuffix(env *snowball.Env, ctx *context) bool {
	var amongVar int32
	v1 := env.Limit - env.Cursor
	if env.Cursor < ctx.iP1 {
		return false
	}
	env.Cursor = ctx.iP1
	v2 := env.LimitBackward
	env.LimitBackward = env.Cursor
	env.Cursor = env.Limit - v1
	env.Ket = env.Cursor
	amongVar = snowball.FindAmongB(env, a2, ctx)
	if amongVar == 0 {
		env.LimitBackward = v2
		return false
	}
	env.Bra = env.Cursor
	if amongVar == 0 {
		env.LimitBackward = v2
		return false
	}
	switch amongVar {
	case 1:
		if !env.SliceDel() {
			return false
		}
	case 2:
		if !env.SliceFrom("lös") {
			return false
		}
	case 3:
		if !env.SliceFrom("full") {
			return false
		}
	}
	env.LimitBackward = v2
	return true
}

// Stem runs the Snowball swedish algorithm over env, mirroring the generated
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
lab1:
	for {
		if !mainSuffix(env, ctx) {
			break lab1
		}
		break lab1
	}
	env.Cursor = env.Limit - v2

	v3 := env.Limit - env.Cursor
lab2:
	for {
		if !consonantPair(env, ctx) {
			break lab2
		}
		break lab2
	}
	env.Cursor = env.Limit - v3

	v4 := env.Limit - env.Cursor
lab3:
	for {
		if !otherSuffix(env, ctx) {
			break lab3
		}
		break lab3
	}
	env.Cursor = env.Limit - v4
	env.Cursor = env.LimitBackward

	return true
}
