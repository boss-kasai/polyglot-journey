/// 月の返済額 `A`、返済年数 `years`、年利 `annual_rate` から借入可能額を求める関数
/// 元の TS では `Math.round` して返しているので、戻り値は整数型 (i64) としています。
pub fn calculate_loan_principal(a_: f64, years: u32, annual_rate: f64) -> i64 {
    // 返済回数 n (年数 * 12)
    let n = years * 12;

    // 月利 i (年利を12で割り、月複利で計算)
    let i = (1.0 + annual_rate).powf(1.0 / 12.0) - 1.0;

    // A = P * ( i / (1 - (1 + i)^(-n)) )
    // → P = A * (1 - (1 + i)^(-n)) / i
    let numerator = 1.0 - (1.0 + i).powf(-(n as f64));
    let principal = a_ * (numerator / i);

    // TS 同様に round
    principal.round() as i64
}
