//! rustgen — the rust-stemmers oracle for go-stemmers.
//!
//! Reads one word per line from stdin and writes its stem per line to stdout,
//! using the rust-stemmers algorithm named by the single CLI argument. This is
//! the ground truth the Go port is diffed against: it generates the committed
//! `res_<lang>.txt` goldens for languages rust-stemmers ships no test vectors
//! for, regenerates any golden after a version bump, and is the oracle the Go
//! fuzz test compares against.
//!
//! Usage:
//!   rustgen <algorithm> < words.txt > stems.txt
//!
//! `<algorithm>` is a case-insensitive Snowball language name (e.g. "english",
//! "danish", "turkish") or its short test-data code ("en", "ger", "no", ...).
//! Input is expected already lowercased, matching the stemmer contract.

use std::io::{self, BufRead, BufWriter, Write};
use std::process::exit;

use rust_stemmers::{Algorithm, Stemmer};

/// Maps a language name or short code to a rust-stemmers [`Algorithm`].
fn algorithm(name: &str) -> Option<Algorithm> {
    let n = name.to_lowercase();
    let a = match n.as_str() {
        "arabic" | "ar" => Algorithm::Arabic,
        "danish" | "da" => Algorithm::Danish,
        "dutch" | "nl" => Algorithm::Dutch,
        "english" | "en" => Algorithm::English,
        "finnish" | "fi" => Algorithm::Finnish,
        "french" | "fr" => Algorithm::French,
        "german" | "ger" | "de" => Algorithm::German,
        "greek" | "el" => Algorithm::Greek,
        "hungarian" | "hu" => Algorithm::Hungarian,
        "italian" | "it" => Algorithm::Italian,
        "norwegian" | "no" | "nb" => Algorithm::Norwegian,
        "portuguese" | "pt" => Algorithm::Portuguese,
        "romanian" | "ro" => Algorithm::Romanian,
        "russian" | "ru" => Algorithm::Russian,
        "spanish" | "es" => Algorithm::Spanish,
        "swedish" | "sv" => Algorithm::Swedish,
        "tamil" | "ta" => Algorithm::Tamil,
        "turkish" | "tr" => Algorithm::Turkish,
        _ => return None,
    };
    Some(a)
}

fn main() {
    let arg = match std::env::args().nth(1) {
        Some(a) => a,
        None => {
            eprintln!("usage: rustgen <algorithm> < words.txt > stems.txt");
            exit(2);
        }
    };
    let algo = match algorithm(&arg) {
        Some(a) => a,
        None => {
            eprintln!("rustgen: unknown algorithm {arg:?}");
            exit(2);
        }
    };
    let stemmer = Stemmer::create(algo);

    let stdin = io::stdin();
    let stdout = io::stdout();
    let mut out = BufWriter::new(stdout.lock());
    for line in stdin.lock().lines() {
        let word = line.expect("read stdin");
        out.write_all(stemmer.stem(&word).as_bytes())
            .expect("write stdout");
        out.write_all(b"\n").expect("write stdout");
    }
    out.flush().expect("flush stdout");
}
