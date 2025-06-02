# main.py

import logging
import traceback

import sqltap
from app.db.database import get_db  # session生成関数
from app.models.models import Company
from app.schemas.schemas import CompanyResponse
from fastapi import Depends, FastAPI, HTTPException
from sqlalchemy import event
from sqlalchemy.engine import Engine
from sqlalchemy.orm import Session

# ======================================
# 🔧 ログ設定
# ======================================

# 通常のクエリログ
logging.basicConfig(
    filename="sqltap_report.log",
    filemode="a",
    format="%(asctime)s - %(levelname)s - %(message)s",
    level=logging.INFO,
)

logger = logging.getLogger(__name__)

# 閾値超過用の警告ログ（別ファイル）
alert_handler = logging.FileHandler("sqltap_alert.log")
alert_handler.setLevel(logging.WARNING)

alert_logger = logging.getLogger("sqltap.alert")
alert_logger.setLevel(logging.WARNING)
alert_logger.addHandler(alert_handler)

# ======================================
# 🚀 FastAPI アプリケーション
# ======================================

app = FastAPI()

# クエリ数の閾値
QUERY_THRESHOLD = 100


@app.get("/companies-n1/{key_word}", response_model=list[CompanyResponse])
def search_companies(key_word: str, db: Session = Depends(get_db)):
    profiler = sqltap.start()  # クエリ記録開始

    companies = (
        db.query(Company).filter(Company.name.ilike(f"%{key_word}%")).limit(500).all()
    )
    print(f"companies: {len(companies)}")

    results = []
    try:
        for company in companies:
            results.append(CompanyResponse.from_orm(company))
    except Exception:
        traceback.print_exc()
        raise HTTPException(status_code=500, detail="Internal error")

    print(f"results: {len(results)}")

    # ======================================
    # 🔍 SQLTAP レポート処理
    # ======================================
    stats = profiler.collect()
    report_text = sqltap.report(stats, report_format="text")
    # クエリの行数をカウント（見出しや空行を除外）
    query_lines = [
        line
        for line in report_text.splitlines()
        if line.strip()
        and not line.startswith("Total")
        and not line.startswith("SQLTap")
    ]
    total_queries = len(query_lines)

    # 通常のレポートログ
    logger.info(
        "🔍 SQLTAP Report (%d queries)\n%s\n🔍 End Report", total_queries, report_text
    )

    # 閾値超過ログ
    if total_queries > QUERY_THRESHOLD:
        alert_logger.warning(
            "⚠️ QUERY COUNT EXCEEDED (%d queries)\n%s", total_queries, report_text
        )

    return results


query_log = []  # クエリを記録するリスト（毎回リセット推奨）


# フック登録（DB接続後のエンジンに対して）
@event.listens_for(Engine, "before_cursor_execute")
def before_cursor_execute(conn, cursor, statement, parameters, context, executemany):
    query_log.append(statement)


# クエリ数閾値
QUERY_LIMIT = 10

app = FastAPI()


@app.get("/companies-n1-hook/{key_word}", response_model=list[CompanyResponse])
def search_companies_hook(key_word: str, db: Session = Depends(get_db)):
    query_log.clear()  # ← 毎回リセットするのが重要

    companies = (
        db.query(Company).filter(Company.name.ilike(f"%{key_word}%")).limit(500).all()
    )

    try:
        results = [CompanyResponse.from_orm(c) for c in companies]
    except Exception:
        import traceback

        traceback.print_exc()
        raise HTTPException(status_code=500, detail="Internal error")

    total_queries = len(query_log)

    print("🔥 Captured SQL Queries:")
    for q in query_log:
        print(q.strip())

    if total_queries > QUERY_LIMIT:
        print(f"⚠️ クエリ数が閾値({QUERY_LIMIT})を超えました: {total_queries}件")

    return results
