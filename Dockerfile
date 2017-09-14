FROM golang:1.9

COPY . /go/src/github.com/lvzhihao/uchat2mq 

WORKDIR /go/src/github.com/lvzhihao/uchat2mq

RUN go-wrapper install

CMD ["go-wrapper", "run", "start"]
