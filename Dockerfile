FROM golang:1.21-alpine as dev

EXPOSE 4000

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main cmd/server/main.go

CMD ["./main"]