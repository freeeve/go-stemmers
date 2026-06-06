// Package turkish is a byte-faithful Go port of rust-stemmers' generated
// Snowball "turkish" stemmer. It produces output identical to rust-stemmers
// 1.2.0's Turkish algorithm; the canonical Snowball turkish vocabulary is the
// conformance oracle.
//
// The control flow mirrors the generated Rust: each Snowball labelled block
// (`'lab: loop { … break 'lab }`) becomes a Go labelled `for { … break LABEL }`,
// the "goto" blocks become nested labelled `for` loops, and the result-code
// dispatch (`else if among_var == N`) becomes a `switch`.
package turkish

import "github.com/freeeve/go-stemmers/internal/snowball"

// context holds the per-run state the generated algorithm threads through its
// routines: whether noun-suffix stemming should continue and a scratch length.
type context struct {
	bContinueStemmingNounSuffixes bool
	iStrlen                       int
}

var a0 = []snowball.Among[context]{
	{Str: "m", SubstringI: -1, Result: -1},
	{Str: "n", SubstringI: -1, Result: -1},
	{Str: "miz", SubstringI: -1, Result: -1},
	{Str: "niz", SubstringI: -1, Result: -1},
	{Str: "muz", SubstringI: -1, Result: -1},
	{Str: "nuz", SubstringI: -1, Result: -1},
	{Str: "mız", SubstringI: -1, Result: -1},
	{Str: "nız", SubstringI: -1, Result: -1},
	{Str: "müz", SubstringI: -1, Result: -1},
	{Str: "nüz", SubstringI: -1, Result: -1},
}

var a1 = []snowball.Among[context]{
	{Str: "leri", SubstringI: -1, Result: -1},
	{Str: "ları", SubstringI: -1, Result: -1},
}

var a2 = []snowball.Among[context]{
	{Str: "ni", SubstringI: -1, Result: -1},
	{Str: "nu", SubstringI: -1, Result: -1},
	{Str: "nı", SubstringI: -1, Result: -1},
	{Str: "nü", SubstringI: -1, Result: -1},
}

var a3 = []snowball.Among[context]{
	{Str: "in", SubstringI: -1, Result: -1},
	{Str: "un", SubstringI: -1, Result: -1},
	{Str: "ın", SubstringI: -1, Result: -1},
	{Str: "ün", SubstringI: -1, Result: -1},
}

var a4 = []snowball.Among[context]{
	{Str: "a", SubstringI: -1, Result: -1},
	{Str: "e", SubstringI: -1, Result: -1},
}

var a5 = []snowball.Among[context]{
	{Str: "na", SubstringI: -1, Result: -1},
	{Str: "ne", SubstringI: -1, Result: -1},
}

var a6 = []snowball.Among[context]{
	{Str: "da", SubstringI: -1, Result: -1},
	{Str: "ta", SubstringI: -1, Result: -1},
	{Str: "de", SubstringI: -1, Result: -1},
	{Str: "te", SubstringI: -1, Result: -1},
}

var a7 = []snowball.Among[context]{
	{Str: "nda", SubstringI: -1, Result: -1},
	{Str: "nde", SubstringI: -1, Result: -1},
}

var a8 = []snowball.Among[context]{
	{Str: "dan", SubstringI: -1, Result: -1},
	{Str: "tan", SubstringI: -1, Result: -1},
	{Str: "den", SubstringI: -1, Result: -1},
	{Str: "ten", SubstringI: -1, Result: -1},
}

var a9 = []snowball.Among[context]{
	{Str: "ndan", SubstringI: -1, Result: -1},
	{Str: "nden", SubstringI: -1, Result: -1},
}

var a10 = []snowball.Among[context]{
	{Str: "la", SubstringI: -1, Result: -1},
	{Str: "le", SubstringI: -1, Result: -1},
}

var a11 = []snowball.Among[context]{
	{Str: "ca", SubstringI: -1, Result: -1},
	{Str: "ce", SubstringI: -1, Result: -1},
}

var a12 = []snowball.Among[context]{
	{Str: "im", SubstringI: -1, Result: -1},
	{Str: "um", SubstringI: -1, Result: -1},
	{Str: "ım", SubstringI: -1, Result: -1},
	{Str: "üm", SubstringI: -1, Result: -1},
}

var a13 = []snowball.Among[context]{
	{Str: "sin", SubstringI: -1, Result: -1},
	{Str: "sun", SubstringI: -1, Result: -1},
	{Str: "sın", SubstringI: -1, Result: -1},
	{Str: "sün", SubstringI: -1, Result: -1},
}

var a14 = []snowball.Among[context]{
	{Str: "iz", SubstringI: -1, Result: -1},
	{Str: "uz", SubstringI: -1, Result: -1},
	{Str: "ız", SubstringI: -1, Result: -1},
	{Str: "üz", SubstringI: -1, Result: -1},
}

var a15 = []snowball.Among[context]{
	{Str: "siniz", SubstringI: -1, Result: -1},
	{Str: "sunuz", SubstringI: -1, Result: -1},
	{Str: "sınız", SubstringI: -1, Result: -1},
	{Str: "sünüz", SubstringI: -1, Result: -1},
}

var a16 = []snowball.Among[context]{
	{Str: "lar", SubstringI: -1, Result: -1},
	{Str: "ler", SubstringI: -1, Result: -1},
}

var a17 = []snowball.Among[context]{
	{Str: "niz", SubstringI: -1, Result: -1},
	{Str: "nuz", SubstringI: -1, Result: -1},
	{Str: "nız", SubstringI: -1, Result: -1},
	{Str: "nüz", SubstringI: -1, Result: -1},
}

var a18 = []snowball.Among[context]{
	{Str: "dir", SubstringI: -1, Result: -1},
	{Str: "tir", SubstringI: -1, Result: -1},
	{Str: "dur", SubstringI: -1, Result: -1},
	{Str: "tur", SubstringI: -1, Result: -1},
	{Str: "dır", SubstringI: -1, Result: -1},
	{Str: "tır", SubstringI: -1, Result: -1},
	{Str: "dür", SubstringI: -1, Result: -1},
	{Str: "tür", SubstringI: -1, Result: -1},
}

var a19 = []snowball.Among[context]{
	{Str: "casına", SubstringI: -1, Result: -1},
	{Str: "cesine", SubstringI: -1, Result: -1},
}

var a20 = []snowball.Among[context]{
	{Str: "di", SubstringI: -1, Result: -1},
	{Str: "ti", SubstringI: -1, Result: -1},
	{Str: "dik", SubstringI: -1, Result: -1},
	{Str: "tik", SubstringI: -1, Result: -1},
	{Str: "duk", SubstringI: -1, Result: -1},
	{Str: "tuk", SubstringI: -1, Result: -1},
	{Str: "dık", SubstringI: -1, Result: -1},
	{Str: "tık", SubstringI: -1, Result: -1},
	{Str: "dük", SubstringI: -1, Result: -1},
	{Str: "tük", SubstringI: -1, Result: -1},
	{Str: "dim", SubstringI: -1, Result: -1},
	{Str: "tim", SubstringI: -1, Result: -1},
	{Str: "dum", SubstringI: -1, Result: -1},
	{Str: "tum", SubstringI: -1, Result: -1},
	{Str: "dım", SubstringI: -1, Result: -1},
	{Str: "tım", SubstringI: -1, Result: -1},
	{Str: "düm", SubstringI: -1, Result: -1},
	{Str: "tüm", SubstringI: -1, Result: -1},
	{Str: "din", SubstringI: -1, Result: -1},
	{Str: "tin", SubstringI: -1, Result: -1},
	{Str: "dun", SubstringI: -1, Result: -1},
	{Str: "tun", SubstringI: -1, Result: -1},
	{Str: "dın", SubstringI: -1, Result: -1},
	{Str: "tın", SubstringI: -1, Result: -1},
	{Str: "dün", SubstringI: -1, Result: -1},
	{Str: "tün", SubstringI: -1, Result: -1},
	{Str: "du", SubstringI: -1, Result: -1},
	{Str: "tu", SubstringI: -1, Result: -1},
	{Str: "dı", SubstringI: -1, Result: -1},
	{Str: "tı", SubstringI: -1, Result: -1},
	{Str: "dü", SubstringI: -1, Result: -1},
	{Str: "tü", SubstringI: -1, Result: -1},
}

var a21 = []snowball.Among[context]{
	{Str: "sa", SubstringI: -1, Result: -1},
	{Str: "se", SubstringI: -1, Result: -1},
	{Str: "sak", SubstringI: -1, Result: -1},
	{Str: "sek", SubstringI: -1, Result: -1},
	{Str: "sam", SubstringI: -1, Result: -1},
	{Str: "sem", SubstringI: -1, Result: -1},
	{Str: "san", SubstringI: -1, Result: -1},
	{Str: "sen", SubstringI: -1, Result: -1},
}

var a22 = []snowball.Among[context]{
	{Str: "miş", SubstringI: -1, Result: -1},
	{Str: "muş", SubstringI: -1, Result: -1},
	{Str: "mış", SubstringI: -1, Result: -1},
	{Str: "müş", SubstringI: -1, Result: -1},
}

var a23 = []snowball.Among[context]{
	{Str: "b", SubstringI: -1, Result: 1},
	{Str: "c", SubstringI: -1, Result: 2},
	{Str: "d", SubstringI: -1, Result: 3},
	{Str: "ğ", SubstringI: -1, Result: 4},
}

var gVowel = []byte{17, 65, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 32, 8, 0, 0, 0, 0, 0, 0, 1}
var gU = []byte{1, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 1}
var gVowel1 = []byte{1, 64, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
var gVowel2 = []byte{17, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 130}
var gVowel3 = []byte{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
var gVowel4 = []byte{17}
var gVowel5 = []byte{65}
var gVowel6 = []byte{65}

// checkVowelHarmony verifies that the vowel immediately before the cursor is
// harmonically compatible with the preceding vowel of the stem.
func checkVowelHarmony(env *snowball.Env, ctx *context) bool {
	v1 := env.Limit - env.Cursor
golab0:
	for {
		v2 := env.Limit - env.Cursor
	lab1:
		for {
			if !env.InGroupingB(gVowel, 97, 305) {
				break lab1
			}
			env.Cursor = env.Limit - v2
			break golab0
		}
		env.Cursor = env.Limit - v2
		if env.Cursor <= env.LimitBackward {
			return false
		}
		env.PreviousChar()
	}
lab2:
	for {
		v3 := env.Limit - env.Cursor
	lab3:
		for {
			if !env.EqSB("a") {
				break lab3
			}
		golab4:
			for {
				v4 := env.Limit - env.Cursor
			lab5:
				for {
					if !env.InGroupingB(gVowel1, 97, 305) {
						break lab5
					}
					env.Cursor = env.Limit - v4
					break golab4
				}
				env.Cursor = env.Limit - v4
				if env.Cursor <= env.LimitBackward {
					break lab3
				}
				env.PreviousChar()
			}
			break lab2
		}
		env.Cursor = env.Limit - v3
	lab6:
		for {
			if !env.EqSB("e") {
				break lab6
			}
		golab7:
			for {
				v5 := env.Limit - env.Cursor
			lab8:
				for {
					if !env.InGroupingB(gVowel2, 101, 252) {
						break lab8
					}
					env.Cursor = env.Limit - v5
					break golab7
				}
				env.Cursor = env.Limit - v5
				if env.Cursor <= env.LimitBackward {
					break lab6
				}
				env.PreviousChar()
			}
			break lab2
		}
		env.Cursor = env.Limit - v3
	lab9:
		for {
			if !env.EqSB("ı") {
				break lab9
			}
		golab10:
			for {
				v6 := env.Limit - env.Cursor
			lab11:
				for {
					if !env.InGroupingB(gVowel3, 97, 305) {
						break lab11
					}
					env.Cursor = env.Limit - v6
					break golab10
				}
				env.Cursor = env.Limit - v6
				if env.Cursor <= env.LimitBackward {
					break lab9
				}
				env.PreviousChar()
			}
			break lab2
		}
		env.Cursor = env.Limit - v3
	lab12:
		for {
			if !env.EqSB("i") {
				break lab12
			}
		golab13:
			for {
				v7 := env.Limit - env.Cursor
			lab14:
				for {
					if !env.InGroupingB(gVowel4, 101, 105) {
						break lab14
					}
					env.Cursor = env.Limit - v7
					break golab13
				}
				env.Cursor = env.Limit - v7
				if env.Cursor <= env.LimitBackward {
					break lab12
				}
				env.PreviousChar()
			}
			break lab2
		}
		env.Cursor = env.Limit - v3
	lab15:
		for {
			if !env.EqSB("o") {
				break lab15
			}
		golab16:
			for {
				v8 := env.Limit - env.Cursor
			lab17:
				for {
					if !env.InGroupingB(gVowel5, 111, 117) {
						break lab17
					}
					env.Cursor = env.Limit - v8
					break golab16
				}
				env.Cursor = env.Limit - v8
				if env.Cursor <= env.LimitBackward {
					break lab15
				}
				env.PreviousChar()
			}
			break lab2
		}
		env.Cursor = env.Limit - v3
	lab18:
		for {
			if !env.EqSB("ö") {
				break lab18
			}
		golab19:
			for {
				v9 := env.Limit - env.Cursor
			lab20:
				for {
					if !env.InGroupingB(gVowel6, 246, 252) {
						break lab20
					}
					env.Cursor = env.Limit - v9
					break golab19
				}
				env.Cursor = env.Limit - v9
				if env.Cursor <= env.LimitBackward {
					break lab18
				}
				env.PreviousChar()
			}
			break lab2
		}
		env.Cursor = env.Limit - v3
	lab21:
		for {
			if !env.EqSB("u") {
				break lab21
			}
		golab22:
			for {
				v10 := env.Limit - env.Cursor
			lab23:
				for {
					if !env.InGroupingB(gVowel5, 111, 117) {
						break lab23
					}
					env.Cursor = env.Limit - v10
					break golab22
				}
				env.Cursor = env.Limit - v10
				if env.Cursor <= env.LimitBackward {
					break lab21
				}
				env.PreviousChar()
			}
			break lab2
		}
		env.Cursor = env.Limit - v3
		if !env.EqSB("ü") {
			return false
		}
	golab24:
		for {
			v11 := env.Limit - env.Cursor
		lab25:
			for {
				if !env.InGroupingB(gVowel6, 246, 252) {
					break lab25
				}
				env.Cursor = env.Limit - v11
				break golab24
			}
			env.Cursor = env.Limit - v11
			if env.Cursor <= env.LimitBackward {
				return false
			}
			env.PreviousChar()
		}
		break lab2
	}
	env.Cursor = env.Limit - v1
	return true
}

// markSuffixWithOptionalNConsonant matches a suffix that may or may not be
// preceded by an "n" buffer consonant.
func markSuffixWithOptionalNConsonant(env *snowball.Env, ctx *context) bool {
lab0:
	for {
		v1 := env.Limit - env.Cursor
	lab1:
		for {
			if !env.EqSB("n") {
				break lab1
			}
			v2 := env.Limit - env.Cursor
			if !env.InGroupingB(gVowel, 97, 305) {
				break lab1
			}
			env.Cursor = env.Limit - v2
			break lab0
		}
		env.Cursor = env.Limit - v1
		v3 := env.Limit - env.Cursor
	lab2:
		for {
			v4 := env.Limit - env.Cursor
			if !env.EqSB("n") {
				break lab2
			}
			env.Cursor = env.Limit - v4
			return false
		}
		env.Cursor = env.Limit - v3
		v5 := env.Limit - env.Cursor
		if env.Cursor <= env.LimitBackward {
			return false
		}
		env.PreviousChar()
		if !env.InGroupingB(gVowel, 97, 305) {
			return false
		}
		env.Cursor = env.Limit - v5
		break lab0
	}
	return true
}

// markSuffixWithOptionalSConsonant matches a suffix that may or may not be
// preceded by an "s" buffer consonant.
func markSuffixWithOptionalSConsonant(env *snowball.Env, ctx *context) bool {
lab0:
	for {
		v1 := env.Limit - env.Cursor
	lab1:
		for {
			if !env.EqSB("s") {
				break lab1
			}
			v2 := env.Limit - env.Cursor
			if !env.InGroupingB(gVowel, 97, 305) {
				break lab1
			}
			env.Cursor = env.Limit - v2
			break lab0
		}
		env.Cursor = env.Limit - v1
		v3 := env.Limit - env.Cursor
	lab2:
		for {
			v4 := env.Limit - env.Cursor
			if !env.EqSB("s") {
				break lab2
			}
			env.Cursor = env.Limit - v4
			return false
		}
		env.Cursor = env.Limit - v3
		v5 := env.Limit - env.Cursor
		if env.Cursor <= env.LimitBackward {
			return false
		}
		env.PreviousChar()
		if !env.InGroupingB(gVowel, 97, 305) {
			return false
		}
		env.Cursor = env.Limit - v5
		break lab0
	}
	return true
}

// markSuffixWithOptionalYConsonant matches a suffix that may or may not be
// preceded by a "y" buffer consonant.
func markSuffixWithOptionalYConsonant(env *snowball.Env, ctx *context) bool {
lab0:
	for {
		v1 := env.Limit - env.Cursor
	lab1:
		for {
			if !env.EqSB("y") {
				break lab1
			}
			v2 := env.Limit - env.Cursor
			if !env.InGroupingB(gVowel, 97, 305) {
				break lab1
			}
			env.Cursor = env.Limit - v2
			break lab0
		}
		env.Cursor = env.Limit - v1
		v3 := env.Limit - env.Cursor
	lab2:
		for {
			v4 := env.Limit - env.Cursor
			if !env.EqSB("y") {
				break lab2
			}
			env.Cursor = env.Limit - v4
			return false
		}
		env.Cursor = env.Limit - v3
		v5 := env.Limit - env.Cursor
		if env.Cursor <= env.LimitBackward {
			return false
		}
		env.PreviousChar()
		if !env.InGroupingB(gVowel, 97, 305) {
			return false
		}
		env.Cursor = env.Limit - v5
		break lab0
	}
	return true
}

// markSuffixWithOptionalUVowel matches a suffix that may or may not be preceded
// by a high vowel from the U group.
func markSuffixWithOptionalUVowel(env *snowball.Env, ctx *context) bool {
lab0:
	for {
		v1 := env.Limit - env.Cursor
	lab1:
		for {
			if !env.InGroupingB(gU, 105, 305) {
				break lab1
			}
			v2 := env.Limit - env.Cursor
			if !env.OutGroupingB(gVowel, 97, 305) {
				break lab1
			}
			env.Cursor = env.Limit - v2
			break lab0
		}
		env.Cursor = env.Limit - v1
		v3 := env.Limit - env.Cursor
	lab2:
		for {
			v4 := env.Limit - env.Cursor
			if !env.InGroupingB(gU, 105, 305) {
				break lab2
			}
			env.Cursor = env.Limit - v4
			return false
		}
		env.Cursor = env.Limit - v3
		v5 := env.Limit - env.Cursor
		if env.Cursor <= env.LimitBackward {
			return false
		}
		env.PreviousChar()
		if !env.OutGroupingB(gVowel, 97, 305) {
			return false
		}
		env.Cursor = env.Limit - v5
		break lab0
	}
	return true
}

// markPossessives marks a possessive suffix (-(U)m/-(U)n/-(U)mUz/-(U)nUz).
func markPossessives(env *snowball.Env, ctx *context) bool {
	if snowball.FindAmongB(env, a0, ctx) == 0 {
		return false
	}
	if !markSuffixWithOptionalUVowel(env, ctx) {
		return false
	}
	return true
}

// markSU marks the 3rd-person possessive -sU suffix.
func markSU(env *snowball.Env, ctx *context) bool {
	if !checkVowelHarmony(env, ctx) {
		return false
	}
	if !env.InGroupingB(gU, 105, 305) {
		return false
	}
	if !markSuffixWithOptionalSConsonant(env, ctx) {
		return false
	}
	return true
}

// markLArI marks the -lArI plural-possessive suffix.
func markLArI(env *snowball.Env, ctx *context) bool {
	if snowball.FindAmongB(env, a1, ctx) == 0 {
		return false
	}
	return true
}

// markYU marks the accusative -(y)U suffix.
func markYU(env *snowball.Env, ctx *context) bool {
	if !checkVowelHarmony(env, ctx) {
		return false
	}
	if !env.InGroupingB(gU, 105, 305) {
		return false
	}
	if !markSuffixWithOptionalYConsonant(env, ctx) {
		return false
	}
	return true
}

// markNU marks the -nU suffix.
func markNU(env *snowball.Env, ctx *context) bool {
	if !checkVowelHarmony(env, ctx) {
		return false
	}
	if snowball.FindAmongB(env, a2, ctx) == 0 {
		return false
	}
	return true
}

// markNUn marks the genitive -(n)Un suffix.
func markNUn(env *snowball.Env, ctx *context) bool {
	if !checkVowelHarmony(env, ctx) {
		return false
	}
	if snowball.FindAmongB(env, a3, ctx) == 0 {
		return false
	}
	if !markSuffixWithOptionalNConsonant(env, ctx) {
		return false
	}
	return true
}

// markYA marks the dative -(y)A suffix.
func markYA(env *snowball.Env, ctx *context) bool {
	if !checkVowelHarmony(env, ctx) {
		return false
	}
	if snowball.FindAmongB(env, a4, ctx) == 0 {
		return false
	}
	if !markSuffixWithOptionalYConsonant(env, ctx) {
		return false
	}
	return true
}

// markNA marks the -nA suffix.
func markNA(env *snowball.Env, ctx *context) bool {
	if !checkVowelHarmony(env, ctx) {
		return false
	}
	if snowball.FindAmongB(env, a5, ctx) == 0 {
		return false
	}
	return true
}

// markDA marks the locative -DA suffix.
func markDA(env *snowball.Env, ctx *context) bool {
	if !checkVowelHarmony(env, ctx) {
		return false
	}
	if snowball.FindAmongB(env, a6, ctx) == 0 {
		return false
	}
	return true
}

// markNdA marks the -ndA suffix.
func markNdA(env *snowball.Env, ctx *context) bool {
	if !checkVowelHarmony(env, ctx) {
		return false
	}
	if snowball.FindAmongB(env, a7, ctx) == 0 {
		return false
	}
	return true
}

// markDAn marks the ablative -DAn suffix.
func markDAn(env *snowball.Env, ctx *context) bool {
	if !checkVowelHarmony(env, ctx) {
		return false
	}
	if snowball.FindAmongB(env, a8, ctx) == 0 {
		return false
	}
	return true
}

// markNdAn marks the -ndAn suffix.
func markNdAn(env *snowball.Env, ctx *context) bool {
	if !checkVowelHarmony(env, ctx) {
		return false
	}
	if snowball.FindAmongB(env, a9, ctx) == 0 {
		return false
	}
	return true
}

// markYlA marks the instrumental -(y)lA suffix.
func markYlA(env *snowball.Env, ctx *context) bool {
	if !checkVowelHarmony(env, ctx) {
		return false
	}
	if snowball.FindAmongB(env, a10, ctx) == 0 {
		return false
	}
	if !markSuffixWithOptionalYConsonant(env, ctx) {
		return false
	}
	return true
}

// markKi marks the relative -ki suffix.
func markKi(env *snowball.Env, ctx *context) bool {
	if !env.EqSB("ki") {
		return false
	}
	return true
}

// markNcA marks the equative -ncA suffix.
func markNcA(env *snowball.Env, ctx *context) bool {
	if !checkVowelHarmony(env, ctx) {
		return false
	}
	if snowball.FindAmongB(env, a11, ctx) == 0 {
		return false
	}
	if !markSuffixWithOptionalNConsonant(env, ctx) {
		return false
	}
	return true
}

// markYUm marks the 1st-person-singular copular -(y)Um suffix.
func markYUm(env *snowball.Env, ctx *context) bool {
	if !checkVowelHarmony(env, ctx) {
		return false
	}
	if snowball.FindAmongB(env, a12, ctx) == 0 {
		return false
	}
	if !markSuffixWithOptionalYConsonant(env, ctx) {
		return false
	}
	return true
}

// markSUn marks the 2nd-person-singular copular -sUn suffix.
func markSUn(env *snowball.Env, ctx *context) bool {
	if !checkVowelHarmony(env, ctx) {
		return false
	}
	if snowball.FindAmongB(env, a13, ctx) == 0 {
		return false
	}
	return true
}

// markYUz marks the 1st-person-plural copular -(y)Uz suffix.
func markYUz(env *snowball.Env, ctx *context) bool {
	if !checkVowelHarmony(env, ctx) {
		return false
	}
	if snowball.FindAmongB(env, a14, ctx) == 0 {
		return false
	}
	if !markSuffixWithOptionalYConsonant(env, ctx) {
		return false
	}
	return true
}

// markSUnUz marks the 2nd-person-plural copular -sUnUz suffix.
func markSUnUz(env *snowball.Env, ctx *context) bool {
	if snowball.FindAmongB(env, a15, ctx) == 0 {
		return false
	}
	return true
}

// markLAr marks the plural -lAr suffix.
func markLAr(env *snowball.Env, ctx *context) bool {
	if !checkVowelHarmony(env, ctx) {
		return false
	}
	if snowball.FindAmongB(env, a16, ctx) == 0 {
		return false
	}
	return true
}

// markNUz marks the -nUz suffix.
func markNUz(env *snowball.Env, ctx *context) bool {
	if !checkVowelHarmony(env, ctx) {
		return false
	}
	if snowball.FindAmongB(env, a17, ctx) == 0 {
		return false
	}
	return true
}

// markDUr marks the copular -DUr suffix.
func markDUr(env *snowball.Env, ctx *context) bool {
	if !checkVowelHarmony(env, ctx) {
		return false
	}
	if snowball.FindAmongB(env, a18, ctx) == 0 {
		return false
	}
	return true
}

// markCAsInA marks the -cAsInA/-cesine adverbial suffix.
func markCAsInA(env *snowball.Env, ctx *context) bool {
	if snowball.FindAmongB(env, a19, ctx) == 0 {
		return false
	}
	return true
}

// markYDU marks the past-tense copular -(y)DU suffix family.
func markYDU(env *snowball.Env, ctx *context) bool {
	if !checkVowelHarmony(env, ctx) {
		return false
	}
	if snowball.FindAmongB(env, a20, ctx) == 0 {
		return false
	}
	if !markSuffixWithOptionalYConsonant(env, ctx) {
		return false
	}
	return true
}

// markYsA marks the conditional copular -(y)sA suffix family.
func markYsA(env *snowball.Env, ctx *context) bool {
	if snowball.FindAmongB(env, a21, ctx) == 0 {
		return false
	}
	if !markSuffixWithOptionalYConsonant(env, ctx) {
		return false
	}
	return true
}

// markYmUs marks the inferential copular -(y)mUş suffix family.
func markYmUs(env *snowball.Env, ctx *context) bool {
	if !checkVowelHarmony(env, ctx) {
		return false
	}
	if snowball.FindAmongB(env, a22, ctx) == 0 {
		return false
	}
	if !markSuffixWithOptionalYConsonant(env, ctx) {
		return false
	}
	return true
}

// markYken marks the -(y)ken adverbial suffix.
func markYken(env *snowball.Env, ctx *context) bool {
	if !env.EqSB("ken") {
		return false
	}
	if !markSuffixWithOptionalYConsonant(env, ctx) {
		return false
	}
	return true
}

// stemNominalVerbSuffixes removes copular/verbal person-and-tense suffixes,
// recording whether noun-suffix stemming should continue afterwards.
func stemNominalVerbSuffixes(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	ctx.bContinueStemmingNounSuffixes = true
lab0:
	for {
		v1 := env.Limit - env.Cursor
	lab1:
		for {
		lab2:
			for {
				v2 := env.Limit - env.Cursor
			lab3:
				for {
					if !markYmUs(env, ctx) {
						break lab3
					}
					break lab2
				}
				env.Cursor = env.Limit - v2
			lab4:
				for {
					if !markYDU(env, ctx) {
						break lab4
					}
					break lab2
				}
				env.Cursor = env.Limit - v2
			lab5:
				for {
					if !markYsA(env, ctx) {
						break lab5
					}
					break lab2
				}
				env.Cursor = env.Limit - v2
				if !markYken(env, ctx) {
					break lab1
				}
				break lab2
			}
			break lab0
		}
		env.Cursor = env.Limit - v1
	lab6:
		for {
			if !markCAsInA(env, ctx) {
				break lab6
			}
		lab7:
			for {
				v3 := env.Limit - env.Cursor
			lab8:
				for {
					if !markSUnUz(env, ctx) {
						break lab8
					}
					break lab7
				}
				env.Cursor = env.Limit - v3
			lab9:
				for {
					if !markLAr(env, ctx) {
						break lab9
					}
					break lab7
				}
				env.Cursor = env.Limit - v3
			lab10:
				for {
					if !markYUm(env, ctx) {
						break lab10
					}
					break lab7
				}
				env.Cursor = env.Limit - v3
			lab11:
				for {
					if !markSUn(env, ctx) {
						break lab11
					}
					break lab7
				}
				env.Cursor = env.Limit - v3
			lab12:
				for {
					if !markYUz(env, ctx) {
						break lab12
					}
					break lab7
				}
				env.Cursor = env.Limit - v3
				break lab7
			}
			if !markYmUs(env, ctx) {
				break lab6
			}
			break lab0
		}
		env.Cursor = env.Limit - v1
	lab13:
		for {
			if !markLAr(env, ctx) {
				break lab13
			}
			env.Bra = env.Cursor
			if !env.SliceDel() {
				return false
			}
			v4 := env.Limit - env.Cursor
		lab14:
			for {
				env.Ket = env.Cursor
			lab15:
				for {
					v5 := env.Limit - env.Cursor
				lab16:
					for {
						if !markDUr(env, ctx) {
							break lab16
						}
						break lab15
					}
					env.Cursor = env.Limit - v5
				lab17:
					for {
						if !markYDU(env, ctx) {
							break lab17
						}
						break lab15
					}
					env.Cursor = env.Limit - v5
				lab18:
					for {
						if !markYsA(env, ctx) {
							break lab18
						}
						break lab15
					}
					env.Cursor = env.Limit - v5
					if !markYmUs(env, ctx) {
						env.Cursor = env.Limit - v4
						break lab14
					}
					break lab15
				}
				break lab14
			}
			ctx.bContinueStemmingNounSuffixes = false
			break lab0
		}
		env.Cursor = env.Limit - v1
	lab19:
		for {
			if !markNUz(env, ctx) {
				break lab19
			}
		lab20:
			for {
				v6 := env.Limit - env.Cursor
			lab21:
				for {
					if !markYDU(env, ctx) {
						break lab21
					}
					break lab20
				}
				env.Cursor = env.Limit - v6
				if !markYsA(env, ctx) {
					break lab19
				}
				break lab20
			}
			break lab0
		}
		env.Cursor = env.Limit - v1
	lab22:
		for {
		lab23:
			for {
				v7 := env.Limit - env.Cursor
			lab24:
				for {
					if !markSUnUz(env, ctx) {
						break lab24
					}
					break lab23
				}
				env.Cursor = env.Limit - v7
			lab25:
				for {
					if !markYUz(env, ctx) {
						break lab25
					}
					break lab23
				}
				env.Cursor = env.Limit - v7
			lab26:
				for {
					if !markSUn(env, ctx) {
						break lab26
					}
					break lab23
				}
				env.Cursor = env.Limit - v7
				if !markYUm(env, ctx) {
					break lab22
				}
				break lab23
			}
			env.Bra = env.Cursor
			if !env.SliceDel() {
				return false
			}
			v8 := env.Limit - env.Cursor
		lab27:
			for {
				env.Ket = env.Cursor
				if !markYmUs(env, ctx) {
					env.Cursor = env.Limit - v8
					break lab27
				}
				break lab27
			}
			break lab0
		}
		env.Cursor = env.Limit - v1
		if !markDUr(env, ctx) {
			return false
		}
		env.Bra = env.Cursor
		if !env.SliceDel() {
			return false
		}
		v9 := env.Limit - env.Cursor
	lab28:
		for {
			env.Ket = env.Cursor
		lab29:
			for {
				v10 := env.Limit - env.Cursor
			lab30:
				for {
					if !markSUnUz(env, ctx) {
						break lab30
					}
					break lab29
				}
				env.Cursor = env.Limit - v10
			lab31:
				for {
					if !markLAr(env, ctx) {
						break lab31
					}
					break lab29
				}
				env.Cursor = env.Limit - v10
			lab32:
				for {
					if !markYUm(env, ctx) {
						break lab32
					}
					break lab29
				}
				env.Cursor = env.Limit - v10
			lab33:
				for {
					if !markSUn(env, ctx) {
						break lab33
					}
					break lab29
				}
				env.Cursor = env.Limit - v10
			lab34:
				for {
					if !markYUz(env, ctx) {
						break lab34
					}
					break lab29
				}
				env.Cursor = env.Limit - v10
				break lab29
			}
			if !markYmUs(env, ctx) {
				env.Cursor = env.Limit - v9
				break lab28
			}
			break lab28
		}
		break lab0
	}
	env.Bra = env.Cursor
	if !env.SliceDel() {
		return false
	}
	return true
}

// stemSuffixChainBeforeKi removes a chain of case/possessive suffixes that may
// precede a relative -ki, recursing to peel further -ki chains.
func stemSuffixChainBeforeKi(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	if !markKi(env, ctx) {
		return false
	}
lab0:
	for {
		v1 := env.Limit - env.Cursor
	lab1:
		for {
			if !markDA(env, ctx) {
				break lab1
			}
			env.Bra = env.Cursor
			if !env.SliceDel() {
				return false
			}
			v2 := env.Limit - env.Cursor
		lab2:
			for {
				env.Ket = env.Cursor
			lab3:
				for {
					v3 := env.Limit - env.Cursor
				lab4:
					for {
						if !markLAr(env, ctx) {
							break lab4
						}
						env.Bra = env.Cursor
						if !env.SliceDel() {
							return false
						}
						v4 := env.Limit - env.Cursor
					lab5:
						for {
							if !stemSuffixChainBeforeKi(env, ctx) {
								env.Cursor = env.Limit - v4
								break lab5
							}
							break lab5
						}
						break lab3
					}
					env.Cursor = env.Limit - v3
					if !markPossessives(env, ctx) {
						env.Cursor = env.Limit - v2
						break lab2
					}
					env.Bra = env.Cursor
					if !env.SliceDel() {
						return false
					}
					v5 := env.Limit - env.Cursor
				lab6:
					for {
						env.Ket = env.Cursor
						if !markLAr(env, ctx) {
							env.Cursor = env.Limit - v5
							break lab6
						}
						env.Bra = env.Cursor
						if !env.SliceDel() {
							return false
						}
						if !stemSuffixChainBeforeKi(env, ctx) {
							env.Cursor = env.Limit - v5
							break lab6
						}
						break lab6
					}
					break lab3
				}
				break lab2
			}
			break lab0
		}
		env.Cursor = env.Limit - v1
	lab7:
		for {
			if !markNUn(env, ctx) {
				break lab7
			}
			env.Bra = env.Cursor
			if !env.SliceDel() {
				return false
			}
			v6 := env.Limit - env.Cursor
		lab8:
			for {
				env.Ket = env.Cursor
			lab9:
				for {
					v7 := env.Limit - env.Cursor
				lab10:
					for {
						if !markLArI(env, ctx) {
							break lab10
						}
						env.Bra = env.Cursor
						if !env.SliceDel() {
							return false
						}
						break lab9
					}
					env.Cursor = env.Limit - v7
				lab11:
					for {
						env.Ket = env.Cursor
					lab12:
						for {
							v8 := env.Limit - env.Cursor
						lab13:
							for {
								if !markPossessives(env, ctx) {
									break lab13
								}
								break lab12
							}
							env.Cursor = env.Limit - v8
							if !markSU(env, ctx) {
								break lab11
							}
							break lab12
						}
						env.Bra = env.Cursor
						if !env.SliceDel() {
							return false
						}
						v9 := env.Limit - env.Cursor
					lab14:
						for {
							env.Ket = env.Cursor
							if !markLAr(env, ctx) {
								env.Cursor = env.Limit - v9
								break lab14
							}
							env.Bra = env.Cursor
							if !env.SliceDel() {
								return false
							}
							if !stemSuffixChainBeforeKi(env, ctx) {
								env.Cursor = env.Limit - v9
								break lab14
							}
							break lab14
						}
						break lab9
					}
					env.Cursor = env.Limit - v7
					if !stemSuffixChainBeforeKi(env, ctx) {
						env.Cursor = env.Limit - v6
						break lab8
					}
					break lab9
				}
				break lab8
			}
			break lab0
		}
		env.Cursor = env.Limit - v1
		if !markNdA(env, ctx) {
			return false
		}
	lab15:
		for {
			v10 := env.Limit - env.Cursor
		lab16:
			for {
				if !markLArI(env, ctx) {
					break lab16
				}
				env.Bra = env.Cursor
				if !env.SliceDel() {
					return false
				}
				break lab15
			}
			env.Cursor = env.Limit - v10
		lab17:
			for {
				if !markSU(env, ctx) {
					break lab17
				}
				env.Bra = env.Cursor
				if !env.SliceDel() {
					return false
				}
				v11 := env.Limit - env.Cursor
			lab18:
				for {
					env.Ket = env.Cursor
					if !markLAr(env, ctx) {
						env.Cursor = env.Limit - v11
						break lab18
					}
					env.Bra = env.Cursor
					if !env.SliceDel() {
						return false
					}
					if !stemSuffixChainBeforeKi(env, ctx) {
						env.Cursor = env.Limit - v11
						break lab18
					}
					break lab18
				}
				break lab15
			}
			env.Cursor = env.Limit - v10
			if !stemSuffixChainBeforeKi(env, ctx) {
				return false
			}
			break lab15
		}
		break lab0
	}
	return true
}

// stemNounSuffixes removes nominal case, possessive and plural suffixes,
// chaining into stemSuffixChainBeforeKi where a -ki may follow.
func stemNounSuffixes(env *snowball.Env, ctx *context) bool {
lab0:
	for {
		v1 := env.Limit - env.Cursor
	lab1:
		for {
			env.Ket = env.Cursor
			if !markLAr(env, ctx) {
				break lab1
			}
			env.Bra = env.Cursor
			if !env.SliceDel() {
				return false
			}
			v2 := env.Limit - env.Cursor
		lab2:
			for {
				if !stemSuffixChainBeforeKi(env, ctx) {
					env.Cursor = env.Limit - v2
					break lab2
				}
				break lab2
			}
			break lab0
		}
		env.Cursor = env.Limit - v1
	lab3:
		for {
			env.Ket = env.Cursor
			if !markNcA(env, ctx) {
				break lab3
			}
			env.Bra = env.Cursor
			if !env.SliceDel() {
				return false
			}
			v3 := env.Limit - env.Cursor
		lab4:
			for {
			lab5:
				for {
					v4 := env.Limit - env.Cursor
				lab6:
					for {
						env.Ket = env.Cursor
						if !markLArI(env, ctx) {
							break lab6
						}
						env.Bra = env.Cursor
						if !env.SliceDel() {
							return false
						}
						break lab5
					}
					env.Cursor = env.Limit - v4
				lab7:
					for {
						env.Ket = env.Cursor
					lab8:
						for {
							v5 := env.Limit - env.Cursor
						lab9:
							for {
								if !markPossessives(env, ctx) {
									break lab9
								}
								break lab8
							}
							env.Cursor = env.Limit - v5
							if !markSU(env, ctx) {
								break lab7
							}
							break lab8
						}
						env.Bra = env.Cursor
						if !env.SliceDel() {
							return false
						}
						v6 := env.Limit - env.Cursor
					lab10:
						for {
							env.Ket = env.Cursor
							if !markLAr(env, ctx) {
								env.Cursor = env.Limit - v6
								break lab10
							}
							env.Bra = env.Cursor
							if !env.SliceDel() {
								return false
							}
							if !stemSuffixChainBeforeKi(env, ctx) {
								env.Cursor = env.Limit - v6
								break lab10
							}
							break lab10
						}
						break lab5
					}
					env.Cursor = env.Limit - v4
					env.Ket = env.Cursor
					if !markLAr(env, ctx) {
						env.Cursor = env.Limit - v3
						break lab4
					}
					env.Bra = env.Cursor
					if !env.SliceDel() {
						return false
					}
					if !stemSuffixChainBeforeKi(env, ctx) {
						env.Cursor = env.Limit - v3
						break lab4
					}
					break lab5
				}
				break lab4
			}
			break lab0
		}
		env.Cursor = env.Limit - v1
	lab11:
		for {
			env.Ket = env.Cursor
		lab12:
			for {
				v7 := env.Limit - env.Cursor
			lab13:
				for {
					if !markNdA(env, ctx) {
						break lab13
					}
					break lab12
				}
				env.Cursor = env.Limit - v7
				if !markNA(env, ctx) {
					break lab11
				}
				break lab12
			}
		lab14:
			for {
				v8 := env.Limit - env.Cursor
			lab15:
				for {
					if !markLArI(env, ctx) {
						break lab15
					}
					env.Bra = env.Cursor
					if !env.SliceDel() {
						return false
					}
					break lab14
				}
				env.Cursor = env.Limit - v8
			lab16:
				for {
					if !markSU(env, ctx) {
						break lab16
					}
					env.Bra = env.Cursor
					if !env.SliceDel() {
						return false
					}
					v9 := env.Limit - env.Cursor
				lab17:
					for {
						env.Ket = env.Cursor
						if !markLAr(env, ctx) {
							env.Cursor = env.Limit - v9
							break lab17
						}
						env.Bra = env.Cursor
						if !env.SliceDel() {
							return false
						}
						if !stemSuffixChainBeforeKi(env, ctx) {
							env.Cursor = env.Limit - v9
							break lab17
						}
						break lab17
					}
					break lab14
				}
				env.Cursor = env.Limit - v8
				if !stemSuffixChainBeforeKi(env, ctx) {
					break lab11
				}
				break lab14
			}
			break lab0
		}
		env.Cursor = env.Limit - v1
	lab18:
		for {
			env.Ket = env.Cursor
		lab19:
			for {
				v10 := env.Limit - env.Cursor
			lab20:
				for {
					if !markNdAn(env, ctx) {
						break lab20
					}
					break lab19
				}
				env.Cursor = env.Limit - v10
				if !markNU(env, ctx) {
					break lab18
				}
				break lab19
			}
		lab21:
			for {
				v11 := env.Limit - env.Cursor
			lab22:
				for {
					if !markSU(env, ctx) {
						break lab22
					}
					env.Bra = env.Cursor
					if !env.SliceDel() {
						return false
					}
					v12 := env.Limit - env.Cursor
				lab23:
					for {
						env.Ket = env.Cursor
						if !markLAr(env, ctx) {
							env.Cursor = env.Limit - v12
							break lab23
						}
						env.Bra = env.Cursor
						if !env.SliceDel() {
							return false
						}
						if !stemSuffixChainBeforeKi(env, ctx) {
							env.Cursor = env.Limit - v12
							break lab23
						}
						break lab23
					}
					break lab21
				}
				env.Cursor = env.Limit - v11
				if !markLArI(env, ctx) {
					break lab18
				}
				break lab21
			}
			break lab0
		}
		env.Cursor = env.Limit - v1
	lab24:
		for {
			env.Ket = env.Cursor
			if !markDAn(env, ctx) {
				break lab24
			}
			env.Bra = env.Cursor
			if !env.SliceDel() {
				return false
			}
			v13 := env.Limit - env.Cursor
		lab25:
			for {
				env.Ket = env.Cursor
			lab26:
				for {
					v14 := env.Limit - env.Cursor
				lab27:
					for {
						if !markPossessives(env, ctx) {
							break lab27
						}
						env.Bra = env.Cursor
						if !env.SliceDel() {
							return false
						}
						v15 := env.Limit - env.Cursor
					lab28:
						for {
							env.Ket = env.Cursor
							if !markLAr(env, ctx) {
								env.Cursor = env.Limit - v15
								break lab28
							}
							env.Bra = env.Cursor
							if !env.SliceDel() {
								return false
							}
							if !stemSuffixChainBeforeKi(env, ctx) {
								env.Cursor = env.Limit - v15
								break lab28
							}
							break lab28
						}
						break lab26
					}
					env.Cursor = env.Limit - v14
				lab29:
					for {
						if !markLAr(env, ctx) {
							break lab29
						}
						env.Bra = env.Cursor
						if !env.SliceDel() {
							return false
						}
						v16 := env.Limit - env.Cursor
					lab30:
						for {
							if !stemSuffixChainBeforeKi(env, ctx) {
								env.Cursor = env.Limit - v16
								break lab30
							}
							break lab30
						}
						break lab26
					}
					env.Cursor = env.Limit - v14
					if !stemSuffixChainBeforeKi(env, ctx) {
						env.Cursor = env.Limit - v13
						break lab25
					}
					break lab26
				}
				break lab25
			}
			break lab0
		}
		env.Cursor = env.Limit - v1
	lab31:
		for {
			env.Ket = env.Cursor
		lab32:
			for {
				v17 := env.Limit - env.Cursor
			lab33:
				for {
					if !markNUn(env, ctx) {
						break lab33
					}
					break lab32
				}
				env.Cursor = env.Limit - v17
				if !markYlA(env, ctx) {
					break lab31
				}
				break lab32
			}
			env.Bra = env.Cursor
			if !env.SliceDel() {
				return false
			}
			v18 := env.Limit - env.Cursor
		lab34:
			for {
			lab35:
				for {
					v19 := env.Limit - env.Cursor
				lab36:
					for {
						env.Ket = env.Cursor
						if !markLAr(env, ctx) {
							break lab36
						}
						env.Bra = env.Cursor
						if !env.SliceDel() {
							return false
						}
						if !stemSuffixChainBeforeKi(env, ctx) {
							break lab36
						}
						break lab35
					}
					env.Cursor = env.Limit - v19
				lab37:
					for {
						env.Ket = env.Cursor
					lab38:
						for {
							v20 := env.Limit - env.Cursor
						lab39:
							for {
								if !markPossessives(env, ctx) {
									break lab39
								}
								break lab38
							}
							env.Cursor = env.Limit - v20
							if !markSU(env, ctx) {
								break lab37
							}
							break lab38
						}
						env.Bra = env.Cursor
						if !env.SliceDel() {
							return false
						}
						v21 := env.Limit - env.Cursor
					lab40:
						for {
							env.Ket = env.Cursor
							if !markLAr(env, ctx) {
								env.Cursor = env.Limit - v21
								break lab40
							}
							env.Bra = env.Cursor
							if !env.SliceDel() {
								return false
							}
							if !stemSuffixChainBeforeKi(env, ctx) {
								env.Cursor = env.Limit - v21
								break lab40
							}
							break lab40
						}
						break lab35
					}
					env.Cursor = env.Limit - v19
					if !stemSuffixChainBeforeKi(env, ctx) {
						env.Cursor = env.Limit - v18
						break lab34
					}
					break lab35
				}
				break lab34
			}
			break lab0
		}
		env.Cursor = env.Limit - v1
	lab41:
		for {
			env.Ket = env.Cursor
			if !markLArI(env, ctx) {
				break lab41
			}
			env.Bra = env.Cursor
			if !env.SliceDel() {
				return false
			}
			break lab0
		}
		env.Cursor = env.Limit - v1
	lab42:
		for {
			if !stemSuffixChainBeforeKi(env, ctx) {
				break lab42
			}
			break lab0
		}
		env.Cursor = env.Limit - v1
	lab43:
		for {
			env.Ket = env.Cursor
		lab44:
			for {
				v22 := env.Limit - env.Cursor
			lab45:
				for {
					if !markDA(env, ctx) {
						break lab45
					}
					break lab44
				}
				env.Cursor = env.Limit - v22
			lab46:
				for {
					if !markYU(env, ctx) {
						break lab46
					}
					break lab44
				}
				env.Cursor = env.Limit - v22
				if !markYA(env, ctx) {
					break lab43
				}
				break lab44
			}
			env.Bra = env.Cursor
			if !env.SliceDel() {
				return false
			}
			v23 := env.Limit - env.Cursor
		lab47:
			for {
				env.Ket = env.Cursor
			lab48:
				for {
					v24 := env.Limit - env.Cursor
				lab49:
					for {
						if !markPossessives(env, ctx) {
							break lab49
						}
						env.Bra = env.Cursor
						if !env.SliceDel() {
							return false
						}
						v25 := env.Limit - env.Cursor
					lab50:
						for {
							env.Ket = env.Cursor
							if !markLAr(env, ctx) {
								env.Cursor = env.Limit - v25
								break lab50
							}
							break lab50
						}
						break lab48
					}
					env.Cursor = env.Limit - v24
					if !markLAr(env, ctx) {
						env.Cursor = env.Limit - v23
						break lab47
					}
					break lab48
				}
				env.Bra = env.Cursor
				if !env.SliceDel() {
					return false
				}
				env.Ket = env.Cursor
				if !stemSuffixChainBeforeKi(env, ctx) {
					env.Cursor = env.Limit - v23
					break lab47
				}
				break lab47
			}
			break lab0
		}
		env.Cursor = env.Limit - v1
		env.Ket = env.Cursor
	lab51:
		for {
			v26 := env.Limit - env.Cursor
		lab52:
			for {
				if !markPossessives(env, ctx) {
					break lab52
				}
				break lab51
			}
			env.Cursor = env.Limit - v26
			if !markSU(env, ctx) {
				return false
			}
			break lab51
		}
		env.Bra = env.Cursor
		if !env.SliceDel() {
			return false
		}
		v27 := env.Limit - env.Cursor
	lab53:
		for {
			env.Ket = env.Cursor
			if !markLAr(env, ctx) {
				env.Cursor = env.Limit - v27
				break lab53
			}
			env.Bra = env.Cursor
			if !env.SliceDel() {
				return false
			}
			if !stemSuffixChainBeforeKi(env, ctx) {
				env.Cursor = env.Limit - v27
				break lab53
			}
			break lab53
		}
		break lab0
	}
	return true
}

// postProcessLastConsonants reverses final-consonant softening, turning a
// terminal b/c/d/ğ back into its hard counterpart p/ç/t/k.
func postProcessLastConsonants(env *snowball.Env, ctx *context) bool {
	var amongVar int32
	env.Ket = env.Cursor
	amongVar = snowball.FindAmongB(env, a23, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	switch amongVar {
	case 1:
		if !env.SliceFrom("p") {
			return false
		}
	case 2:
		if !env.SliceFrom("ç") {
			return false
		}
	case 3:
		if !env.SliceFrom("t") {
			return false
		}
	case 4:
		if !env.SliceFrom("k") {
			return false
		}
	}
	return true
}

// appendUToStemsEndingWithDOrG re-inserts the harmonic high vowel that was
// dropped from a stem ending in d or g (e.g. restoring the dropped vowel).
func appendUToStemsEndingWithDOrG(env *snowball.Env, ctx *context) bool {
	v1 := env.Limit - env.Cursor
lab0:
	for {
		v2 := env.Limit - env.Cursor
	lab1:
		for {
			if !env.EqSB("d") {
				break lab1
			}
			break lab0
		}
		env.Cursor = env.Limit - v2
		if !env.EqSB("g") {
			return false
		}
		break lab0
	}
	env.Cursor = env.Limit - v1
lab2:
	for {
		v3 := env.Limit - env.Cursor
	lab3:
		for {
			v4 := env.Limit - env.Cursor
		golab4:
			for {
				v5 := env.Limit - env.Cursor
			lab5:
				for {
					if !env.InGroupingB(gVowel, 97, 305) {
						break lab5
					}
					env.Cursor = env.Limit - v5
					break golab4
				}
				env.Cursor = env.Limit - v5
				if env.Cursor <= env.LimitBackward {
					break lab3
				}
				env.PreviousChar()
			}
		lab6:
			for {
				v6 := env.Limit - env.Cursor
			lab7:
				for {
					if !env.EqSB("a") {
						break lab7
					}
					break lab6
				}
				env.Cursor = env.Limit - v6
				if !env.EqSB("ı") {
					break lab3
				}
				break lab6
			}
			env.Cursor = env.Limit - v4
			c := env.Cursor
			env.Insert(env.Cursor, env.Cursor, "ı")
			env.Cursor = c
			break lab2
		}
		env.Cursor = env.Limit - v3
	lab8:
		for {
			v7 := env.Limit - env.Cursor
		golab9:
			for {
				v8 := env.Limit - env.Cursor
			lab10:
				for {
					if !env.InGroupingB(gVowel, 97, 305) {
						break lab10
					}
					env.Cursor = env.Limit - v8
					break golab9
				}
				env.Cursor = env.Limit - v8
				if env.Cursor <= env.LimitBackward {
					break lab8
				}
				env.PreviousChar()
			}
		lab11:
			for {
				v9 := env.Limit - env.Cursor
			lab12:
				for {
					if !env.EqSB("e") {
						break lab12
					}
					break lab11
				}
				env.Cursor = env.Limit - v9
				if !env.EqSB("i") {
					break lab8
				}
				break lab11
			}
			env.Cursor = env.Limit - v7
			c := env.Cursor
			env.Insert(env.Cursor, env.Cursor, "i")
			env.Cursor = c
			break lab2
		}
		env.Cursor = env.Limit - v3
	lab13:
		for {
			v10 := env.Limit - env.Cursor
		golab14:
			for {
				v11 := env.Limit - env.Cursor
			lab15:
				for {
					if !env.InGroupingB(gVowel, 97, 305) {
						break lab15
					}
					env.Cursor = env.Limit - v11
					break golab14
				}
				env.Cursor = env.Limit - v11
				if env.Cursor <= env.LimitBackward {
					break lab13
				}
				env.PreviousChar()
			}
		lab16:
			for {
				v12 := env.Limit - env.Cursor
			lab17:
				for {
					if !env.EqSB("o") {
						break lab17
					}
					break lab16
				}
				env.Cursor = env.Limit - v12
				if !env.EqSB("u") {
					break lab13
				}
				break lab16
			}
			env.Cursor = env.Limit - v10
			c := env.Cursor
			env.Insert(env.Cursor, env.Cursor, "u")
			env.Cursor = c
			break lab2
		}
		env.Cursor = env.Limit - v3
		v13 := env.Limit - env.Cursor
	golab18:
		for {
			v14 := env.Limit - env.Cursor
		lab19:
			for {
				if !env.InGroupingB(gVowel, 97, 305) {
					break lab19
				}
				env.Cursor = env.Limit - v14
				break golab18
			}
			env.Cursor = env.Limit - v14
			if env.Cursor <= env.LimitBackward {
				return false
			}
			env.PreviousChar()
		}
	lab20:
		for {
			v15 := env.Limit - env.Cursor
		lab21:
			for {
				if !env.EqSB("ö") {
					break lab21
				}
				break lab20
			}
			env.Cursor = env.Limit - v15
			if !env.EqSB("ü") {
				return false
			}
			break lab20
		}
		env.Cursor = env.Limit - v13
		c := env.Cursor
		env.Insert(env.Cursor, env.Cursor, "ü")
		env.Cursor = c
		break lab2
	}
	return true
}

// moreThanOneSyllableWord succeeds only when the word contains at least two
// vowels, the minimum length the algorithm will attempt to stem.
func moreThanOneSyllableWord(env *snowball.Env, ctx *context) bool {
	v1 := env.Cursor
	v2 := 2
replab0:
	for {
		v3 := env.Cursor
	lab1:
		for once := 0; once < 1; once++ {
		golab2:
			for {
			lab3:
				for {
					if !env.InGrouping(gVowel, 97, 305) {
						break lab3
					}
					break golab2
				}
				if env.Cursor >= env.Limit {
					break lab1
				}
				env.NextChar()
			}
			v2--
			continue replab0
		}
		env.Cursor = v3
		break replab0
	}
	if v2 > 0 {
		return false
	}
	env.Cursor = v1
	return true
}

// isReservedWord recognises the protected words "ad" and "soyad", which must
// not be stemmed.
func isReservedWord(env *snowball.Env, ctx *context) bool {
lab0:
	for {
		v1 := env.Cursor
	lab1:
		for {
			v2 := env.Cursor
		golab2:
			for {
			lab3:
				for {
					if !env.EqS("ad") {
						break lab3
					}
					break golab2
				}
				if env.Cursor >= env.Limit {
					break lab1
				}
				env.NextChar()
			}
			ctx.iStrlen = 2
			if ctx.iStrlen != env.Limit {
				break lab1
			}
			env.Cursor = v2
			break lab0
		}
		env.Cursor = v1
		v4 := env.Cursor
	golab4:
		for {
		lab5:
			for {
				if !env.EqS("soyad") {
					break lab5
				}
				break golab4
			}
			if env.Cursor >= env.Limit {
				return false
			}
			env.NextChar()
		}
		ctx.iStrlen = 5
		if ctx.iStrlen != env.Limit {
			return false
		}
		env.Cursor = v4
		break lab0
	}
	return true
}

// postlude bails on reserved words, then re-inserts dropped vowels and undoes
// consonant softening over the stemmed result.
func postlude(env *snowball.Env, ctx *context) bool {
	v1 := env.Cursor
lab0:
	for {
		if !isReservedWord(env, ctx) {
			break lab0
		}
		return false
	}
	env.Cursor = v1
	env.LimitBackward = env.Cursor
	env.Cursor = env.Limit
	v2 := env.Limit - env.Cursor
lab1:
	for {
		if !appendUToStemsEndingWithDOrG(env, ctx) {
			break lab1
		}
		break lab1
	}
	env.Cursor = env.Limit - v2
	v3 := env.Limit - env.Cursor
lab2:
	for {
		if !postProcessLastConsonants(env, ctx) {
			break lab2
		}
		break lab2
	}
	env.Cursor = env.Limit - v3
	env.Cursor = env.LimitBackward
	return true
}

// Stem runs the Snowball turkish algorithm over env, mirroring the generated
// `stem` entry point. It returns false when the word is too short or when a
// required stage fails, matching rust-stemmers' behaviour.
func Stem(env *snowball.Env) bool {
	ctx := &context{}
	if !moreThanOneSyllableWord(env, ctx) {
		return false
	}
	env.LimitBackward = env.Cursor
	env.Cursor = env.Limit
	v1 := env.Limit - env.Cursor
lab0:
	for {
		if !stemNominalVerbSuffixes(env, ctx) {
			break lab0
		}
		break lab0
	}
	env.Cursor = env.Limit - v1
	if !ctx.bContinueStemmingNounSuffixes {
		return false
	}
	v2 := env.Limit - env.Cursor
lab1:
	for {
		if !stemNounSuffixes(env, ctx) {
			break lab1
		}
		break lab1
	}
	env.Cursor = env.Limit - v2
	env.Cursor = env.LimitBackward
	if !postlude(env, ctx) {
		return false
	}
	return true
}
