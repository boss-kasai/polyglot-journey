# models.py
from sqlalchemy import TIMESTAMP, Column, ForeignKey, Integer, String, func
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import declarative_base, relationship

Base = declarative_base()


class Prefecture(Base):
    __tablename__ = "prefectures"

    id = Column(Integer, primary_key=True, index=True)
    name = Column(String, unique=True, nullable=False)
    created_at = Column(TIMESTAMP, server_default=func.now())
    updated_at = Column(TIMESTAMP, server_default=func.now(), onupdate=func.now())

    companies = relationship("Company", back_populates="prefecture")


class PostalCode(Base):
    __tablename__ = "postal_codes"

    id = Column(Integer, primary_key=True, index=True)
    code = Column(String, unique=True, nullable=False)
    created_at = Column(TIMESTAMP, server_default=func.now())
    updated_at = Column(TIMESTAMP, server_default=func.now(), onupdate=func.now())

    companies = relationship("Company", back_populates="postal_code")


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

    prefecture = relationship("Prefecture", back_populates="companies")
    postal_code = relationship("PostalCode", back_populates="companies")
