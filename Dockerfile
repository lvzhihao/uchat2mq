FROM golang:1.9 as builder
WORKDIR /go/src/github.com/lvzhihao/uchat2mq
COPY . . 
RUN make build 

FROM alpine:latest  
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /usr/local/uchat2mq
COPY --from=builder /go/src/github.com/lvzhihao/uchat2mq/uchat2mq .
COPY ./docker-entrypoint.sh  .
ENV PATH /usr/local/uchat2mq:$PATH
RUN chmod +x /usr/local/uchat2mq/docker-entrypoint.sh
ENTRYPOINT ["/usr/local/uchat2mq/docker-entrypoint.sh"]
