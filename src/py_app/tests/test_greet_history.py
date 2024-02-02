from py_app.app import app
from fastapi.testclient import TestClient


def test_get_history():
    client = TestClient(app)

    response = client.get("/greet/history")

    assert response.status_code == 200

    assert response.headers["content-type"] == "application/json"

    # Проверка содержимого ответа
    history_data = response.json()
    assert isinstance(history_data, list)
    assert len(history_data) > 0
