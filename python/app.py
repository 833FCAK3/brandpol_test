from fastapi import FastAPI, Depends
from sqlalchemy.orm import Session
from database import get_postgres_db
from models import Greeting

app = FastAPI()


@app.get("/greet")
def greet(name: str = "пользователь", db: Session = Depends(get_postgres_db)):
    name = name.capitalize()
    db_greeting = Greeting(name=name)
    db.add(db_greeting)
    db.commit()
    return f"Привет, {name}, от Python!"


@app.get("/greet/history")
def get_history(db: Session = Depends(get_postgres_db)):
    return db.query(Greeting).all()
