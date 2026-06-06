// Package hungarian is a byte-faithful Go port of rust-stemmers' generated
// Snowball "hungarian" stemmer. It produces output identical to rust-stemmers
// 1.2.0's Hungarian algorithm; the canonical Snowball hungarian vocabulary is
// the conformance oracle.
//
// The control flow mirrors the generated Rust: each Snowball labelled block
// (`'lab: loop { … break 'lab }`) becomes a Go labelled `for { … break LABEL }`,
// and the result-code dispatch (`else if among_var == N`) becomes a `switch`.
package hungarian

import "github.com/freeeve/go-stemmers/internal/snowball"

// context holds the per-run state the generated algorithm threads through its
// routines: the R1 region mark.
type context struct {
	iP1 int
}

var a0 = []snowball.Among[context]{
	{Str: "cs", SubstringI: -1, Result: -1},
	{Str: "dzs", SubstringI: -1, Result: -1},
	{Str: "gy", SubstringI: -1, Result: -1},
	{Str: "ly", SubstringI: -1, Result: -1},
	{Str: "ny", SubstringI: -1, Result: -1},
	{Str: "sz", SubstringI: -1, Result: -1},
	{Str: "ty", SubstringI: -1, Result: -1},
	{Str: "zs", SubstringI: -1, Result: -1},
}

var a1 = []snowball.Among[context]{
	{Str: "á", SubstringI: -1, Result: 1},
	{Str: "é", SubstringI: -1, Result: 2},
}

var a2 = []snowball.Among[context]{
	{Str: "bb", SubstringI: -1, Result: -1},
	{Str: "cc", SubstringI: -1, Result: -1},
	{Str: "dd", SubstringI: -1, Result: -1},
	{Str: "ff", SubstringI: -1, Result: -1},
	{Str: "gg", SubstringI: -1, Result: -1},
	{Str: "jj", SubstringI: -1, Result: -1},
	{Str: "kk", SubstringI: -1, Result: -1},
	{Str: "ll", SubstringI: -1, Result: -1},
	{Str: "mm", SubstringI: -1, Result: -1},
	{Str: "nn", SubstringI: -1, Result: -1},
	{Str: "pp", SubstringI: -1, Result: -1},
	{Str: "rr", SubstringI: -1, Result: -1},
	{Str: "ccs", SubstringI: -1, Result: -1},
	{Str: "ss", SubstringI: -1, Result: -1},
	{Str: "zzs", SubstringI: -1, Result: -1},
	{Str: "tt", SubstringI: -1, Result: -1},
	{Str: "vv", SubstringI: -1, Result: -1},
	{Str: "ggy", SubstringI: -1, Result: -1},
	{Str: "lly", SubstringI: -1, Result: -1},
	{Str: "nny", SubstringI: -1, Result: -1},
	{Str: "tty", SubstringI: -1, Result: -1},
	{Str: "ssz", SubstringI: -1, Result: -1},
	{Str: "zz", SubstringI: -1, Result: -1},
}

var a3 = []snowball.Among[context]{
	{Str: "al", SubstringI: -1, Result: 1},
	{Str: "el", SubstringI: -1, Result: 2},
}

var a4 = []snowball.Among[context]{
	{Str: "ba", SubstringI: -1, Result: -1},
	{Str: "ra", SubstringI: -1, Result: -1},
	{Str: "be", SubstringI: -1, Result: -1},
	{Str: "re", SubstringI: -1, Result: -1},
	{Str: "ig", SubstringI: -1, Result: -1},
	{Str: "nak", SubstringI: -1, Result: -1},
	{Str: "nek", SubstringI: -1, Result: -1},
	{Str: "val", SubstringI: -1, Result: -1},
	{Str: "vel", SubstringI: -1, Result: -1},
	{Str: "ul", SubstringI: -1, Result: -1},
	{Str: "ből", SubstringI: -1, Result: -1},
	{Str: "ről", SubstringI: -1, Result: -1},
	{Str: "től", SubstringI: -1, Result: -1},
	{Str: "nál", SubstringI: -1, Result: -1},
	{Str: "nél", SubstringI: -1, Result: -1},
	{Str: "ból", SubstringI: -1, Result: -1},
	{Str: "ról", SubstringI: -1, Result: -1},
	{Str: "tól", SubstringI: -1, Result: -1},
	{Str: "ül", SubstringI: -1, Result: -1},
	{Str: "n", SubstringI: -1, Result: -1},
	{Str: "an", SubstringI: 19, Result: -1},
	{Str: "ban", SubstringI: 20, Result: -1},
	{Str: "en", SubstringI: 19, Result: -1},
	{Str: "ben", SubstringI: 22, Result: -1},
	{Str: "képpen", SubstringI: 22, Result: -1},
	{Str: "on", SubstringI: 19, Result: -1},
	{Str: "ön", SubstringI: 19, Result: -1},
	{Str: "képp", SubstringI: -1, Result: -1},
	{Str: "kor", SubstringI: -1, Result: -1},
	{Str: "t", SubstringI: -1, Result: -1},
	{Str: "at", SubstringI: 29, Result: -1},
	{Str: "et", SubstringI: 29, Result: -1},
	{Str: "ként", SubstringI: 29, Result: -1},
	{Str: "anként", SubstringI: 32, Result: -1},
	{Str: "enként", SubstringI: 32, Result: -1},
	{Str: "onként", SubstringI: 32, Result: -1},
	{Str: "ot", SubstringI: 29, Result: -1},
	{Str: "ért", SubstringI: 29, Result: -1},
	{Str: "öt", SubstringI: 29, Result: -1},
	{Str: "hez", SubstringI: -1, Result: -1},
	{Str: "hoz", SubstringI: -1, Result: -1},
	{Str: "höz", SubstringI: -1, Result: -1},
	{Str: "vá", SubstringI: -1, Result: -1},
	{Str: "vé", SubstringI: -1, Result: -1},
}

var a5 = []snowball.Among[context]{
	{Str: "án", SubstringI: -1, Result: 2},
	{Str: "én", SubstringI: -1, Result: 1},
	{Str: "ánként", SubstringI: -1, Result: 3},
}

var a6 = []snowball.Among[context]{
	{Str: "stul", SubstringI: -1, Result: 2},
	{Str: "astul", SubstringI: 0, Result: 1},
	{Str: "ástul", SubstringI: 0, Result: 3},
	{Str: "stül", SubstringI: -1, Result: 2},
	{Str: "estül", SubstringI: 3, Result: 1},
	{Str: "éstül", SubstringI: 3, Result: 4},
}

var a7 = []snowball.Among[context]{
	{Str: "á", SubstringI: -1, Result: 1},
	{Str: "é", SubstringI: -1, Result: 2},
}

var a8 = []snowball.Among[context]{
	{Str: "k", SubstringI: -1, Result: 7},
	{Str: "ak", SubstringI: 0, Result: 4},
	{Str: "ek", SubstringI: 0, Result: 6},
	{Str: "ok", SubstringI: 0, Result: 5},
	{Str: "ák", SubstringI: 0, Result: 1},
	{Str: "ék", SubstringI: 0, Result: 2},
	{Str: "ök", SubstringI: 0, Result: 3},
}

var a9 = []snowball.Among[context]{
	{Str: "éi", SubstringI: -1, Result: 7},
	{Str: "áéi", SubstringI: 0, Result: 6},
	{Str: "ééi", SubstringI: 0, Result: 5},
	{Str: "é", SubstringI: -1, Result: 9},
	{Str: "ké", SubstringI: 3, Result: 4},
	{Str: "aké", SubstringI: 4, Result: 1},
	{Str: "eké", SubstringI: 4, Result: 1},
	{Str: "oké", SubstringI: 4, Result: 1},
	{Str: "áké", SubstringI: 4, Result: 3},
	{Str: "éké", SubstringI: 4, Result: 2},
	{Str: "öké", SubstringI: 4, Result: 1},
	{Str: "éé", SubstringI: 3, Result: 8},
}

var a10 = []snowball.Among[context]{
	{Str: "a", SubstringI: -1, Result: 18},
	{Str: "ja", SubstringI: 0, Result: 17},
	{Str: "d", SubstringI: -1, Result: 16},
	{Str: "ad", SubstringI: 2, Result: 13},
	{Str: "ed", SubstringI: 2, Result: 13},
	{Str: "od", SubstringI: 2, Result: 13},
	{Str: "ád", SubstringI: 2, Result: 14},
	{Str: "éd", SubstringI: 2, Result: 15},
	{Str: "öd", SubstringI: 2, Result: 13},
	{Str: "e", SubstringI: -1, Result: 18},
	{Str: "je", SubstringI: 9, Result: 17},
	{Str: "nk", SubstringI: -1, Result: 4},
	{Str: "unk", SubstringI: 11, Result: 1},
	{Str: "ánk", SubstringI: 11, Result: 2},
	{Str: "énk", SubstringI: 11, Result: 3},
	{Str: "ünk", SubstringI: 11, Result: 1},
	{Str: "uk", SubstringI: -1, Result: 8},
	{Str: "juk", SubstringI: 16, Result: 7},
	{Str: "ájuk", SubstringI: 17, Result: 5},
	{Str: "ük", SubstringI: -1, Result: 8},
	{Str: "jük", SubstringI: 19, Result: 7},
	{Str: "éjük", SubstringI: 20, Result: 6},
	{Str: "m", SubstringI: -1, Result: 12},
	{Str: "am", SubstringI: 22, Result: 9},
	{Str: "em", SubstringI: 22, Result: 9},
	{Str: "om", SubstringI: 22, Result: 9},
	{Str: "ám", SubstringI: 22, Result: 10},
	{Str: "ém", SubstringI: 22, Result: 11},
	{Str: "o", SubstringI: -1, Result: 18},
	{Str: "á", SubstringI: -1, Result: 19},
	{Str: "é", SubstringI: -1, Result: 20},
}

var a11 = []snowball.Among[context]{
	{Str: "id", SubstringI: -1, Result: 10},
	{Str: "aid", SubstringI: 0, Result: 9},
	{Str: "jaid", SubstringI: 1, Result: 6},
	{Str: "eid", SubstringI: 0, Result: 9},
	{Str: "jeid", SubstringI: 3, Result: 6},
	{Str: "áid", SubstringI: 0, Result: 7},
	{Str: "éid", SubstringI: 0, Result: 8},
	{Str: "i", SubstringI: -1, Result: 15},
	{Str: "ai", SubstringI: 7, Result: 14},
	{Str: "jai", SubstringI: 8, Result: 11},
	{Str: "ei", SubstringI: 7, Result: 14},
	{Str: "jei", SubstringI: 10, Result: 11},
	{Str: "ái", SubstringI: 7, Result: 12},
	{Str: "éi", SubstringI: 7, Result: 13},
	{Str: "itek", SubstringI: -1, Result: 24},
	{Str: "eitek", SubstringI: 14, Result: 21},
	{Str: "jeitek", SubstringI: 15, Result: 20},
	{Str: "éitek", SubstringI: 14, Result: 23},
	{Str: "ik", SubstringI: -1, Result: 29},
	{Str: "aik", SubstringI: 18, Result: 26},
	{Str: "jaik", SubstringI: 19, Result: 25},
	{Str: "eik", SubstringI: 18, Result: 26},
	{Str: "jeik", SubstringI: 21, Result: 25},
	{Str: "áik", SubstringI: 18, Result: 27},
	{Str: "éik", SubstringI: 18, Result: 28},
	{Str: "ink", SubstringI: -1, Result: 20},
	{Str: "aink", SubstringI: 25, Result: 17},
	{Str: "jaink", SubstringI: 26, Result: 16},
	{Str: "eink", SubstringI: 25, Result: 17},
	{Str: "jeink", SubstringI: 28, Result: 16},
	{Str: "áink", SubstringI: 25, Result: 18},
	{Str: "éink", SubstringI: 25, Result: 19},
	{Str: "aitok", SubstringI: -1, Result: 21},
	{Str: "jaitok", SubstringI: 32, Result: 20},
	{Str: "áitok", SubstringI: -1, Result: 22},
	{Str: "im", SubstringI: -1, Result: 5},
	{Str: "aim", SubstringI: 35, Result: 4},
	{Str: "jaim", SubstringI: 36, Result: 1},
	{Str: "eim", SubstringI: 35, Result: 4},
	{Str: "jeim", SubstringI: 38, Result: 1},
	{Str: "áim", SubstringI: 35, Result: 2},
	{Str: "éim", SubstringI: 35, Result: 3},
}

var gV = []byte{17, 65, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 17, 36, 10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1}

// markRegions sets the R1 region boundary (iP1): the position after the first
// non-vowel following the initial vowel sequence (or after a leading digraph).
func markRegions(env *snowball.Env, ctx *context) bool {
	ctx.iP1 = env.Limit
lab0:
	for {
		v1 := env.Cursor
	lab1:
		for {
			if !env.InGrouping(gV, 97, 369) {
				break lab1
			}
		golab2:
			for {
				v2 := env.Cursor
				for {
					if !env.OutGrouping(gV, 97, 369) {
						break
					}
					env.Cursor = v2
					break golab2
				}
				env.Cursor = v2
				if env.Cursor >= env.Limit {
					break lab1
				}
				env.NextChar()
			}
		lab4:
			for {
				v3 := env.Cursor
				for {
					if snowball.FindAmong(env, a0, ctx) == 0 {
						break
					}
					break lab4
				}
				env.Cursor = v3
				if env.Cursor >= env.Limit {
					break lab1
				}
				env.NextChar()
				break lab4
			}
			ctx.iP1 = env.Cursor
			break lab0
		}
		env.Cursor = v1
		if !env.OutGrouping(gV, 97, 369) {
			return false
		}
	golab6:
		for {
			for {
				if !env.InGrouping(gV, 97, 369) {
					break
				}
				break golab6
			}
			if env.Cursor >= env.Limit {
				return false
			}
			env.NextChar()
		}
		ctx.iP1 = env.Cursor
		break lab0
	}
	return true
}

// r1 reports whether the cursor is within region R1.
func r1(env *snowball.Env, ctx *context) bool {
	return ctx.iP1 <= env.Cursor
}

// vEnding rewrites a terminal á/é (left after a case suffix was removed) back to
// the plain vowel a/e when in R1.
func vEnding(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a1, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !r1(env, ctx) {
		return false
	}
	if amongVar == 0 {
		return false
	}
	switch amongVar {
	case 1:
		if !env.SliceFrom("a") {
			return false
		}
	case 2:
		if !env.SliceFrom("e") {
			return false
		}
	}
	return true
}

// double tests (without consuming) whether a doubled consonant sits before the
// cursor.
func double(env *snowball.Env, ctx *context) bool {
	v1 := env.Limit - env.Cursor
	if snowball.FindAmongB(env, a2, ctx) == 0 {
		return false
	}
	env.Cursor = env.Limit - v1
	return true
}

// undouble removes one half of a doubled consonant immediately before the
// cursor.
func undouble(env *snowball.Env, ctx *context) bool {
	if env.Cursor <= env.LimitBackward {
		return false
	}
	env.PreviousChar()
	env.Ket = env.Cursor
	c := env.ByteIndexForHop(-1)
	if env.LimitBackward > c || c > env.Limit {
		return false
	}
	env.Cursor = c
	env.Bra = env.Cursor
	if !env.SliceDel() {
		return false
	}
	return true
}

// instrum handles the instrumental -val/-vel suffixes (a3), undoubling the
// assimilated consonant.
func instrum(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a3, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !r1(env, ctx) {
		return false
	}
	if amongVar == 0 {
		return false
	}
	switch amongVar {
	case 1:
		if !double(env, ctx) {
			return false
		}
	case 2:
		if !double(env, ctx) {
			return false
		}
	}
	if !env.SliceDel() {
		return false
	}
	if !undouble(env, ctx) {
		return false
	}
	return true
}

// caseFn deletes the grammatical case suffixes (a4) in R1 and then restores a
// preceding á/é via vEnding.
func caseFn(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	if snowball.FindAmongB(env, a4, ctx) == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !r1(env, ctx) {
		return false
	}
	if !env.SliceDel() {
		return false
	}
	if !vEnding(env, ctx) {
		return false
	}
	return true
}

// caseSpecial handles the -án/-én/-ánként case endings (a5), rewriting them to
// a/e.
func caseSpecial(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a5, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !r1(env, ctx) {
		return false
	}
	if amongVar == 0 {
		return false
	}
	switch amongVar {
	case 1:
		if !env.SliceFrom("e") {
			return false
		}
	case 2:
		if !env.SliceFrom("a") {
			return false
		}
	case 3:
		if !env.SliceFrom("a") {
			return false
		}
	}
	return true
}

// caseOther handles the -stul/-stül case-ending family (a6).
func caseOther(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a6, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !r1(env, ctx) {
		return false
	}
	if amongVar == 0 {
		return false
	}
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
		if !env.SliceFrom("a") {
			return false
		}
	case 4:
		if !env.SliceFrom("e") {
			return false
		}
	}
	return true
}

// factive handles the factive -á/-é endings (a7), undoubling the assimilated
// consonant.
func factive(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a7, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !r1(env, ctx) {
		return false
	}
	if amongVar == 0 {
		return false
	}
	switch amongVar {
	case 1:
		if !double(env, ctx) {
			return false
		}
	case 2:
		if !double(env, ctx) {
			return false
		}
	}
	if !env.SliceDel() {
		return false
	}
	if !undouble(env, ctx) {
		return false
	}
	return true
}

// plural handles the plural -k suffixes (a8), normalising long vowels to a/e.
func plural(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a8, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !r1(env, ctx) {
		return false
	}
	if amongVar == 0 {
		return false
	}
	switch amongVar {
	case 1:
		if !env.SliceFrom("a") {
			return false
		}
	case 2:
		if !env.SliceFrom("e") {
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
		if !env.SliceDel() {
			return false
		}
	case 6:
		if !env.SliceDel() {
			return false
		}
	case 7:
		if !env.SliceDel() {
			return false
		}
	}
	return true
}

// owned handles the possessive -é (owned-thing) endings (a9).
func owned(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a9, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !r1(env, ctx) {
		return false
	}
	if amongVar == 0 {
		return false
	}
	switch amongVar {
	case 1:
		if !env.SliceDel() {
			return false
		}
	case 2:
		if !env.SliceFrom("e") {
			return false
		}
	case 3:
		if !env.SliceFrom("a") {
			return false
		}
	case 4:
		if !env.SliceDel() {
			return false
		}
	case 5:
		if !env.SliceFrom("e") {
			return false
		}
	case 6:
		if !env.SliceFrom("a") {
			return false
		}
	case 7:
		if !env.SliceDel() {
			return false
		}
	case 8:
		if !env.SliceFrom("e") {
			return false
		}
	case 9:
		if !env.SliceDel() {
			return false
		}
	}
	return true
}

// singOwner handles the singular-owner possessive suffixes (a10).
func singOwner(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a10, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !r1(env, ctx) {
		return false
	}
	if amongVar == 0 {
		return false
	}
	switch amongVar {
	case 1:
		if !env.SliceDel() {
			return false
		}
	case 2:
		if !env.SliceFrom("a") {
			return false
		}
	case 3:
		if !env.SliceFrom("e") {
			return false
		}
	case 4:
		if !env.SliceDel() {
			return false
		}
	case 5:
		if !env.SliceFrom("a") {
			return false
		}
	case 6:
		if !env.SliceFrom("e") {
			return false
		}
	case 7:
		if !env.SliceDel() {
			return false
		}
	case 8:
		if !env.SliceDel() {
			return false
		}
	case 9:
		if !env.SliceDel() {
			return false
		}
	case 10:
		if !env.SliceFrom("a") {
			return false
		}
	case 11:
		if !env.SliceFrom("e") {
			return false
		}
	case 12:
		if !env.SliceDel() {
			return false
		}
	case 13:
		if !env.SliceDel() {
			return false
		}
	case 14:
		if !env.SliceFrom("a") {
			return false
		}
	case 15:
		if !env.SliceFrom("e") {
			return false
		}
	case 16:
		if !env.SliceDel() {
			return false
		}
	case 17:
		if !env.SliceDel() {
			return false
		}
	case 18:
		if !env.SliceDel() {
			return false
		}
	case 19:
		if !env.SliceFrom("a") {
			return false
		}
	case 20:
		if !env.SliceFrom("e") {
			return false
		}
	}
	return true
}

// plurOwner handles the plural-owner possessive suffixes (a11).
func plurOwner(env *snowball.Env, ctx *context) bool {
	env.Ket = env.Cursor
	amongVar := snowball.FindAmongB(env, a11, ctx)
	if amongVar == 0 {
		return false
	}
	env.Bra = env.Cursor
	if !r1(env, ctx) {
		return false
	}
	if amongVar == 0 {
		return false
	}
	switch amongVar {
	case 1:
		if !env.SliceDel() {
			return false
		}
	case 2:
		if !env.SliceFrom("a") {
			return false
		}
	case 3:
		if !env.SliceFrom("e") {
			return false
		}
	case 4:
		if !env.SliceDel() {
			return false
		}
	case 5:
		if !env.SliceDel() {
			return false
		}
	case 6:
		if !env.SliceDel() {
			return false
		}
	case 7:
		if !env.SliceFrom("a") {
			return false
		}
	case 8:
		if !env.SliceFrom("e") {
			return false
		}
	case 9:
		if !env.SliceDel() {
			return false
		}
	case 10:
		if !env.SliceDel() {
			return false
		}
	case 11:
		if !env.SliceDel() {
			return false
		}
	case 12:
		if !env.SliceFrom("a") {
			return false
		}
	case 13:
		if !env.SliceFrom("e") {
			return false
		}
	case 14:
		if !env.SliceDel() {
			return false
		}
	case 15:
		if !env.SliceDel() {
			return false
		}
	case 16:
		if !env.SliceDel() {
			return false
		}
	case 17:
		if !env.SliceDel() {
			return false
		}
	case 18:
		if !env.SliceFrom("a") {
			return false
		}
	case 19:
		if !env.SliceFrom("e") {
			return false
		}
	case 20:
		if !env.SliceDel() {
			return false
		}
	case 21:
		if !env.SliceDel() {
			return false
		}
	case 22:
		if !env.SliceFrom("a") {
			return false
		}
	case 23:
		if !env.SliceFrom("e") {
			return false
		}
	case 24:
		if !env.SliceDel() {
			return false
		}
	case 25:
		if !env.SliceDel() {
			return false
		}
	case 26:
		if !env.SliceDel() {
			return false
		}
	case 27:
		if !env.SliceFrom("a") {
			return false
		}
	case 28:
		if !env.SliceFrom("e") {
			return false
		}
	case 29:
		if !env.SliceDel() {
			return false
		}
	}
	return true
}

// Stem runs the Snowball hungarian algorithm over env, mirroring the generated
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
	instrum(env, ctx)
	env.Cursor = env.Limit - v2

	v3 := env.Limit - env.Cursor
	caseFn(env, ctx)
	env.Cursor = env.Limit - v3

	v4 := env.Limit - env.Cursor
	caseSpecial(env, ctx)
	env.Cursor = env.Limit - v4

	v5 := env.Limit - env.Cursor
	caseOther(env, ctx)
	env.Cursor = env.Limit - v5

	v6 := env.Limit - env.Cursor
	factive(env, ctx)
	env.Cursor = env.Limit - v6

	v7 := env.Limit - env.Cursor
	owned(env, ctx)
	env.Cursor = env.Limit - v7

	v8 := env.Limit - env.Cursor
	singOwner(env, ctx)
	env.Cursor = env.Limit - v8

	v9 := env.Limit - env.Cursor
	plurOwner(env, ctx)
	env.Cursor = env.Limit - v9

	v10 := env.Limit - env.Cursor
	plural(env, ctx)
	env.Cursor = env.Limit - v10

	env.Cursor = env.LimitBackward
	return true
}
