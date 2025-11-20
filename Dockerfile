# syntax=docker/dockerfile:1

FROM golang:latest

WORKDIR "/github.com/phthallo/fanfi.cx"

COPY go.mod go.sum ./

COPY main.go ./

COPY internal ./internal

COPY pkg ./pkg

RUN go get 

RUN CGO_ENABLED=0 GOOS=linux go build -o /fanficx

CMD ["/fanficx"]
