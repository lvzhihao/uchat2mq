FROM golang:1.9

COPY . /go/src/github.com/lvzhihao/uchat2mq 

WORKDIR /go/src/github.com/lvzhihao/uchat2mq

RUN rm /go/src/github.com/lvzhihao/uchat2mq/.uchat2mq.yaml
RUN go-wrapper install

CMD ["go-wrapper", "run", "receive"]
