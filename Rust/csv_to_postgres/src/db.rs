use std::sync::Arc;
use tokio::sync::Semaphore;
use tokio_postgres::Client;
use crate::models::Company;
use tokio_postgres::NoTls;


pub async fn init_db(url: &str) -> Result<(Client, impl std::future::Future<Output = Result<(), tokio_postgres::Error>>), Box<dyn std::error::Error>> {
    let (client, connection) = tokio_postgres::connect(url, NoTls).await?;
    Ok((client, connection))
}


const BATCH_SIZE: usize = 1000;
const MAX_PARALLEL: usize = 10;

pub async fn insert_companies_parallel(
    client: Client,
    companies: Vec<Company>,
) -> Result<(), Box<dyn std::error::Error>> {
    let client = Arc::new(client);
    let semaphore = Arc::new(Semaphore::new(MAX_PARALLEL));

    let mut handles = Vec::new();

    for chunk in companies.chunks(BATCH_SIZE) {
        let permit = semaphore.clone().acquire_owned().await?;
        let client = client.clone();
        let chunk = chunk.to_vec(); // chunkは借用なのでVecにして渡す

        let handle = tokio::spawn(async move {
            let _permit = permit; // Drop時に自動でpermit返却
            let mut query = String::from("INSERT INTO companies (name, postal_code, prefecture, address, contact_name, phone) VALUES ");
            let mut params: Vec<&(dyn tokio_postgres::types::ToSql + Sync)> = Vec::new();

            for (i, company) in chunk.iter().enumerate() {
                if i > 0 {
                    query.push_str(", ");
                }

                let base = i * 6;
                query.push_str(&format!(
                    "(${}, ${}, ${}, ${}, ${}, ${})",
                    base + 1, base + 2, base + 3, base + 4, base + 5, base + 6
                ));

                params.push(&company.name);
                params.push(&company.postal_code);
                params.push(&company.prefecture);
                params.push(&company.address);
                params.push(&company.contact_name);
                params.push(&company.phone);
            }

            client.execute(&query, &params).await?;
            Ok::<(), tokio_postgres::Error>(())
        });

        handles.push(handle);
    }

    for handle in handles {
        handle.await??; // Joinと?でエラー伝播
    }

    Ok(())
}
