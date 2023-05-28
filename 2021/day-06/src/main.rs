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
    let mut ages = extract_input(&lines);
    let mut day = 0u32;
    let mut amount_days = 80;

    loop {
        day += 1;

        let mut new_ages_first: Vec<i32> = vec![];
        let mut new_ages_last: Vec<i32> = vec![];

        for current_age in ages {
            if current_age == 0 {
                new_ages_first.push(6);
                new_ages_last.push(8);
            } else {
                new_ages_first.push(current_age - 1);
            }
        }
        ages = vec![];
        ages.append(&mut new_ages_first);
        ages.append(&mut new_ages_last);

        if day == amount_days {
            break;
        }
    }

    println!("After {} days there are {} fishes", amount_days, ages.len());
}

fn second() {
    let lines = lines_from_file("files/input.txt").expect("Could not load lines");
    let mut map = lines.get(0).unwrap()
        .split(',')
        .fold([0; 9], |mut map, n| {
            map[n.parse::<usize>().unwrap()] += 1;
            map
        });

    (1..256).for_each(|day| map[(day + 7) % 9] += map[day % 9]);

    println!("{}", map.iter().sum::<usize>());
}

fn extract_input(lines: &Vec<String>) -> Vec<i32> {
    return lines.get(0).unwrap().split(",").map(|number| number.parse().unwrap()).collect();
}

fn lines_from_file(filename: impl AsRef<Path>) -> io::Result<Vec<String>> {
    return BufReader::new(File::open(filename)?).lines().collect();
}
