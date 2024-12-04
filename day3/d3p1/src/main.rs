use std::fs::File;
use std::io::{self, BufRead};
use regex::Regex;

fn file_reader(filename: &str) -> io::Result<String> {
    // Open the file
    let file: File = File::open(filename)?;

    // Initialize Buffered Reader
    let reader = io::BufReader::new(file);

    // Output vector
    let output = reader
        .lines()
        .map(|line| line.unwrap())
        .collect::<String>();

    Ok(output)
}


fn regex(text: String) -> i32{
    // mul\(    Matches "mul("
    // (-?\d+)  Matches and captures the first +/- integer, followed by 1 or more digits
    // ,        Matches the comma between the two integers
    // (-?\d+)  Matches and captures the second +/- integer, followed by 1 or more digits
    // \)       Matches the closing parenthesis

    let pattern: &str = r"mul\((-?\d+),(-?\d+)\)";

    let re = Regex::new(pattern).unwrap();

    let mut sum: i32 = 0;
    for matched in re.captures_iter(&text) {

        let num1: i32 = matched[1].parse().unwrap();
        let num2: i32 = matched[2].parse().unwrap();

        sum += num1 * num2;
    }

    sum
}

fn main() {
    let filename: &str = "day3/day-3-input.txt";

    //Match for verifying that the file could be read
    match file_reader(filename) {
        Ok(output) => {
            let answer = regex(output);
            println!("{}", answer);
        }
        Err(e) => eprintln!("Error reading file: {}", e),
    }
}
