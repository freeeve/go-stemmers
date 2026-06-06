//! bench — a simple wall-clock throughput probe for rust-stemmers, so the Go
//! port's `go test -bench` numbers can be compared against the reference Rust on
//! the same corpus and the same work (stem every word, repeated).
//!
//! Usage: bench <algorithm> <corpus-file> [passes]
//! Reports ns/word and words/s over `passes` sweeps of the corpus (default 200).
//! Not Criterion — just a black_box'd loop — but enough to gauge the gap.

use std::hint::black_box;
use std::time::Instant;

use rust_stemmers::{Algorithm, Stemmer};

fn main() {
    let mut args = std::env::args().skip(1);
    let algo = args.next().unwrap_or_else(|| "english".to_string());
    let path = args
        .next()
        .expect("usage: bench <algorithm> <corpus-file> [passes]");
    let passes: u64 = args.next().and_then(|s| s.parse().ok()).unwrap_or(200);

    let stemmer = Stemmer::create(match algo.as_str() {
        "english" | "en" => Algorithm::English,
        "russian" | "ru" => Algorithm::Russian,
        "turkish" | "tr" => Algorithm::Turkish,
        "arabic" | "ar" => Algorithm::Arabic,
        "german" | "ger" | "de" => Algorithm::German,
        other => panic!("unsupported algorithm for bench: {other}"),
    });

    let text = std::fs::read_to_string(&path).expect("read corpus");
    let words: Vec<&str> = text.lines().filter(|l| !l.is_empty()).collect();

    // Warm up (caches, branch predictor).
    for w in &words {
        black_box(stemmer.stem(black_box(w)));
    }

    let start = Instant::now();
    let mut n: u64 = 0;
    for _ in 0..passes {
        for w in &words {
            black_box(stemmer.stem(black_box(w)));
            n += 1;
        }
    }
    let elapsed = start.elapsed();
    let ns_per = elapsed.as_nanos() as f64 / n as f64;
    eprintln!(
        "rust {algo}: {} words x {passes} passes = {n} stems in {:.2?}",
        words.len(),
        elapsed
    );
    println!(
        "rust {algo}: {ns_per:.1} ns/word  {:.0} words/s",
        1e9 / ns_per
    );
}
