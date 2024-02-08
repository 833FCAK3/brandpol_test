include .env

all: build migrate up

test: testpy testgo

build:
	docker-compose build --no-cache

down:
	docker-compose down

up:
	docker-compose up -d

upp:
	docker-compose up

migrate:
	docker-compose run py_app bash -c '/wait && alembic upgrade head'

testgo:
	docker-compose run go_app go test -v

testpy:
	docker-compose run py_app pytest -v

devgo:
	docker-compose run --rm --volume=${CURDIR}/src/go_app:/src/go_app go_app sh -c "/wait && go run app.go config.go"

devtestgo:
	docker-compose run --rm --volume=${CURDIR}/src/go_app:/src/go_app go_app sh -c "/wait && go test -v"

devpy:
	docker-compose run --rm --volume=${CURDIR}/src/py_app:/src/py_app py_app bash -c "uvicorn app:app --host 0.0.0.0 --port ${PY_PORT}"

devtestpy:
	docker-compose run --rm --volume=${CURDIR}/src/py_app:/src/py_app py_app pytest -v
