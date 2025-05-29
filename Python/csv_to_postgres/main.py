import csv
from time import time

import psycopg

CSV_FILE = "companies_100k.csv"
DB_URL = "postgresql://user:password@localhost:5432/my_rust_db"


def load_csv(file_path):
    with open(file_path, newline="", encoding="utf-8") as f:
        reader = csv.DictReader(f)
        return list(reader)


def insert_into_db(conn, company):
    with conn.cursor() as cur:
        cur.execute(
            """
            INSERT INTO companies (name, postal_code, prefecture, address, contact_name, phone)
            VALUES (%s, %s, %s, %s, %s, %s)
            """,
            (
                company["企業名"],
                company["郵便番号"],
                company["都道府県"],
                company["住所"],
                company["担当者氏名"],
                company["電話番号"],
            ),
        )


def main():
    start_total = time()
    print("⏳ CSV読み込み中...")
    start_csv = time()
    companies = load_csv(CSV_FILE)
    print(f"✅ CSV読み込み完了: {time() - start_csv:.2f}s")

    print("⏳ DB接続中...")
    conn = psycopg.connect(DB_URL)
    conn.autocommit = True  # テスト用（本番はトランザクション管理推奨）

    print("⏳ DBにインサート中...")
    start_insert = time()
    for c in companies:
        insert_into_db(conn, c)
    print(f"✅ DB挿入完了: {time() - start_insert:.2f}s")

    print(f"🏁 全体処理時間: {time() - start_total:.2f}s")
    conn.close()


if __name__ == "__main__":
    main()
