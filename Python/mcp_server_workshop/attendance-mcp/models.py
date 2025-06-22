from datetime import datetime

from sqlmodel import Field, SQLModel, create_engine


# モデルを定義
class Attendance(SQLModel, table=True):
    """
    出勤・退勤のモデル
    Attributes:
        id (int | None): 出勤・退勤のID
        punch_in (datetime): 出勤時間
        punch_out (datetime | None): 退勤時間
        created_at (datetime): レコード作成日時
        updated_at (datetime | None): レコード更新日時
    """

    id: int | None = Field(default=None, primary_key=True)
    punch_in: datetime
    punch_out: datetime | None = Field(default=None)
    created_at: datetime = Field(default_factory=datetime.now)
    updated_at: datetime = Field(nullable=True)


# データベースの接続設定
sqlite_database_url = "sqlite:///./database.db"
connect_args = {"check_same_thread": False}
engine = create_engine(sqlite_database_url, connect_args=connect_args, echo=True)
