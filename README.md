# brandpol_test

## Start
Клонировать репозиторий:
````
>>> git clone https://github.com/833FCAK3/brandpol_test.git
````
- Установить [docker](https://docs.docker.com/engine/install/)
- Установить [docker-compose](https://docs.docker.com/compose/install/)
- На Windows для удобства Установить [MAKE Windows](http://gnuwin32.sourceforge.net/packages/make.htm)

---

## Запуск проекта
- В .env файле настроить порты приложений и баз данных
- Запустить. 
````
>>> make
````

*На Windows, если не поставил make:
```
>>> docker-compose build
>>> docker-compose run app bash -c '/wait && alembic upgrade head'
>>> docker-compose up -d
```

- Тесты.
````
>>> test
````

*Отдельно python и go приложения тестятся командами make testpy и make testgo.

---

## Структура приложения
* `/go_app` - приложение на go, использует фреймворк Echo для реализации API
* `/py_app` - приложение на python, использует фреймворк FastAPI для реализации API
* `/go_postgres` - содержит Dockerfile базы данных Postgresql, используемой go сервисом
* `/py_postgres` - содержит Dockerfile базы данных Postgresql, используемой python сервисом

---

## Эндпоинты

### На питоне
Все эндпойнты находятся в файле `py_app\app.py`

* `GET /greet` - эндпойнт приветствия пользователя. Сохраняет отправленное сообщение, предоставленное имя и дату + время запроса.
    Args:
        name (str): имя пользователя, default="пользователь"
        db (Session): соединение с базой данных
    Returns:
        Объект Response со статусным кодом 200 и string сообщением
    Examples:
        Request:
            http://127.0.0.1:8000/greet?name=alex
        Response:
            Привет, Пользователь, от Python!
        

* `GET /greet/history` - эндпойнт запроса истории приветствий пользователя.
    Args:
        db (Session): соединение с базой данных
    Returns:
        Объект Response со статусным кодом 200 и лист с json объектами, представляющими приветствия
    Examples:
        Request:
            http://127.0.0.1:8000/greet/history
        Response:
            [
                {
                    "name": "Lewa",
                    "date": "2024-02-01T11:53:18.598003",
                    "id": 1
                },
                {
                    "name": "Lewa",
                    "date": "2024-02-01T12:15:01.859239",
                    "id": 2
                }
            ]
---

### На Go
Все эндпойнты находятся в файле `go_app\app.go`

* `GET /greet` - эндпойнт приветствия пользователя. Сохраняет отправленное сообщение и дату + время запроса.
    Returns:
        Объект response со статусным кодом 200 и string объектом приветствия
    Examples:
        Request:
            http://localhost:8080/greet
        Response:
            Привет от Go!

* `GET /greet/history` - эндпойнт запроса истории приветствий пользователя.
    Returns:
        String response со статусным кодом 200. Текст - лист истории приветствий с json объектами, представляющими отдельные приветствия.
    Examples:
        Request:
            http://localhost:8080/greet/history
        Response:
            [{"id":2,"created_at":"2024-02-01T09:19:59.674128Z","message":"Привет от Go!"},{"id":1,"created_at":"2024-02-01T09:19:23.186561Z","message":"Привет от Go!"}]

* `GET /python_greet` - эндпойнт обращения к python эндпойнту /greet. Сохраняет отправленное сообщение, имя и дату + время запроса.
    Args:
        name (str): имя пользователя, обязателен
    Returns:
        Объект Response со статусным кодом 200 и string объектом приветствия
        Объект Response со статусным кодом 400 и уточняющим string объектом
        Объект Response со статусным кодом 500 и уточняющим string объектом
    Examples:
        Request:
            http://localhost:8080/python_greet?name=lexa
        Response:
            Привет, Lexa, от Python!

* `GET /python_greet_history` - эндпойнт обращения к python эндпойнту /greet/history. Сохраняет отправленное сообщение, имя и дату + время запроса.
    Returns:
        Объект Response со статусным кодом 200 и string объектом истории приветствий в виде листа с json объектами, представляющими отдельные приветствия
        Объект Response со статусным кодом 500 и уточняющим string объектом
    Examples:
        Request:
            http://localhost:8080/python_greet_history
        Response:
            [{"name": "Lexa", "date": "2024-02-01T14:06:05.428881", "id": 1}]
