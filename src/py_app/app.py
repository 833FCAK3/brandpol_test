import json

from fastapi import Depends, FastAPI
from fastapi.encoders import jsonable_encoder
from sqlalchemy.orm import Session
from starlette.responses import Response

from database import get_postgres_db
from models import Greeting

app = FastAPI()


@app.get("/greet")
def greet(name: str = "пользователь", db: Session = Depends(get_postgres_db)) -> Response:
    name = name.capitalize()
    db_greeting = Greeting(name=name)
    db.add(db_greeting)
    db.commit()
    data = f"Привет, {name}, от Python!"
    return Response(status_code=200, content=data, media_type="text/plain")


@app.get("/greet/history")
def get_history(db: Session = Depends(get_postgres_db)) -> Response:
    data = json.dumps(jsonable_encoder(db.query(Greeting).all()), ensure_ascii=False)
    return Response(status_code=200, content=data, media_type="application/json")
