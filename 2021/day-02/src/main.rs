use std::{io};
use std::fs::File;
use std::io::{BufRead, BufReader, Lines, Result};
use std::path::Path;

fn main() {
    first();
    second();
}

fn first() {
    let lines = lines_from_file("files/input.txt").expect("Could not load lines");
    let mut horizontal = 0;
    let mut depth = 0;

    for (_pos, e) in lines.iter().enumerate() {

        if e.starts_with("forward") {
            horizontal  = horizontal + e.replace("forward ", "").parse::<i32>().unwrap();

        }

        if e.starts_with("up") {
            depth  = depth - e.replace("up ", "").parse::<i32>().unwrap();

        }

        if e.starts_with("down") {
            depth  = depth + e.replace("down ", "").parse::<i32>().unwrap();

        }
    }
    println!("#1 horizontal: {}", horizontal);
    println!("#1 depth: {}", depth);
    println!("#1 result: {}", horizontal * depth);
}

fn second() {
    let lines = lines_from_file("files/input.txt").expect("Could not load lines");
    let mut horizontal = 0;
    let mut aim = 0;
    let mut depth = 0;

    for (_pos, e) in lines.iter().enumerate() {

        if e.starts_with("forward") {
            let forward = e.replace("forward ", "").parse::<i32>().unwrap();
            horizontal  = horizontal + forward;
            depth = depth + forward * aim;
        }

        if e.starts_with("up") {
            aim  = aim - e.replace("up ", "").parse::<i32>().unwrap();

        }

        if e.starts_with("down") {
            aim  = aim + e.replace("down ", "").parse::<i32>().unwrap();

        }
    }
    println!("#2 horizontal: {}", horizontal);
    println!("#2 depth: {}", depth);
    println!("#2 result: {}", horizontal * depth);
}

fn lines_from_file(filename: impl AsRef<Path>) -> io::Result<Vec<String>> {
    BufReader::new(File::open(filename)?).lines().collect()
}
