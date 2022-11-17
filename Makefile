CONTAINERID = $(shell docker ps -a | grep avito_avito | cut -b 1-12)

all:
	docker-compose up

serv:
	go run main.go
stop:
	docker stop $(CONTAINERID)
