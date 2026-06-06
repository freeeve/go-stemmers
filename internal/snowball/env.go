package snowball

import (
	"bytes"
	"unicode/utf8"
)

// Env is the mutable state a Snowball program operates on: the working string
// plus the cursor/limit/bra/ket byte offsets the generated code manipulates. It
// is a byte-faithful port of rust-stemmers' `SnowballEnv`.
//
// Indices are byte offsets into Current (UTF-8), and every offset the algorithm
// produces lands on a UTF-8 character boundary, exactly as the Rust runtime
// maintains. Offsets are kept as signed int so the boundary-scan and hop
// arithmetic mirror the i32 casts in the Rust source without unsigned underflow.
type Env struct {
	// Current is the working string as raw UTF-8 bytes; slice operations
	// rewrite it in place via replaceS.
	Current []byte
	// Cursor is the current byte position the program reads from / writes at.
	Cursor int
	// Limit is the forward end of the region under consideration.
	Limit int
	// LimitBackward is the backward end (set when the program runs backwards).
	LimitBackward int
	// Bra and Ket bracket the slice a replace/delete operates on.
	Bra int
	Ket int
}

// NewEnv creates an Env over value, mirroring `SnowballEnv::create`.
func NewEnv(value string) *Env {
	b := []byte(value)
	return &Env{
		Current:       b,
		Cursor:        0,
		Limit:         len(b),
		LimitBackward: 0,
		Bra:           0,
		Ket:           len(b),
	}
}

// GetCurrent returns the working string, mirroring `get_current`.
func (e *Env) GetCurrent() string {
	return string(e.Current)
}

// RuneCount returns the number of Unicode characters in the working string,
// mirroring `self.current.chars().count()` (used by some algorithms to gate on
// word length).
func (e *Env) RuneCount() int {
	return utf8.RuneCount(e.Current)
}

// isCharBoundary reports whether i is a UTF-8 character boundary of Current,
// matching Rust's `str::is_char_boundary` (true at 0 and len, false past len, a
// non-continuation byte otherwise).
func (e *Env) isCharBoundary(i int) bool {
	if i == 0 || i == len(e.Current) {
		return true
	}
	if i < 0 || i > len(e.Current) {
		return false
	}
	return e.Current[i]&0xC0 != 0x80
}

// replaceS replaces Current[bra:ket] with s and fixes up limit/cursor, mirroring
// `replace_s`. Returns the length adjustment (len(s) - (ket-bra)).
func (e *Env) replaceS(bra, ket int, s string) int {
	adjustment := len(s) - (ket - bra)
	result := make([]byte, 0, len(e.Current)+adjustment)
	result = append(result, e.Current[:bra]...)
	result = append(result, s...)
	result = append(result, e.Current[ket:]...)
	e.Limit += adjustment
	if e.Cursor >= ket {
		e.Cursor += adjustment
	} else if e.Cursor > bra {
		e.Cursor = bra
	}
	e.Current = result
	return adjustment
}

// EqS checks whether s is at the cursor (forwards); if so it advances the cursor
// past s (to the next char boundary) and returns true. Mirrors `eq_s`.
func (e *Env) EqS(s string) bool {
	if e.Cursor >= e.Limit {
		return false
	}
	if bytes.HasPrefix(e.Current[e.Cursor:], []byte(s)) {
		e.Cursor += len(s)
		for !e.isCharBoundary(e.Cursor) {
			e.Cursor++
		}
		return true
	}
	return false
}

// EqSB checks whether s is immediately before the cursor (backwards); if so it
// moves the cursor back to the start of s and returns true. Mirrors `eq_s_b`.
func (e *Env) EqSB(s string) bool {
	if e.Cursor-e.LimitBackward < len(s) {
		return false
	}
	if !e.isCharBoundary(e.Cursor-len(s)) || !bytes.HasPrefix(e.Current[e.Cursor-len(s):], []byte(s)) {
		return false
	}
	e.Cursor -= len(s)
	return true
}

// SliceFrom replaces Current[bra:ket] with s and returns true. Mirrors
// `slice_from`.
func (e *Env) SliceFrom(s string) bool {
	e.replaceS(e.Bra, e.Ket, s)
	return true
}

// SliceDel deletes Current[bra:ket]. Mirrors `slice_del`.
func (e *Env) SliceDel() bool {
	return e.SliceFrom("")
}

// SliceTo returns Current[bra:ket]. Mirrors `slice_to`.
func (e *Env) SliceTo() string {
	return string(e.Current[e.Bra:e.Ket])
}

// Insert splices s into Current[bra:ket] and shifts the env's own bra/ket by the
// resulting adjustment when they sit at/after bra. Mirrors `insert`.
func (e *Env) Insert(bra, ket int, s string) {
	adjustment := e.replaceS(bra, ket, s)
	if bra <= e.Bra {
		e.Bra += adjustment
	}
	if bra <= e.Ket {
		e.Ket += adjustment
	}
}

// NextChar advances the cursor to the next character boundary. Mirrors
// `next_char`.
func (e *Env) NextChar() {
	e.Cursor++
	for !e.isCharBoundary(e.Cursor) {
		e.Cursor++
	}
}

// PreviousChar moves the cursor to the previous character boundary. Mirrors
// `previous_char`.
func (e *Env) PreviousChar() {
	e.Cursor--
	for !e.isCharBoundary(e.Cursor) {
		e.Cursor--
	}
}

// ByteIndexForHop returns the byte index delta characters from the cursor
// (delta may be negative), mirroring `byte_index_for_hop`. The caller validates
// the result against the limits before committing it to the cursor.
func (e *Env) ByteIndexForHop(delta int) int {
	if delta > 0 {
		res := e.Cursor
		for delta > 0 {
			res++
			delta--
			for res <= len(e.Current) && !e.isCharBoundary(res) {
				res++
			}
		}
		return res
	} else if delta < 0 {
		res := e.Cursor
		for delta < 0 {
			res--
			delta++
			for res >= 0 && !e.isCharBoundary(res) {
				res--
			}
		}
		return res
	}
	return e.Cursor
}

// runeAtCursor decodes the rune at the cursor, mirroring
// `self.current[self.cursor..].chars().next()`.
func (e *Env) runeAtCursor() (rune, bool) {
	if e.Cursor >= len(e.Current) {
		return 0, false
	}
	r, _ := utf8.DecodeRune(e.Current[e.Cursor:])
	return r, true
}

// inGrouping reports whether the rune at the cursor is in the grouping bitfield
// chars (codepoints min..=max); on success it advances one char. Mirrors
// `in_grouping`.
func (e *Env) InGrouping(chars []byte, min, max uint32) bool {
	if e.Cursor >= e.Limit {
		return false
	}
	if r, ok := e.runeAtCursor(); ok {
		ch := uint32(r)
		if ch > max || ch < min {
			return false
		}
		ch -= min
		if chars[ch>>3]&(0x1<<(ch&0x7)) == 0 {
			return false
		}
		e.NextChar()
		return true
	}
	return false
}

// InGroupingB is the backward form of [Env.InGrouping]. Mirrors `in_grouping_b`.
func (e *Env) InGroupingB(chars []byte, min, max uint32) bool {
	if e.Cursor <= e.LimitBackward {
		return false
	}
	e.PreviousChar()
	if r, ok := e.runeAtCursor(); ok {
		ch := uint32(r)
		e.NextChar()
		if ch > max || ch < min {
			return false
		}
		ch -= min
		if chars[ch>>3]&(0x1<<(ch&0x7)) == 0 {
			return false
		}
		e.PreviousChar()
		return true
	}
	return false
}

// OutGrouping reports whether the rune at the cursor is NOT in the grouping; on
// success it advances one char. Mirrors `out_grouping`.
func (e *Env) OutGrouping(chars []byte, min, max uint32) bool {
	if e.Cursor >= e.Limit {
		return false
	}
	if r, ok := e.runeAtCursor(); ok {
		ch := uint32(r)
		if ch > max || ch < min {
			e.NextChar()
			return true
		}
		ch -= min
		if chars[ch>>3]&(0x1<<(ch&0x7)) == 0 {
			e.NextChar()
			return true
		}
	}
	return false
}

// OutGroupingB is the backward form of [Env.OutGrouping]. Mirrors
// `out_grouping_b`.
func (e *Env) OutGroupingB(chars []byte, min, max uint32) bool {
	if e.Cursor <= e.LimitBackward {
		return false
	}
	e.PreviousChar()
	if r, ok := e.runeAtCursor(); ok {
		ch := uint32(r)
		e.NextChar()
		if ch > max || ch < min {
			e.PreviousChar()
			return true
		}
		ch -= min
		if chars[ch>>3]&(0x1<<(ch&0x7)) == 0 {
			e.PreviousChar()
			return true
		}
	}
	return false
}

// minInt returns the smaller of a and b.
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// FindAmong searches the among table forwards from the cursor and returns the
// matched entry's Result (0 if none), running any associated routine. It is a
// byte-faithful port of `find_among`.
func FindAmong[T any](e *Env, amongs []Among[T], ctx *T) int32 {
	i := int32(0)
	j := int32(len(amongs))

	c := e.Cursor
	l := e.Limit

	commonI := 0
	commonJ := 0

	firstKeyInspected := false
	for {
		k := i + ((j - i) >> 1)
		diff := 0
		common := minInt(commonI, commonJ)
		w := &amongs[k]
		for lvar := common; lvar < len(w.Str); lvar++ {
			if c+common == l {
				diff = -1
				break
			}
			diff = int(e.Current[c+common]) - int(w.Str[lvar])
			if diff != 0 {
				break
			}
			common++
		}
		if diff < 0 {
			j = k
			commonJ = common
		} else {
			i = k
			commonI = common
		}
		if j-i <= 1 {
			if i > 0 {
				break
			}
			if j == i {
				break
			}
			if firstKeyInspected {
				break
			}
			firstKeyInspected = true
		}
	}

	for {
		w := &amongs[i]
		if commonI >= len(w.Str) {
			e.Cursor = c + len(w.Str)
			if w.Method != nil {
				res := w.Method(e, ctx)
				e.Cursor = c + len(w.Str)
				if res {
					return w.Result
				}
			} else {
				return w.Result
			}
		}
		i = w.SubstringI
		if i < 0 {
			return 0
		}
	}
}

// FindAmongB searches the among table backwards from the cursor and returns the
// matched entry's Result (0 if none), running any associated routine. It is a
// byte-faithful port of `find_among_b`.
func FindAmongB[T any](e *Env, amongs []Among[T], ctx *T) int32 {
	i := int32(0)
	j := int32(len(amongs))

	c := e.Cursor
	lb := e.LimitBackward

	commonI := 0
	commonJ := 0

	firstKeyInspected := false
	for {
		k := i + ((j - i) >> 1)
		diff := 0
		common := minInt(commonI, commonJ)
		w := &amongs[k]
		for lvar := len(w.Str) - common - 1; lvar >= 0; lvar-- {
			if c-common == lb {
				diff = -1
				break
			}
			diff = int(e.Current[c-common-1]) - int(w.Str[lvar])
			if diff != 0 {
				break
			}
			common++
		}
		if diff < 0 {
			j = k
			commonJ = common
		} else {
			i = k
			commonI = common
		}
		if j-i <= 1 {
			if i > 0 {
				break
			}
			if j == i {
				break
			}
			if firstKeyInspected {
				break
			}
			firstKeyInspected = true
		}
	}

	for {
		w := &amongs[i]
		if commonI >= len(w.Str) {
			e.Cursor = c - len(w.Str)
			if w.Method != nil {
				res := w.Method(e, ctx)
				e.Cursor = c - len(w.Str)
				if res {
					return w.Result
				}
			} else {
				return w.Result
			}
		}
		i = w.SubstringI
		if i < 0 {
			return 0
		}
	}
}
