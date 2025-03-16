use actix_web::{get, post, web, HttpResponse, Responder};
pub mod fizzbuzz;
pub mod models;
pub mod prime;
pub mod four_arithmetic_operations; //ファイルを指定
use crate::models::BmiResponse;


#[get("/health")]
pub async fn health_check() -> impl Responder {
    HttpResponse::Ok().body("Ok")
}

#[get("/hello")]
pub async fn hello() -> impl Responder {
    HttpResponse::Ok().body("Hello, Actix-web!")
}

#[get("/fizzbuzz/{num}")]
pub async fn fizzbuzz_endpoint(path: web::Path<String>) -> impl Responder {
    // まずは文字列として受け取る
    let raw = path.into_inner();
    // ここで手動で parse する
    match raw.parse::<u32>() {
        Ok(num) => {
            // 正常に u32 化できたので fizzbuzz などの処理
            let result = fizzbuzz::fizzbuzz(num);
            HttpResponse::Ok().body(result)
        }
        Err(e) => {
            // ここで 400 を返すなど自由にエラーハンドリング可能
            HttpResponse::BadRequest().body(format!("'{}' は数値変換に失敗: {}", raw, e))
        }
    }
}

#[post("/bmi")]
pub async fn bmi_endpoint(req_body: web::Json<models::BmiRequest>) -> impl Responder {
    let height = req_body.height/100.0;
    let weight = req_body.weight;
    let bmi_row = weight / height.powi(2);
    let bmi = (bmi_row * 100.0).round() / 100.0;
    let category = match bmi {
        n if n < 18.5 => "Underweight",
        n if n < 24.9 => "Normal weight",
        n if n < 29.9 => "Overweight",
        _ => "Obesity",
    };
    // 構造体を生成
    let response = BmiResponse {
        bmi,
        category: category.to_string(),
    };

    // JSONとして返す
    HttpResponse::Ok().json(response)
}

#[get("/fibo/{num}")]
pub async fn fibonacci_endpoint(path: web::Path<String>) -> impl Responder {
    let raw = path.into_inner();
    match raw.parse::<u32>() {
        Ok(num) => {
            if num > 186 {
                return HttpResponse::BadRequest().body("上限値オーバー: 186が最大です");
            }
            let result = fizzbuzz::fibonacci(num);
            HttpResponse::Ok().body(result.to_string())
        }
        Err(_) => {
            HttpResponse::BadRequest().body(format!("Invalid number: {}", raw))
        }
    }
}

#[get("/prime/{num}")]
pub async fn prime_endpoint(path: web::Path<String>) -> impl Responder {
    let raw = path.into_inner();
    match raw.parse::<usize>() {
        Ok(num) => {
            if num > 203280221 {
                return HttpResponse::BadRequest().body("上限値オーバー: 203,280,221が最大です");
            }
            if num > 1000{
                let result = prime::nth_primes(num);
                return HttpResponse::Ok().json(result);
            }
            let result = prime::getPrime(num);
            HttpResponse::Ok().json(result)
        }
        Err(_) => {
            HttpResponse::BadRequest().body(format!("Invalid number: {}", raw))
        }
    }
}

#[get("/four_arithmetic_operations/{str}")]
pub async fn four_arithmetic_operations_endpoint(path: web::Path<String>) -> impl Responder {
    let raw = path.into_inner();
    match four_arithmetic_operations::evaluate_expression(&raw) {
        Ok(result) => HttpResponse::Ok().body(result.to_string()),
        Err(err) => HttpResponse::BadRequest().body(err),
    }
}