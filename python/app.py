from fastapi import FastAPI

app = FastAPI()

@app.get("/greet")
def greet(name: str = "пользователь"):
    return f"Привет, {name.capitalize()}, от Python!"

@app.get("/greet/history")
def get_history():
    pass
