// Package english is a byte-faithful Go port of rust-stemmers' generated
// Snowball "english" stemmer (Porter2). It produces output identical to
// rust-stemmers 1.2.0's English algorithm; the canonical Snowball english
// vocabulary is the conformance oracle.
//
// The control flow mirrors the generated Rust: each Snowball labelled block
// (`'lab: loop { … break 'lab }`) becomes a Go labelled `for { … break LABEL }`,
// and the result-code dispatch (`else if among_var == N`) becomes a `switch`.
package english

import "github.com/freeeve/go-stemmers/internal/snowball"

// context holds the per-run state the generated algorithm threads through its
// routines: whether a vowel y was upper-cased to Y, and the R1/R2 region marks.
type context struct {
	bYfound bool
	iP2     int
	iP1     int
}

var a0 = []snowball.Among[context]{
	{Str: "arsen", SubstringI: -1, Result: -1},
	{Str: "commun", SubstringI: -1, Result: -1},
	{Str: "gener", SubstringI: -1, Result: -1},
}

var a1 = []snowball.Among[context]{
	{Str: "'", SubstringI: -1, Result: 1},
	{Str: "'s'", SubstringI: 0, Result: 1},
	{Str: "'s", SubstringI: -1, Result: 1},
}

var a2 = []snowball.Among[context]{
	{Str: "ied", SubstringI: -1, Result: 2},
	{Str: "s", SubstringI: -1, Result: 3},
	{Str: "ies", SubstringI: 1, Result: 2},
	{Str: "sses", SubstringI: 1, Result: 1},
	{Str: "ss", SubstringI: 1, Result: -1},
	{Str: "us", SubstringI: 1, Result: -1},
}

var a3 = []snowball.Among[context]{
	{Str: "", SubstringI: -1, Result: 3},
	{Str: "bb", SubstringI: 0, Result: 2},
	{Str: "dd", SubstringI: 0, Result: 2},
	{Str: "ff", SubstringI: 0, Result: 2},
	{Str: "gg", SubstringI: 0, Result: 2},
	{Str: "bl", SubstringI: 0, Result: 1},
	{Str: "mm", SubstringI: 0, Result: 2},
	{Str: "nn", SubstringI: 0, Result: 2},
	{Str: "pp", SubstringI: 0, Result: 2},
	{Str: "rr", SubstringI: 0, Result: 2},
	{Str: "at", SubstringI: 0, Result: 1},
	{Str: "tt", SubstringI: 0, Result: 2},
	{Str: "iz", SubstringI: 0, Result: 1},
}

var a4 = []snowball.Among[context]{
	{Str: "ed", SubstringI: -1, Result: 2},
	{Str: "eed", SubstringI: 0, Result: 1},
	{Str: "ing", SubstringI: -1, Result: 2},
	{Str: "edly", SubstringI: -1, Result: 2},
	{Str: "eedly", SubstringI: 3, Result: 1},
	{Str: "ingly", SubstringI: -1, Result: 2},
}

var a5 = []snowball.Among[context]{
	{Str: "anci", SubstringI: -1, Result: 3},
	{Str: "enci", SubstringI: -1, Result: 2},
	{Str: "ogi", SubstringI: -1, Result: 13},
	{Str: "li", SubstringI: -1, Result: 16},
	{Str: "bli", SubstringI: 3, Result: 12},
	{Str: "abli", SubstringI: 4, Result: 4},
	{Str: "alli", SubstringI: 3, Result: 8},
	{Str: "fulli", SubstringI: 3, Result: 14},
	{Str: "lessli", SubstringI: 3, Result: 15},
	{Str: "ousli", SubstringI: 3, Result: 10},
	{Str: "entli", SubstringI: 3, Result: 5},
	{Str: "aliti", SubstringI: -1, Result: 8},
	{Str: "biliti", SubstringI: -1, Result: 12},
	{Str: "iviti", SubstringI: -1, Result: 11},
	{Str: "tional", SubstringI: -1, Result: 1},
	{Str: "ational", SubstringI: 14, Result: 7},
	{Str: "alism", SubstringI: -1, Result: 8},
	{Str: "ation", SubstringI: -1, Result: 7},
	{Str: "ization", SubstringI: 17, Result: 6},
	{Str: "izer", SubstringI: -1, Result: 6},
	{Str: "ator", SubstringI: -1, Result: 7},
	{Str: "iveness", SubstringI: -1, Result: 11},
	{Str: "fulness", SubstringI: -1, Result: 9},
	{Str: "ousness", SubstringI: -1, Result: 10},
}

var a6 = []snowball.Among[context]{
	{Str: "icate", SubstringI: -1, Result: 4},
	{Str: "ative", SubstringI: -1, Result: 6},
	{Str: "alize", SubstringI: -1, Result: 3},
	{Str: "iciti", SubstringI: -1, Result: 4},
	{Str: "ical", SubstringI: -1, Result: 4},
	{Str: "tional", SubstringI: -1, Result: 1},
	{Str: "ational", SubstringI: 5, Result: 2},
	{Str: "ful", SubstringI: -1, Result: 5},
	{Str: "ness", SubstringI: -1, Result: 5},
}

var a7 = []snowball.Among[context]{
	{Str: "ic", SubstringI: -1, Result: 1},
	{Str: "ance", SubstringI: -1, Result: 1},
	{Str: "ence", SubstringI: -1, Result: 1},
	{Str: "able", SubstringI: -1, Result: 1},
	{Str: "ible", SubstringI: -1, Result: 1},
	{Str: "ate", SubstringI: -1, Result: 1},
	{Str: "ive", SubstringI: -1, Result: 1},
	{Str: "ize", SubstringI: -1, Result: 1},
	{Str: "iti", SubstringI: -1, Result: 1},
	{Str: "al", SubstringI: -1, Result: 1},
	{Str: "ism", SubstringI: -1, Result: 1},
	{Str: "ion", SubstringI: -1, Result: 2},
	{Str: "er", SubstringI: -1, Result: 1},
	{Str: "ous", SubstringI: -1, Result: 1},
	{Str: "ant", SubstringI: -1, Result: 1},
	{Str: "ent", SubstringI: -1, Result: 1},
	{Str: "ment", SubstringI: 15, Result: 1},
	{Str: "ement", SubstringI: 16, Result: 1},
}

var a8 = []snowball.Among[context]{
	{Str: "e", SubstringI: -1, Result: 1},
	{Str: "l", SubstringI: -1, Result: 2},
}

var a9 = []snowball.Among[context]{
	{Str: "succeed", SubstringI: -1, Result: -1},
	{Str: "proceed", SubstringI: -1, Result: -1},
	{Str: "exceed", SubstringI: -1, Result: -1},
	{Str: "canning", SubstringI: -1, Result: -1},
	{Str: "inning", SubstringI: -1, Result: -1},
	{Str: "earring", SubstringI: -1, Result: -1},
	{Str: "herring", SubstringI: -1, Result: -1},
	{Str: "outing", SubstringI: -1, Result: -1},
}

var a10 = []snowball.Among[context]{
	{Str: "andes", SubstringI: -1, Result: -1},
	{Str: "atlas", SubstringI: -1, Result: -1},
	{Str: "bias", SubstringI: -1, Result: -1},
	{Str: "cosmos", SubstringI: -1, Result: -1},
	{Str: "dying", SubstringI: -1, Result: 3},
	{Str: "early", SubstringI: -1, Result: 9},
	{Str: "gently", SubstringI: -1, Result: 7},
	{Str: "howe", SubstringI: -1, Result: -1},
	{Str: "idly", SubstringI: -1, Result: 6},
	{Str: "lying", SubstringI: -1, Result: 4},
	{Str: "news", SubstringI: -1, Result: -1},
	{Str: "only", SubstringI: -1, Result: 10},
	{Str: "singly", SubstringI: -1, Result: 11},
	{Str: "skies", SubstringI: -1, Result: 2},
	{Str: "skis", SubstringI: -1, Result: 1},
	{Str: "sky", SubstringI: -1, Result: -1},
	{Str: "tying", SubstringI: -1, Result: 5},
	{Str: "ugly", SubstringI: -1, Result: 8},
}

var gV = []byte{17, 65, 16, 1}
var gVWXY = []byte{1, 17, 65, 208, 1}
var gValidLI = []byte{55, 141, 2}

// prelude removes a leading apostrophe and upper-cases consonant-following y to
// Y so step 1c can treat it as a consonant.
func prelude(env *snowball.Env, ctx *context) bool {
	ctx.bYfound = false
	v1 := env.Cursor
lab0:
	for {
		env.Bra = env.Cursor
		if !env.EqS("'") {
			break lab0
		}
		env.Ket = env.Cursor
		if !env.SliceDel() {
			return false
		}
		break lab0
	}
	env.Cursor = v1
	v2 := env.Cursor
lab1:
	for {
		env.Bra = env.Cursor
		if !env.EqS("y") {
			break lab1
		}
		env.Ket = env.Cursor
		if !env.SliceFrom("Y") {
			return false
		}
		ctx.bYfound = true
		break lab1
	}
	env.Cursor = v2
	v3 := env.Cursor
replab3:
	for {
		v4 := env.Cursor
	lab4:
		for once := 0; once < 1; once++ {
		golab5:
			for {
				v5 := env.Cursor
			lab6:
				for {
					if !env.InGrouping(gV, 97, 121) {
						break lab6
					}
					env.Bra = env.Cursor
					if !env.EqS("y") {
						break lab6
					}
					env.Ket = env.Cursor
					env.Cursor = v5
					break golab5
				}
				env.Cursor = v5
				if env.Cursor >= env.Limit {
					break lab4
				}
				env.NextChar()
			}
			if !env.SliceFrom("Y") {
				return false
			}
			ctx.bYfound = true
			continue replab3
		}
		env.Cursor = v4
		break replab3
	}
	env.Cursor = v3
	return true
}

// markRegions sets the R1 (iP1) and R2 (iP2) region boundaries, with the special
// prefixes in a0 forcing R1 past them.
func markRegions(env *snowball.Env, ctx *context) bool {
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
				if snowball.FindAmong(env, a0, ctx) == 0 {
					break lab2
				}
				break lab1
			}
			env.Cursor = v2
		golab3:
			for {
				for {
					if !env.InGrouping(gV, 97, 121) {
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
					if !env.OutGrouping(gV, 97, 121) {
						break
					}
					break golab5
				}
				if env.Cursor >= env.Limit {
					break lab0
				}
				env.NextChar()
			}
			break lab1
		}
		ctx.iP1 = env.Cursor
	golab7:
		for {
			for {
				if !env.InGrouping(gV, 97, 121) {
					break
				}
				break golab7
			}
			if env.Cursor >= env.Limit {
				break lab0
			}
			env.NextChar()
		}
	golab9:
		for {
			for {
				if !env.OutGrouping(gV, 97, 121) {
					break
				}
				break golab9
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

// shortv reports whether the cursor sits just after a short syllable (the
// condition that blocks a final-e deletion in step 5).
func shortv(env *snowball.Env, ctx *context) bool {
lab0:
	for {
		v1 := env.Limit - env.Cursor
	lab1:
		for {
			if !env.OutGroupingB(gVWXY, 89, 121) {
				break lab1
			}
			if !env.InGroupingB(gV, 97, 121) {
				break lab1
			}
			if !env.OutGroupingB(gV, 97, 121) {
				break lab1
			}
			break lab0
		}
		env.Cursor = env.Limit - v1
		if !env.OutGroupingB(gV, 97, 121) {
			return false
		}
		if !env.InGroupingB(gV, 97, 121) {
			return false
		}
		if env.Cursor > env.LimitBackward {
			return false
		}
		break lab0
	}
	return true
}

// r1 reports whether the cursor is within region R1.
func r1(env *snowball.Env, ctx *context) bool {
	return ctx.iP1 <= env.Cursor
}

// r2 reports whether the cursor is within region R2.
func r2(env *snowball.Env, ctx *context) bool {
	return ctx.iP2 <= env.Cursor
}

// step1a handles plural/possessive endings (apostrophe-s, -sses, -ied, -s …).
func step1a(env *snowball.Env, ctx *context) bool {
	var amongVar int32
	v1 := env.Limit - env.Cursor
lab0:
	for {
		env.Ket = env.Cursor
		amongVar = snowball.FindAmongB(env, a1, ctx)
		if amongVar == 0 {
			env.Cursor = env.Limit - v1
			break lab0
		}
		env.Bra = env.Cursor
		switch amongVar {
		case 1:
			if !env.SliceDel() {
				return false
			}
		}
		break lab0
	}
	env.Ket = env.Cursor
	amongVar = snowball.FindAmongB(env, a2, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	switch amongVar {
	case 1:
		if !env.SliceFrom("ss") {
			return false
		}
	case 2:
	lab1:
		for {
			v2 := env.Limit - env.Cursor
		lab2:
			for {
				c := env.ByteIndexForHop(-2)
				if env.LimitBackward > c || c > env.Limit {
					break lab2
				}
				env.Cursor = c
				if !env.SliceFrom("i") {
					return false
				}
				break lab1
			}
			env.Cursor = env.Limit - v2
			if !env.SliceFrom("ie") {
				return false
			}
			break lab1
		}
	case 3:
		if env.Cursor <= env.LimitBackward {
			return false
		}
		env.PreviousChar()
	golab3:
		for {
			for {
				if !env.InGroupingB(gV, 97, 121) {
					break
				}
				break golab3
			}
			if env.Cursor <= env.LimitBackward {
				return false
			}
			env.PreviousChar()
		}
		if !env.SliceDel() {
			return false
		}
	}
	return true
}

// step1b handles -ed/-ing endings, with the post-deletion fix-ups in a3.
func step1b(env *snowball.Env, ctx *context) bool {
	var amongVar int32
	env.Ket = env.Cursor
	amongVar = snowball.FindAmongB(env, a4, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	switch amongVar {
	case 1:
		if !r1(env, ctx) {
			return false
		}
		if !env.SliceFrom("ee") {
			return false
		}
	case 2:
		v1 := env.Limit - env.Cursor
	golab0:
		for {
			for {
				if !env.InGroupingB(gV, 97, 121) {
					break
				}
				break golab0
			}
			if env.Cursor <= env.LimitBackward {
				return false
			}
			env.PreviousChar()
		}
		env.Cursor = env.Limit - v1
		if !env.SliceDel() {
			return false
		}
		v3 := env.Limit - env.Cursor
		amongVar = snowball.FindAmongB(env, a3, ctx)
		if amongVar == 0 {
			return false
		}
		env.Cursor = env.Limit - v3
		switch amongVar {
		case 1:
			c := env.Cursor
			env.Insert(env.Cursor, env.Cursor, "e")
			env.Cursor = c
		case 2:
			env.Ket = env.Cursor
			if env.Cursor <= env.LimitBackward {
				return false
			}
			env.PreviousChar()
			env.Bra = env.Cursor
			if !env.SliceDel() {
				return false
			}
		case 3:
			if env.Cursor != ctx.iP1 {
				return false
			}
			v4 := env.Limit - env.Cursor
			if !shortv(env, ctx) {
				return false
			}
			env.Cursor = env.Limit - v4
			c := env.Cursor
			env.Insert(env.Cursor, env.Cursor, "e")
			env.Cursor = c
		}
	}
	return true
}

// step1c turns a terminal y/Y after a consonant into i.
func step1c(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
lab0:
	for {
		v1 := env.Limit - env.Cursor
	lab1:
		for {
			if !env.EqSB("y") {
				break lab1
			}
			break lab0
		}
		env.Cursor = env.Limit - v1
		if !env.EqSB("Y") {
			return false
		}
		break lab0
	}
	env.Bra = env.Cursor
	if !env.OutGroupingB(gV, 97, 121) {
		return false
	}
	v2 := env.Limit - env.Cursor
lab2:
	for {
		if env.Cursor > env.LimitBackward {
			break lab2
		}
		return false
	}
	env.Cursor = env.Limit - v2
	if !env.SliceFrom("i") {
		return false
	}
	return true
}

// step2 rewrites the longer derivational suffixes in a5 when in R1.
func step2(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a5, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !r1(env, ctx) {
		return false
	}
	switch amongVar {
	case 1:
		if !env.SliceFrom("tion") {
			return false
		}
	case 2:
		if !env.SliceFrom("ence") {
			return false
		}
	case 3:
		if !env.SliceFrom("ance") {
			return false
		}
	case 4:
		if !env.SliceFrom("able") {
			return false
		}
	case 5:
		if !env.SliceFrom("ent") {
			return false
		}
	case 6:
		if !env.SliceFrom("ize") {
			return false
		}
	case 7:
		if !env.SliceFrom("ate") {
			return false
		}
	case 8:
		if !env.SliceFrom("al") {
			return false
		}
	case 9:
		if !env.SliceFrom("ful") {
			return false
		}
	case 10:
		if !env.SliceFrom("ous") {
			return false
		}
	case 11:
		if !env.SliceFrom("ive") {
			return false
		}
	case 12:
		if !env.SliceFrom("ble") {
			return false
		}
	case 13:
		if !env.EqSB("l") {
			return false
		}
		if !env.SliceFrom("og") {
			return false
		}
	case 14:
		if !env.SliceFrom("ful") {
			return false
		}
	case 15:
		if !env.SliceFrom("less") {
			return false
		}
	case 16:
		if !env.InGroupingB(gValidLI, 99, 116) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	}
	return true
}

// step3 rewrites the suffixes in a6 when in R1 (and -ative only in R2).
func step3(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a6, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !r1(env, ctx) {
		return false
	}
	switch amongVar {
	case 1:
		if !env.SliceFrom("tion") {
			return false
		}
	case 2:
		if !env.SliceFrom("ate") {
			return false
		}
	case 3:
		if !env.SliceFrom("al") {
			return false
		}
	case 4:
		if !env.SliceFrom("ic") {
			return false
		}
	case 5:
		if !env.SliceDel() {
			return false
		}
	case 6:
		if !r2(env, ctx) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	}
	return true
}

// step4 deletes the residual derivational suffixes in a7 when in R2.
func step4(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a7, ctx)
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
	lab0:
		for {
			v1 := env.Limit - env.Cursor
		lab1:
			for {
				if !env.EqSB("s") {
					break lab1
				}
				break lab0
			}
			env.Cursor = env.Limit - v1
			if !env.EqSB("t") {
				return false
			}
			break lab0
		}
		if !env.SliceDel() {
			return false
		}
	}
	return true
}

// step5 removes a final e (in R2, or R1 when not a short syllable) and a final l
// after l in R2.
func step5(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a8, ctx)
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
				if !r2(env, ctx) {
					break lab1
				}
				break lab0
			}
			env.Cursor = env.Limit - v1
			if !r1(env, ctx) {
				return false
			}
			v2 := env.Limit - env.Cursor
		lab2:
			for {
				if !shortv(env, ctx) {
					break lab2
				}
				return false
			}
			env.Cursor = env.Limit - v2
			break lab0
		}
		if !env.SliceDel() {
			return false
		}
	case 2:
		if !r2(env, ctx) {
			return false
		}
		if !env.EqSB("l") {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	}
	return true
}

// exception2 recognises invariant words (succeed, herring, outing …) that must
// not pass through steps 1b–5.
func exception2(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a9, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if env.Cursor > env.LimitBackward {
		return false
	}
	return true
}

// exception1 handles the whole-word exceptions in a10 (skis→ski, dying→die, plus
// invariants like atlas/news/sky).
func exception1(env *snowball.Env, ctx *context) bool {
	env.Bra = env.Cursor
	amongVar := snowball.FindAmong(env, a10, ctx)
	if amongVar == 0 {
		return false
	}
	env.Ket = env.Cursor
	if env.Cursor < env.Limit {
		return false
	}
	switch amongVar {
	case 1:
		if !env.SliceFrom("ski") {
			return false
		}
	case 2:
		if !env.SliceFrom("sky") {
			return false
		}
	case 3:
		if !env.SliceFrom("die") {
			return false
		}
	case 4:
		if !env.SliceFrom("lie") {
			return false
		}
	case 5:
		if !env.SliceFrom("tie") {
			return false
		}
	case 6:
		if !env.SliceFrom("idl") {
			return false
		}
	case 7:
		if !env.SliceFrom("gentl") {
			return false
		}
	case 8:
		if !env.SliceFrom("ugli") {
			return false
		}
	case 9:
		if !env.SliceFrom("earli") {
			return false
		}
	case 10:
		if !env.SliceFrom("onli") {
			return false
		}
	case 11:
		if !env.SliceFrom("singl") {
			return false
		}
	}
	return true
}

// postlude restores any Y introduced by the prelude back to lowercase y.
func postlude(env *snowball.Env, ctx *context) bool {
	if !ctx.bYfound {
		return false
	}
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
					env.Bra = env.Cursor
					if !env.EqS("Y") {
						break lab3
					}
					env.Ket = env.Cursor
					env.Cursor = v2
					break golab2
				}
				env.Cursor = v2
				if env.Cursor >= env.Limit {
					break lab1
				}
				env.NextChar()
			}
			if !env.SliceFrom("y") {
				return false
			}
			continue replab0
		}
		env.Cursor = v1
		break replab0
	}
	return true
}

// Stem runs the Snowball english (Porter2) algorithm over env, mirroring the
// generated `stem` entry point. It always returns true; the result is the
// mutated env.
func Stem(env *snowball.Env) bool {
	ctx := &context{}
lab0:
	for {
		v1 := env.Cursor
	lab1:
		for {
			if !exception1(env, ctx) {
				break lab1
			}
			break lab0
		}
		env.Cursor = v1
	lab2:
		for {
			v2 := env.Cursor
		lab3:
			for {
				c := env.ByteIndexForHop(3)
				if 0 > c || c > env.Limit {
					break lab3
				}
				env.Cursor = c
				break lab2
			}
			env.Cursor = v2
			break lab0
		}
		env.Cursor = v1

		v3 := env.Cursor
		prelude(env, ctx)
		env.Cursor = v3

		v4 := env.Cursor
		markRegions(env, ctx)
		env.Cursor = v4

		env.LimitBackward = env.Cursor
		env.Cursor = env.Limit

		v5 := env.Limit - env.Cursor
		step1a(env, ctx)
		env.Cursor = env.Limit - v5

	lab7:
		for {
			v6 := env.Limit - env.Cursor
		lab8:
			for {
				if !exception2(env, ctx) {
					break lab8
				}
				break lab7
			}
			env.Cursor = env.Limit - v6

			v7 := env.Limit - env.Cursor
			step1b(env, ctx)
			env.Cursor = env.Limit - v7

			v8 := env.Limit - env.Cursor
			step1c(env, ctx)
			env.Cursor = env.Limit - v8

			v9 := env.Limit - env.Cursor
			step2(env, ctx)
			env.Cursor = env.Limit - v9

			v10 := env.Limit - env.Cursor
			step3(env, ctx)
			env.Cursor = env.Limit - v10

			v11 := env.Limit - env.Cursor
			step4(env, ctx)
			env.Cursor = env.Limit - v11

			v12 := env.Limit - env.Cursor
			step5(env, ctx)
			env.Cursor = env.Limit - v12

			break lab7
		}
		env.Cursor = env.LimitBackward

		v13 := env.Cursor
		postlude(env, ctx)
		env.Cursor = v13

		break lab0
	}
	return true
}
