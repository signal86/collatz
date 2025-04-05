fn collatz(mut n: u64) -> bool {
    while n != 1 {
        if n % 2 == 0 {
            n /= 2;
        } else {
            n = (n * 3) + 1;
        }
    }
    true
}

fn main() {
    for n in 1000000..5000001 {
        collatz(n);
    }
    println!("all numbers checked");
}
