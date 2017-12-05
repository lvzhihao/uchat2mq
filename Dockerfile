FROM golang:1.9

WORKDIR /go/src/github.com/lvzhihao/uchat2mq

COPY . . 

RUN go-wrapper install && \
    rm -rf *

CMD ["go-wrapper", "run", "receive"]
