use actix_web::{test, http::StatusCode, App};
use hello_actix_web::{health_check, hello, fizzbuzz_endpoint, bmi_endpoint, fibonacci_endpoint, prime_endpoint};
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

#[actix_web::test]
async fn test_fibonacci_endpoint() {
    let app = test::init_service(
        App::new()
            .service(fibonacci_endpoint)
    ).await;

    // 正常パターン (例: /fibonacci/10 -> 55)
    let req = test::TestRequest::get().uri("/fibo/10").to_request();
    let resp = test::call_service(&app, req).await;
    assert_eq!(resp.status(), StatusCode::OK);
    let body = test::read_body(resp).await;
    assert_eq!(body, "55");

    // 異常パターン (数値変換エラー)
    let req_err = test::TestRequest::get().uri("/fibo/abc").to_request();
    let resp_err = test::call_service(&app, req_err).await;
    assert_eq!(resp_err.status(), StatusCode::BAD_REQUEST);
    let err_body = test::read_body(resp_err).await;
    assert!(std::str::from_utf8(&err_body).unwrap().contains("Invalid number: abc"));

    // 異常パターン (数値が大きすぎる)
    let req_err = test::TestRequest::get().uri("/fibo/187").to_request();
    let resp_err = test::call_service(&app, req_err).await;
    assert_eq!(resp_err.status(), StatusCode::BAD_REQUEST);
    let err_body = test::read_body(resp_err).await;
    assert!(std::str::from_utf8(&err_body).unwrap().contains("上限値オーバー: 186が最大です"));
}

#[actix_web::test]
async fn test_prime_endpoint(){
    let app = test::init_service(
        App::new()
            .service(prime_endpoint)
    ).await;

    // 正常パターン (例: /prime/5 -> [2, 3, 5, 7, 11])
    let req = test::TestRequest::get().uri("/prime/5").to_request();
    let resp = test::call_service(&app, req).await;
    assert_eq!(resp.status(), StatusCode::OK);
    let body = test::read_body(resp).await;
    let body_json: Value = serde_json::from_slice(&body).unwrap();
    assert_eq!(body_json, json!([2, 3, 5, 7, 11]));

    // 異常パターン (数値変換エラー)
    let req_err = test::TestRequest::get().uri("/prime/abc").to_request();
    let resp_err = test::call_service(&app, req_err).await;
    assert_eq!(resp_err.status(), StatusCode::BAD_REQUEST);
    let err_body = test::read_body(resp_err).await;
    assert!(std::str::from_utf8(&err_body).unwrap().contains("Invalid number: abc"));

    // 異常パターン (数値が大きすぎる)
    let req_err = test::TestRequest::get().uri("/prime/203280222").to_request();
    let resp_err = test::call_service(&app, req_err).await;
    assert_eq!(resp_err.status(), StatusCode::BAD_REQUEST);
    let err_body = test::read_body(resp_err).await;
    assert!(std::str::from_utf8(&err_body).unwrap().contains("上限値オーバー: 203,280,221が最大です"));
}