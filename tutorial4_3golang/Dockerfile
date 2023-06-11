FROM golang:latest
WORKDIR /app

COPY go.mod .

COPY . .

RUN go build -o api

CMD ["./api"]


