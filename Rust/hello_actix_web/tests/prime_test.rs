use hello_actix_web::prime::get_prime;
use hello_actix_web::prime::nth_primes;

#[test]
fn test_prime() {
    let cases = [
        (1, vec![2]),
        (2, vec![2, 3]),
        (3, vec![2, 3, 5]),
        (4, vec![2, 3, 5, 7]),
        (5, vec![2, 3, 5, 7, 11]),
    ];
    for (input, expected) in cases {
        assert_eq!(get_prime(input), expected);
    }
}

#[test]
fn test_nth_primes() {
    let cases = [
        (10, vec![2, 3, 5, 7, 11, 13, 17, 19, 23, 29]),
        (
            50,
            vec![
                2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79,
                83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167,
                173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229,
            ],
        ),
    ];
    for (input, expected) in cases {
        assert_eq!(nth_primes(input), expected);
    }
}
