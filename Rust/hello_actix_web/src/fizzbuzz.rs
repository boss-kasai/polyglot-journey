use std::collections::HashMap;

pub fn fizzbuzz(num: u32) -> String {
    if num % 15 == 0 {
        "FizzBuzz".to_string()
    } else if num % 3 == 0 {
        "Fizz".to_string()
    } else if num % 5 == 0 {
        "Buzz".to_string()
    } else {
        num.to_string()
    }
}

fn fib_memo(num: u32, memo: &mut HashMap<u32, u128>) -> u128 {
    if let Some(&result) = memo.get(&num) {
        return result;
    }
    let result = if num <= 1 {
        num as u128
    } else {
        fib_memo(num - 1, memo) + fib_memo(num - 2, memo)
    };
    memo.insert(num, result);
    result
}

pub fn fibonacci(num: u32) -> u128 {
    let mut memo = HashMap::new();
    fib_memo(num, &mut memo)
}