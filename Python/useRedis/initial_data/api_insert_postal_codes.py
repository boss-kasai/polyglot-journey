import csv

import requests

API_URL = "http://localhost:8000/postal_code"
CSV_FILE_PATH = "utf_ken_all_2.csv"  # パスは必要に応じて調整


def extract_postal_codes(csv_file: str) -> set[str]:
    postal_codes = set()
    with open(csv_file, newline="", encoding="utf-8") as f:
        reader = csv.reader(f)
        for row in reader:
            if len(row) >= 3:
                postal_code = row[2].strip().replace('"', "")
                postal_codes.add(postal_code)
    return postal_codes


def register_postal_codes(postal_codes: set[str]):
    for code in sorted(postal_codes):
        res = requests.post(API_URL, json={"code": code})
        if res.status_code == 200:
            print(f"[✓] 登録成功: {code}")
        elif res.status_code == 400 and "すでに存在" in res.text:
            print(f"[=] 重複スキップ: {code}")
        else:
            print(f"[✗] 登録失敗: {code} → {res.status_code}: {res.text}")


if __name__ == "__main__":
    postal_codes = extract_postal_codes(CSV_FILE_PATH)
    print(f"抽出件数: {len(postal_codes)} 件")
    register_postal_codes(postal_codes)
