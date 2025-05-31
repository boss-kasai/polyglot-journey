from datetime import datetime

from pydantic import BaseModel


class PrefectureBase(BaseModel):
    name: str


class PrefectureCreate(PrefectureBase):
    pass


class PrefectureUpdate(PrefectureBase):
    pass


class PrefectureOut(PrefectureBase):
    id: int
    created_at: datetime
    updated_at: datetime

    class Config:
        orm_mode = True
