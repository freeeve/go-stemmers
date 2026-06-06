// Package norwegian is a byte-faithful Go port of rust-stemmers' generated
// Snowball "norwegian" stemmer. It produces output identical to
// rust-stemmers 1.2.0's Norwegian algorithm.
//
// The control flow mirrors the generated Rust: each Snowball labelled block
// (`'lab: loop { … break 'lab }`) becomes a Go labelled `for { … break LABEL }`,
// and the result-code dispatch (`else if among_var == N`) becomes a `switch`.
package norwegian

import "github.com/freeeve/go-stemmers/internal/snowball"

// context holds the per-run state the generated algorithm threads through its
// routines: the iX hop mark and the R1 (iP1) region boundary.
type context struct {
	iX  int
	iP1 int
}

var a0 = []snowball.Among[context]{
	{Str: "a", SubstringI: -1, Result: 1},
	{Str: "e", SubstringI: -1, Result: 1},
	{Str: "ede", SubstringI: 1, Result: 1},
	{Str: "ande", SubstringI: 1, Result: 1},
	{Str: "ende", SubstringI: 1, Result: 1},
	{Str: "ane", SubstringI: 1, Result: 1},
	{Str: "ene", SubstringI: 1, Result: 1},
	{Str: "hetene", SubstringI: 6, Result: 1},
	{Str: "erte", SubstringI: 1, Result: 3},
	{Str: "en", SubstringI: -1, Result: 1},
	{Str: "heten", SubstringI: 9, Result: 1},
	{Str: "ar", SubstringI: -1, Result: 1},
	{Str: "er", SubstringI: -1, Result: 1},
	{Str: "heter", SubstringI: 12, Result: 1},
	{Str: "s", SubstringI: -1, Result: 2},
	{Str: "as", SubstringI: 14, Result: 1},
	{Str: "es", SubstringI: 14, Result: 1},
	{Str: "edes", SubstringI: 16, Result: 1},
	{Str: "endes", SubstringI: 16, Result: 1},
	{Str: "enes", SubstringI: 16, Result: 1},
	{Str: "hetenes", SubstringI: 19, Result: 1},
	{Str: "ens", SubstringI: 14, Result: 1},
	{Str: "hetens", SubstringI: 21, Result: 1},
	{Str: "ers", SubstringI: 14, Result: 1},
	{Str: "ets", SubstringI: 14, Result: 1},
	{Str: "et", SubstringI: -1, Result: 1},
	{Str: "het", SubstringI: 25, Result: 1},
	{Str: "ert", SubstringI: -1, Result: 3},
	{Str: "ast", SubstringI: -1, Result: 1},
}

var a1 = []snowball.Among[context]{
	{Str: "dt", SubstringI: -1, Result: -1},
	{Str: "vt", SubstringI: -1, Result: -1},
}

var a2 = []snowball.Among[context]{
	{Str: "leg", SubstringI: -1, Result: 1},
	{Str: "eleg", SubstringI: 0, Result: 1},
	{Str: "ig", SubstringI: -1, Result: 1},
	{Str: "eig", SubstringI: 2, Result: 1},
	{Str: "lig", SubstringI: 2, Result: 1},
	{Str: "elig", SubstringI: 4, Result: 1},
	{Str: "els", SubstringI: -1, Result: 1},
	{Str: "lov", SubstringI: -1, Result: 1},
	{Str: "elov", SubstringI: 7, Result: 1},
	{Str: "slov", SubstringI: 7, Result: 1},
	{Str: "hetslov", SubstringI: 9, Result: 1},
}

var gV = []byte{17, 65, 16, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 48, 0, 128}

var gSending = []byte{119, 125, 149, 1}

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

// mainSuffix strips the inflectional endings in a0 within R1, with the
// s-ending guard for the -s case and the -ert→-er rewrite.
func mainSuffix(env *snowball.Env, ctx *context) bool {
	var amongVar int32
	if env.Cursor < ctx.iP1 {
		return false
	}
	v2 := env.LimitBackward
	env.LimitBackward = ctx.iP1
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
	lab0:
		for {
			v3 := env.Limit - env.Cursor
		lab1:
			for {
				if !env.InGroupingB(gSending, 98, 122) {
					break lab1
				}
				break lab0
			}
			env.Cursor = env.Limit - v3
			if !env.EqSB("k") {
				return false
			}
			if !env.OutGroupingB(gV, 97, 248) {
				return false
			}
			break lab0
		}
		if !env.SliceDel() {
			return false
		}
	case 3:
		if !env.SliceFrom("er") {
			return false
		}
	}
	return true
}

// consonantPair deletes the trailing consonant of a -dt/-vt pair within R1.
func consonantPair(env *snowball.Env, ctx *context) bool {
	v1 := env.Limit - env.Cursor
	if env.Cursor < ctx.iP1 {
		return false
	}
	v3 := env.LimitBackward
	env.LimitBackward = ctx.iP1
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

// otherSuffix deletes the derivational suffixes in a2 within R1.
func otherSuffix(env *snowball.Env, ctx *context) bool {
	if env.Cursor < ctx.iP1 {
		return false
	}

	v2 := env.LimitBackward
	env.LimitBackward = ctx.iP1
	env.Ket = env.Cursor

	if snowball.FindAmongB(env, a2, ctx) == 0 {
		env.LimitBackward = v2
		return false
	}
	env.Bra = env.Cursor
	env.LimitBackward = v2

	if !env.SliceDel() {
		return false
	}

	return true
}

// Stem runs the Snowball norwegian algorithm over env, mirroring the generated
// `stem` entry point. It always returns true; the result is the mutated env.
func Stem(env *snowball.Env) bool {
	ctx := &context{}

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
	env.Cursor = env.LimitBackward

	return true
}
