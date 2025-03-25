use actix_web::{test, App};
use hello_actix_web::{char_count_endpoint};
use hello_actix_web::models::TextRequest;

#[actix_web::test]
async fn test_char_endpoint_count() {
    let app = test::init_service(
        App::new().service(char_count_endpoint)
    ).await;

    let req_body = TextRequest {
        text: "こんにちはRust".to_string(),
    };

    let req = test::TestRequest::post()
        .uri("/char_count")
        .set_json(&req_body)
        .to_request();

    let resp_body = test::call_and_read_body(&app, req).await;
    println!("raw body = {:?}", String::from_utf8_lossy(&resp_body));

    let resp_json: serde_json::Value =
        serde_json::from_slice(&resp_body).expect("Failed to parse response JSON");

    assert_eq!(resp_json["length"], 9);
}