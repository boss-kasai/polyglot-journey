import csv
import random
from io import TextIOWrapper

from app.db.deps import get_db
from app.get_cache import get_postal_code_id, get_prefecture_id
from app.models.company import Company
from app.models.postal_code import PostalCode
from app.models.prefecture import Prefecture
from app.schemas.postal_code import PostalCodeCreate
from app.schemas.prefecture import PrefectureCreate
from fastapi import Depends, FastAPI, File, HTTPException, UploadFile
from sqlalchemy.orm import Session
from sqlalchemy.sql import func

app = FastAPI()


@app.get("/health")
def check_health():
    return {"message": "I am alive!"}


# 都道府県の追加
@app.post("/prefecture")
def create_prefecture(prefecture_in: PrefectureCreate, db: Session = Depends(get_db)):
    # 重複チェック
    existing = (
        db.query(Prefecture).filter(Prefecture.name == prefecture_in.name).first()
    )
    if existing:
        raise HTTPException(status_code=400, detail="都道府県はすでに存在します。")

    # 登録処理
    prefecture = Prefecture(name=prefecture_in.name)
    db.add(prefecture)
    db.commit()
    db.refresh(prefecture)
    return prefecture


# 都道府県の一覧取得
@app.get("/prefectures")
def get_prefectures(db: Session = Depends(get_db)):
    prefectures = db.query(Prefecture).all()
    return prefectures


# 郵便番号の追加
@app.post("/postal_code")
def create_postal_code(postal_code_in: PostalCodeCreate, db: Session = Depends(get_db)):
    # 重複チェック
    existing = (
        db.query(PostalCode).filter(PostalCode.code == postal_code_in.code).first()
    )
    if existing:
        raise HTTPException(status_code=400, detail="郵便番号はすでに存在します。")

    # 登録処理
    postal_code = PostalCode(code=postal_code_in.code)
    db.add(postal_code)
    db.commit()
    db.refresh(postal_code)
    return postal_code


# 郵便番号の検索取得
@app.get("/postal_codes/{code}")
def get_postal_code(code: str, db: Session = Depends(get_db)):
    results = db.query(PostalCode).filter(PostalCode.code.like(f"%{code}%")).all()
    if not results:
        raise HTTPException(
            status_code=404, detail="該当する郵便番号が見つかりません。"
        )
    if len(results) > 20:
        raise HTTPException(
            status_code=400,
            detail="該当する郵便番号が多すぎます。20件以下で検索してください。",
        )
    return results


@app.get("/validate_postal_code/{code}")
def validate_postal_code(code: str, db: Session = Depends(get_db)):
    exists = db.query(PostalCode).filter(PostalCode.code == code).first()

    if exists:
        return {"exists": True, "code": code}

    # 存在しない場合 → ランダムな郵便番号を取得
    max_id = db.query(func.max(PostalCode.id)).scalar()
    suggested = None
    for _ in range(10):  # 最大10回まで試行
        rand_id = random.randint(1, max_id)
        suggestion = db.query(PostalCode).filter(PostalCode.id == rand_id).first()
        if suggestion:
            suggested = suggestion.code
            break

    return {"exists": False, "suggested": suggested}


# シンプルな企業の取り込み
@app.post("/simple_companies/upload")
def upload_companies(file: UploadFile = File(...), db: Session = Depends(get_db)):
    print("simple_companies/upload called")
    reader = csv.DictReader(TextIOWrapper(file.file, encoding="utf-8"))
    inserted = 0
    failed_rows = []

    # ヘッダーをプリント
    headers = reader.fieldnames
    print("📎 CSV Headers:", headers)  # コンソールに出力

    for idx, row in enumerate(reader, start=2):  # 2行目からデータ開始
        try:
            name = row["企業名"]
            postal_code_str = (row["郵便番号"].strip()).replace("-", "")
            prefecture_name = row["都道府県"].strip()
            address = row["住所"]
            contact_name = row["担当者氏名"]
            phone = row["電話番号"]

            # 外部キーID解決
            postal_code = (
                db.query(PostalCode).filter(PostalCode.code == postal_code_str).first()
            )
            if not postal_code:
                raise HTTPException(
                    status_code=400,
                    detail=f"{idx}行目: 郵便番号が見つかりません -> {postal_code_str}",
                )

            prefecture = (
                db.query(Prefecture).filter(Prefecture.name == prefecture_name).first()
            )
            if not prefecture:
                raise HTTPException(
                    status_code=400,
                    detail=f"{idx}行目: 都道府県が見つかりません -> {prefecture_name}",
                )

            company = Company(
                name=name,
                postal_code_id=postal_code.id,
                prefecture_id=prefecture.id,
                address=address,
                contact_name=contact_name,
                phone=phone,
            )
            db.add(company)
            db.commit()
            db.refresh(company)
            inserted += 1

        except Exception as e:
            db.rollback()
            failed_rows.append({"row": idx, "reason": str(e)})

    return {"inserted": inserted, "failed": len(failed_rows), "errors": failed_rows}


# バルクインサートの企業の取り込み
BATCH_SIZE = 1000  # バルクのサイズ


@app.post("/bulk_simple_companies/upload")
def bulk_upload_companies(file: UploadFile = File(...), db: Session = Depends(get_db)):
    print("simple_companies/upload called")
    reader = csv.DictReader(TextIOWrapper(file.file, encoding="utf-8"))
    inserted = 0
    failed_rows = []
    companies_batch = []

    headers = reader.fieldnames
    print("📎 CSV Headers:", headers)

    for idx, row in enumerate(reader, start=2):
        try:
            name = row["企業名"]
            postal_code_str = (row["郵便番号"].strip()).replace("-", "")
            prefecture_name = row["都道府県"].strip()
            address = row["住所"]
            contact_name = row["担当者氏名"]
            phone = row["電話番号"]

            # 外部キー解決
            postal_code = (
                db.query(PostalCode).filter(PostalCode.code == postal_code_str).first()
            )
            if not postal_code:
                raise ValueError(
                    f"{idx}行目: 郵便番号が見つかりません -> {postal_code_str}"
                )

            prefecture = (
                db.query(Prefecture).filter(Prefecture.name == prefecture_name).first()
            )
            if not prefecture:
                raise ValueError(
                    f"{idx}行目: 都道府県が見つかりません -> {prefecture_name}"
                )

            company = Company(
                name=name,
                postal_code_id=postal_code.id,
                prefecture_id=prefecture.id,
                address=address,
                contact_name=contact_name,
                phone=phone,
            )
            companies_batch.append(company)

            # 1000件ごとにまとめてバルク挿入
            if len(companies_batch) >= BATCH_SIZE:
                db.bulk_save_objects(companies_batch)
                db.commit()
                inserted += len(companies_batch)
                companies_batch.clear()

        except Exception as e:
            db.rollback()
            failed_rows.append({"row": idx, "reason": str(e)})

    # 残りが1000件未満だった場合
    if companies_batch:
        try:
            db.bulk_save_objects(companies_batch)
            db.commit()
            inserted += len(companies_batch)
        except Exception as e:
            db.rollback()
            failed_rows.append({"row": "final_batch", "reason": str(e)})

    return {"inserted": inserted, "failed": len(failed_rows), "errors": failed_rows}


# redisを使用したsingle companyのインサート
@app.post("/redis_companies/upload")
def upload_redis_companies(file: UploadFile = File(...), db: Session = Depends(get_db)):
    reader = csv.DictReader(TextIOWrapper(file.file, encoding="utf-8"))
    inserted = 0
    failed_rows = []

    for idx, row in enumerate(reader, start=2):
        try:
            name = row["企業名"]
            postal_code_str = row["郵便番号"].strip()
            prefecture_name = row["都道府県"].strip()
            address = row["住所"]
            contact_name = row["担当者氏名"]
            phone = row["電話番号"]

            # Redisキャッシュ経由の外部キーID解決
            postal_code_id = get_postal_code_id(postal_code_str, db)
            prefecture_id = get_prefecture_id(prefecture_name, db)

            company = Company(
                name=name,
                postal_code_id=postal_code_id,
                prefecture_id=prefecture_id,
                address=address,
                contact_name=contact_name,
                phone=phone,
            )
            db.add(company)
            db.commit()
            db.refresh(company)
            inserted += 1

        except Exception as e:
            db.rollback()
            failed_rows.append({"row": idx, "reason": str(e)})

    return {"inserted": inserted, "failed": len(failed_rows), "errors": failed_rows}


@app.post("/bulk_redis_companies/upload")
def bulk_upload_redis_companies(
    file: UploadFile = File(...), db: Session = Depends(get_db)
):
    reader = csv.DictReader(TextIOWrapper(file.file, encoding="utf-8"))
    inserted = 0
    failed_rows = []
    companies_batch = []

    for idx, row in enumerate(reader, start=2):
        try:
            name = row["企業名"]
            postal_code_str = row["郵便番号"].strip().replace("-", "")
            prefecture_name = row["都道府県"].strip()
            address = row["住所"]
            contact_name = row["担当者氏名"]
            phone = row["電話番号"]

            # Redisキャッシュから外部キーID取得
            postal_code_id = get_postal_code_id(postal_code_str, db)
            prefecture_id = get_prefecture_id(prefecture_name, db)

            company = Company(
                name=name,
                postal_code_id=postal_code_id,
                prefecture_id=prefecture_id,
                address=address,
                contact_name=contact_name,
                phone=phone,
            )
            companies_batch.append(company)

            # バルクインサート実行
            if len(companies_batch) >= BATCH_SIZE:
                db.bulk_save_objects(companies_batch)
                db.commit()
                inserted += len(companies_batch)
                companies_batch.clear()

        except Exception as e:
            db.rollback()
            failed_rows.append({"row": idx, "reason": str(e)})

    # 残りがあれば最後にまとめて挿入
    if companies_batch:
        try:
            db.bulk_save_objects(companies_batch)
            db.commit()
            inserted += len(companies_batch)
        except Exception as e:
            db.rollback()
            failed_rows.append({"row": "final_batch", "reason": str(e)})

    return {"inserted": inserted, "failed": len(failed_rows), "errors": failed_rows}
