// Package snowball is the runtime the generated Snowball stemmers execute
// against — a byte-faithful Go port of rust-stemmers' `snowball` module
// (snowball_env.rs + among.rs). Each language package (internal/<lang>) ports
// the corresponding generated algorithm file and drives an [Env] through this
// runtime, so the Go stemmers produce output identical to rust-stemmers.
package snowball

// Among is one entry of a Snowball "among" table: a search string, a
// back-pointer to the entry that is its longest proper prefix (SubstringI, or
// -1 when there is none), a result code the algorithm switches on, and an
// optional associated routine that must succeed for the entry to match.
//
// It mirrors rust-stemmers' `Among<T>(&'static str, i32, i32, Option<method>)`.
// The generic parameter T is the per-language context the routine reads/writes.
type Among[T any] struct {
	// Str is the search string compared against the input (forward in
	// [FindAmong], reversed in [FindAmongB]).
	Str string
	// SubstringI is the index of the among entry that is Str's longest proper
	// prefix, or -1 if none. Used to walk shorter candidates after a partial
	// match during the linear-probe phase of the search.
	SubstringI int32
	// Result is the value returned to the algorithm when this entry matches.
	Result int32
	// Method is the optional associated routine; when non-nil it must return
	// true for the entry to be accepted. Nil for entries with no routine.
	Method func(*Env, *T) bool
}
