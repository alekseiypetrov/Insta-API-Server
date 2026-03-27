FROM golang:1.25.6

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o post-service ./services/post-service/cmd/main.go
CMD ["./post-service"]