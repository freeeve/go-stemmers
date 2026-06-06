// Package danish is a byte-faithful Go port of rust-stemmers' generated
// Snowball "danish" stemmer. It produces output identical to rust-stemmers
// 1.2.0's Danish algorithm.
//
// The control flow mirrors the generated Rust: each Snowball labelled block
// (`'lab: loop { … break 'lab }`) becomes a Go labelled `for { … break LABEL }`,
// and the result-code dispatch (`else if among_var == N`) becomes a `switch`.
package danish

import "github.com/freeeve/go-stemmers/internal/snowball"

// context holds the per-run state the generated algorithm threads through its
// routines: the iX hop mark, the R1 (iP1) region boundary, and the captured
// final consonant (sCh) used by undouble.
type context struct {
	iX  int
	iP1 int
	sCh string
}

var a0 = []snowball.Among[context]{
	{Str: "hed", SubstringI: -1, Result: 1},
	{Str: "ethed", SubstringI: 0, Result: 1},
	{Str: "ered", SubstringI: -1, Result: 1},
	{Str: "e", SubstringI: -1, Result: 1},
	{Str: "erede", SubstringI: 3, Result: 1},
	{Str: "ende", SubstringI: 3, Result: 1},
	{Str: "erende", SubstringI: 5, Result: 1},
	{Str: "ene", SubstringI: 3, Result: 1},
	{Str: "erne", SubstringI: 3, Result: 1},
	{Str: "ere", SubstringI: 3, Result: 1},
	{Str: "en", SubstringI: -1, Result: 1},
	{Str: "heden", SubstringI: 10, Result: 1},
	{Str: "eren", SubstringI: 10, Result: 1},
	{Str: "er", SubstringI: -1, Result: 1},
	{Str: "heder", SubstringI: 13, Result: 1},
	{Str: "erer", SubstringI: 13, Result: 1},
	{Str: "s", SubstringI: -1, Result: 2},
	{Str: "heds", SubstringI: 16, Result: 1},
	{Str: "es", SubstringI: 16, Result: 1},
	{Str: "endes", SubstringI: 18, Result: 1},
	{Str: "erendes", SubstringI: 19, Result: 1},
	{Str: "enes", SubstringI: 18, Result: 1},
	{Str: "ernes", SubstringI: 18, Result: 1},
	{Str: "eres", SubstringI: 18, Result: 1},
	{Str: "ens", SubstringI: 16, Result: 1},
	{Str: "hedens", SubstringI: 24, Result: 1},
	{Str: "erens", SubstringI: 24, Result: 1},
	{Str: "ers", SubstringI: 16, Result: 1},
	{Str: "ets", SubstringI: 16, Result: 1},
	{Str: "erets", SubstringI: 28, Result: 1},
	{Str: "et", SubstringI: -1, Result: 1},
	{Str: "eret", SubstringI: 30, Result: 1},
}

var a1 = []snowball.Among[context]{
	{Str: "gd", SubstringI: -1, Result: -1},
	{Str: "dt", SubstringI: -1, Result: -1},
	{Str: "gt", SubstringI: -1, Result: -1},
	{Str: "kt", SubstringI: -1, Result: -1},
}

var a2 = []snowball.Among[context]{
	{Str: "ig", SubstringI: -1, Result: 1},
	{Str: "lig", SubstringI: 0, Result: 1},
	{Str: "elig", SubstringI: 1, Result: 1},
	{Str: "els", SubstringI: -1, Result: 1},
	{Str: "løst", SubstringI: -1, Result: 2},
}

var gV = []byte{17, 65, 16, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 48, 0, 128}

var gSEnding = []byte{239, 254, 42, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 16}

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
			if !env.InGrouping(gV, 97, 248) {
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
			if !env.OutGrouping(gV, 97, 248) {
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

// mainSuffix strips the inflectional endings in a0 within R1, deleting the
// -s ending only when it follows a valid s-ending consonant.
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
	switch amongVar {
	case 1:
		if !env.SliceDel() {
			return false
		}
	case 2:
		if !env.InGroupingB(gSEnding, 97, 229) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	}
	return true
}

// consonantPair deletes the trailing consonant of a -gd/-dt/-gt/-kt pair found
// within R1.
func consonantPair(env *snowball.Env, ctx *context) bool {
	v1 := env.Limit - env.Cursor
	v2 := env.Limit - env.Cursor
	if env.Cursor < ctx.iP1 {
		return false
	}
	env.Cursor = ctx.iP1
	v3 := env.LimitBackward
	env.LimitBackward = env.Cursor
	env.Cursor = env.Limit - v2
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a1, ctx) == 0 {
		env.LimitBackward = v3
		return false
	}
	env.Bra = env.Cursor
	env.LimitBackward = v3
	env.Cursor = env.Limit - v1
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

// otherSuffix strips the standalone -st(ig) sequence and the derivational
// suffixes in a2 within R1, re-running consonantPair after the -ig family.
func otherSuffix(env *snowball.Env, ctx *context) bool {
	var amongVar int32
	v1 := env.Limit - env.Cursor
lab0:
	for {
		env.Ket = env.Cursor
		if !env.EqSB("st") {
			break lab0
		}
		env.Bra = env.Cursor
		if !env.EqSB("ig") {
			break lab0
		}
		if !env.SliceDel() {
			return false
		}
		break lab0
	}
	env.Cursor = env.Limit - v1
	v2 := env.Limit - env.Cursor
	if env.Cursor < ctx.iP1 {
		return false
	}
	env.Cursor = ctx.iP1
	v3 := env.LimitBackward
	env.LimitBackward = env.Cursor
	env.Cursor = env.Limit - v2
	env.Ket = env.Cursor
	amongVar = snowball.FindAmongB(env, a2, ctx)
	if amongVar == 0 {
		env.LimitBackward = v3
		return false
	}
	env.Bra = env.Cursor
	env.LimitBackward = v3
	switch amongVar {
	case 1:
		if !env.SliceDel() {
			return false
		}
		v4 := env.Limit - env.Cursor
	lab1:
		for {
			if !consonantPair(env, ctx) {
				break lab1
			}
			break lab1
		}
		env.Cursor = env.Limit - v4
	case 2:
		if !env.SliceFrom("løs") {
			return false
		}
	}
	return true
}

// undouble removes a doubled final consonant within R1, capturing the final
// consonant (sCh) and deleting it when the preceding character matches.
func undouble(env *snowball.Env, ctx *context) bool {
	v1 := env.Limit - env.Cursor
	if env.Cursor < ctx.iP1 {
		return false
	}
	env.Cursor = ctx.iP1
	v2 := env.LimitBackward
	env.LimitBackward = env.Cursor
	env.Cursor = env.Limit - v1
	env.Ket = env.Cursor
	if !env.OutGroupingB(gV, 97, 248) {
		env.LimitBackward = v2
		return false
	}
	env.Bra = env.Cursor
	ctx.sCh = env.SliceTo()
	if ctx.sCh == "" {
		return false
	}
	env.LimitBackward = v2
	if !env.EqSB(ctx.sCh) {
		return false
	}
	if !env.SliceDel() {
		return false
	}
	return true
}

// Stem runs the Snowball danish algorithm over env, mirroring the generated
// `stem` entry point. It always returns true; the result is the mutated env.
func Stem(env *snowball.Env) bool {
	ctx, _ := env.Scratch.(*context)
	if ctx == nil {
		ctx = &context{}
		env.Scratch = ctx
	}
	*ctx = context{}

	v1 := env.Cursor
	markRegions(env, ctx)
	env.Cursor = v1
	env.LimitBackward = env.Cursor
	env.Cursor = env.Limit

	v2 := env.Limit - env.Cursor
	mainSuffix(env, ctx)
	env.Cursor = env.Limit - v2

	v3 := env.Limit - env.Cursor
	consonantPair(env, ctx)
	env.Cursor = env.Limit - v3

	v4 := env.Limit - env.Cursor
	otherSuffix(env, ctx)
	env.Cursor = env.Limit - v4

	v5 := env.Limit - env.Cursor
	undouble(env, ctx)
	env.Cursor = env.Limit - v5
	env.Cursor = env.LimitBackward

	return true
}
