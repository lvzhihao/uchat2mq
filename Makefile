OS := $(shell uname)

build: */*.go
	go build 

receive:
	./uchat2mq receive

migrate: build
	./uchat2mq migrate

docker-build:
	sudo docker build -t edwinlll/uchat2mq:latest .

docker-push:
	sudo docker push edwinlll/uchat2mq:latest
