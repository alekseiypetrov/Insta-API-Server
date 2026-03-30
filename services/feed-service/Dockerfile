FROM golang:1.25.6

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o feed-service ./services/feed-service/cmd/main.go
CMD ["./feed-service"]