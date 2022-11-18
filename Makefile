

all: build
	docker-compose up

build:
	docker-compose build

serv:
	go run main.go


