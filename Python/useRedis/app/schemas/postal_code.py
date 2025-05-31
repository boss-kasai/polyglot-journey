from datetime import datetime

from pydantic import BaseModel


class PostalCodeBase(BaseModel):
    code: str


class PostalCodeCreate(PostalCodeBase):
    pass


class PostalCodeUpdate(PostalCodeBase):
    pass


class PostalCodeOut(PostalCodeBase):
    id: int
    created_at: datetime
    updated_at: datetime

    class Config:
        orm_mode = True
