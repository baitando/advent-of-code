use std::{io};
use std::borrow::Borrow;
use std::fs::File;
use std::io::{BufRead, BufReader, Lines, Result};
use std::path::Path;

fn main() {
    first();
}

fn first() {
    let lines = lines_from_file("files/input.txt").expect("Could not load lines");
    let input = extract_input(&lines);

    let mut boards = extract_boards(&lines);
    for current in &input {
        for board in &mut boards {
            board.add_value(&current);
        }
    }

    for number in input {}
}

fn extract_input(lines: Vector<String>) -> [[i32;]]

fn lines_from_file(filename: impl AsRef<Path>) -> io::Result<Vec<String>> {
    BufReader::new(File::open(filename)?).lines().collect()
}
