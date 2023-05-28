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
    let positions = extract_input(&lines);

    let min = get_min(positions.borrow());
    let max = get_max(positions.borrow());

    let mut min_pos = None;
    let mut min_distance = None;
    for target_pos in min.clone()..max.clone() {
        let distance = calc_distance(&target_pos, positions.borrow());
        if min_distance == None {
            min_distance = Some(distance);
            min_pos = Some(target_pos);
        } else {
            if distance < min_distance.unwrap() {
                min_distance = Some(distance);
                min_pos = Some(target_pos);
            }
        }
    }
    println!("#1 Minimum distance is on target position {}. Total distance is {}", min_pos.unwrap(), min_distance.unwrap());
}

fn second() {
    let lines = lines_from_file("files/input.txt").expect("Could not load lines");
    let positions = extract_input(&lines);

    let min = get_min(positions.borrow());
    let max = get_max(positions.borrow());

    let mut min_pos = None;
    let mut min_distance = None;
    for target_pos in min.clone()..max.clone() {
        let distance = calc_fuel(&target_pos, positions.borrow());
        if min_distance == None {
            min_distance = Some(distance);
            min_pos = Some(target_pos);
        } else {
            if distance < min_distance.unwrap() {
                min_distance = Some(distance);
                min_pos = Some(target_pos);
            }
        }
    }
    println!("#2 Minimum distance is on target position {}. Total distance is {}", min_pos.unwrap(), min_distance.unwrap());
}

fn calc_distance(target_pos: &i32, positions: &Vec<i32>) -> i32 {
    let mut distance: i32 = 0;
    for position in positions {
        distance += (position - target_pos).abs();
    }
    return distance;
}

fn calc_fuel(target_pos: &i32, positions: &Vec<i32>) -> i32 {
    let mut total_fuel: i32 = 0;
    for position in positions {
        let distance = (position - target_pos).abs();
        let mut fuel =  0;
        for value in 1..distance.clone()+1 {
            fuel += value;
        }
        total_fuel += fuel;
    }
    return total_fuel;
}

fn get_min(positions: &Vec<i32>) -> &i32 {
    let mut current_min = positions.get(0).unwrap();
    for position in positions {
        if position < current_min {
            current_min = position;
        }
    }
    return current_min;
}

fn get_max(positions: &Vec<i32>) -> &i32 {
    let mut current_max = positions.get(0).unwrap();
    for position in positions {
        if position > current_max {
            current_max = position;
        }
    }
    return current_max;
}

fn extract_input(lines: &Vec<String>) -> Vec<i32> {
    return lines.get(0).unwrap().split(",").map(|number| number.parse().unwrap()).collect();
}

fn lines_from_file(filename: impl AsRef<Path>) -> io::Result<Vec<String>> {
    return BufReader::new(File::open(filename)?).lines().collect();
}
