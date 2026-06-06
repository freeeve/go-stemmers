// Package arabic is a byte-faithful Go port of rust-stemmers' generated
// Snowball "arabic" stemmer. It produces output identical to rust-stemmers
// 1.2.0's Arabic algorithm.
//
// The control flow mirrors the generated Rust: each Snowball labelled block
// (`'lab: loop { … break 'lab }`) becomes a Go labelled `for { … break LABEL }`,
// and the result-code dispatch (`else if among_var == N`) becomes a `switch`.
// Arabic-script literals are pasted verbatim as UTF-8.
package arabic

import "github.com/freeeve/go-stemmers/internal/snowball"

// context holds the per-run state the generated algorithm threads through its
// routines: the is-noun/is-verb/is-defined classification flags and the word
// length (in characters) used to gate suffix/prefix removal.
type context struct {
	bIsDefined bool
	bIsVerb    bool
	bIsNoun    bool
	iWordLen   int
}

var a0 = []snowball.Among[context]{
	{Str: "!", SubstringI: -1, Result: 3},
	{Str: "\"", SubstringI: -1, Result: 3},
	{Str: "%", SubstringI: -1, Result: 3},
	{Str: "*", SubstringI: -1, Result: 3},
	{Str: ",", SubstringI: -1, Result: 3},
	{Str: ".", SubstringI: -1, Result: 3},
	{Str: "/", SubstringI: -1, Result: 3},
	{Str: ":", SubstringI: -1, Result: 3},
	{Str: ";", SubstringI: -1, Result: 3},
	{Str: "?", SubstringI: -1, Result: 3},
	{Str: "\\", SubstringI: -1, Result: 3},
	{Str: "،", SubstringI: -1, Result: 4},
	{Str: "؛", SubstringI: -1, Result: 4},
	{Str: "؟", SubstringI: -1, Result: 4},
	{Str: "ـ", SubstringI: -1, Result: 2},
	{Str: "ً", SubstringI: -1, Result: 1},
	{Str: "ٌ", SubstringI: -1, Result: 1},
	{Str: "ٍ", SubstringI: -1, Result: 1},
	{Str: "َ", SubstringI: -1, Result: 1},
	{Str: "ُ", SubstringI: -1, Result: 1},
	{Str: "ِ", SubstringI: -1, Result: 1},
	{Str: "ّ", SubstringI: -1, Result: 1},
	{Str: "ْ", SubstringI: -1, Result: 1},
	{Str: "٠", SubstringI: -1, Result: 5},
	{Str: "١", SubstringI: -1, Result: 6},
	{Str: "٢", SubstringI: -1, Result: 7},
	{Str: "٣", SubstringI: -1, Result: 8},
	{Str: "٤", SubstringI: -1, Result: 9},
	{Str: "٥", SubstringI: -1, Result: 10},
	{Str: "٦", SubstringI: -1, Result: 11},
	{Str: "٧", SubstringI: -1, Result: 12},
	{Str: "٨", SubstringI: -1, Result: 13},
	{Str: "٩", SubstringI: -1, Result: 14},
	{Str: "٪", SubstringI: -1, Result: 15},
	{Str: "٫", SubstringI: -1, Result: 15},
	{Str: "٬", SubstringI: -1, Result: 15},
	{Str: "ﺀ", SubstringI: -1, Result: 16},
	{Str: "ﺁ", SubstringI: -1, Result: 20},
	{Str: "ﺂ", SubstringI: -1, Result: 20},
	{Str: "ﺃ", SubstringI: -1, Result: 17},
	{Str: "ﺄ", SubstringI: -1, Result: 17},
	{Str: "ﺅ", SubstringI: -1, Result: 21},
	{Str: "ﺆ", SubstringI: -1, Result: 21},
	{Str: "ﺇ", SubstringI: -1, Result: 18},
	{Str: "ﺈ", SubstringI: -1, Result: 18},
	{Str: "ﺉ", SubstringI: -1, Result: 19},
	{Str: "ﺊ", SubstringI: -1, Result: 19},
	{Str: "ﺋ", SubstringI: -1, Result: 19},
	{Str: "ﺌ", SubstringI: -1, Result: 19},
	{Str: "ﺍ", SubstringI: -1, Result: 22},
	{Str: "ﺎ", SubstringI: -1, Result: 22},
	{Str: "ﺏ", SubstringI: -1, Result: 23},
	{Str: "ﺐ", SubstringI: -1, Result: 23},
	{Str: "ﺑ", SubstringI: -1, Result: 23},
	{Str: "ﺒ", SubstringI: -1, Result: 23},
	{Str: "ﺓ", SubstringI: -1, Result: 24},
	{Str: "ﺔ", SubstringI: -1, Result: 24},
	{Str: "ﺕ", SubstringI: -1, Result: 25},
	{Str: "ﺖ", SubstringI: -1, Result: 25},
	{Str: "ﺗ", SubstringI: -1, Result: 25},
	{Str: "ﺘ", SubstringI: -1, Result: 25},
	{Str: "ﺙ", SubstringI: -1, Result: 26},
	{Str: "ﺚ", SubstringI: -1, Result: 26},
	{Str: "ﺛ", SubstringI: -1, Result: 26},
	{Str: "ﺜ", SubstringI: -1, Result: 26},
	{Str: "ﺝ", SubstringI: -1, Result: 27},
	{Str: "ﺞ", SubstringI: -1, Result: 27},
	{Str: "ﺟ", SubstringI: -1, Result: 27},
	{Str: "ﺠ", SubstringI: -1, Result: 27},
	{Str: "ﺡ", SubstringI: -1, Result: 28},
	{Str: "ﺢ", SubstringI: -1, Result: 28},
	{Str: "ﺣ", SubstringI: -1, Result: 28},
	{Str: "ﺤ", SubstringI: -1, Result: 28},
	{Str: "ﺥ", SubstringI: -1, Result: 29},
	{Str: "ﺦ", SubstringI: -1, Result: 29},
	{Str: "ﺧ", SubstringI: -1, Result: 29},
	{Str: "ﺨ", SubstringI: -1, Result: 29},
	{Str: "ﺩ", SubstringI: -1, Result: 30},
	{Str: "ﺪ", SubstringI: -1, Result: 30},
	{Str: "ﺫ", SubstringI: -1, Result: 31},
	{Str: "ﺬ", SubstringI: -1, Result: 31},
	{Str: "ﺭ", SubstringI: -1, Result: 32},
	{Str: "ﺮ", SubstringI: -1, Result: 32},
	{Str: "ﺯ", SubstringI: -1, Result: 33},
	{Str: "ﺰ", SubstringI: -1, Result: 33},
	{Str: "ﺱ", SubstringI: -1, Result: 34},
	{Str: "ﺲ", SubstringI: -1, Result: 34},
	{Str: "ﺳ", SubstringI: -1, Result: 34},
	{Str: "ﺴ", SubstringI: -1, Result: 34},
	{Str: "ﺵ", SubstringI: -1, Result: 35},
	{Str: "ﺶ", SubstringI: -1, Result: 35},
	{Str: "ﺷ", SubstringI: -1, Result: 35},
	{Str: "ﺸ", SubstringI: -1, Result: 35},
	{Str: "ﺹ", SubstringI: -1, Result: 36},
	{Str: "ﺺ", SubstringI: -1, Result: 36},
	{Str: "ﺻ", SubstringI: -1, Result: 36},
	{Str: "ﺼ", SubstringI: -1, Result: 36},
	{Str: "ﺽ", SubstringI: -1, Result: 37},
	{Str: "ﺾ", SubstringI: -1, Result: 37},
	{Str: "ﺿ", SubstringI: -1, Result: 37},
	{Str: "ﻀ", SubstringI: -1, Result: 37},
	{Str: "ﻁ", SubstringI: -1, Result: 38},
	{Str: "ﻂ", SubstringI: -1, Result: 38},
	{Str: "ﻃ", SubstringI: -1, Result: 38},
	{Str: "ﻄ", SubstringI: -1, Result: 38},
	{Str: "ﻅ", SubstringI: -1, Result: 39},
	{Str: "ﻆ", SubstringI: -1, Result: 39},
	{Str: "ﻇ", SubstringI: -1, Result: 39},
	{Str: "ﻈ", SubstringI: -1, Result: 39},
	{Str: "ﻉ", SubstringI: -1, Result: 40},
	{Str: "ﻊ", SubstringI: -1, Result: 40},
	{Str: "ﻋ", SubstringI: -1, Result: 40},
	{Str: "ﻌ", SubstringI: -1, Result: 40},
	{Str: "ﻍ", SubstringI: -1, Result: 41},
	{Str: "ﻎ", SubstringI: -1, Result: 41},
	{Str: "ﻏ", SubstringI: -1, Result: 41},
	{Str: "ﻐ", SubstringI: -1, Result: 41},
	{Str: "ﻑ", SubstringI: -1, Result: 42},
	{Str: "ﻒ", SubstringI: -1, Result: 42},
	{Str: "ﻓ", SubstringI: -1, Result: 42},
	{Str: "ﻔ", SubstringI: -1, Result: 42},
	{Str: "ﻕ", SubstringI: -1, Result: 43},
	{Str: "ﻖ", SubstringI: -1, Result: 43},
	{Str: "ﻗ", SubstringI: -1, Result: 43},
	{Str: "ﻘ", SubstringI: -1, Result: 43},
	{Str: "ﻙ", SubstringI: -1, Result: 44},
	{Str: "ﻚ", SubstringI: -1, Result: 44},
	{Str: "ﻛ", SubstringI: -1, Result: 44},
	{Str: "ﻜ", SubstringI: -1, Result: 44},
	{Str: "ﻝ", SubstringI: -1, Result: 45},
	{Str: "ﻞ", SubstringI: -1, Result: 45},
	{Str: "ﻟ", SubstringI: -1, Result: 45},
	{Str: "ﻠ", SubstringI: -1, Result: 45},
	{Str: "ﻡ", SubstringI: -1, Result: 46},
	{Str: "ﻢ", SubstringI: -1, Result: 46},
	{Str: "ﻣ", SubstringI: -1, Result: 46},
	{Str: "ﻤ", SubstringI: -1, Result: 46},
	{Str: "ﻥ", SubstringI: -1, Result: 47},
	{Str: "ﻦ", SubstringI: -1, Result: 47},
	{Str: "ﻧ", SubstringI: -1, Result: 47},
	{Str: "ﻨ", SubstringI: -1, Result: 47},
	{Str: "ﻩ", SubstringI: -1, Result: 48},
	{Str: "ﻪ", SubstringI: -1, Result: 48},
	{Str: "ﻫ", SubstringI: -1, Result: 48},
	{Str: "ﻬ", SubstringI: -1, Result: 48},
	{Str: "ﻭ", SubstringI: -1, Result: 49},
	{Str: "ﻮ", SubstringI: -1, Result: 49},
	{Str: "ﻯ", SubstringI: -1, Result: 50},
	{Str: "ﻰ", SubstringI: -1, Result: 50},
	{Str: "ﻱ", SubstringI: -1, Result: 51},
	{Str: "ﻲ", SubstringI: -1, Result: 51},
	{Str: "ﻳ", SubstringI: -1, Result: 51},
	{Str: "ﻴ", SubstringI: -1, Result: 51},
	{Str: "ﻵ", SubstringI: -1, Result: 55},
	{Str: "ﻶ", SubstringI: -1, Result: 55},
	{Str: "ﻷ", SubstringI: -1, Result: 53},
	{Str: "ﻸ", SubstringI: -1, Result: 53},
	{Str: "ﻹ", SubstringI: -1, Result: 54},
	{Str: "ﻺ", SubstringI: -1, Result: 54},
	{Str: "ﻻ", SubstringI: -1, Result: 52},
	{Str: "ﻼ", SubstringI: -1, Result: 52},
}

var a1 = []snowball.Among[context]{
	{Str: "آ", SubstringI: -1, Result: 1},
	{Str: "أ", SubstringI: -1, Result: 1},
	{Str: "ؤ", SubstringI: -1, Result: 2},
	{Str: "إ", SubstringI: -1, Result: 1},
	{Str: "ئ", SubstringI: -1, Result: 3},
}

var a2 = []snowball.Among[context]{
	{Str: "آ", SubstringI: -1, Result: 1},
	{Str: "أ", SubstringI: -1, Result: 1},
	{Str: "ؤ", SubstringI: -1, Result: 2},
	{Str: "إ", SubstringI: -1, Result: 1},
	{Str: "ئ", SubstringI: -1, Result: 3},
}

var a3 = []snowball.Among[context]{
	{Str: "ال", SubstringI: -1, Result: 2},
	{Str: "بال", SubstringI: -1, Result: 1},
	{Str: "كال", SubstringI: -1, Result: 1},
	{Str: "لل", SubstringI: -1, Result: 2},
}

var a4 = []snowball.Among[context]{
	{Str: "أآ", SubstringI: -1, Result: 2},
	{Str: "أأ", SubstringI: -1, Result: 1},
	{Str: "أؤ", SubstringI: -1, Result: 3},
	{Str: "أإ", SubstringI: -1, Result: 5},
	{Str: "أا", SubstringI: -1, Result: 4},
}

var a5 = []snowball.Among[context]{
	{Str: "ف", SubstringI: -1, Result: 1},
	{Str: "و", SubstringI: -1, Result: 2},
}

var a6 = []snowball.Among[context]{
	{Str: "ال", SubstringI: -1, Result: 2},
	{Str: "بال", SubstringI: -1, Result: 1},
	{Str: "كال", SubstringI: -1, Result: 1},
	{Str: "لل", SubstringI: -1, Result: 2},
}

var a7 = []snowball.Among[context]{
	{Str: "ب", SubstringI: -1, Result: 1},
	{Str: "بب", SubstringI: 0, Result: 2},
	{Str: "كك", SubstringI: -1, Result: 3},
}

var a8 = []snowball.Among[context]{
	{Str: "سأ", SubstringI: -1, Result: 4},
	{Str: "ست", SubstringI: -1, Result: 2},
	{Str: "سن", SubstringI: -1, Result: 3},
	{Str: "سي", SubstringI: -1, Result: 1},
}

var a9 = []snowball.Among[context]{
	{Str: "تست", SubstringI: -1, Result: 1},
	{Str: "نست", SubstringI: -1, Result: 1},
	{Str: "يست", SubstringI: -1, Result: 1},
}

var a10 = []snowball.Among[context]{
	{Str: "ك", SubstringI: -1, Result: 1},
	{Str: "كم", SubstringI: -1, Result: 2},
	{Str: "هم", SubstringI: -1, Result: 2},
	{Str: "هن", SubstringI: -1, Result: 2},
	{Str: "ه", SubstringI: -1, Result: 1},
	{Str: "ي", SubstringI: -1, Result: 1},
	{Str: "كما", SubstringI: -1, Result: 3},
	{Str: "هما", SubstringI: -1, Result: 3},
	{Str: "نا", SubstringI: -1, Result: 2},
	{Str: "ها", SubstringI: -1, Result: 2},
}

var a11 = []snowball.Among[context]{
	{Str: "ن", SubstringI: -1, Result: 1},
}

var a12 = []snowball.Among[context]{
	{Str: "و", SubstringI: -1, Result: 1},
	{Str: "ي", SubstringI: -1, Result: 1},
	{Str: "ا", SubstringI: -1, Result: 1},
}

var a13 = []snowball.Among[context]{
	{Str: "ات", SubstringI: -1, Result: 1},
}

var a14 = []snowball.Among[context]{
	{Str: "ت", SubstringI: -1, Result: 1},
}

var a15 = []snowball.Among[context]{
	{Str: "ة", SubstringI: -1, Result: 1},
}

var a16 = []snowball.Among[context]{
	{Str: "ي", SubstringI: -1, Result: 1},
}

var a17 = []snowball.Among[context]{
	{Str: "ك", SubstringI: -1, Result: 1},
	{Str: "كم", SubstringI: -1, Result: 2},
	{Str: "هم", SubstringI: -1, Result: 2},
	{Str: "كن", SubstringI: -1, Result: 2},
	{Str: "هن", SubstringI: -1, Result: 2},
	{Str: "ه", SubstringI: -1, Result: 1},
	{Str: "كمو", SubstringI: -1, Result: 3},
	{Str: "ني", SubstringI: -1, Result: 2},
	{Str: "كما", SubstringI: -1, Result: 3},
	{Str: "هما", SubstringI: -1, Result: 3},
	{Str: "نا", SubstringI: -1, Result: 2},
	{Str: "ها", SubstringI: -1, Result: 2},
}

var a18 = []snowball.Among[context]{
	{Str: "ن", SubstringI: -1, Result: 2},
	{Str: "ون", SubstringI: 0, Result: 4},
	{Str: "ين", SubstringI: 0, Result: 4},
	{Str: "ان", SubstringI: 0, Result: 4},
	{Str: "تن", SubstringI: 0, Result: 3},
	{Str: "ي", SubstringI: -1, Result: 2},
	{Str: "ا", SubstringI: -1, Result: 2},
	{Str: "تما", SubstringI: 6, Result: 5},
	{Str: "نا", SubstringI: 6, Result: 3},
	{Str: "تا", SubstringI: 6, Result: 3},
	{Str: "ت", SubstringI: -1, Result: 1},
}

var a19 = []snowball.Among[context]{
	{Str: "تم", SubstringI: -1, Result: 1},
	{Str: "وا", SubstringI: -1, Result: 1},
}

var a20 = []snowball.Among[context]{
	{Str: "و", SubstringI: -1, Result: 1},
	{Str: "تمو", SubstringI: 0, Result: 2},
}

var a21 = []snowball.Among[context]{
	{Str: "ى", SubstringI: -1, Result: 1},
}

// normalizePre normalizes the surface form before stemming: it strips
// punctuation, tatweel and harakat, maps Arabic-Indic digits to ASCII, and
// folds Arabic presentation forms back to their base letters.
func normalizePre(env *snowball.Env, ctx *context) bool {
	var amongVar int32
	runeCount := env.RuneCount()
	for i := 0; i < runeCount; i++ {
	lab0:
		for {
			v2 := env.Cursor
		lab1:
			for {
				env.Bra = env.Cursor
				amongVar = snowball.FindAmong(env, a0, ctx)
				if amongVar == 0 {
					break lab1
				}
				env.Ket = env.Cursor
				switch amongVar {
				case 1:
					if !env.SliceDel() {
						return false
					}
				case 2:
					if !env.SliceDel() {
						return false
					}
				case 3:
					if !env.SliceDel() {
						return false
					}
				case 4:
					if !env.SliceDel() {
						return false
					}
				case 5:
					if !env.SliceFrom("0") {
						return false
					}
				case 6:
					if !env.SliceFrom("1") {
						return false
					}
				case 7:
					if !env.SliceFrom("2") {
						return false
					}
				case 8:
					if !env.SliceFrom("3") {
						return false
					}
				case 9:
					if !env.SliceFrom("4") {
						return false
					}
				case 10:
					if !env.SliceFrom("5") {
						return false
					}
				case 11:
					if !env.SliceFrom("6") {
						return false
					}
				case 12:
					if !env.SliceFrom("7") {
						return false
					}
				case 13:
					if !env.SliceFrom("8") {
						return false
					}
				case 14:
					if !env.SliceFrom("9") {
						return false
					}
				case 15:
					if !env.SliceDel() {
						return false
					}
				case 16:
					if !env.SliceFrom("ء") {
						return false
					}
				case 17:
					if !env.SliceFrom("أ") {
						return false
					}
				case 18:
					if !env.SliceFrom("إ") {
						return false
					}
				case 19:
					if !env.SliceFrom("ئ") {
						return false
					}
				case 20:
					if !env.SliceFrom("آ") {
						return false
					}
				case 21:
					if !env.SliceFrom("ؤ") {
						return false
					}
				case 22:
					if !env.SliceFrom("ا") {
						return false
					}
				case 23:
					if !env.SliceFrom("ب") {
						return false
					}
				case 24:
					if !env.SliceFrom("ة") {
						return false
					}
				case 25:
					if !env.SliceFrom("ت") {
						return false
					}
				case 26:
					if !env.SliceFrom("ث") {
						return false
					}
				case 27:
					if !env.SliceFrom("ج") {
						return false
					}
				case 28:
					if !env.SliceFrom("ح") {
						return false
					}
				case 29:
					if !env.SliceFrom("خ") {
						return false
					}
				case 30:
					if !env.SliceFrom("د") {
						return false
					}
				case 31:
					if !env.SliceFrom("ذ") {
						return false
					}
				case 32:
					if !env.SliceFrom("ر") {
						return false
					}
				case 33:
					if !env.SliceFrom("ز") {
						return false
					}
				case 34:
					if !env.SliceFrom("س") {
						return false
					}
				case 35:
					if !env.SliceFrom("ش") {
						return false
					}
				case 36:
					if !env.SliceFrom("ص") {
						return false
					}
				case 37:
					if !env.SliceFrom("ض") {
						return false
					}
				case 38:
					if !env.SliceFrom("ط") {
						return false
					}
				case 39:
					if !env.SliceFrom("ظ") {
						return false
					}
				case 40:
					if !env.SliceFrom("ع") {
						return false
					}
				case 41:
					if !env.SliceFrom("غ") {
						return false
					}
				case 42:
					if !env.SliceFrom("ف") {
						return false
					}
				case 43:
					if !env.SliceFrom("ق") {
						return false
					}
				case 44:
					if !env.SliceFrom("ك") {
						return false
					}
				case 45:
					if !env.SliceFrom("ل") {
						return false
					}
				case 46:
					if !env.SliceFrom("م") {
						return false
					}
				case 47:
					if !env.SliceFrom("ن") {
						return false
					}
				case 48:
					if !env.SliceFrom("ه") {
						return false
					}
				case 49:
					if !env.SliceFrom("و") {
						return false
					}
				case 50:
					if !env.SliceFrom("ى") {
						return false
					}
				case 51:
					if !env.SliceFrom("ي") {
						return false
					}
				case 52:
					if !env.SliceFrom("لا") {
						return false
					}
				case 53:
					if !env.SliceFrom("لأ") {
						return false
					}
				case 54:
					if !env.SliceFrom("لإ") {
						return false
					}
				case 55:
					if !env.SliceFrom("لآ") {
						return false
					}
				}
				break lab0
			}
			env.Cursor = v2
			if env.Cursor >= env.Limit {
				return false
			}
			env.NextChar()
			break lab0
		}
	}
	return true
}

// normalizePost runs the post-stemming clean-up: it folds a trailing hamza
// carrier to a bare hamza (backwards), then walks the word fixing the
// hamza-on-carrier forms left by stemming back to plain alef/waw/ya.
func normalizePost(env *snowball.Env, ctx *context) bool {
	var amongVar int32
	v1 := env.Cursor
lab0:
	for {
		env.LimitBackward = env.Cursor
		env.Cursor = env.Limit
		env.Ket = env.Cursor
		amongVar = snowball.FindAmongB(env, a1, ctx)
		if amongVar == 0 {
			break lab0
		}
		env.Bra = env.Cursor
		switch amongVar {
		case 1:
			if !env.SliceFrom("ء") {
				return false
			}
		case 2:
			if !env.SliceFrom("ء") {
				return false
			}
		case 3:
			if !env.SliceFrom("ء") {
				return false
			}
		}
		env.Cursor = env.LimitBackward
		break lab0
	}
	env.Cursor = v1
	v2 := env.Cursor
lab1:
	for {
		for i := 0; i < ctx.iWordLen; i++ {
		lab2:
			for {
				v4 := env.Cursor
			lab3:
				for {
					env.Bra = env.Cursor
					amongVar = snowball.FindAmong(env, a2, ctx)
					if amongVar == 0 {
						break lab3
					}
					env.Ket = env.Cursor
					switch amongVar {
					case 1:
						if !env.SliceFrom("ا") {
							return false
						}
					case 2:
						if !env.SliceFrom("و") {
							return false
						}
					case 3:
						if !env.SliceFrom("ي") {
							return false
						}
					}
					break lab2
				}
				env.Cursor = v4
				if env.Cursor >= env.Limit {
					break lab1
				}
				env.NextChar()
				break lab2
			}
		}
		break lab1
	}
	env.Cursor = v2
	return true
}

// checks1 detects the definite-article noun prefixes (al-, bil-, kal-, lil-)
// and, when the word is long enough, classifies the word as a defined noun.
func checks1(env *snowball.Env, ctx *context) bool {
	ctx.iWordLen = env.RuneCount()
	env.Bra = env.Cursor
	amongVar := snowball.FindAmong(env, a3, ctx)
	if amongVar == 0 {
		return false
	}
	env.Ket = env.Cursor
	switch amongVar {
	case 1:
		if !(ctx.iWordLen > 4) {
			return false
		}
		ctx.bIsNoun = true
		ctx.bIsVerb = false
		ctx.bIsDefined = true
	case 2:
		if !(ctx.iWordLen > 3) {
			return false
		}
		ctx.bIsNoun = true
		ctx.bIsVerb = false
		ctx.bIsDefined = true
	}
	return true
}

// prefixStep1 normalizes hamza-bearing word-initial sequences.
func prefixStep1(env *snowball.Env, ctx *context) bool {
	ctx.iWordLen = env.RuneCount()
	env.Bra = env.Cursor
	amongVar := snowball.FindAmong(env, a4, ctx)
	if amongVar == 0 {
		return false
	}
	env.Ket = env.Cursor
	switch amongVar {
	case 1:
		if !(ctx.iWordLen > 3) {
			return false
		}
		if !env.SliceFrom("أ") {
			return false
		}
	case 2:
		if !(ctx.iWordLen > 3) {
			return false
		}
		if !env.SliceFrom("آ") {
			return false
		}
	case 3:
		if !(ctx.iWordLen > 3) {
			return false
		}
		if !env.SliceFrom("أ") {
			return false
		}
	case 4:
		if !(ctx.iWordLen > 3) {
			return false
		}
		if !env.SliceFrom("ا") {
			return false
		}
	case 5:
		if !(ctx.iWordLen > 3) {
			return false
		}
		if !env.SliceFrom("إ") {
			return false
		}
	}
	return true
}

// prefixStep2 removes the conjunction prefixes fa- and wa- (guarding against
// the words fa and wa themselves).
func prefixStep2(env *snowball.Env, ctx *context) bool {
	ctx.iWordLen = env.RuneCount()
	v1 := env.Cursor
lab0:
	for {
		if !env.EqS("فا") {
			break lab0
		}
		return false
	}
	env.Cursor = v1
	v2 := env.Cursor
lab1:
	for {
		if !env.EqS("وا") {
			break lab1
		}
		return false
	}
	env.Cursor = v2
	env.Bra = env.Cursor
	amongVar := snowball.FindAmong(env, a5, ctx)
	if amongVar == 0 {
		return false
	}
	env.Ket = env.Cursor
	switch amongVar {
	case 1:
		if !(ctx.iWordLen > 3) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	case 2:
		if !(ctx.iWordLen > 3) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	}
	return true
}

// prefixStep3aNoun removes the definite-article noun prefixes.
func prefixStep3aNoun(env *snowball.Env, ctx *context) bool {
	ctx.iWordLen = env.RuneCount()
	env.Bra = env.Cursor
	amongVar := snowball.FindAmong(env, a6, ctx)
	if amongVar == 0 {
		return false
	}
	env.Ket = env.Cursor
	switch amongVar {
	case 1:
		if !(ctx.iWordLen > 5) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	case 2:
		if !(ctx.iWordLen > 4) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	}
	return true
}

// prefixStep3bNoun removes single-letter noun prefixes (guarding ba-) and
// restores doubled bi-/ka- prefixes.
func prefixStep3bNoun(env *snowball.Env, ctx *context) bool {
	ctx.iWordLen = env.RuneCount()
	v1 := env.Cursor
lab0:
	for {
		if !env.EqS("با") {
			break lab0
		}
		return false
	}
	env.Cursor = v1
	env.Bra = env.Cursor
	amongVar := snowball.FindAmong(env, a7, ctx)
	if amongVar == 0 {
		return false
	}
	env.Ket = env.Cursor
	switch amongVar {
	case 1:
		if !(ctx.iWordLen > 3) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	case 2:
		if !(ctx.iWordLen > 3) {
			return false
		}
		if !env.SliceFrom("ب") {
			return false
		}
	case 3:
		if !(ctx.iWordLen > 3) {
			return false
		}
		if !env.SliceFrom("ك") {
			return false
		}
	}
	return true
}

// prefixStep3Verb rewrites the imperfective verb prefixes back to their base
// letter.
func prefixStep3Verb(env *snowball.Env, ctx *context) bool {
	ctx.iWordLen = env.RuneCount()
	env.Bra = env.Cursor
	amongVar := snowball.FindAmong(env, a8, ctx)
	if amongVar == 0 {
		return false
	}
	env.Ket = env.Cursor
	switch amongVar {
	case 1:
		if !(ctx.iWordLen > 4) {
			return false
		}
		if !env.SliceFrom("ي") {
			return false
		}
	case 2:
		if !(ctx.iWordLen > 4) {
			return false
		}
		if !env.SliceFrom("ت") {
			return false
		}
	case 3:
		if !(ctx.iWordLen > 4) {
			return false
		}
		if !env.SliceFrom("ن") {
			return false
		}
	case 4:
		if !(ctx.iWordLen > 4) {
			return false
		}
		if !env.SliceFrom("أ") {
			return false
		}
	}
	return true
}

// prefixStep4Verb recognises the verbal istif'al prefix and marks the word as
// a verb.
func prefixStep4Verb(env *snowball.Env, ctx *context) bool {
	ctx.iWordLen = env.RuneCount()
	env.Bra = env.Cursor
	amongVar := snowball.FindAmong(env, a9, ctx)
	if amongVar == 0 {
		return false
	}
	env.Ket = env.Cursor
	switch amongVar {
	case 1:
		if !(ctx.iWordLen > 4) {
			return false
		}
		ctx.bIsVerb = true
		ctx.bIsNoun = false
		if !env.SliceFrom("است") {
			return false
		}
	}
	return true
}

// suffixNounStep1a removes the attached pronoun noun suffixes.
func suffixNounStep1a(env *snowball.Env, ctx *context) bool {
	ctx.iWordLen = env.RuneCount()
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a10, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	switch amongVar {
	case 1:
		if !(ctx.iWordLen >= 4) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	case 2:
		if !(ctx.iWordLen >= 5) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	case 3:
		if !(ctx.iWordLen >= 6) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	}
	return true
}

// suffixNounStep1b removes a trailing noon noun suffix.
func suffixNounStep1b(env *snowball.Env, ctx *context) bool {
	ctx.iWordLen = env.RuneCount()
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a11, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	switch amongVar {
	case 1:
		if !(ctx.iWordLen > 5) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	}
	return true
}

// suffixNounStep2a removes a trailing waw/ya/alef noun suffix.
func suffixNounStep2a(env *snowball.Env, ctx *context) bool {
	ctx.iWordLen = env.RuneCount()
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a12, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	switch amongVar {
	case 1:
		if !(ctx.iWordLen > 4) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	}
	return true
}

// suffixNounStep2b removes the trailing -at noun suffix.
func suffixNounStep2b(env *snowball.Env, ctx *context) bool {
	ctx.iWordLen = env.RuneCount()
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a13, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	switch amongVar {
	case 1:
		if !(ctx.iWordLen >= 5) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	}
	return true
}

// suffixNounStep2c1 removes a trailing ta noun suffix.
func suffixNounStep2c1(env *snowball.Env, ctx *context) bool {
	ctx.iWordLen = env.RuneCount()
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a14, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	switch amongVar {
	case 1:
		if !(ctx.iWordLen >= 4) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	}
	return true
}

// suffixNounStep2c2 removes the trailing ta marbuta noun suffix.
func suffixNounStep2c2(env *snowball.Env, ctx *context) bool {
	ctx.iWordLen = env.RuneCount()
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a15, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	switch amongVar {
	case 1:
		if !(ctx.iWordLen >= 4) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	}
	return true
}

// suffixNounStep3 removes a trailing ya noun suffix.
func suffixNounStep3(env *snowball.Env, ctx *context) bool {
	ctx.iWordLen = env.RuneCount()
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a16, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	switch amongVar {
	case 1:
		if !(ctx.iWordLen >= 3) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	}
	return true
}

// suffixVerbStep1 removes the attached pronoun verb suffixes.
func suffixVerbStep1(env *snowball.Env, ctx *context) bool {
	ctx.iWordLen = env.RuneCount()
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a17, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	switch amongVar {
	case 1:
		if !(ctx.iWordLen >= 4) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	case 2:
		if !(ctx.iWordLen >= 5) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	case 3:
		if !(ctx.iWordLen >= 6) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	}
	return true
}

// suffixVerbStep2a removes the verb conjugation suffixes (noon/waw/ya/alef
// families).
func suffixVerbStep2a(env *snowball.Env, ctx *context) bool {
	ctx.iWordLen = env.RuneCount()
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a18, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	switch amongVar {
	case 1:
		if !(ctx.iWordLen >= 4) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	case 2:
		if !(ctx.iWordLen >= 4) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	case 3:
		if !(ctx.iWordLen >= 5) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	case 4:
		if !(ctx.iWordLen > 5) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	case 5:
		if !(ctx.iWordLen >= 6) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	}
	return true
}

// suffixVerbStep2b removes the -tum/-wa verb suffixes.
func suffixVerbStep2b(env *snowball.Env, ctx *context) bool {
	ctx.iWordLen = env.RuneCount()
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a19, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	switch amongVar {
	case 1:
		if !(ctx.iWordLen >= 5) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	}
	return true
}

// suffixVerbStep2c removes the -tumuu/-waw verb suffixes.
func suffixVerbStep2c(env *snowball.Env, ctx *context) bool {
	ctx.iWordLen = env.RuneCount()
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a20, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	switch amongVar {
	case 1:
		if !(ctx.iWordLen >= 4) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	case 2:
		if !(ctx.iWordLen >= 6) {
			return false
		}
		if !env.SliceDel() {
			return false
		}
	}
	return true
}

// suffixAllAlefMaqsura rewrites a trailing alef maqsura to ya.
func suffixAllAlefMaqsura(env *snowball.Env, ctx *context) bool {
	ctx.iWordLen = env.RuneCount()
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a21, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	switch amongVar {
	case 1:
		if !env.SliceFrom("ي") {
			return false
		}
	}
	return true
}

// Stem runs the Snowball arabic algorithm over env, mirroring the generated
// `stem` entry point. It always returns true; the result is the mutated env.
func Stem(env *snowball.Env) bool {
	ctx, _ := env.Scratch.(*context)
	if ctx == nil {
		ctx = &context{}
		env.Scratch = ctx
	}
	*ctx = context{}
	ctx.bIsNoun = true
	ctx.bIsVerb = true
	ctx.bIsDefined = false
	v1 := env.Cursor
lab0:
	for {
		if !checks1(env, ctx) {
			break lab0
		}
		break lab0
	}
	env.Cursor = v1
	v2 := env.Cursor
lab1:
	for {
		if !normalizePre(env, ctx) {
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
				if !ctx.bIsVerb {
					break lab4
				}
			lab5:
				for {
					v5 := env.Limit - env.Cursor
				lab6:
					for {
						v6 := 1
					replab7:
						for {
							v7 := env.Limit - env.Cursor
						lab8:
							for once := 0; once < 1; once++ {
								if !suffixVerbStep1(env, ctx) {
									break lab8
								}
								v6--
								continue replab7
							}
							env.Cursor = env.Limit - v7
							break replab7
						}
						if v6 > 0 {
							break lab6
						}
					lab9:
						for {
							v8 := env.Limit - env.Cursor
						lab10:
							for {
								if !suffixVerbStep2a(env, ctx) {
									break lab10
								}
								break lab9
							}
							env.Cursor = env.Limit - v8
						lab11:
							for {
								if !suffixVerbStep2c(env, ctx) {
									break lab11
								}
								break lab9
							}
							env.Cursor = env.Limit - v8
							if env.Cursor <= env.LimitBackward {
								break lab6
							}
							env.PreviousChar()
							break lab9
						}
						break lab5
					}
					env.Cursor = env.Limit - v5
				lab12:
					for {
						if !suffixVerbStep2b(env, ctx) {
							break lab12
						}
						break lab5
					}
					env.Cursor = env.Limit - v5
					if !suffixVerbStep2a(env, ctx) {
						break lab4
					}
					break lab5
				}
				break lab3
			}
			env.Cursor = env.Limit - v4
		lab13:
			for {
				if !ctx.bIsNoun {
					break lab13
				}
				v9 := env.Limit - env.Cursor
			lab14:
				for {
				lab15:
					for {
						v10 := env.Limit - env.Cursor
					lab16:
						for {
							if !suffixNounStep2c2(env, ctx) {
								break lab16
							}
							break lab15
						}
						env.Cursor = env.Limit - v10
					lab17:
						for {
						lab18:
							for {
								if !ctx.bIsDefined {
									break lab18
								}
								break lab17
							}
							if !suffixNounStep1a(env, ctx) {
								break lab17
							}
						lab19:
							for {
								v12 := env.Limit - env.Cursor
							lab20:
								for {
									if !suffixNounStep2a(env, ctx) {
										break lab20
									}
									break lab19
								}
								env.Cursor = env.Limit - v12
							lab21:
								for {
									if !suffixNounStep2b(env, ctx) {
										break lab21
									}
									break lab19
								}
								env.Cursor = env.Limit - v12
							lab22:
								for {
									if !suffixNounStep2c1(env, ctx) {
										break lab22
									}
									break lab19
								}
								env.Cursor = env.Limit - v12
								if env.Cursor <= env.LimitBackward {
									break lab17
								}
								env.PreviousChar()
								break lab19
							}
							break lab15
						}
						env.Cursor = env.Limit - v10
					lab23:
						for {
							if !suffixNounStep1b(env, ctx) {
								break lab23
							}
						lab24:
							for {
								v13 := env.Limit - env.Cursor
							lab25:
								for {
									if !suffixNounStep2a(env, ctx) {
										break lab25
									}
									break lab24
								}
								env.Cursor = env.Limit - v13
							lab26:
								for {
									if !suffixNounStep2b(env, ctx) {
										break lab26
									}
									break lab24
								}
								env.Cursor = env.Limit - v13
								if !suffixNounStep2c1(env, ctx) {
									break lab23
								}
								break lab24
							}
							break lab15
						}
						env.Cursor = env.Limit - v10
					lab27:
						for {
						lab28:
							for {
								if !ctx.bIsDefined {
									break lab28
								}
								break lab27
							}
							if !suffixNounStep2a(env, ctx) {
								break lab27
							}
							break lab15
						}
						env.Cursor = env.Limit - v10
						if !suffixNounStep2b(env, ctx) {
							env.Cursor = env.Limit - v9
							break lab14
						}
						break lab15
					}
					break lab14
				}
				if !suffixNounStep3(env, ctx) {
					break lab13
				}
				break lab3
			}
			env.Cursor = env.Limit - v4
			if !suffixAllAlefMaqsura(env, ctx) {
				break lab2
			}
			break lab3
		}
		break lab2
	}
	env.Cursor = env.Limit - v3
	env.Cursor = env.LimitBackward
	v15 := env.Cursor
lab29:
	for {
		v16 := env.Cursor
	lab30:
		for {
			if !prefixStep1(env, ctx) {
				env.Cursor = v16
				break lab30
			}
			break lab30
		}
		v17 := env.Cursor
	lab31:
		for {
			if !prefixStep2(env, ctx) {
				env.Cursor = v17
				break lab31
			}
			break lab31
		}
	lab32:
		for {
			v18 := env.Cursor
		lab33:
			for {
				if !prefixStep3aNoun(env, ctx) {
					break lab33
				}
				break lab32
			}
			env.Cursor = v18
		lab34:
			for {
				if !ctx.bIsNoun {
					break lab34
				}
				if !prefixStep3bNoun(env, ctx) {
					break lab34
				}
				break lab32
			}
			env.Cursor = v18
			if !ctx.bIsVerb {
				break lab29
			}
			v19 := env.Cursor
		lab35:
			for {
				if !prefixStep3Verb(env, ctx) {
					env.Cursor = v19
					break lab35
				}
				break lab35
			}
			if !prefixStep4Verb(env, ctx) {
				break lab29
			}
			break lab32
		}
		break lab29
	}
	env.Cursor = v15
	v20 := env.Cursor
lab36:
	for {
		if !normalizePost(env, ctx) {
			break lab36
		}
		break lab36
	}
	env.Cursor = v20
	return true
}
