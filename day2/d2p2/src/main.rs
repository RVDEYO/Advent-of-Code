use std::fs::File;
use std::io::{self, BufRead};

fn file_reader(filename: &str) -> io::Result<Vec<Vec<i32>>> {
    // Open the file
    let file: File = File::open(filename)?;

    // Initialize Buffered Reader
    let reader = io::BufReader::new(file);

    // Output vector
    let mut output: Vec<Vec<i32>> = Vec::new();

    // Iterate over each line
    for line in reader.lines() {
        let line = line?;
        
        let row: Vec<i32> = line
            .split_whitespace()
            .filter_map(|x| x.parse::<i32>().ok())
            .collect();
        output.push(row);
    }
    Ok(output)
}

fn double_safe_checker(report: &Vec<i32>) -> bool {
    let mut is_safe: bool = false;

    for i in 0..report.len() {

        let mut  temp = report.clone();
        
        temp.remove(i);
        if safe_checker(&temp) {
            is_safe = true
        }
    }
    is_safe
}

fn safe_checker(report: &Vec<i32>) -> bool {
    let mut is_safe: bool = true;
    
    // None == Neutral 1 == Increasing, -1 == Decreasing
    let mut direction: Option<i32> = None;

    for i in 0..report.len()-1 {
        if report[i] - report [i+1] >= 1 && report[i] - report[i+1] <= 3 {
            if direction == None {
                direction = Some(1);
            } else if direction == Some(-1) {
                is_safe = false;
                break;
            }
        } else if report[i] - report [i+1] <= -1 && report[i] - report[i+1] >= -3 {
            if direction == None {
                direction = Some(-1);
            } else if direction == Some(1) {
                is_safe = false;
                break;
            }

        } else {
            is_safe = false;
            break;
        }
    }
    is_safe
}

fn main() {
    let filename: &str = "day2/day-2-input.txt";
    // let filename: &str = "day2/test.txt";

    //Match for verifying that the file could be read
    match file_reader(filename) {
        Ok(output) => {
            let mut answer: i32 = 0;

            for i in 0..output.len() {
                if safe_checker(&output[i]) {
                    answer += 1;
                } else if double_safe_checker(&output[i]) {
                    answer += 1;
                }
            }
            println!("Answer: {}", answer);
        }
        Err(e) => eprintln!("Error reading file: {}", e),
    }    
}
