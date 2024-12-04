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
    let mut is_enabled: bool = true;

    // mul\(        Matches "mul("
    // (-?\d+)      Matches and captures the first +/- integer, followed by 1 or more digits
    // ,            Matches the comma between the two integers
    // (-?\d+)      Matches and captures the second +/- integer, followed by 1 or more digits
    // \)           Matches the closing parenthesis
    // |(do\(\))    Or Matches "do()"
    // |(don't\(\)) Or Matches "don't()"

    let pattern: &str = r"mul\((-?\d+),(-?\d+)\)|(do\(\))|(don't\(\))";

    let re = Regex::new(pattern).unwrap();

    let mut sum: i32 = 0;
    for matched in re.captures_iter(&text) {
    
        // Checks if the current match correlates to the "do()" pattern
        if matched.get(3).is_some() {
            is_enabled = true;
        // Checks if the current match correlates to the "don't()" pattern
    } else if matched.get(4).is_some() {
            is_enabled = false;
        }

        // Checks if the current match correlates to the mul(num1,num2) pattern
        if let (Some(_num1), Some(_num2)) = (matched.get(1), matched.get(2)) {
            // Converting the captured strings into integers
            let num1: i32 = matched[1].parse().unwrap();
            let num2: i32 = matched[2].parse().unwrap();
        
            if is_enabled == true {
                sum += num1 * num2;
            }
        }
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
