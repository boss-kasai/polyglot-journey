use actix_web::{App, HttpServer};
mod fizzbuzz;
use hello_actix_web::{health_check, hello, fizzbuzz_endpoint, bmi_endpoint, fibonacci_endpoint, prime_endpoint};

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .service(health_check)
            .service(hello)
            .service(fizzbuzz_endpoint)
            .service(bmi_endpoint)
            .service(fibonacci_endpoint)
            .service(prime_endpoint)
    })
    .bind(("127.0.0.1", 8080))?
    .run()
    .await
}