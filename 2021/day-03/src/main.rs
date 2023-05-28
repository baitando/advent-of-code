use std::{io};
use std::borrow::Borrow;
use std::fs::File;
use std::io::{BufRead, BufReader, Lines, Result};
use std::path::Path;

fn main() {
    first();
    second();
}

fn first() {
    let lines = lines_from_file("files/input.txt").expect("Could not load lines");
    let entry_count = lines.len();
    let bit_count = lines.get(0).unwrap().len();
    let mut gamma_rate_binary: String = "".to_owned();
    let mut epsilon_rate_binary: String = "".to_owned();

    for bit in 0..bit_count {
        let mut current_bit_count1 = 0;
        for (_pos, e) in lines.iter().enumerate() {
            if e.as_bytes()[bit] as char == '1' {
                current_bit_count1 = current_bit_count1 + 1;
            }
        }
        let mut gamma = "0";
        let mut epsilon = "1";
        if current_bit_count1 > entry_count / 2 {
            gamma = "1";
            epsilon = "0";
        }

        gamma_rate_binary.push_str(gamma);
        epsilon_rate_binary.push_str(epsilon);
    }
    let gamma_rate = isize::from_str_radix(gamma_rate_binary.as_str(), 2).unwrap();
    let epsilon_rate = isize::from_str_radix(epsilon_rate_binary.as_str(), 2).unwrap();
    println!("#1 gamma rate is {} which is converted {}", gamma_rate_binary, gamma_rate);
    println!("#1 epsilon rate is {} which is converted {}", epsilon_rate_binary, epsilon_rate);
    println!("#1 result is {}", gamma_rate * epsilon_rate);
}

fn second() {
    let lines = lines_from_file("files/input.txt").expect("Could not load lines");
    let bit_count = lines.get(0).unwrap().len();

    let mut remaining: Vec<String> = lines.to_vec();
    for bit in 0..bit_count {
        let entry_count = remaining.len() as i32;
        let count1 = count_1_in_position(&remaining, bit);
        let count0 = entry_count - count1;

        let mut bit_mask = '0';
        if count0 == count1 {
            bit_mask = '1';
        } else if count1 > count0 {
            bit_mask = '1';
        }
        remaining = filter_by_value(&remaining, bit, bit_mask);
        if remaining.len() == 1 {
            break;
        }
    }
    let oxygen_binary = remaining.get(0).unwrap().to_owned();

    remaining= lines.to_vec();
    for bit in 0..bit_count {
        let entry_count = remaining.len() as i32;
        let count1 = count_1_in_position(&remaining, bit);
        let count0 = entry_count - count1;

        let mut bit_mask = '0';
        if count0 == count1 {
            bit_mask = '0';
        } else if count1 < count0 {
            bit_mask = '1';
        }
        remaining = filter_by_value(&remaining, bit, bit_mask);
        if remaining.len() == 1 {
            break;
        }
    }
    let co2_binary = remaining.get(0).unwrap().to_owned();

    let co2 = isize::from_str_radix(co2_binary.as_str(), 2).unwrap();
    let oxygen = isize::from_str_radix(oxygen_binary.as_str(), 2).unwrap();

    println!("#2 oxygen is {} which is converted {}", oxygen_binary, oxygen);
    println!("#2 CO2 is {} which is converted {}", co2_binary, co2);
    println!("#2 result is {}", co2 * oxygen);
}

fn filter_by_value(lines: &Vec<String>, position: usize, value: char) -> Vec<String> {
    let mut result: Vec<String> = vec![];
    for (_pos, e) in lines.iter().enumerate() {
        if e.as_bytes()[position] as char == value {
            result.push(e.to_owned());
        }
    }
    return result;
}

fn count_1_in_position(lines: &Vec<String>, position: usize) -> i32 {
    let mut current_bit_count1 = 0;
    for (_pos, e) in lines.iter().enumerate() {
        if e.as_bytes()[position] as char == '1' {
            current_bit_count1 = current_bit_count1 + 1;
        }
    }
    return current_bit_count1;
}

fn lines_from_file(filename: impl AsRef<Path>) -> io::Result<Vec<String>> {
    BufReader::new(File::open(filename)?).lines().collect()
}
