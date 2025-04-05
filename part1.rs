use std::fs;

fn main() {
    let raw = match fs::read_to_string("in.txt") {
        Ok(raw) => raw,
        Err(_) => panic!("uhoh"),
    };

    let raw_vector: Vec<_> = raw.split_whitespace().collect();
    raw_vector.join(" ");
    //println!("{:?}", v);
    let mut vec = Vec::new();
    let iter = raw_vector.iter();
    for i in iter {
        let a: i32 = i.parse().unwrap();
        vec.push(a);
    }
    //println!("{:?}", vec);
    let mut total_indexes: usize = 0;
    let mut distances = Vec::new();
    println!("{:?}", vec);
    while total_indexes <= (vec.len() / 2) - 1 {
        let mut smallest_even: i32 = vec[0];
        let mut smallest_odd: i32 = vec[1];
        let mut prev_even: usize = 0;
        let mut prev_odd: usize = 1;
        for i in 0..vec.len() {
            if i % 2 == 0 {
                if smallest_even >= vec[i] {
                    vec[prev_even] = smallest_even;
                    prev_even = i;
                    smallest_even = vec[i];
                    vec[i] = 99999999;
                }
            }
            else if i % 2 == 1 {
                if smallest_odd >= vec[i] {
                    vec[prev_odd] = smallest_odd;
                    prev_odd = i;
                    smallest_odd = vec[i];
                    vec[i] = 99999999;
                }
            }
        }
        println!("{:?}", vec);
        total_indexes += 1;
        distances.push((smallest_even - smallest_odd).abs());
    }
    println!("{:?}", distances);
    println!("{}", distances.iter().sum::<i32>());
}
