from app import app
from fastapi.testclient import TestClient


def test_greet():
    client = TestClient(app)

    # Тест с указанием имени
    response = client.get("/greet?name=John")
    assert response.status_code == 200
    assert response.text == "Привет, John, от Python!"

    # Тест без указания имени
    response_default = client.get("/greet")
    assert response_default.status_code == 200
    assert response_default.text == "Привет, Пользователь, от Python!"
