use bcrypt::{hash, verify};
use hello_actix_web::hash_check::get_hash;

#[test]
fn test_hash_check() {
    let password = "secret_password";
    let hashed = hash(password, 10).unwrap();
    assert!(
        verify(password, &hashed).unwrap(),
        "パスワードが一致していません。"
    );
    let wrong_password = "wrong_password";
    assert!(
        !verify(wrong_password, &hashed).unwrap(),
        "誤ったパスワードで検証が通ってしまいました。"
    );
}
