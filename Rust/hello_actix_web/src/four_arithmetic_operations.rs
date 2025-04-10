use meval::eval_str;
use urlencoding::decode;

pub fn evaluate_expression(expression: &str) -> Result<f64, String> {
    let decoded_expr = decode(expression)
        .map_err(|_| "Failed to decode URL".to_string())?
        .into_owned(); // `Cow<str>` を `String` に変換
    eval_str(&decoded_expr).map_err(|_e| format!("Invalid expression format"))
}
