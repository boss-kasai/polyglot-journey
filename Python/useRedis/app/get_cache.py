from app.models.postal_code import PostalCode
from app.models.prefecture import Prefecture
from app.redis_client import redis_client
from sqlalchemy.orm import Session


def get_prefecture_id(name: str, db: Session) -> int:
    key = f"prefecture:{name}"
    cached = redis_client.get(key)
    if cached:
        return int(cached)

    prefecture = db.query(Prefecture).filter(Prefecture.name == name).first()
    if not prefecture:
        raise ValueError(f"都道府県が見つかりません -> {name}")

    redis_client.set(key, prefecture.id, ex=86400)  # 有効期限を1日（86400秒）に設定
    return prefecture.id


def get_postal_code_id(code: str, db: Session) -> int:
    key = f"postal_code:{code}"
    cached = redis_client.get(key)
    if cached:
        return int(cached)

    postal_code = db.query(PostalCode).filter(PostalCode.code == code).first()
    if not postal_code:
        raise ValueError(f"郵便番号が見つかりません -> {code}")

    redis_client.set(key, postal_code.id, ex=86400)  # 有効期限を1日（86400秒）に設定
    return postal_code.id
