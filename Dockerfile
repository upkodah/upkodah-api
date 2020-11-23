FROM golang:1.14.2 as builder

ENV SRC_DIR=$GOPATH/src/github.com/upkodah/upkodah-api

COPY . $SRC_DIR

WORKDIR $SRC_DIR
RUN CGO_ENABLED=0 GOOS=linux go build -o $GOPATH/bin/upkodah-api

FROM docker.io/library/ubuntu


WORKDIR /root/
COPY --from=builder /go/bin/upkodah-api .
EXPOSE $HTTP_PORT
CMD ["./upkodah-api"]