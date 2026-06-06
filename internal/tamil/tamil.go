// Package tamil is a byte-faithful Go port of rust-stemmers' generated Snowball
// "tamil" stemmer. It produces output identical to rust-stemmers 1.2.0's Tamil
// algorithm.
//
// The control flow mirrors the generated Rust: each Snowball labelled block
// (`'lab: loop { … break 'lab }`) becomes a Go labelled `for { … break LABEL }`,
// each `repeat`/`do` wrapper its corresponding loop, and the among-table search
// the runtime FindAmong/FindAmongB helpers.
package tamil

import "github.com/freeeve/go-stemmers/internal/snowball"

// context holds the per-run state the generated algorithm threads through its
// routines: the character length and the boolean match flags.
type context struct {
	iLength             int
	bFoundWrongEnding   bool
	bFoundVetrumaiUrupu bool
	bFoundAMatch        bool
}

var a0 = []snowball.Among[context]{
	{Str: "க", SubstringI: -1, Result: -1},
	{Str: "ங", SubstringI: -1, Result: -1},
	{Str: "ச", SubstringI: -1, Result: -1},
	{Str: "ஞ", SubstringI: -1, Result: -1},
	{Str: "த", SubstringI: -1, Result: -1},
	{Str: "ந", SubstringI: -1, Result: -1},
	{Str: "ப", SubstringI: -1, Result: -1},
	{Str: "ம", SubstringI: -1, Result: -1},
	{Str: "ய", SubstringI: -1, Result: -1},
	{Str: "வ", SubstringI: -1, Result: -1},
}

var a1 = []snowball.Among[context]{
	{Str: "ந்த்", SubstringI: -1, Result: -1},
	{Str: "ந்", SubstringI: -1, Result: -1},
	{Str: "ந்த", SubstringI: -1, Result: -1},
}

var a2 = []snowball.Among[context]{
	{Str: "ீ", SubstringI: -1, Result: -1},
	{Str: "ை", SubstringI: -1, Result: -1},
	{Str: "ி", SubstringI: -1, Result: -1},
}

var a3 = []snowball.Among[context]{
	{Str: "க", SubstringI: -1, Result: -1},
	{Str: "ச", SubstringI: -1, Result: -1},
	{Str: "ட", SubstringI: -1, Result: -1},
	{Str: "த", SubstringI: -1, Result: -1},
	{Str: "ப", SubstringI: -1, Result: -1},
	{Str: "ற", SubstringI: -1, Result: -1},
}

var a4 = []snowball.Among[context]{
	{Str: "க", SubstringI: -1, Result: -1},
	{Str: "ச", SubstringI: -1, Result: -1},
	{Str: "ட", SubstringI: -1, Result: -1},
	{Str: "த", SubstringI: -1, Result: -1},
	{Str: "ப", SubstringI: -1, Result: -1},
	{Str: "ற", SubstringI: -1, Result: -1},
}

var a5 = []snowball.Among[context]{
	{Str: "க", SubstringI: -1, Result: -1},
	{Str: "ச", SubstringI: -1, Result: -1},
	{Str: "ட", SubstringI: -1, Result: -1},
	{Str: "த", SubstringI: -1, Result: -1},
	{Str: "ப", SubstringI: -1, Result: -1},
	{Str: "ற", SubstringI: -1, Result: -1},
}

var a6 = []snowball.Among[context]{
	{Str: "ய", SubstringI: -1, Result: -1},
	{Str: "ர", SubstringI: -1, Result: -1},
	{Str: "ல", SubstringI: -1, Result: -1},
	{Str: "ள", SubstringI: -1, Result: -1},
	{Str: "ழ", SubstringI: -1, Result: -1},
	{Str: "வ", SubstringI: -1, Result: -1},
}

var a7 = []snowball.Among[context]{
	{Str: "ங", SubstringI: -1, Result: -1},
	{Str: "ஞ", SubstringI: -1, Result: -1},
	{Str: "ண", SubstringI: -1, Result: -1},
	{Str: "ந", SubstringI: -1, Result: -1},
	{Str: "ன", SubstringI: -1, Result: -1},
	{Str: "ம", SubstringI: -1, Result: -1},
}

var a8 = []snowball.Among[context]{
	{Str: "வ்", SubstringI: -1, Result: -1},
	{Str: "ய", SubstringI: -1, Result: -1},
	{Str: "வ", SubstringI: -1, Result: -1},
}

var a9 = []snowball.Among[context]{
	{Str: "ீ", SubstringI: -1, Result: -1},
	{Str: "ு", SubstringI: -1, Result: -1},
	{Str: "ூ", SubstringI: -1, Result: -1},
	{Str: "ெ", SubstringI: -1, Result: -1},
	{Str: "ே", SubstringI: -1, Result: -1},
	{Str: "ை", SubstringI: -1, Result: -1},
	{Str: "ா", SubstringI: -1, Result: -1},
	{Str: "ி", SubstringI: -1, Result: -1},
}

var a10 = []snowball.Among[context]{
	{Str: "ீ", SubstringI: -1, Result: -1},
	{Str: "ு", SubstringI: -1, Result: -1},
	{Str: "ூ", SubstringI: -1, Result: -1},
	{Str: "ெ", SubstringI: -1, Result: -1},
	{Str: "ே", SubstringI: -1, Result: -1},
	{Str: "ை", SubstringI: -1, Result: -1},
	{Str: "ா", SubstringI: -1, Result: -1},
	{Str: "ி", SubstringI: -1, Result: -1},
}

var a11 = []snowball.Among[context]{
	{Str: "அ", SubstringI: -1, Result: -1},
	{Str: "இ", SubstringI: -1, Result: -1},
	{Str: "உ", SubstringI: -1, Result: -1},
}

var a12 = []snowball.Among[context]{
	{Str: "க", SubstringI: -1, Result: -1},
	{Str: "ங", SubstringI: -1, Result: -1},
	{Str: "ச", SubstringI: -1, Result: -1},
	{Str: "ஞ", SubstringI: -1, Result: -1},
	{Str: "த", SubstringI: -1, Result: -1},
	{Str: "ந", SubstringI: -1, Result: -1},
	{Str: "ப", SubstringI: -1, Result: -1},
	{Str: "ம", SubstringI: -1, Result: -1},
	{Str: "ய", SubstringI: -1, Result: -1},
	{Str: "வ", SubstringI: -1, Result: -1},
}

var a13 = []snowball.Among[context]{
	{Str: "க", SubstringI: -1, Result: -1},
	{Str: "ச", SubstringI: -1, Result: -1},
	{Str: "ட", SubstringI: -1, Result: -1},
	{Str: "த", SubstringI: -1, Result: -1},
	{Str: "ப", SubstringI: -1, Result: -1},
	{Str: "ற", SubstringI: -1, Result: -1},
}

var a14 = []snowball.Among[context]{
	{Str: "ே", SubstringI: -1, Result: -1},
	{Str: "ோ", SubstringI: -1, Result: -1},
	{Str: "ா", SubstringI: -1, Result: -1},
}

var a15 = []snowball.Among[context]{
	{Str: "பி", SubstringI: -1, Result: -1},
	{Str: "வி", SubstringI: -1, Result: -1},
}

var a16 = []snowball.Among[context]{
	{Str: "ீ", SubstringI: -1, Result: -1},
	{Str: "ு", SubstringI: -1, Result: -1},
	{Str: "ூ", SubstringI: -1, Result: -1},
	{Str: "ெ", SubstringI: -1, Result: -1},
	{Str: "ே", SubstringI: -1, Result: -1},
	{Str: "ை", SubstringI: -1, Result: -1},
	{Str: "ா", SubstringI: -1, Result: -1},
	{Str: "ி", SubstringI: -1, Result: -1},
}

var a17 = []snowball.Among[context]{
	{Str: "பட்டு", SubstringI: -1, Result: -1},
	{Str: "விட்டு", SubstringI: -1, Result: -1},
	{Str: "படு", SubstringI: -1, Result: -1},
	{Str: "விடு", SubstringI: -1, Result: -1},
	{Str: "பட்டது", SubstringI: -1, Result: -1},
	{Str: "ெல்லாம்", SubstringI: -1, Result: -1},
	{Str: "பட்ட", SubstringI: -1, Result: -1},
	{Str: "பட்டண", SubstringI: -1, Result: -1},
	{Str: "தான", SubstringI: -1, Result: -1},
	{Str: "படிதான", SubstringI: 8, Result: -1},
	{Str: "குரிய", SubstringI: -1, Result: -1},
	{Str: "படி", SubstringI: -1, Result: -1},
	{Str: "பற்றி", SubstringI: -1, Result: -1},
}

var a18 = []snowball.Among[context]{
	{Str: "க", SubstringI: -1, Result: -1},
	{Str: "ச", SubstringI: -1, Result: -1},
	{Str: "ட", SubstringI: -1, Result: -1},
	{Str: "த", SubstringI: -1, Result: -1},
	{Str: "ப", SubstringI: -1, Result: -1},
	{Str: "ற", SubstringI: -1, Result: -1},
}

var a19 = []snowball.Among[context]{
	{Str: "க", SubstringI: -1, Result: -1},
	{Str: "ச", SubstringI: -1, Result: -1},
	{Str: "ட", SubstringI: -1, Result: -1},
	{Str: "த", SubstringI: -1, Result: -1},
	{Str: "ப", SubstringI: -1, Result: -1},
	{Str: "ற", SubstringI: -1, Result: -1},
}

var a20 = []snowball.Among[context]{
	{Str: "ீ", SubstringI: -1, Result: -1},
	{Str: "ு", SubstringI: -1, Result: -1},
	{Str: "ூ", SubstringI: -1, Result: -1},
	{Str: "ெ", SubstringI: -1, Result: -1},
	{Str: "ே", SubstringI: -1, Result: -1},
	{Str: "ை", SubstringI: -1, Result: -1},
	{Str: "ா", SubstringI: -1, Result: -1},
	{Str: "ி", SubstringI: -1, Result: -1},
}

var a21 = []snowball.Among[context]{
	{Str: "ீ", SubstringI: -1, Result: -1},
	{Str: "ு", SubstringI: -1, Result: -1},
	{Str: "ூ", SubstringI: -1, Result: -1},
	{Str: "ெ", SubstringI: -1, Result: -1},
	{Str: "ே", SubstringI: -1, Result: -1},
	{Str: "ை", SubstringI: -1, Result: -1},
	{Str: "ா", SubstringI: -1, Result: -1},
	{Str: "ி", SubstringI: -1, Result: -1},
}

var a22 = []snowball.Among[context]{
	{Str: "படு", SubstringI: -1, Result: -1},
	{Str: "கொண்டிர்", SubstringI: -1, Result: -1},
}

var a23 = []snowball.Among[context]{
	{Str: "அ", SubstringI: -1, Result: -1},
	{Str: "ஆ", SubstringI: -1, Result: -1},
	{Str: "இ", SubstringI: -1, Result: -1},
	{Str: "ஈ", SubstringI: -1, Result: -1},
	{Str: "உ", SubstringI: -1, Result: -1},
	{Str: "ஊ", SubstringI: -1, Result: -1},
	{Str: "எ", SubstringI: -1, Result: -1},
	{Str: "ஏ", SubstringI: -1, Result: -1},
	{Str: "ஐ", SubstringI: -1, Result: -1},
	{Str: "ஒ", SubstringI: -1, Result: -1},
	{Str: "ஓ", SubstringI: -1, Result: -1},
	{Str: "ஔ", SubstringI: -1, Result: -1},
}

var a24 = []snowball.Among[context]{
	{Str: "ீ", SubstringI: -1, Result: -1},
	{Str: "ு", SubstringI: -1, Result: -1},
	{Str: "ூ", SubstringI: -1, Result: -1},
	{Str: "ெ", SubstringI: -1, Result: -1},
	{Str: "ே", SubstringI: -1, Result: -1},
	{Str: "ை", SubstringI: -1, Result: -1},
	{Str: "ா", SubstringI: -1, Result: -1},
	{Str: "ி", SubstringI: -1, Result: -1},
}

var a25 = []snowball.Among[context]{
	{Str: "கின்ற்", SubstringI: -1, Result: -1},
	{Str: "ாநின்ற்", SubstringI: -1, Result: -1},
	{Str: "கிற்", SubstringI: -1, Result: -1},
	{Str: "கின்ற", SubstringI: -1, Result: -1},
	{Str: "ாநின்ற", SubstringI: -1, Result: -1},
	{Str: "கிற", SubstringI: -1, Result: -1},
}

// hasMinLength records the character length and reports whether the word is
// longer than four characters.
func hasMinLength(env *snowball.Env, ctx *context) bool {
	ctx.iLength = env.RuneCount()
	return ctx.iLength > 4
}

// fixVaStart normalises a leading va-vowel ligature back to its bare vowel.
func fixVaStart(env *snowball.Env, ctx *context) bool {
lab0:
	for {
		v1 := env.Cursor
	lab1:
		for {
			v2 := env.Cursor
			v3 := env.Cursor
		lab2:
			for {
				if !env.EqS("வோ") {
					env.Cursor = v3
					break lab2
				}
				break lab2
			}
			env.Cursor = v2
			env.Bra = env.Cursor
			if !env.EqS("வோ") {
				break lab1
			}
			env.Ket = env.Cursor
			if !env.SliceFrom("ஓ") {
				return false
			}
			break lab0
		}
		env.Cursor = v1
	lab3:
		for {
			v4 := env.Cursor
			v5 := env.Cursor
		lab4:
			for {
				if !env.EqS("வொ") {
					env.Cursor = v5
					break lab4
				}
				break lab4
			}
			env.Cursor = v4
			env.Bra = env.Cursor
			if !env.EqS("வொ") {
				break lab3
			}
			env.Ket = env.Cursor
			if !env.SliceFrom("ஒ") {
				return false
			}
			break lab0
		}
		env.Cursor = v1
	lab5:
		for {
			v6 := env.Cursor
			v7 := env.Cursor
		lab6:
			for {
				if !env.EqS("வு") {
					env.Cursor = v7
					break lab6
				}
				break lab6
			}
			env.Cursor = v6
			env.Bra = env.Cursor
			if !env.EqS("வு") {
				break lab5
			}
			env.Ket = env.Cursor
			if !env.SliceFrom("உ") {
				return false
			}
			break lab0
		}
		env.Cursor = v1
		v8 := env.Cursor
		v9 := env.Cursor
	lab7:
		for {
			if !env.EqS("வூ") {
				env.Cursor = v9
				break lab7
			}
			break lab7
		}
		env.Cursor = v8
		env.Bra = env.Cursor
		if !env.EqS("வூ") {
			return false
		}
		env.Ket = env.Cursor
		if !env.SliceFrom("ஊ") {
			return false
		}
		break lab0
	}
	return true
}

// fixEndings repeatedly applies fixEnding while it keeps flagging a wrong ending.
func fixEndings(env *snowball.Env, ctx *context) bool {
	ctx.bFoundWrongEnding = true
replab0:
	for {
		v1 := env.Cursor
	lab1:
		for once := 0; once < 1; once++ {
			if !ctx.bFoundWrongEnding {
				break lab1
			}
			v2 := env.Cursor
		lab2:
			for {
				if !fixEnding(env, ctx) {
					break lab2
				}
				break lab2
			}
			env.Cursor = v2
			continue replab0
		}
		env.Cursor = v1
		break replab0
	}
	return true
}

// removeQuestionPrefixes strips a leading interrogative prefix.
func removeQuestionPrefixes(env *snowball.Env, ctx *context) bool {
	env.Bra = env.Cursor
	if !env.EqS("எ") {
		return false
	}
	if snowball.FindAmong(env, a0, ctx) == 0 {
		return false
	}
	if !env.EqS("்") {
		return false
	}
	env.Ket = env.Cursor
	if !env.SliceDel() {
		return false
	}
	v1 := env.Cursor
lab0:
	for {
		if !fixVaStart(env, ctx) {
			break lab0
		}
		break lab0
	}
	env.Cursor = v1
	return true
}

// fixEnding rewrites or deletes a trailing consonant cluster left malformed by a
// preceding suffix removal.
func fixEnding(env *snowball.Env, ctx *context) bool {
	ctx.bFoundWrongEnding = false
	ctx.iLength = env.RuneCount()
	if !(ctx.iLength > 3) {
		return false
	}
	env.LimitBackward = env.Cursor
	env.Cursor = env.Limit
lab0:
	for {
		v1 := env.Limit - env.Cursor
	lab1:
		for {
			env.Ket = env.Cursor
			if snowball.FindAmongB(env, a1, ctx) == 0 {
				break lab1
			}
			env.Bra = env.Cursor
			if !env.SliceDel() {
				return false
			}
			break lab0
		}
		env.Cursor = env.Limit - v1
	lab2:
		for {
			env.Ket = env.Cursor
			if !env.EqSB("ய்") {
				break lab2
			}
			v2 := env.Limit - env.Cursor
			if snowball.FindAmongB(env, a2, ctx) == 0 {
				break lab2
			}
			env.Cursor = env.Limit - v2
			env.Bra = env.Cursor
			if !env.SliceDel() {
				return false
			}
			break lab0
		}
		env.Cursor = env.Limit - v1
	lab3:
		for {
			env.Ket = env.Cursor
		lab4:
			for {
				v3 := env.Limit - env.Cursor
			lab5:
				for {
					if !env.EqSB("ட்ப்") {
						break lab5
					}
					break lab4
				}
				env.Cursor = env.Limit - v3
				if !env.EqSB("ட்க்") {
					break lab3
				}
				break lab4
			}
			env.Bra = env.Cursor
			if !env.SliceFrom("ள்") {
				return false
			}
			break lab0
		}
		env.Cursor = env.Limit - v1
	lab6:
		for {
			env.Ket = env.Cursor
			if !env.EqSB("ன்ற்") {
				break lab6
			}
			env.Bra = env.Cursor
			if !env.SliceFrom("ல்") {
				return false
			}
			break lab0
		}
		env.Cursor = env.Limit - v1
	lab7:
		for {
			env.Ket = env.Cursor
			if !env.EqSB("ற்க்") {
				break lab7
			}
			env.Bra = env.Cursor
			if !env.SliceFrom("ல்") {
				return false
			}
			break lab0
		}
		env.Cursor = env.Limit - v1
	lab8:
		for {
			env.Ket = env.Cursor
			if !env.EqSB("ட்ட்") {
				break lab8
			}
			env.Bra = env.Cursor
			if !env.SliceFrom("டு") {
				return false
			}
			break lab0
		}
		env.Cursor = env.Limit - v1
	lab9:
		for {
			if !ctx.bFoundVetrumaiUrupu {
				break lab9
			}
			env.Ket = env.Cursor
			if !env.EqSB("த்த்") {
				break lab9
			}
			v4 := env.Limit - env.Cursor
			v5 := env.Limit - env.Cursor
		lab10:
			for {
				if !env.EqSB("ை") {
					break lab10
				}
				break lab9
			}
			env.Cursor = env.Limit - v5
			env.Cursor = env.Limit - v4
			env.Bra = env.Cursor
			if !env.SliceFrom("ம்") {
				return false
			}
			env.Bra = env.Cursor
			break lab0
		}
		env.Cursor = env.Limit - v1
	lab11:
		for {
			env.Ket = env.Cursor
		lab12:
			for {
				v6 := env.Limit - env.Cursor
			lab13:
				for {
					if !env.EqSB("ுக்") {
						break lab13
					}
					break lab12
				}
				env.Cursor = env.Limit - v6
				if !env.EqSB("ுக்க்") {
					break lab11
				}
				break lab12
			}
			env.Bra = env.Cursor
			if !env.SliceFrom("்") {
				return false
			}
			break lab0
		}
		env.Cursor = env.Limit - v1
	lab14:
		for {
			env.Ket = env.Cursor
			if !env.EqSB("்") {
				break lab14
			}
			if snowball.FindAmongB(env, a3, ctx) == 0 {
				break lab14
			}
			if !env.EqSB("்") {
				break lab14
			}
			if snowball.FindAmongB(env, a4, ctx) == 0 {
				break lab14
			}
			env.Bra = env.Cursor
			if !env.SliceDel() {
				return false
			}
			break lab0
		}
		env.Cursor = env.Limit - v1
	lab15:
		for {
			env.Ket = env.Cursor
			if !env.EqSB("ுக்") {
				break lab15
			}
			env.Bra = env.Cursor
			if !env.SliceFrom("்") {
				return false
			}
			break lab0
		}
		env.Cursor = env.Limit - v1
	lab16:
		for {
			env.Ket = env.Cursor
			if !env.EqSB("்") {
				break lab16
			}
			if snowball.FindAmongB(env, a5, ctx) == 0 {
				break lab16
			}
			env.Bra = env.Cursor
			if !env.SliceDel() {
				return false
			}
			break lab0
		}
		env.Cursor = env.Limit - v1
	lab17:
		for {
			env.Ket = env.Cursor
			if !env.EqSB("்") {
				break lab17
			}
		lab18:
			for {
				v7 := env.Limit - env.Cursor
			lab19:
				for {
					if snowball.FindAmongB(env, a6, ctx) == 0 {
						break lab19
					}
					break lab18
				}
				env.Cursor = env.Limit - v7
				if snowball.FindAmongB(env, a7, ctx) == 0 {
					break lab17
				}
				break lab18
			}
			if !env.EqSB("்") {
				break lab17
			}
			env.Bra = env.Cursor
			if !env.SliceFrom("்") {
				return false
			}
			break lab0
		}
		env.Cursor = env.Limit - v1
	lab20:
		for {
			env.Ket = env.Cursor
			if snowball.FindAmongB(env, a8, ctx) == 0 {
				break lab20
			}
			env.Bra = env.Cursor
			if !env.SliceDel() {
				return false
			}
			break lab0
		}
		env.Cursor = env.Limit - v1
	lab21:
		for {
			env.Ket = env.Cursor
			if !env.EqSB("னு") {
				break lab21
			}
			v8 := env.Limit - env.Cursor
			v9 := env.Limit - env.Cursor
		lab22:
			for {
				if snowball.FindAmongB(env, a9, ctx) == 0 {
					break lab22
				}
				break lab21
			}
			env.Cursor = env.Limit - v9
			env.Cursor = env.Limit - v8
			env.Bra = env.Cursor
			if !env.SliceDel() {
				return false
			}
			break lab0
		}
		env.Cursor = env.Limit - v1
	lab23:
		for {
			env.Ket = env.Cursor
			if !env.EqSB("ங்") {
				break lab23
			}
			v10 := env.Limit - env.Cursor
			v11 := env.Limit - env.Cursor
		lab24:
			for {
				if !env.EqSB("ை") {
					break lab24
				}
				break lab23
			}
			env.Cursor = env.Limit - v11
			env.Cursor = env.Limit - v10
			env.Bra = env.Cursor
			if !env.SliceFrom("ம்") {
				return false
			}
			break lab0
		}
		env.Cursor = env.Limit - v1
	lab25:
		for {
			env.Ket = env.Cursor
			if !env.EqSB("ங்") {
				break lab25
			}
			env.Bra = env.Cursor
			if !env.SliceDel() {
				return false
			}
			break lab0
		}
		env.Cursor = env.Limit - v1
		env.Ket = env.Cursor
		if !env.EqSB("்") {
			return false
		}
		v12 := env.Limit - env.Cursor
	lab26:
		for {
			v13 := env.Limit - env.Cursor
		lab27:
			for {
				if snowball.FindAmongB(env, a10, ctx) == 0 {
					break lab27
				}
				break lab26
			}
			env.Cursor = env.Limit - v13
			if !env.EqSB("்") {
				return false
			}
			break lab26
		}
		env.Cursor = env.Limit - v12
		env.Bra = env.Cursor
		if !env.SliceDel() {
			return false
		}
		break lab0
	}
	env.Cursor = env.LimitBackward
	ctx.bFoundWrongEnding = true
	return true
}

// removePronounPrefixes strips a leading pronoun prefix.
func removePronounPrefixes(env *snowball.Env, ctx *context) bool {
	ctx.bFoundAMatch = false
	env.Bra = env.Cursor
	if snowball.FindAmong(env, a11, ctx) == 0 {
		return false
	}
	if snowball.FindAmong(env, a12, ctx) == 0 {
		return false
	}
	if !env.EqS("்") {
		return false
	}
	env.Ket = env.Cursor
	if !env.SliceDel() {
		return false
	}
	ctx.bFoundAMatch = true
	v1 := env.Cursor
lab0:
	for {
		if !fixVaStart(env, ctx) {
			break lab0
		}
		break lab0
	}
	env.Cursor = v1
	return true
}

// removePluralSuffix strips a trailing plural marker.
func removePluralSuffix(env *snowball.Env, ctx *context) bool {
	ctx.bFoundAMatch = false
	env.LimitBackward = env.Cursor
	env.Cursor = env.Limit
lab0:
	for {
		v1 := env.Limit - env.Cursor
	lab1:
		for {
			env.Ket = env.Cursor
			if !env.EqSB("ுங்கள்") {
				break lab1
			}
			v2 := env.Limit - env.Cursor
			v3 := env.Limit - env.Cursor
		lab2:
			for {
				if snowball.FindAmongB(env, a13, ctx) == 0 {
					break lab2
				}
				break lab1
			}
			env.Cursor = env.Limit - v3
			env.Cursor = env.Limit - v2
			env.Bra = env.Cursor
			if !env.SliceFrom("்") {
				return false
			}
			break lab0
		}
		env.Cursor = env.Limit - v1
	lab3:
		for {
			env.Ket = env.Cursor
			if !env.EqSB("ற்கள்") {
				break lab3
			}
			env.Bra = env.Cursor
			if !env.SliceFrom("ல்") {
				return false
			}
			break lab0
		}
		env.Cursor = env.Limit - v1
	lab4:
		for {
			env.Ket = env.Cursor
			if !env.EqSB("ட்கள்") {
				break lab4
			}
			env.Bra = env.Cursor
			if !env.SliceFrom("ள்") {
				return false
			}
			break lab0
		}
		env.Cursor = env.Limit - v1
		env.Ket = env.Cursor
		if !env.EqSB("கள்") {
			return false
		}
		env.Bra = env.Cursor
		if !env.SliceDel() {
			return false
		}
		break lab0
	}
	ctx.bFoundAMatch = true
	env.Cursor = env.LimitBackward
	return true
}

// removeQuestionSuffixes strips a trailing interrogative suffix.
func removeQuestionSuffixes(env *snowball.Env, ctx *context) bool {
	if !hasMinLength(env, ctx) {
		return false
	}
	ctx.bFoundAMatch = false
	env.LimitBackward = env.Cursor
	env.Cursor = env.Limit
	v1 := env.Limit - env.Cursor
lab0:
	for {
		env.Ket = env.Cursor
		if snowball.FindAmongB(env, a14, ctx) == 0 {
			break lab0
		}
		env.Bra = env.Cursor
		if !env.SliceFrom("்") {
			return false
		}
		ctx.bFoundAMatch = true
		break lab0
	}
	env.Cursor = env.Limit - v1
	env.Cursor = env.LimitBackward
	v2 := env.Cursor
lab1:
	for {
		if !fixEndings(env, ctx) {
			break lab1
		}
		break lab1
	}
	env.Cursor = v2
	return true
}

// removeCommandSuffixes strips a trailing imperative suffix.
func removeCommandSuffixes(env *snowball.Env, ctx *context) bool {
	if !hasMinLength(env, ctx) {
		return false
	}
	ctx.bFoundAMatch = false
	env.LimitBackward = env.Cursor
	env.Cursor = env.Limit
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a15, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !env.SliceDel() {
		return false
	}
	ctx.bFoundAMatch = true
	env.Cursor = env.LimitBackward
	return true
}

// removeUm strips a trailing -um suffix and repairs the resulting ending.
func removeUm(env *snowball.Env, ctx *context) bool {
	ctx.bFoundAMatch = false
	if !hasMinLength(env, ctx) {
		return false
	}
	env.LimitBackward = env.Cursor
	env.Cursor = env.Limit
	env.Ket = env.Cursor
	if !env.EqSB("ும்") {
		return false
	}
	env.Bra = env.Cursor
	if !env.SliceFrom("்") {
		return false
	}
	ctx.bFoundAMatch = true
	env.Cursor = env.LimitBackward
	v1 := env.Cursor
lab0:
	for {
		if !fixEnding(env, ctx) {
			break lab0
		}
		break lab0
	}
	env.Cursor = v1
	return true
}

// removeCommonWordEndings strips frequently-occurring postpositional endings.
func removeCommonWordEndings(env *snowball.Env, ctx *context) bool {
	ctx.bFoundAMatch = false
	if !hasMinLength(env, ctx) {
		return false
	}
	env.LimitBackward = env.Cursor
	env.Cursor = env.Limit
lab0:
	for {
		v1 := env.Limit - env.Cursor
	lab1:
		for {
			v2 := env.Limit - env.Cursor
			env.Ket = env.Cursor
		lab2:
			for {
				v3 := env.Limit - env.Cursor
			lab3:
				for {
					if !env.EqSB("ுடன்") {
						break lab3
					}
					break lab2
				}
				env.Cursor = env.Limit - v3
			lab4:
				for {
					if !env.EqSB("ில்லை") {
						break lab4
					}
					break lab2
				}
				env.Cursor = env.Limit - v3
			lab5:
				for {
					if !env.EqSB("ிடம்") {
						break lab5
					}
					break lab2
				}
				env.Cursor = env.Limit - v3
			lab6:
				for {
					if !env.EqSB("ின்றி") {
						break lab6
					}
					break lab2
				}
				env.Cursor = env.Limit - v3
			lab7:
				for {
					if !env.EqSB("ாகி") {
						break lab7
					}
					break lab2
				}
				env.Cursor = env.Limit - v3
			lab8:
				for {
					if !env.EqSB("ாகிய") {
						break lab8
					}
					break lab2
				}
				env.Cursor = env.Limit - v3
			lab9:
				for {
					if !env.EqSB("ென்று") {
						break lab9
					}
					break lab2
				}
				env.Cursor = env.Limit - v3
			lab10:
				for {
					if !env.EqSB("ுள்ள") {
						break lab10
					}
					break lab2
				}
				env.Cursor = env.Limit - v3
			lab11:
				for {
					if !env.EqSB("ுடைய") {
						break lab11
					}
					break lab2
				}
				env.Cursor = env.Limit - v3
			lab12:
				for {
					if !env.EqSB("ுடை") {
						break lab12
					}
					break lab2
				}
				env.Cursor = env.Limit - v3
			lab13:
				for {
					if !env.EqSB("ெனும்") {
						break lab13
					}
					break lab2
				}
				env.Cursor = env.Limit - v3
			lab14:
				for {
					if !env.EqSB("ல்ல") {
						break lab14
					}
					v4 := env.Limit - env.Cursor
					v5 := env.Limit - env.Cursor
				lab15:
					for {
						if snowball.FindAmongB(env, a16, ctx) == 0 {
							break lab15
						}
						break lab14
					}
					env.Cursor = env.Limit - v5
					env.Cursor = env.Limit - v4
					break lab2
				}
				env.Cursor = env.Limit - v3
			lab16:
				for {
					if !env.EqSB("ென") {
						break lab16
					}
					break lab2
				}
				env.Cursor = env.Limit - v3
				if !env.EqSB("ாகி") {
					break lab1
				}
				break lab2
			}
			env.Bra = env.Cursor
			if !env.SliceFrom("்") {
				return false
			}
			ctx.bFoundAMatch = true
			env.Cursor = env.Limit - v2
			break lab0
		}
		env.Cursor = env.Limit - v1
		v6 := env.Limit - env.Cursor
		env.Ket = env.Cursor
		if snowball.FindAmongB(env, a17, ctx) == 0 {
			return false
		}
		env.Bra = env.Cursor
		if !env.SliceDel() {
			return false
		}
		ctx.bFoundAMatch = true
		env.Cursor = env.Limit - v6
		break lab0
	}
	env.Cursor = env.LimitBackward
	v7 := env.Cursor
lab17:
	for {
		if !fixEndings(env, ctx) {
			break lab17
		}
		break lab17
	}
	env.Cursor = v7
	return true
}

// removeVetrumaiUrupukal strips trailing case (vetrumai urupu) markers.
func removeVetrumaiUrupukal(env *snowball.Env, ctx *context) bool {
	ctx.bFoundAMatch = false
	ctx.bFoundVetrumaiUrupu = false
	if !hasMinLength(env, ctx) {
		return false
	}
	env.LimitBackward = env.Cursor
	env.Cursor = env.Limit
lab0:
	for {
		v1 := env.Limit - env.Cursor
	lab1:
		for {
			v2 := env.Limit - env.Cursor
			env.Ket = env.Cursor
			if !env.EqSB("னை") {
				break lab1
			}
			env.Bra = env.Cursor
			if !env.SliceDel() {
				return false
			}
			env.Cursor = env.Limit - v2
			break lab0
		}
		env.Cursor = env.Limit - v1
	lab2:
		for {
			v3 := env.Limit - env.Cursor
			env.Ket = env.Cursor
		lab3:
			for {
				v4 := env.Limit - env.Cursor
			lab4:
				for {
				lab5:
					for {
						v5 := env.Limit - env.Cursor
					lab6:
						for {
							if !env.EqSB("ினை") {
								break lab6
							}
							break lab5
						}
						env.Cursor = env.Limit - v5
						if !env.EqSB("ை") {
							break lab4
						}
						break lab5
					}
					v6 := env.Limit - env.Cursor
					v7 := env.Limit - env.Cursor
				lab7:
					for {
						if snowball.FindAmongB(env, a18, ctx) == 0 {
							break lab7
						}
						break lab4
					}
					env.Cursor = env.Limit - v7
					env.Cursor = env.Limit - v6
					break lab3
				}
				env.Cursor = env.Limit - v4
				if !env.EqSB("ை") {
					break lab2
				}
				v8 := env.Limit - env.Cursor
				if snowball.FindAmongB(env, a19, ctx) == 0 {
					break lab2
				}
				if !env.EqSB("்") {
					break lab2
				}
				env.Cursor = env.Limit - v8
				break lab3
			}
			env.Bra = env.Cursor
			if !env.SliceFrom("்") {
				return false
			}
			env.Cursor = env.Limit - v3
			break lab0
		}
		env.Cursor = env.Limit - v1
	lab8:
		for {
			v9 := env.Limit - env.Cursor
			env.Ket = env.Cursor
		lab9:
			for {
				v10 := env.Limit - env.Cursor
			lab10:
				for {
					if !env.EqSB("ொடு") {
						break lab10
					}
					break lab9
				}
				env.Cursor = env.Limit - v10
			lab11:
				for {
					if !env.EqSB("ோடு") {
						break lab11
					}
					break lab9
				}
				env.Cursor = env.Limit - v10
			lab12:
				for {
					if !env.EqSB("ில்") {
						break lab12
					}
					break lab9
				}
				env.Cursor = env.Limit - v10
			lab13:
				for {
					if !env.EqSB("ிற்") {
						break lab13
					}
					break lab9
				}
				env.Cursor = env.Limit - v10
			lab14:
				for {
					if !env.EqSB("ின்") {
						break lab14
					}
					v11 := env.Limit - env.Cursor
					v12 := env.Limit - env.Cursor
				lab15:
					for {
						if !env.EqSB("ம") {
							break lab15
						}
						break lab14
					}
					env.Cursor = env.Limit - v12
					env.Cursor = env.Limit - v11
					break lab9
				}
				env.Cursor = env.Limit - v10
			lab16:
				for {
					if !env.EqSB("ின்று") {
						break lab16
					}
					break lab9
				}
				env.Cursor = env.Limit - v10
			lab17:
				for {
					if !env.EqSB("ிருந்து") {
						break lab17
					}
					break lab9
				}
				env.Cursor = env.Limit - v10
			lab18:
				for {
					if !env.EqSB("விட") {
						break lab18
					}
					break lab9
				}
				env.Cursor = env.Limit - v10
			lab19:
				for {
					if !(ctx.iLength >= 7) {
						break lab19
					}
					if !env.EqSB("ிடம்") {
						break lab19
					}
					break lab9
				}
				env.Cursor = env.Limit - v10
			lab20:
				for {
					if !env.EqSB("ால்") {
						break lab20
					}
					break lab9
				}
				env.Cursor = env.Limit - v10
			lab21:
				for {
					if !env.EqSB("ுடை") {
						break lab21
					}
					break lab9
				}
				env.Cursor = env.Limit - v10
			lab22:
				for {
					if !env.EqSB("ாமல்") {
						break lab22
					}
					break lab9
				}
				env.Cursor = env.Limit - v10
			lab23:
				for {
					if !env.EqSB("ல்") {
						break lab23
					}
					v13 := env.Limit - env.Cursor
					v14 := env.Limit - env.Cursor
				lab24:
					for {
						if snowball.FindAmongB(env, a20, ctx) == 0 {
							break lab24
						}
						break lab23
					}
					env.Cursor = env.Limit - v14
					env.Cursor = env.Limit - v13
					break lab9
				}
				env.Cursor = env.Limit - v10
				if !env.EqSB("ுள்") {
					break lab8
				}
				break lab9
			}
			env.Bra = env.Cursor
			if !env.SliceFrom("்") {
				return false
			}
			env.Cursor = env.Limit - v9
			break lab0
		}
		env.Cursor = env.Limit - v1
	lab25:
		for {
			v15 := env.Limit - env.Cursor
			env.Ket = env.Cursor
		lab26:
			for {
				v16 := env.Limit - env.Cursor
			lab27:
				for {
					if !env.EqSB("கண்") {
						break lab27
					}
					break lab26
				}
				env.Cursor = env.Limit - v16
			lab28:
				for {
					if !env.EqSB("முன்") {
						break lab28
					}
					break lab26
				}
				env.Cursor = env.Limit - v16
			lab29:
				for {
					if !env.EqSB("மேல்") {
						break lab29
					}
					break lab26
				}
				env.Cursor = env.Limit - v16
			lab30:
				for {
					if !env.EqSB("மேற்") {
						break lab30
					}
					break lab26
				}
				env.Cursor = env.Limit - v16
			lab31:
				for {
					if !env.EqSB("கீழ்") {
						break lab31
					}
					break lab26
				}
				env.Cursor = env.Limit - v16
			lab32:
				for {
					if !env.EqSB("பின்") {
						break lab32
					}
					break lab26
				}
				env.Cursor = env.Limit - v16
				if !env.EqSB("து") {
					break lab25
				}
				v17 := env.Limit - env.Cursor
				v18 := env.Limit - env.Cursor
			lab33:
				for {
					if snowball.FindAmongB(env, a21, ctx) == 0 {
						break lab33
					}
					break lab25
				}
				env.Cursor = env.Limit - v18
				env.Cursor = env.Limit - v17
				break lab26
			}
			env.Bra = env.Cursor
			if !env.SliceDel() {
				return false
			}
			env.Cursor = env.Limit - v15
			break lab0
		}
		env.Cursor = env.Limit - v1
		v19 := env.Limit - env.Cursor
		env.Ket = env.Cursor
		if !env.EqSB("ீ") {
			return false
		}
		env.Bra = env.Cursor
		if !env.SliceFrom("ி") {
			return false
		}
		env.Cursor = env.Limit - v19
		break lab0
	}
	ctx.bFoundAMatch = true
	ctx.bFoundVetrumaiUrupu = true
	v20 := env.Limit - env.Cursor
lab34:
	for {
		env.Ket = env.Cursor
		if !env.EqSB("ின்") {
			break lab34
		}
		env.Bra = env.Cursor
		if !env.SliceFrom("்") {
			return false
		}
		break lab34
	}
	env.Cursor = env.Limit - v20
	env.Cursor = env.LimitBackward
	v21 := env.Cursor
lab35:
	for {
		if !fixEndings(env, ctx) {
			break lab35
		}
		break lab35
	}
	env.Cursor = v21
	return true
}

// removeTenseSuffixes repeatedly applies removeTenseSuffix while it keeps
// matching.
func removeTenseSuffixes(env *snowball.Env, ctx *context) bool {
	ctx.bFoundAMatch = true
replab0:
	for {
		v1 := env.Cursor
	lab1:
		for once := 0; once < 1; once++ {
			if !ctx.bFoundAMatch {
				break lab1
			}
			v2 := env.Cursor
		lab2:
			for {
				if !removeTenseSuffix(env, ctx) {
					break lab2
				}
				break lab2
			}
			env.Cursor = v2
			continue replab0
		}
		env.Cursor = v1
		break replab0
	}
	return true
}

// removeTenseSuffix strips a single trailing tense/person verb suffix.
func removeTenseSuffix(env *snowball.Env, ctx *context) bool {
	ctx.bFoundAMatch = false
	if !hasMinLength(env, ctx) {
		return false
	}
	env.LimitBackward = env.Cursor
	env.Cursor = env.Limit
	v1 := env.Limit - env.Cursor
lab0:
	for {
	lab1:
		for {
			v2 := env.Limit - env.Cursor
		lab2:
			for {
				v3 := env.Limit - env.Cursor
				env.Ket = env.Cursor
				if snowball.FindAmongB(env, a22, ctx) == 0 {
					break lab2
				}
				env.Bra = env.Cursor
				if !env.SliceDel() {
					return false
				}
				ctx.bFoundAMatch = true
				env.Cursor = env.Limit - v3
				break lab1
			}
			env.Cursor = env.Limit - v2
		lab3:
			for {
				v4 := env.Limit - env.Cursor
				env.Ket = env.Cursor
			lab4:
				for {
					v5 := env.Limit - env.Cursor
				lab5:
					for {
						if !env.EqSB("மார்") {
							break lab5
						}
						break lab4
					}
					env.Cursor = env.Limit - v5
				lab6:
					for {
						if !env.EqSB("மின்") {
							break lab6
						}
						break lab4
					}
					env.Cursor = env.Limit - v5
				lab7:
					for {
						if !env.EqSB("னன்") {
							break lab7
						}
						break lab4
					}
					env.Cursor = env.Limit - v5
				lab8:
					for {
						if !env.EqSB("னான்") {
							break lab8
						}
						break lab4
					}
					env.Cursor = env.Limit - v5
				lab9:
					for {
						if !env.EqSB("னாள்") {
							break lab9
						}
						break lab4
					}
					env.Cursor = env.Limit - v5
				lab10:
					for {
						if !env.EqSB("னார்") {
							break lab10
						}
						break lab4
					}
					env.Cursor = env.Limit - v5
				lab11:
					for {
						if !env.EqSB("வன்") {
							break lab11
						}
						v6 := env.Limit - env.Cursor
						v7 := env.Limit - env.Cursor
					lab12:
						for {
							if snowball.FindAmongB(env, a23, ctx) == 0 {
								break lab12
							}
							break lab11
						}
						env.Cursor = env.Limit - v7
						env.Cursor = env.Limit - v6
						break lab4
					}
					env.Cursor = env.Limit - v5
				lab13:
					for {
						if !env.EqSB("னள்") {
							break lab13
						}
						break lab4
					}
					env.Cursor = env.Limit - v5
				lab14:
					for {
						if !env.EqSB("வள்") {
							break lab14
						}
						break lab4
					}
					env.Cursor = env.Limit - v5
				lab15:
					for {
						if !env.EqSB("னர்") {
							break lab15
						}
						break lab4
					}
					env.Cursor = env.Limit - v5
				lab16:
					for {
						if !env.EqSB("வர்") {
							break lab16
						}
						break lab4
					}
					env.Cursor = env.Limit - v5
				lab17:
					for {
						if !env.EqSB("ன") {
							break lab17
						}
						break lab4
					}
					env.Cursor = env.Limit - v5
				lab18:
					for {
						if !env.EqSB("ப") {
							break lab18
						}
						break lab4
					}
					env.Cursor = env.Limit - v5
				lab19:
					for {
						if !env.EqSB("க") {
							break lab19
						}
						break lab4
					}
					env.Cursor = env.Limit - v5
				lab20:
					for {
						if !env.EqSB("த") {
							break lab20
						}
						break lab4
					}
					env.Cursor = env.Limit - v5
				lab21:
					for {
						if !env.EqSB("ய") {
							break lab21
						}
						break lab4
					}
					env.Cursor = env.Limit - v5
				lab22:
					for {
						if !env.EqSB("பன்") {
							break lab22
						}
						break lab4
					}
					env.Cursor = env.Limit - v5
				lab23:
					for {
						if !env.EqSB("பள்") {
							break lab23
						}
						break lab4
					}
					env.Cursor = env.Limit - v5
				lab24:
					for {
						if !env.EqSB("பர்") {
							break lab24
						}
						break lab4
					}
					env.Cursor = env.Limit - v5
				lab25:
					for {
						if !env.EqSB("து") {
							break lab25
						}
						v8 := env.Limit - env.Cursor
						v9 := env.Limit - env.Cursor
					lab26:
						for {
							if snowball.FindAmongB(env, a24, ctx) == 0 {
								break lab26
							}
							break lab25
						}
						env.Cursor = env.Limit - v9
						env.Cursor = env.Limit - v8
						break lab4
					}
					env.Cursor = env.Limit - v5
				lab27:
					for {
						if !env.EqSB("ிற்று") {
							break lab27
						}
						break lab4
					}
					env.Cursor = env.Limit - v5
				lab28:
					for {
						if !env.EqSB("பம்") {
							break lab28
						}
						break lab4
					}
					env.Cursor = env.Limit - v5
				lab29:
					for {
						if !env.EqSB("னம்") {
							break lab29
						}
						break lab4
					}
					env.Cursor = env.Limit - v5
				lab30:
					for {
						if !env.EqSB("தும்") {
							break lab30
						}
						break lab4
					}
					env.Cursor = env.Limit - v5
				lab31:
					for {
						if !env.EqSB("றும்") {
							break lab31
						}
						break lab4
					}
					env.Cursor = env.Limit - v5
				lab32:
					for {
						if !env.EqSB("கும்") {
							break lab32
						}
						break lab4
					}
					env.Cursor = env.Limit - v5
				lab33:
					for {
						if !env.EqSB("னென்") {
							break lab33
						}
						break lab4
					}
					env.Cursor = env.Limit - v5
				lab34:
					for {
						if !env.EqSB("னை") {
							break lab34
						}
						break lab4
					}
					env.Cursor = env.Limit - v5
					if !env.EqSB("வை") {
						break lab3
					}
					break lab4
				}
				env.Bra = env.Cursor
				if !env.SliceDel() {
					return false
				}
				ctx.bFoundAMatch = true
				env.Cursor = env.Limit - v4
				break lab1
			}
			env.Cursor = env.Limit - v2
		lab35:
			for {
				v10 := env.Limit - env.Cursor
				env.Ket = env.Cursor
			lab36:
				for {
					v11 := env.Limit - env.Cursor
				lab37:
					for {
						if !env.EqSB("ான்") {
							break lab37
						}
						v12 := env.Limit - env.Cursor
						v13 := env.Limit - env.Cursor
					lab38:
						for {
							if !env.EqSB("ச") {
								break lab38
							}
							break lab37
						}
						env.Cursor = env.Limit - v13
						env.Cursor = env.Limit - v12
						break lab36
					}
					env.Cursor = env.Limit - v11
				lab39:
					for {
						if !env.EqSB("ாள்") {
							break lab39
						}
						break lab36
					}
					env.Cursor = env.Limit - v11
				lab40:
					for {
						if !env.EqSB("ார்") {
							break lab40
						}
						break lab36
					}
					env.Cursor = env.Limit - v11
				lab41:
					for {
						if !env.EqSB("ேன்") {
							break lab41
						}
						break lab36
					}
					env.Cursor = env.Limit - v11
				lab42:
					for {
						if !env.EqSB("ா") {
							break lab42
						}
						break lab36
					}
					env.Cursor = env.Limit - v11
				lab43:
					for {
						if !env.EqSB("ாம்") {
							break lab43
						}
						break lab36
					}
					env.Cursor = env.Limit - v11
				lab44:
					for {
						if !env.EqSB("ெம்") {
							break lab44
						}
						break lab36
					}
					env.Cursor = env.Limit - v11
				lab45:
					for {
						if !env.EqSB("ேம்") {
							break lab45
						}
						break lab36
					}
					env.Cursor = env.Limit - v11
				lab46:
					for {
						if !env.EqSB("ோம்") {
							break lab46
						}
						break lab36
					}
					env.Cursor = env.Limit - v11
				lab47:
					for {
						if !env.EqSB("கும்") {
							break lab47
						}
						break lab36
					}
					env.Cursor = env.Limit - v11
				lab48:
					for {
						if !env.EqSB("தும்") {
							break lab48
						}
						break lab36
					}
					env.Cursor = env.Limit - v11
				lab49:
					for {
						if !env.EqSB("டும்") {
							break lab49
						}
						break lab36
					}
					env.Cursor = env.Limit - v11
				lab50:
					for {
						if !env.EqSB("றும்") {
							break lab50
						}
						break lab36
					}
					env.Cursor = env.Limit - v11
				lab51:
					for {
						if !env.EqSB("ாய்") {
							break lab51
						}
						break lab36
					}
					env.Cursor = env.Limit - v11
				lab52:
					for {
						if !env.EqSB("னென்") {
							break lab52
						}
						break lab36
					}
					env.Cursor = env.Limit - v11
				lab53:
					for {
						if !env.EqSB("னிர்") {
							break lab53
						}
						break lab36
					}
					env.Cursor = env.Limit - v11
				lab54:
					for {
						if !env.EqSB("ீர்") {
							break lab54
						}
						break lab36
					}
					env.Cursor = env.Limit - v11
					if !env.EqSB("ீயர்") {
						break lab35
					}
					break lab36
				}
				env.Bra = env.Cursor
				if !env.SliceFrom("்") {
					return false
				}
				ctx.bFoundAMatch = true
				env.Cursor = env.Limit - v10
				break lab1
			}
			env.Cursor = env.Limit - v2
			v14 := env.Limit - env.Cursor
			env.Ket = env.Cursor
		lab55:
			for {
				v15 := env.Limit - env.Cursor
			lab56:
				for {
					if !env.EqSB("கு") {
						break lab56
					}
					break lab55
				}
				env.Cursor = env.Limit - v15
				if !env.EqSB("து") {
					break lab0
				}
				break lab55
			}
			v16 := env.Limit - env.Cursor
			if !env.EqSB("்") {
				break lab0
			}
			env.Cursor = env.Limit - v16
			env.Bra = env.Cursor
			if !env.SliceDel() {
				return false
			}
			ctx.bFoundAMatch = true
			env.Cursor = env.Limit - v14
			break lab1
		}
		break lab0
	}
	env.Cursor = env.Limit - v1
	v17 := env.Limit - env.Cursor
lab57:
	for {
		env.Ket = env.Cursor
		if snowball.FindAmongB(env, a25, ctx) == 0 {
			break lab57
		}
		env.Bra = env.Cursor
		if !env.SliceDel() {
			return false
		}
		ctx.bFoundAMatch = true
		break lab57
	}
	env.Cursor = env.Limit - v17
	env.Cursor = env.LimitBackward
	v18 := env.Cursor
lab58:
	for {
		if !fixEndings(env, ctx) {
			break lab58
		}
		break lab58
	}
	env.Cursor = v18
	return true
}

// Stem runs the Snowball tamil algorithm over env, mirroring the generated
// `stem` entry point. It returns true unless the word is too short.
func Stem(env *snowball.Env) bool {
	ctx, _ := env.Scratch.(*context)
	if ctx == nil {
		ctx = &context{}
		env.Scratch = ctx
	}
	*ctx = context{}
	ctx.bFoundVetrumaiUrupu = false
	v1 := env.Cursor
lab0:
	for {
		if !fixEnding(env, ctx) {
			break lab0
		}
		break lab0
	}
	env.Cursor = v1
	if !hasMinLength(env, ctx) {
		return false
	}
	v2 := env.Cursor
lab1:
	for {
		if !removeQuestionPrefixes(env, ctx) {
			break lab1
		}
		break lab1
	}
	env.Cursor = v2
	v3 := env.Cursor
lab2:
	for {
		if !removePronounPrefixes(env, ctx) {
			break lab2
		}
		break lab2
	}
	env.Cursor = v3
	v4 := env.Cursor
lab3:
	for {
		if !removeQuestionSuffixes(env, ctx) {
			break lab3
		}
		break lab3
	}
	env.Cursor = v4
	v5 := env.Cursor
lab4:
	for {
		if !removeUm(env, ctx) {
			break lab4
		}
		break lab4
	}
	env.Cursor = v5
	v6 := env.Cursor
lab5:
	for {
		if !removeCommonWordEndings(env, ctx) {
			break lab5
		}
		break lab5
	}
	env.Cursor = v6
	v7 := env.Cursor
lab6:
	for {
		if !removeVetrumaiUrupukal(env, ctx) {
			break lab6
		}
		break lab6
	}
	env.Cursor = v7
	v8 := env.Cursor
lab7:
	for {
		if !removePluralSuffix(env, ctx) {
			break lab7
		}
		break lab7
	}
	env.Cursor = v8
	v9 := env.Cursor
lab8:
	for {
		if !removeCommandSuffixes(env, ctx) {
			break lab8
		}
		break lab8
	}
	env.Cursor = v9
	v10 := env.Cursor
lab9:
	for {
		if !removeTenseSuffixes(env, ctx) {
			break lab9
		}
		break lab9
	}
	env.Cursor = v10
	return true
}
