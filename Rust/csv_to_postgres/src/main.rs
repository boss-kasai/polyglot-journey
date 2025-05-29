use std::time::Instant;
mod config;
mod csv_loader;
mod models;
mod db;

use config::get_db_url;
use crate::db::{init_db, insert_companies_parallel};
use csv_loader::load_csv;
use models::Company;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let db_url = get_db_url();
    let (client, connection) = init_db(&db_url).await?;

    tokio::spawn(async move {
        if let Err(e) = connection.await {
            eprintln!("DB接続エラー: {}", e);
        }
    });

    println!("⏳ CSV読み込み中...");
    let companies = load_csv("companies_100k.csv")?;

    println!("⏳ DB挿入開始...");
    let start_insert = Instant::now();
    insert_companies_parallel(client, companies).await?;
    println!("✅ DB挿入完了: {:.2?}", start_insert.elapsed());

    Ok(())
}
