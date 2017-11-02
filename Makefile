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

docker-ccr: 
	sudo docker tag edwinlll/uchat2mq:latest ccr.ccs.tencentyun.com/wdwd/uchat2mq
	sudo docker push ccr.ccs.tencentyun.com/wdwd/uchat2mq

docker-uhub:
	sudo docker tag edwinlll/uchat2mq:latest uhub.service.ucloud.cn/mmzs/uchat2mq
	sudo docker push uhub.service.ucloud.cn/mmzs/uchat2mq
