from contextlib import asynccontextmanager
from datetime import datetime
from typing import Annotated, List

from fastapi import Depends, FastAPI, HTTPException, Query
from models import Attendance, engine
from sqlmodel import Session, SQLModel, select


def create_db_and_tables():
    SQLModel.metadata.create_all(engine)


def get_session():
    with Session(engine) as session:
        yield session


SessionDep = Annotated[Session, Depends(get_session)]


@asynccontextmanager
async def lifespan(app: FastAPI):
    create_db_and_tables()
    yield
    # Cleanup code can be added here if needed


app = FastAPI(lifespan=lifespan)


@app.post("/attendances/", response_model=Attendance)
def create_attendance(session: SessionDep) -> Attendance:
    attendance = Attendance(punch_in=datetime.now())
    session.add(attendance)
    session.commit()
    session.refresh(attendance)
    return attendance


@app.get("/attendances/", response_model=List[Attendance])
def read_attendances(
    session: SessionDep,
    skip: int = Query(0, ge=0),
    limit: int = Query(10, ge=1, le=100),
) -> List[Attendance]:
    statement = select(Attendance).offset(skip).limit(limit)
    attendances = session.exec(statement).all()
    return attendances


@app.get("/attendances/{attendance_id}", response_model=Attendance)
def read_attendance(attendance_id: int, session: SessionDep) -> Attendance:
    statement = select(Attendance).where(Attendance.id == attendance_id)
    attendance = session.exec(statement).first()
    if not attendance:
        raise HTTPException(status_code=404, detail="Attendance not found")
    return attendance


@app.put("/attendances/{attendance_id}", response_model=Attendance)
def update_attendance(attendance_id: int, session: SessionDep) -> Attendance:
    statement = select(Attendance).where(Attendance.id == attendance_id)
    attendance = session.exec(statement).first()
    if not attendance:
        raise HTTPException(status_code=404, detail="Attendance not found")

    attendance.punch_out = datetime.now()
    attendance.updated_at = datetime.now()
    session.add(attendance)
    session.commit()
    session.refresh(attendance)
    return attendance


@app.delete("/attendances/{attendance_id}", response_model=Attendance)
def delete_attendance(attendance_id: int, session: SessionDep) -> Attendance:
    statement = select(Attendance).where(Attendance.id == attendance_id)
    attendance = session.exec(statement).first()
    if not attendance:
        raise HTTPException(status_code=404, detail="Attendance not found")

    session.delete(attendance)
    session.commit()
    return attendance


from fastapi_mcp import FastApiMCP

mcp = FastApiMCP(
    app,
    name="My API MCP",
    description="My API description",
)

# Mount the MCP server directly to your FastAPI app
mcp.mount()
