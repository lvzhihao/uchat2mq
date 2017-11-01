FROM golang:1.9

WORKDIR /go/src/github.com/lvzhihao/uchat2mq

COPY . . 

RUN rm -f .uchat2mq.yaml
RUN go-wrapper install

CMD ["go-wrapper", "run", "receive"]
