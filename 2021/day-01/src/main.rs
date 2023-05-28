use std::{io};
use std::fs::File;
use std::io::{BufRead, BufReader, Lines, Result};
use std::path::Path;

fn main() {
    first();
    second();
}

fn lines_from_file(filename: impl AsRef<Path>) -> io::Result<Vec<String>> {
    BufReader::new(File::open(filename)?).lines().collect()
}

fn second() {
    let lines = lines_from_file("files/input.txt").expect("Could not load lines");
    let mut increases = 0;
    let mut previous: Option<i32> = None;
    for (pos, e) in lines.iter().enumerate() {
        let mut window = e.parse::<i32>().unwrap();

        if lines.get(pos + 1).is_some() {
            window = window + lines.get(pos + 1).unwrap().parse::<i32>().unwrap();
        }

        if lines.get(pos + 2).is_some() {
            window = window + lines.get(pos + 2).unwrap().parse::<i32>().unwrap();
        }

        if previous.is_some() && window > previous.unwrap() {
            increases = increases + 1;
        }
        previous = Some(window);
    }
    println!("Increases #2: {}", increases);
}

fn first() {
    let filename = "files/input.txt";

    if let Ok(lines) = read_lines(filename) {
        let mut previous: Option<i32> = None;
        let mut increases = 0;

        for line in lines {
            if let Ok(ip) = line {
                let current = ip.parse::<i32>().unwrap();

                if previous.is_some() && current > previous.unwrap() {
                    increases = increases + 1;
                }
                previous = Some(current);
            }
        }
        println!("Increases #1: {}", increases);
    }
}

fn read_lines<P>(filename: P) -> Result<Lines<BufReader<File>>>
    where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}
