from app.db.database import Base
from sqlalchemy import TIMESTAMP, Column, ForeignKey, Integer, String, func


class Company(Base):
    __tablename__ = "companies"

    id = Column(Integer, primary_key=True, index=True)
    name = Column(String, nullable=False)
    postal_code_id = Column(Integer, ForeignKey("postal_codes.id"))
    prefecture_id = Column(Integer, ForeignKey("prefectures.id"))
    address = Column(String)
    contact_name = Column(String)
    phone = Column(String)
    created_at = Column(TIMESTAMP, server_default=func.now())
    updated_at = Column(TIMESTAMP, server_default=func.now(), onupdate=func.now())
