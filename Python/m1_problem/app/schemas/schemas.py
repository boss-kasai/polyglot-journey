# schemas.py
from pydantic import BaseModel


class CompanyResponse(BaseModel):
    id: int
    name: str
    prefecture: str
    postal_code: str

    @classmethod
    def from_orm(cls, company):
        return cls(
            id=company.id,
            name=company.name,
            prefecture=company.prefecture.name,  # 🔥 ここでN+1
            postal_code=company.postal_code.code,  # 🔥 ここでもN+1
        )
