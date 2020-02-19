FROM golang:1.13.8-alpine3.11

RUN mkdir -p /go/src/github.com/rk23/hinge
COPY . /go/src/github.com/rk23/hinge
WORKDIR /go/src/github.com/rk23/hinge

RUN apk update && apk add dep git
RUN dep ensure --vendor-only
RUN go build -o ./bin/api ./cmd/api/ 

CMD ["./bin/api"]