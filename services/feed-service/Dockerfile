FROM alpine:latest

WORKDIR /app

COPY feed-service .
COPY .env .
RUN chmod +x feed-service

CMD ["./feed-service"]