use actix_web::{test, http::StatusCode, App};
use hello_actix_web::{health_check, hello, fizzbuzz_endpoint, bmi_endpoint};
use serde_json::json;
use serde_json::Value;

#[actix_web::test]
async fn test_fizzbuzz_endpoint() {
    let app = test::init_service(
        App::new()
            .service(health_check)
            .service(hello)
            .service(fizzbuzz_endpoint)
    ).await;

    // 正常パターン (例: /fizzbuzz/3 -> Fizz)
    let req = test::TestRequest::get().uri("/fizzbuzz/3").to_request();
    let resp = test::call_service(&app, req).await;
    assert_eq!(resp.status(), StatusCode::OK);
    let body = test::read_body(resp).await;
    assert_eq!(body, "Fizz");

    // 異常パターン (数値変換エラー)
    let req_err = test::TestRequest::get().uri("/fizzbuzz/abc").to_request();
    let resp_err = test::call_service(&app, req_err).await;
    assert_eq!(resp_err.status(), StatusCode::BAD_REQUEST);
    let err_body = test::read_body(resp_err).await;
    // エラーメッセージや形式は実装に合わせて確認
    assert!(std::str::from_utf8(&err_body).unwrap().contains("abc"));
}

#[actix_web::test]
async fn test_bmi_endpoint() {
    let app = test::init_service(
        App::new()
            .service(bmi_endpoint)
    ).await;

    // 正常パターン (例: {"height": 169.5, "weight": 60} -> BMI: 20.88) 小数点3桁目を四捨五入
    let req = test::TestRequest::post()
        .uri("/bmi")
        .set_json(&json!({"height": 169.5, "weight": 60})
    ).to_request();
    let resp = test::call_service(&app, req).await;
    assert_eq!(resp.status(), StatusCode::OK);
    let body = test::read_body(resp).await;
    let body_json: Value = serde_json::from_slice(&body).unwrap();
    assert_eq!(body_json["bmi"].as_f64().unwrap(), 20.88);
    assert_eq!(body_json["category"].as_str().unwrap(), "Normal weight");
}