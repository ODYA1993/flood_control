.PHONY: build run clean

docker_up:
	docker-compose up -d

docker_stop:
	docker stop my-redis
