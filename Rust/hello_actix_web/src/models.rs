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