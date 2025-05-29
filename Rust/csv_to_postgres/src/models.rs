use serde::Deserialize;

#[derive(Debug, Deserialize, Clone)] // ← これが必要！
pub struct Company {
    #[serde(rename = "企業名")]
    pub name: String,
    #[serde(rename = "郵便番号")]
    pub postal_code: String,
    #[serde(rename = "都道府県")]
    pub prefecture: String,
    #[serde(rename = "住所")]
    pub address: String,
    #[serde(rename = "担当者氏名")]
    pub contact_name: String,
    #[serde(rename = "電話番号")]
    pub phone: String,
}
