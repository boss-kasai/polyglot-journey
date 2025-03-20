use hello_actix_web::loan::calculate_loan_principal;

#[test]
fn test_calculate_loan_principal() {
    let cases = [
        (50000.0, 10, 0.03, 5188120),
        (70000.0, 20, 0.02, 13860660),
        (50000.0, 10, 0.00001, 5999698),
    ];
    for (A, years, annual_rate, expected) in cases {
        assert_eq!(calculate_loan_principal(A, years, annual_rate), expected);
    }
}
