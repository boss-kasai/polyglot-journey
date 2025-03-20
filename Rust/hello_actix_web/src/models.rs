// filepath: /Users/cocolo/development/polyglot-journey/Rust/hello_actix_web/src/models.rs
use serde::Deserialize;
use serde::Serialize;

#[derive(Deserialize)]
pub struct BmiRequest {
    pub height: f32,
    pub weight: f32,
}

#[derive(Serialize)]
pub struct BmiResponse {
    pub bmi: f32,
    pub category: String,
}

#[derive(Deserialize)]
pub struct LoanRequest {
    pub monthly_payment: f64,
    pub years: u32,
    pub annual_rate: f64,
}
