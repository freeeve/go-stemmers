// Package finnish is a byte-faithful Go port of rust-stemmers' generated
// Snowball "finnish" stemmer. It produces output identical to rust-stemmers
// 1.2.0's Finnish algorithm; the canonical Snowball finnish vocabulary is the
// conformance oracle.
//
// The control flow mirrors the generated Rust: each Snowball labelled block
// (`'lab: loop { … break 'lab }`) becomes a Go labelled `for { … break LABEL }`,
// and the result-code dispatch (`else if among_var == N`) becomes a `switch`.
// Finnish is the one algorithm whose among tables carry associated routines (the
// Among.Method field): the A_6 entries reference the long and vi helpers.
package finnish

import "github.com/freeeve/go-stemmers/internal/snowball"

// context holds the per-run state the generated algorithm threads through its
// routines: whether a case ending was removed, the captured slice used by tidy,
// and the R1/R2 region marks.
type context struct {
	bEndingRemoved bool
	sX             string
	iP2            int
	iP1            int
}

var a0 = []snowball.Among[context]{
	{Str: "pa", SubstringI: -1, Result: 1},
	{Str: "sti", SubstringI: -1, Result: 2},
	{Str: "kaan", SubstringI: -1, Result: 1},
	{Str: "han", SubstringI: -1, Result: 1},
	{Str: "kin", SubstringI: -1, Result: 1},
	{Str: "hän", SubstringI: -1, Result: 1},
	{Str: "kään", SubstringI: -1, Result: 1},
	{Str: "ko", SubstringI: -1, Result: 1},
	{Str: "pä", SubstringI: -1, Result: 1},
	{Str: "kö", SubstringI: -1, Result: 1},
}

var a1 = []snowball.Among[context]{
	{Str: "lla", SubstringI: -1, Result: -1},
	{Str: "na", SubstringI: -1, Result: -1},
	{Str: "ssa", SubstringI: -1, Result: -1},
	{Str: "ta", SubstringI: -1, Result: -1},
	{Str: "lta", SubstringI: 3, Result: -1},
	{Str: "sta", SubstringI: 3, Result: -1},
}

var a2 = []snowball.Among[context]{
	{Str: "llä", SubstringI: -1, Result: -1},
	{Str: "nä", SubstringI: -1, Result: -1},
	{Str: "ssä", SubstringI: -1, Result: -1},
	{Str: "tä", SubstringI: -1, Result: -1},
	{Str: "ltä", SubstringI: 3, Result: -1},
	{Str: "stä", SubstringI: 3, Result: -1},
}

var a3 = []snowball.Among[context]{
	{Str: "lle", SubstringI: -1, Result: -1},
	{Str: "ine", SubstringI: -1, Result: -1},
}

var a4 = []snowball.Among[context]{
	{Str: "nsa", SubstringI: -1, Result: 3},
	{Str: "mme", SubstringI: -1, Result: 3},
	{Str: "nne", SubstringI: -1, Result: 3},
	{Str: "ni", SubstringI: -1, Result: 2},
	{Str: "si", SubstringI: -1, Result: 1},
	{Str: "an", SubstringI: -1, Result: 4},
	{Str: "en", SubstringI: -1, Result: 6},
	{Str: "än", SubstringI: -1, Result: 5},
	{Str: "nsä", SubstringI: -1, Result: 3},
}

var a5 = []snowball.Among[context]{
	{Str: "aa", SubstringI: -1, Result: -1},
	{Str: "ee", SubstringI: -1, Result: -1},
	{Str: "ii", SubstringI: -1, Result: -1},
	{Str: "oo", SubstringI: -1, Result: -1},
	{Str: "uu", SubstringI: -1, Result: -1},
	{Str: "ää", SubstringI: -1, Result: -1},
	{Str: "öö", SubstringI: -1, Result: -1},
}

var a6 = []snowball.Among[context]{
	{Str: "a", SubstringI: -1, Result: 8},
	{Str: "lla", SubstringI: 0, Result: -1},
	{Str: "na", SubstringI: 0, Result: -1},
	{Str: "ssa", SubstringI: 0, Result: -1},
	{Str: "ta", SubstringI: 0, Result: -1},
	{Str: "lta", SubstringI: 4, Result: -1},
	{Str: "sta", SubstringI: 4, Result: -1},
	{Str: "tta", SubstringI: 4, Result: 9},
	{Str: "lle", SubstringI: -1, Result: -1},
	{Str: "ine", SubstringI: -1, Result: -1},
	{Str: "ksi", SubstringI: -1, Result: -1},
	{Str: "n", SubstringI: -1, Result: 7},
	{Str: "han", SubstringI: 11, Result: 1},
	{Str: "den", SubstringI: 11, Result: -1, Method: vi},
	{Str: "seen", SubstringI: 11, Result: -1, Method: long},
	{Str: "hen", SubstringI: 11, Result: 2},
	{Str: "tten", SubstringI: 11, Result: -1, Method: vi},
	{Str: "hin", SubstringI: 11, Result: 3},
	{Str: "siin", SubstringI: 11, Result: -1, Method: vi},
	{Str: "hon", SubstringI: 11, Result: 4},
	{Str: "hän", SubstringI: 11, Result: 5},
	{Str: "hön", SubstringI: 11, Result: 6},
	{Str: "ä", SubstringI: -1, Result: 8},
	{Str: "llä", SubstringI: 22, Result: -1},
	{Str: "nä", SubstringI: 22, Result: -1},
	{Str: "ssä", SubstringI: 22, Result: -1},
	{Str: "tä", SubstringI: 22, Result: -1},
	{Str: "ltä", SubstringI: 26, Result: -1},
	{Str: "stä", SubstringI: 26, Result: -1},
	{Str: "ttä", SubstringI: 26, Result: 9},
}

var a7 = []snowball.Among[context]{
	{Str: "eja", SubstringI: -1, Result: -1},
	{Str: "mma", SubstringI: -1, Result: 1},
	{Str: "imma", SubstringI: 1, Result: -1},
	{Str: "mpa", SubstringI: -1, Result: 1},
	{Str: "impa", SubstringI: 3, Result: -1},
	{Str: "mmi", SubstringI: -1, Result: 1},
	{Str: "immi", SubstringI: 5, Result: -1},
	{Str: "mpi", SubstringI: -1, Result: 1},
	{Str: "impi", SubstringI: 7, Result: -1},
	{Str: "ejä", SubstringI: -1, Result: -1},
	{Str: "mmä", SubstringI: -1, Result: 1},
	{Str: "immä", SubstringI: 10, Result: -1},
	{Str: "mpä", SubstringI: -1, Result: 1},
	{Str: "impä", SubstringI: 12, Result: -1},
}

var a8 = []snowball.Among[context]{
	{Str: "i", SubstringI: -1, Result: -1},
	{Str: "j", SubstringI: -1, Result: -1},
}

var a9 = []snowball.Among[context]{
	{Str: "mma", SubstringI: -1, Result: 1},
	{Str: "imma", SubstringI: 0, Result: -1},
}

var gAEI = []byte{17, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8}
var gV1 = []byte{17, 65, 16, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 0, 32}
var gV2 = []byte{17, 65, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 0, 32}
var gParticleEnd = []byte{17, 97, 24, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 0, 32}

// markRegions sets the R1 (iP1) and R2 (iP2) region boundaries from the first
// vowel/non-vowel transitions.
func markRegions(env *snowball.Env, ctx *context) bool {
	ctx.iP1 = env.Limit
	ctx.iP2 = env.Limit
golab0:
	for {
		v1 := env.Cursor
	lab1:
		for {
			if !env.InGrouping(gV1, 97, 246) {
				break lab1
			}
			env.Cursor = v1
			break golab0
		}
		env.Cursor = v1
		if env.Cursor >= env.Limit {
			return false
		}
		env.NextChar()
	}
golab2:
	for {
	lab3:
		for {
			if !env.OutGrouping(gV1, 97, 246) {
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
golab4:
	for {
		v3 := env.Cursor
	lab5:
		for {
			if !env.InGrouping(gV1, 97, 246) {
				break lab5
			}
			env.Cursor = v3
			break golab4
		}
		env.Cursor = v3
		if env.Cursor >= env.Limit {
			return false
		}
		env.NextChar()
	}
golab6:
	for {
	lab7:
		for {
			if !env.OutGrouping(gV1, 97, 246) {
				break lab7
			}
			break golab6
		}
		if env.Cursor >= env.Limit {
			return false
		}
		env.NextChar()
	}
	ctx.iP2 = env.Cursor
	return true
}

// r2 reports whether the cursor is within region R2.
func r2(env *snowball.Env, ctx *context) bool {
	return ctx.iP2 <= env.Cursor
}

// particleEtc removes clitic particles (kin, kaan, han …) within R1, gated by a
// closing-vowel or R2 check depending on the matched entry.
func particleEtc(env *snowball.Env, ctx *context) bool {
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
		if !env.InGroupingB(gParticleEnd, 97, 246) {
			return false
		}
	case 2:
		if !r2(env, ctx) {
			return false
		}
	}
	if !env.SliceDel() {
		return false
	}
	return true
}

// possessive removes possessive suffixes (-ni, -si, -nsa …) within R1, with the
// per-suffix fix-ups in a1/a2/a3.
func possessive(env *snowball.Env, ctx *context) bool {
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
	amongVar = snowball.FindAmongB(env, a4, ctx)
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
			if !env.EqSB("k") {
				break lab0
			}
			return false
		}
		env.Cursor = env.Limit - v3
		if !env.SliceDel() {
			return false
		}
	case 2:
		if !env.SliceDel() {
			return false
		}
		env.Ket = env.Cursor
		if !env.EqSB("kse") {
			return false
		}
		env.Bra = env.Cursor
		if !env.SliceFrom("ksi") {
			return false
		}
	case 3:
		if !env.SliceDel() {
			return false
		}
	case 4:
		if snowball.FindAmongB(env, a1, ctx) == 0 {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	case 5:
		if snowball.FindAmongB(env, a2, ctx) == 0 {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	case 6:
		if snowball.FindAmongB(env, a3, ctx) == 0 {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	}
	return true
}

// long reports whether a long vowel pair (aa, ee, … öö) precedes the cursor.
func long(env *snowball.Env, ctx *context) bool {
	if snowball.FindAmongB(env, a5, ctx) == 0 {
		return false
	}
	return true
}

// vi reports whether the ending is "i" preceded by a V2-group vowel.
func vi(env *snowball.Env, ctx *context) bool {
	if !env.EqSB("i") {
		return false
	}
	if !env.InGroupingB(gV2, 97, 246) {
		return false
	}
	return true
}

// caseEnding removes grammatical case endings (a6) within R1, applying the
// vowel/long-vowel guards per result code and marking bEndingRemoved.
func caseEnding(env *snowball.Env, ctx *context) bool {
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
	amongVar = snowball.FindAmongB(env, a6, ctx)
	if amongVar == 0 {
		env.LimitBackward = v2
		return false
	}
	env.Bra = env.Cursor
	env.LimitBackward = v2
	switch amongVar {
	case 1:
		if !env.EqSB("a") {
			return false
		}
	case 2:
		if !env.EqSB("e") {
			return false
		}
	case 3:
		if !env.EqSB("i") {
			return false
		}
	case 4:
		if !env.EqSB("o") {
			return false
		}
	case 5:
		if !env.EqSB("ä") {
			return false
		}
	case 6:
		if !env.EqSB("ö") {
			return false
		}
	case 7:
		v3 := env.Limit - env.Cursor
	lab0:
		for {
			v4 := env.Limit - env.Cursor
		lab1:
			for {
				v5 := env.Limit - env.Cursor
			lab2:
				for {
					if !long(env, ctx) {
						break lab2
					}
					break lab1
				}
				env.Cursor = env.Limit - v5
				if !env.EqSB("ie") {
					env.Cursor = env.Limit - v3
					break lab0
				}
				break lab1
			}
			env.Cursor = env.Limit - v4
			if env.Cursor <= env.LimitBackward {
				env.Cursor = env.Limit - v3
				break lab0
			}
			env.PreviousChar()
			env.Bra = env.Cursor
			break lab0
		}
	case 8:
		if !env.InGroupingB(gV1, 97, 246) {
			return false
		}
		if !env.OutGroupingB(gV1, 97, 246) {
			return false
		}
	case 9:
		if !env.EqSB("e") {
			return false
		}
	}
	if !env.SliceDel() {
		return false
	}
	ctx.bEndingRemoved = true
	return true
}

// otherEndings removes comparative/superlative endings (a7) within R2, blocking
// the "-po" exception.
func otherEndings(env *snowball.Env, ctx *context) bool {
	var amongVar int32
	v1 := env.Limit - env.Cursor
	if env.Cursor < ctx.iP2 {
		return false
	}
	env.Cursor = ctx.iP2
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
		v3 := env.Limit - env.Cursor
	lab0:
		for {
			if !env.EqSB("po") {
				break lab0
			}
			return false
		}
		env.Cursor = env.Limit - v3
	}
	if !env.SliceDel() {
		return false
	}
	return true
}

// iPlural removes a final i/j plural marker within R1.
func iPlural(env *snowball.Env, ctx *context) bool {
	v1 := env.Limit - env.Cursor
	if env.Cursor < ctx.iP1 {
		return false
	}
	env.Cursor = ctx.iP1
	v2 := env.LimitBackward
	env.LimitBackward = env.Cursor
	env.Cursor = env.Limit - v1
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a8, ctx) == 0 {
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

// tPlural removes a t-plural marker within R1, then an associated comparative
// ending (a9) within R2.
func tPlural(env *snowball.Env, ctx *context) bool {
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
	if !env.EqSB("t") {
		env.LimitBackward = v2
		return false
	}
	env.Bra = env.Cursor
	v3 := env.Limit - env.Cursor
	if !env.InGroupingB(gV1, 97, 246) {
		env.LimitBackward = v2
		return false
	}
	env.Cursor = env.Limit - v3
	if !env.SliceDel() {
		return false
	}
	env.LimitBackward = v2
	v4 := env.Limit - env.Cursor
	if env.Cursor < ctx.iP2 {
		return false
	}
	env.Cursor = ctx.iP2
	v5 := env.LimitBackward
	env.LimitBackward = env.Cursor
	env.Cursor = env.Limit - v4
	env.Ket = env.Cursor
	amongVar = snowball.FindAmongB(env, a9, ctx)
	if amongVar == 0 {
		env.LimitBackward = v5
		return false
	}
	env.Bra = env.Cursor
	env.LimitBackward = v5
	switch amongVar {
	case 1:
		v6 := env.Limit - env.Cursor
	lab0:
		for {
			if !env.EqSB("po") {
				break lab0
			}
			return false
		}
		env.Cursor = env.Limit - v6
	}
	if !env.SliceDel() {
		return false
	}
	return true
}

// tidy normalises the stem after ending removal: it shortens a doubled vowel,
// strips a stray a/ä or j, and removes a final duplicated consonant captured via
// the sX slice.
func tidy(env *snowball.Env, ctx *context) bool {
	v1 := env.Limit - env.Cursor
	if env.Cursor < ctx.iP1 {
		return false
	}
	env.Cursor = ctx.iP1
	v2 := env.LimitBackward
	env.LimitBackward = env.Cursor
	env.Cursor = env.Limit - v1
	v3 := env.Limit - env.Cursor
lab0:
	for {
		v4 := env.Limit - env.Cursor
		if !long(env, ctx) {
			break lab0
		}
		env.Cursor = env.Limit - v4
		env.Ket = env.Cursor
		if env.Cursor <= env.LimitBackward {
			break lab0
		}
		env.PreviousChar()
		env.Bra = env.Cursor
		if !env.SliceDel() {
			return false
		}
		break lab0
	}
	env.Cursor = env.Limit - v3
	v5 := env.Limit - env.Cursor
lab1:
	for {
		env.Ket = env.Cursor
		if !env.InGroupingB(gAEI, 97, 228) {
			break lab1
		}
		env.Bra = env.Cursor
		if !env.OutGroupingB(gV1, 97, 246) {
			break lab1
		}
		if !env.SliceDel() {
			return false
		}
		break lab1
	}
	env.Cursor = env.Limit - v5
	v6 := env.Limit - env.Cursor
lab2:
	for {
		env.Ket = env.Cursor
		if !env.EqSB("j") {
			break lab2
		}
		env.Bra = env.Cursor
	lab3:
		for {
			v7 := env.Limit - env.Cursor
		lab4:
			for {
				if !env.EqSB("o") {
					break lab4
				}
				break lab3
			}
			env.Cursor = env.Limit - v7
			if !env.EqSB("u") {
				break lab2
			}
			break lab3
		}
		if !env.SliceDel() {
			return false
		}
		break lab2
	}
	env.Cursor = env.Limit - v6
	v8 := env.Limit - env.Cursor
lab5:
	for {
		env.Ket = env.Cursor
		if !env.EqSB("o") {
			break lab5
		}
		env.Bra = env.Cursor
		if !env.EqSB("j") {
			break lab5
		}
		if !env.SliceDel() {
			return false
		}
		break lab5
	}
	env.Cursor = env.Limit - v8
	env.LimitBackward = v2
golab6:
	for {
		v9 := env.Limit - env.Cursor
	lab7:
		for {
			if !env.OutGroupingB(gV1, 97, 246) {
				break lab7
			}
			env.Cursor = env.Limit - v9
			break golab6
		}
		env.Cursor = env.Limit - v9
		if env.Cursor <= env.LimitBackward {
			return false
		}
		env.PreviousChar()
	}
	env.Ket = env.Cursor
	if env.Cursor <= env.LimitBackward {
		return false
	}
	env.PreviousChar()
	env.Bra = env.Cursor
	ctx.sX = env.SliceTo()
	if ctx.sX == "" {
		return false
	}
	if !env.EqSB(ctx.sX) {
		return false
	}
	if !env.SliceDel() {
		return false
	}
	return true
}

// Stem runs the Snowball finnish algorithm over env, mirroring the generated
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
	ctx.bEndingRemoved = false
	env.LimitBackward = env.Cursor
	env.Cursor = env.Limit
	v2 := env.Limit - env.Cursor
lab1:
	for {
		if !particleEtc(env, ctx) {
			break lab1
		}
		break lab1
	}
	env.Cursor = env.Limit - v2
	v3 := env.Limit - env.Cursor
lab2:
	for {
		if !possessive(env, ctx) {
			break lab2
		}
		break lab2
	}
	env.Cursor = env.Limit - v3
	v4 := env.Limit - env.Cursor
lab3:
	for {
		if !caseEnding(env, ctx) {
			break lab3
		}
		break lab3
	}
	env.Cursor = env.Limit - v4
	v5 := env.Limit - env.Cursor
lab4:
	for {
		if !otherEndings(env, ctx) {
			break lab4
		}
		break lab4
	}
	env.Cursor = env.Limit - v5
lab5:
	for {
		v6 := env.Limit - env.Cursor
	lab6:
		for {
			if !ctx.bEndingRemoved {
				break lab6
			}
			v7 := env.Limit - env.Cursor
		lab7:
			for {
				if !iPlural(env, ctx) {
					break lab7
				}
				break lab7
			}
			env.Cursor = env.Limit - v7
			break lab5
		}
		env.Cursor = env.Limit - v6
		v8 := env.Limit - env.Cursor
	lab8:
		for {
			if !tPlural(env, ctx) {
				break lab8
			}
			break lab8
		}
		env.Cursor = env.Limit - v8
		break lab5
	}
	v9 := env.Limit - env.Cursor
lab9:
	for {
		if !tidy(env, ctx) {
			break lab9
		}
		break lab9
	}
	env.Cursor = env.Limit - v9
	env.Cursor = env.LimitBackward
	return true
}
