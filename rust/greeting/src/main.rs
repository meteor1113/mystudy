use std::{fs::File, thread, time::Duration};

fn main() {
    let a = 12;
    println!("a is {}", add(a, 4));

    println!("Hello, world!");
    let _s = String::from("ccc");

    let x = 1;
    let sum = |y| x + y;
    assert_eq!(3, sum(2));

    let _f = File::open("hello.txt");

    let intensity = 10;
    let random_number = 7;
    workout(intensity, random_number);
}

fn add(a: i32, b: i32) -> i32 {
    a + b
}

fn workout(intensity: u32, random_number: u32) {
    let action = || {
        println!("muuuu.....");
        thread::sleep(Duration::from_secs(2));
        intensity
    };

    if intensity < 25 {
        println!(
            "今天活力满满，先做 {} 个俯卧撑!",
            action()
        );
        println!(
            "旁边有妹子在看，俯卧撑太low，再来 {} 组卧推!",
            action()
        );
    } else if random_number == 3 {
        println!("昨天练过度了，今天还是休息下吧！");
    } else {
        println!(
            "昨天练过度了，今天干干有氧，跑步 {} 分钟!",
            action()
        );
    }
}
