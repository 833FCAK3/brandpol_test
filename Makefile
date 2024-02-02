all: build migrate up

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

test:
	docker-compose run py_app pytest
