FROM golang:1.25.6

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o user-service ./services/user-service/cmd/main.go
CMD ["./user-service"]