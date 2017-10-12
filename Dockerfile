FROM golang:1.9.1

COPY . /go/src/github.com/lvzhihao/uchat2mq 

WORKDIR /go/src/github.com/lvzhihao/uchat2mq

RUN rm -f /go/src/github.com/lvzhihao/uchat2mq/.uchat2mq.yaml
RUN go-wrapper install

CMD ["go-wrapper", "run", "receive"]
