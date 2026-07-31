.PHONY: build up down stop restart go

build:
	docker compose build

up:
	docker compose up -d

down:
	docker compose down

stop:
	docker compose stop

restart:
	docker compose restart

go:
	go build ./... && docker compose restart api
