use actix_web::{get, web, HttpResponse, Responder};
pub mod fizzbuzz;


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