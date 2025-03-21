use bcrypt::{hash, BcryptError};

/// plain_password を bcrypt でハッシュ化し，ハッシュ文字列を返す関数
/// コスト（ラウンド数）は 10 とする。
pub fn get_hash(plain_password: &str) -> Result<String, BcryptError> {
    // コスト（ラウンド数）
    let cost = 10;

    // bcrypt ハッシュを生成（内部でランダムなソルトが付与されます）
    let hashed = hash(plain_password, cost)?;

    Ok(hashed)
}
