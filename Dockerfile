FROM golang:latest

RUN mkdir /app
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o main main.go

EXPOSE 8080

CMD ["/app/main"]

