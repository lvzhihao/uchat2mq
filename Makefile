OS := $(shell uname)
GOOS := $(shell echo "${OS}" | tr '[:upper:]' '[:lower:]')
COMMIT_BRANCH := $(shell git symbolic-ref --short -q HEAD)
COMMIT_HASH := $(shell git rev-parse HEAD 2>/dev/null)
COMMIT_DATE := $(shell git log --date=iso --pretty=format:"%cd" -1 2>/dev/null)
BUILD_DATE := $(shell date +"%F %T %z")
DOCKER_VERSION := $(shell expr substr "${COMMIT_HASH}" 1 6)-${COMMIT_BRANCH}-$(shell git log --date=iso --pretty=format:"%ct" -1 2>/dev/null)
IMPORT_PATH := github.com/lvzhihao/uchat2mq/cmd

build: */*.go
	CGO_ENABLED=0 GOOS=${GOOS} go build -a \
	-installsuffix cgo \
	-ldflags "-X \"${IMPORT_PATH}.BuildBranch=${COMMIT_BRANCH}\" -X \"${IMPORT_PATH}.BuildVersion=${COMMIT_HASH}\" -X \"${IMPORT_PATH}.BuildDate=${BUILD_DATE}\"" .

info:
	@echo ${GOOS}
	@echo ${COMMIT_BRANCH}
	@echo ${COMMIT_HASH}
	@echo ${COMMIT_DATE}
	@echo ${BUILD_DATE}
	@echo ${DOCKER_VERSION}

receive: build
	./uchat2mq receive

migrate: build
	./uchat2mq migrate

docker-build:
	sudo docker build -t edwinlll/uchat2mq:${DOCKER_VERSION} .

docker-latest: docker-build
	sudo docker tag edwinlll/uchat2mq:${DOCKER_VERSION} edwinlll/uchat2mq:latest

docker-push:
	sudo docker push edwinlll/uchat2mq:latest

docker-ccr: 
	sudo docker tag edwinlll/uchat2mq:latest ccr.ccs.tencentyun.com/wdwd/uchat2mq:latest
	sudo docker push ccr.ccs.tencentyun.com/wdwd/uchat2mq:latest
	sudo docker rmi ccr.ccs.tencentyun.com/wdwd/uchat2mq:latest

docker-ali:
	sudo docker tag edwinlll/uchat2mq:latest registry.cn-hangzhou.aliyuncs.com/weishangye/uchat2mq:latest
	sudo docker push registry.cn-hangzhou.aliyuncs.com/weishangye/uchat2mq:latest
	sudo docker rmi registry.cn-hangzhou.aliyuncs.com/weishangye/uchat2mq:latest

docker-wdwd:
	sudo docker tag edwinlll/uchat2mq:latest docker.wdwd.com/wxsq/uchat2mq:latest
	sudo docker push docker.wdwd.com/wxsq/uchat2mq:latest
	sudo docker rmi docker.wdwd.com/wxsq/uchat2mq:latest
