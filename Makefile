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
	docker-compose run py_app pytest
