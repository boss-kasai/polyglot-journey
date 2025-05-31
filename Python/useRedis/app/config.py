from pydantic import BaseSettings


class Settings(BaseSettings):
    POSTGRES_HOST: str = "localhost"
    POSTGRES_PORT: int = 5432
    POSTGRES_DB: str = "company_db"
    POSTGRES_USER: str = "user"
    POSTGRES_PASSWORD: str = "password"

    class Config:
        env_file = ".env"


settings = Settings()
