import csv

import requests

# APIのエンドポイント
API_URL = "http://localhost:8080/postal_codes"

# 送信間隔（秒）
SLEEP_TIME = 0.5

# CSVファイルのパス
CSV_FILE_PATH = "utf_ken_all.csv"

# 列番号の指定 (0-based index)
POSTAL_CODE_COL = 2  # 郵便番号が入っている列
ADDRESS_COL = [6, 7, 8]  # 住所が入っている列 -> 7 + 8 + 9


def read_postal_codes(file_path):
    """CSVファイルから郵便番号データを読み込む（ヘッダーなし）"""
    with open(file_path, newline="", encoding="utf-8") as csvfile:
        reader = csv.reader(csvfile)
        for row in reader:
            # 列の範囲チェックを修正
            if len(row) > max(POSTAL_CODE_COL, max(ADDRESS_COL)):
                my_address = ""
                for i in ADDRESS_COL:
                    my_address += row[i].strip()
                # ちょっと整理
                ## my_addressに（が含まれていたら、それ以降を削除
                if "（" in my_address:
                    my_address = my_address.split("（")[0]
                ## "以下に掲載がない"が含まれていたら、それ以降を削除
                if "以下に掲載がない" in my_address:
                    my_address = my_address.split("以下に掲載がない")[0]
                yield {
                    "postal_code": row[POSTAL_CODE_COL].strip(),
                    "address": my_address,
                }
            else:
                print(f"無効な行をスキップ: {row}")


def send_postal_code(data):
    """GoのAPIにPOSTリクエストを送信"""
    try:
        response = requests.post(API_URL, json=data, timeout=5)
        if response.status_code == 201:
            print(f"成功: {data['postal_code']}")
        else:
            print(
                f"失敗: {data['postal_code']} - {response.status_code} {response.text}"
            )
    except requests.RequestException as e:
        print(f"エラー: {data['postal_code']} - {e}")


def main():
    """メイン処理"""
    for postal_code_data in read_postal_codes(CSV_FILE_PATH):
        print(postal_code_data)
        send_postal_code(postal_code_data)
        # time.sleep(SLEEP_TIME)  # 連続リクエストを防ぐための待機


if __name__ == "__main__":
    main()
