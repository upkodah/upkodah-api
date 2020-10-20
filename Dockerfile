FROM golang:1.14.2 as builder

COPY main.go .
CMD go run main.go


ENV SRC_DIR $GOPATH/src/github.com/upkodah/upkodah-api

COPY . $SRC_DIR

WORKDIR $SRC_DIR
RUN CGO_ENABLED=0 GOOS=linux go build -o $GOPATH/bin/upkodah-api

FROM docker.io/library/ubuntu

RUN useradd -ms /bin/bash pack
ADD bin /usr/bin/
ADD https://storage.googleapis.com/kubernetes-release/release/v1.15.0/bin/linux/amd64/kubectl /usr/bin/
RUN chmod +x /usr/bin/kubectl


WORKDIR /root/
COPY --from=builder /go/bin/upkodah-api .
CMD ["./upkodah-api"]
