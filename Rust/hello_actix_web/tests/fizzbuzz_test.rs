use hello_actix_web::fizzbuzz::fizzbuzz;

#[test]
fn test_fizzbuzz() {
    let cases = [
        (3, "Fizz"),
        (5, "Buzz"),
        (15, "FizzBuzz"),
        (7, "7"),
    ];
    for (input, expected) in cases {
        assert_eq!(fizzbuzz(input), expected);
    }
}
