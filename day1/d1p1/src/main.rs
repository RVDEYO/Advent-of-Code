use std::fs::File;
use std::io::{self, BufRead};


// Takes in filename, reads the file line by line and outputs two sorted int vectors
fn file_reader(filename: &str) -> io::Result<(Vec<i32>, Vec<i32>)> {
    // Open the file
    let file = File::open(filename)?;

    //Create buffered reader
    let reader = io::BufReader::new(file);

    //Create first and second list
    let mut first_list: Vec<i32> = Vec::new();
    let mut second_list: Vec<i32> = Vec::new();

    //Iterate over each line
    for line in reader.lines() {
        let line = line?;
        
        // Splits each line by off of "  "
        if let Some((one, two)) = line.split_once("  ") {
            // Convert (Parse) the strings into ints
            if let(Ok(one_num), Ok(two_num)) = (one.trim().parse::<i32>(), two.trim().parse::<i32>()) {
                first_list.push(one_num);
                second_list.push(two_num);
            }
       } else {
        eprintln!("Error: Couldn't parse line: {}", line);
       }
    }

    // Sorts both vectors
    first_list.sort();
    second_list.sort();
    Ok((first_list, second_list))
}

fn distance_calc(first_list: &Vec<i32>, second_list: &Vec<i32>) -> Vec<i32> {
    //Initialize new vector
    let mut dist_vec: Vec<i32> = Vec::new();

    // Iterate through both lists
    for i in 0..first_list.len() {
        //Checks to see which number is greater
        if first_list[i] >= second_list[i] {
            dist_vec.push(first_list[i] - second_list[i])
        } else {
            dist_vec.push(second_list[i] - first_list[i])
        }
    }

    //Returns dist_vec
    dist_vec
}

fn total_distance(dist_vec: &Vec<i32>) -> i32 {
    let mut total: i32 = 0;
    for i in 0..dist_vec.len() {
        total += dist_vec[i]
    }
    total
}

fn main() {
    let filename = "day1/day-1-input.txt";

    // Match to make sure the file could be read
    match file_reader(filename) {
        Ok((first_list, second_list)) => {
            let dist_vec = distance_calc(&first_list, &second_list);
            let answer = total_distance(&dist_vec);
            println!("{}", answer);
        }
        Err(e) => eprintln!("Error reading file: {}", e),
    }
}

