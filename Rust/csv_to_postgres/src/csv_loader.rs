use crate::models::Company;
use csv::ReaderBuilder;
use std::error::Error;
use std::fs::File;
use std::path::Path;

pub fn load_csv<P: AsRef<Path>>(path: P) -> Result<Vec<Company>, Box<dyn Error>> {
    let file = File::open(path)?;
    let mut rdr = ReaderBuilder::new().has_headers(true).from_reader(file);
    let mut companies = Vec::new();

    for result in rdr.deserialize() {
        let company: Company = result?;
        companies.push(company);
    }

    Ok(companies)
}
