use std::fs::File;

fn main() {
    let a = 12;
    println!("a is {}", add(a, 4));

    println!("Hello, world!");
    let _s = String::from("ccc");

    let _f = File::open("hello.txt");
}

fn add(a: i32, b: i32) -> i32 {
    a + b
}
