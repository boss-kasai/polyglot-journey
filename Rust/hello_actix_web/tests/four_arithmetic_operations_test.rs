use actix_web::{test, http::StatusCode, App};
use hello_actix_web::{health_check, hello, four_arithmetic_operations_endpoint};

#[actix_web::test]
async fn test_four_arithmetic_operations_endpoint() {
    let app = test::init_service(
        App::new()
            .service(health_check)
            .service(hello)
            .service(four_arithmetic_operations_endpoint)
    ).await;

    // 正常パターン (例: /four_arithmetic_operations/3+5 -> 8)
    let req = test::TestRequest::get().uri("/four_arithmetic_operations/3+5").to_request();
    let resp = test::call_service(&app, req).await;
    assert_eq!(resp.status(), StatusCode::OK);
    let body = test::read_body(resp).await;
    assert_eq!(body, "8");

    // 正常パターン (例: /four_arithmetic_operations/3-5 -> -2)
    let req = test::TestRequest::get().uri("/four_arithmetic_operations/3-5").to_request();
    let resp = test::call_service(&app, req).await;
    assert_eq!(resp.status(), StatusCode::OK);
    let body = test::read_body(resp).await;
    assert_eq!(body, "-2");

    // 正常パターン (例: /four_arithmetic_operations/3*5-5+10/2 -> 10)
    let req = test::TestRequest::get().uri("/four_arithmetic_operations/10%2F2").to_request();
    let resp = test::call_service(&app, req).await;
    assert_eq!(resp.status(), StatusCode::OK);
    let body = test::read_body(resp).await;
    assert_eq!(body, "5");

    // 正常パターン (例: /four_arithmetic_operations/3%2A5 -> 15)
    let req = test::TestRequest::get().uri("/four_arithmetic_operations/3%2A5").to_request();
    let resp = test::call_service(&app, req).await;
    assert_eq!(resp.status(), StatusCode::OK);
    let body = test::read_body(resp).await;
    assert_eq!(body, "15");

    //正常パターン (例: /four_arithmetic_operations/3%2A5%2A2 -> 30)
    let req = test::TestRequest::get().uri("/four_arithmetic_operations/3%2A5%2A2").to_request();
    let resp = test::call_service(&app, req).await;
    assert_eq!(resp.status(), StatusCode::OK);
    let body = test::read_body(resp).await;
    assert_eq!(body, "30");

    // 異常パターン (例: /four_arithmetic_operations/3%10-5+5%2A -> 400)
    let req = test::TestRequest::get().uri("/four_arithmetic_operations/3%10-5+5%2A").to_request();
    let resp = test::call_service(&app, req).await;
    assert_eq!(resp.status(), StatusCode::BAD_REQUEST);
    let body = test::read_body(resp).await;
    assert_eq!(body, "Invalid expression format");
}
