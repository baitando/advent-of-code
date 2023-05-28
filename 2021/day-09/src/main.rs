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
    let heights = extract_input(&lines);

    for row_index in 0..heights.len() {
        for column_index in 0..heights.get(row_index).unwrap().len() {
            print!("{}", heights.get(row_index).unwrap().get(column_index).unwrap());
        }
        println!();
    }
}

fn extract_input(lines: &Vec<String>) -> Vec<Vec<i32>> {
    let mut rows: Vec<Vec<i32>> = vec![];
    for line in lines {
        let mut row: Vec<i32> = vec![];
        for height in line.chars() {
            row.push(height.to_digit(10).unwrap() as i32);
        }
        rows.push(row);
    }
    return rows;
}

fn is_lowest(row_index: usize, column_index: usize, heights: Vec<Vec<i32>>) -> bool {
    let mut is_lowest = true;
    let mut value_to_check = heights.get(row_index).unwrap().get(column_index).unwrap();
    is_lowest = is_lowest &&
        !(heights.len() > 1 &&
            heights.get(row_index - 1).len > 1 &&
            heights.get(row_index - 1).unwrap().get(column_index - 1).unwrap() < value_to_check
        );
    // TODO

    return is_lowest;
}

fn lines_from_file(filename: impl AsRef<Path>) -> io::Result<Vec<String>> {
    return BufReader::new(File::open(filename)?).lines().collect();
}
