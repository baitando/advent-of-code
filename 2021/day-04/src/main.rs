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

fn extract_input(lines: &Vec<String>) -> Vec<i32> {
    return lines.get(0).unwrap().split(",").map(|number| number.parse().unwrap()).collect();
}

fn extract_boards(lines: &Vec<String>) -> Vec<Board> {
    let mut result: Vec<Board> = Vec::new();
    let (_left, right) = lines.split_at(2);

    let mut current_board = Board::new();
    let mut current_row = 0;

    for line in right {
        if current_row == 5 {
            result.push(current_board);
            current_board = Board::new();
            current_row = 0;
            continue;
        }

        let mut cleaned_line = line.replace("  ", " ");
        cleaned_line = cleaned_line.trim().to_string();

        let line_numbers: Vec<i32> = cleaned_line.split(" ").map(|number| number.parse().unwrap()).collect();
        for (current_column, value) in line_numbers.iter().enumerate() {
            current_board.update_value(current_row, current_column, *value);
        }
        current_row = current_row + 1;
    }

    return result;
}

fn lines_from_file(filename: impl AsRef<Path>) -> io::Result<Vec<String>> {
    BufReader::new(File::open(filename)?).lines().collect()
}

#[derive(Copy, Clone)]
struct Board {
    values: [[i32; 5]; 5],
    hits: [[bool; 5]; 5],
}

impl Board {
    fn new() -> Self {
        Board {
            values: [[0i32; 5]; 5],
            hits: [[false; 5]; 5],
        }
    }

    fn update_value(&mut self, x_pos: usize, y_pos: usize, value: i32) {
        self.values[x_pos][y_pos] = value;
    }

    fn add_value(&mut self, value: &i32) {
        for (row_index, row_value) in self.values.iter().enumerate() {
            for (column_index, cell_value) in row_value.iter().enumerate() {
                if cell_value.eq(value) {
                    println!("Found a match for {}", cell_value);
                    self.hits[column_index][row_index];
                }
            }
        }
    }

    fn check_rows(&mut self) {
        for (row_index, row_value) in self.values.iter().enumerate() {
            for (column_index, cell_value) in row_value.iter().enumerate() {
                if cell_value
            }
        }
    }

    fn check_columns(&mut self) {

    }
}
