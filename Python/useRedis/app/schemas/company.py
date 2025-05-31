from datetime import datetime

from pydantic import BaseModel


class CompanyCreate(BaseModel):
    name: str
    postal_code_id: int
    prefecture_id: int
    address: str | None = None
    contact_name: str | None = None
    phone: str | None = None


class CompanyOut(CompanyCreate):
    id: int
    created_at: datetime
    updated_at: datetime

    class Config:
        orm_mode = True
