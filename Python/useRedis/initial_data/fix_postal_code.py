import csv
import os
import random
import sys

# プロジェクトルートをモジュール検索パスに追加
sys.path.append(os.path.join(os.path.dirname(__file__), ".."))
from app.models.postal_code import PostalCode
from app.models.prefecture import Prefecture
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker

# 例: from models import PostalCode, Prefecture

DATABASE_URL = "postgresql+psycopg://user:password@localhost:5432/company_db"
# DB接続
engine = create_engine(DATABASE_URL)
SessionLocal = sessionmaker(bind=engine)

# ファイル名
INPUT_CSV = "initial_data/companies_100k.csv"
OUTPUT_CSV = "initial_data/companies_fixed.csv"


def load_postal_code_cache(db):
    codes = db.query(PostalCode.code, PostalCode.id).all()
    return {code: id for code, id in codes}


def load_prefecture_cache(db):
    prefs = db.query(Prefecture.name, Prefecture.id).all()
    return {name: id for name, id in prefs}


def get_random_value(d: dict):
    return random.choice(list(d.keys()))


def process_csv():
    db = SessionLocal()

    postal_code_cache = load_postal_code_cache(db)
    prefecture_cache = load_prefecture_cache(db)

    print(
        f"📦 郵便番号件数: {len(postal_code_cache)}, 都道府県件数: {len(prefecture_cache)}"
    )

    with (
        open(INPUT_CSV, newline="", encoding="utf-8") as infile,
        open(OUTPUT_CSV, "w", newline="", encoding="utf-8") as outfile,
    ):
        reader = csv.DictReader(infile)
        writer = csv.DictWriter(outfile, fieldnames=reader.fieldnames)
        writer.writeheader()

        for idx, row in enumerate(reader, start=2):
            original_postal_code = row["郵便番号"].strip()
            original_prefecture = row["都道府県"].strip()

            if original_postal_code not in postal_code_cache:
                suggested = get_random_value(postal_code_cache)
                print(
                    f"📍 {idx}行目: 郵便番号 '{original_postal_code}' → '{suggested}'"
                )
                row["郵便番号"] = suggested

            if original_prefecture not in prefecture_cache:
                suggested = get_random_value(prefecture_cache)
                print(f"🏙️ {idx}行目: 都道府県 '{original_prefecture}' → '{suggested}'")
                row["都道府県"] = suggested

            writer.writerow(row)

    db.close()
    print(f"✅ 修正済みCSVを書き出しました: {OUTPUT_CSV}")


if __name__ == "__main__":
    process_csv()
