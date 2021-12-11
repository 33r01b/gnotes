all: up run stop

up:
	docker run --rm --name gnotes-redis -p 6379:6379 -d redis:6.2-alpine

stop:
	docker stop gnotes-redis 

run:
	go run ./main.go
