
from sqlalchemy import Column, Integer, String, DateTime, func
from database import Base

class greetings(Base):
    __tablename__ = "greetings"

    id = Column(Integer, primary_key=True, index=True)
    name = Column(String, index=True)
    date = Column(DateTime, default=func.now())
