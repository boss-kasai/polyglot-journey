import csv

import httpx

INPUT_FILE = "companies_100k.csv"
OUTPUT_FILE = "companies_fixed.csv"
API_BASE = "http://localhost:8000"


def validate_postal_code(code: str) -> str:
    try:
        url = f"{API_BASE}/validate_postal_code/{code}"
        response = httpx.get(url, timeout=5.0)
        response.raise_for_status()
        result = response.json()
        if result.get("exists"):
            return code
        else:
            return result.get("suggested")
    except Exception as e:
        print(f"❌ API error for {code}: {e}")
        return "0600000"  # エラー時は元のコードを使う


def fix_csv_via_api():
    with (
        open(INPUT_FILE, newline="", encoding="utf-8") as infile,
        open(OUTPUT_FILE, "w", newline="", encoding="utf-8") as outfile,
    ):
        reader = csv.reader(infile)
        writer = csv.writer(outfile)

        headers = next(reader)
        writer.writerow(headers)

        for idx, row in enumerate(reader, start=2):
            original_code = row[1].strip()
            fixed_code = validate_postal_code(original_code)
            if original_code != fixed_code:
                print(f"🔁 Row {idx}: {original_code} → {fixed_code}")
                row[1] = fixed_code
            writer.writerow(row)

    print(f"✅ 修正済みCSVを {OUTPUT_FILE} に保存しました。")


if __name__ == "__main__":
    fix_csv_via_api()
