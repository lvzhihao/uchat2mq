FROM golang:1.9 as builder
WORKDIR /go/src/github.com/lvzhihao/uchat2mq
COPY . . 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /usr/local/uchat2mq
COPY --from=builder /go/src/github.com/lvzhihao/uchat2mq/uchat2mq .
ENV PATH /usr/local/uchat2mq:$PATH
