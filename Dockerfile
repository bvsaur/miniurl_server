FROM --platform=linux/amd64 golang:1.18.3-alpine

COPY ./ /go/src/github.com/bvsaur/miniurl
WORKDIR /go/src/github.com/bvsaur/miniurl

COPY ./.env.prod .env

RUN go mod download
RUN go build -o miniurl cmd/miniurl/main.go

CMD ./miniurl

EXPOSE 8080