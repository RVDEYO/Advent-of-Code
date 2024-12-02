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

fn similarity_calc(first_list: &Vec<i32>, second_list: &Vec<i32>) -> Vec<i32> {
    //Initialize new vector
    let mut sim_vec: Vec<i32> = Vec::new();

    // Iterate through both lists
    for i in 0..first_list.len() {
        let mut counter: i32 = 0;
        for j in 0..second_list.len() {
            if first_list[i] == second_list[j] {
                counter += 1;
            }
        }
        sim_vec.push(first_list[i]*counter);

    }

    //Returns sim_vec
    sim_vec
}

fn total_similarity(sim_vec: &Vec<i32>) -> i32 {
    let mut total: i32 = 0;
    for i in 0..sim_vec.len() {
        total += sim_vec[i]
    }
    total
}

fn main() {
    let filename = "day1/day-1-input.txt";

    // Match to make sure the file could be read
    match file_reader(filename) {
        Ok((first_list, second_list)) => {
            let sim_vec = similarity_calc(&first_list, &second_list);
            let answer = total_similarity(&sim_vec);
            println!("{}", answer);
        }
        Err(e) => eprintln!("Error reading file: {}", e),
    }
}

