// use std::env::var;
use itertools::Itertools;
use std::collections::HashMap;

fn analyse(line: &str) -> (u32, u32) {
    let mut count2 = 0;
    let mut count3 = 0;
    let mut frequency: HashMap<char, u32> = HashMap::new();
    for letter in line.chars() {
        let entry = frequency.entry(letter).or_insert(0);
        *entry += 1;
        if *entry == 2 {
            count2 += 1;
        } else if *entry == 3 {
            count2 -= 1;
            count3 += 1;
        } else if *entry == 4 {
            count3 -= 1
        }
    }
    let a = if count2 > 0 { 1 } else { 0 };
    let b = if count3 > 0 { 1 } else { 0 };
    (a, b)
}

pub fn part1(input: String) {
    let mut count2 = 0;
    let mut count3 = 0;
    input.lines().for_each(|line| {
        let (a, b) = analyse(line);
        count2 += a;
        count3 += b;
    });
    println!(
        "count2={} count3={} res={}",
        count2,
        count3,
        count2 * count3
    );
}

fn differ_by_1(word1: &str, word2: &str) -> bool {
    let chars1 = word1.chars();
    let chars2 = word2.chars();
    let mut count = 0;

    chars1.zip(chars2).for_each(|(c1, c2)| {
        if c1 != c2 {
            count += 1;
            if count > 1 {
                return;
            }
        }
    });

    if count == 1 {
        return true;
    }
    false
}

pub fn part2(input: String) {
    let (a, b) = input
        .lines()
        .tuple_combinations()
        .find(|(a, b)| differ_by_1(a, b))
        .expect("no pair found");
    println!("{} {}", a, b);

    // iosnxmfkpabcjpdywvrtaqhluy iosnxmfkpabcjpdywvrtawhluy

    // let lines: Vec<&str> = input.lines().collect();
    // for i in 0..lines.len() {
    //     for j in i + 1..lines.len() {
    //         let line1 = lines.get(i).unwrap();
    //         let line2 = lines.get(j).unwrap();
    //         if differ_by_1(line1, line2) {
    //             println!("{} {}", line1, line2);
    //         }
    //     }
    // }
}

pub fn part11(input: String) {
    let mut frequencies = [0u8; 256];
    let (mut twos, mut threes) = (0, 0);
    for line in input.lines() {
        for f in frequencies.iter_mut() {
            *f = 0;
        }
        for b in line.as_bytes().iter().map(|&b| b as usize) {
            frequencies[b] = frequencies[b].saturating_add(1);
        }
        if frequencies.iter().any(|&f| f == 2) {
            twos += 1;
        }
        if frequencies.iter().any(|&f| f == 3) {
            threes += 1;
        }
    }
    println!("{}", twos * threes);
}
