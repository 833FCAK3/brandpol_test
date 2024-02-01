up:
	docker-compose up -d

migrate:
	docker-compose run --volume=${PWD}/python:/python app bash -c '/wait && alembic upgrade head'

makemigrations:
	docker-compose run --rm --no-deps --volume=${PWD}/python:/python app bash -c '/wait && alembic revision --autogenerate -m $n'
	sudo chown -R ${USER} python/app/alembic/versions
