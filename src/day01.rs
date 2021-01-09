// use std::env::var;
use std::collections::HashSet;

fn values(input: String) -> Vec<i32> {
    input.lines().map(|line| line.parse().unwrap()).collect()
}

pub fn part1(input: String) {
    let values = values(input);
    println!("{}", values.iter().sum::<i32>());
}

pub fn part2(input: String) {
    let values = values(input);
    let mut set = HashSet::new();
    let mut sum = 0;
    set.insert(0);
    loop {
        for value in values.iter() {
            sum += value;
            if set.contains(&sum) {
                println!("twice: {}", sum);
                return;
            }
            set.insert(sum);
        }
    }
}
