# app/db/database.py

from sqlalchemy import create_engine
from sqlalchemy.orm import Session, declarative_base, sessionmaker

# DB接続設定
DATABASE_URL = "postgresql+psycopg://user:password@localhost:5432/company_db"

engine = create_engine(DATABASE_URL, echo=True)

# セッション設定
SessionLocal = sessionmaker(bind=engine, autoflush=False, autocommit=False)

# ベースクラス
Base = declarative_base()


# FastAPI依存関係用のセッション取得関数
def get_db() -> Session:
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()
