all: build down migrate up

build:
	docker-compose build

down:
	docker-compose down

up:
	docker-compose up -d

migrate:
	docker-compose run py_app bash -c '/wait && alembic upgrade head'
