from app.db.database import Base
from sqlalchemy import TIMESTAMP, Column, Integer, String, func


class PostalCode(Base):
    __tablename__ = "postal_codes"

    id = Column(Integer, primary_key=True, index=True)
    code = Column(String, unique=True, nullable=False)
    created_at = Column(TIMESTAMP, server_default=func.now())
    updated_at = Column(TIMESTAMP, server_default=func.now(), onupdate=func.now())
