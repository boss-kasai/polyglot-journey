# app/init_db.py
from app.db.database import Base, engine


def create_tables():
    Base.metadata.create_all(bind=engine)


if __name__ == "__main__":
    create_tables()
