FROM alpine:latest

WORKDIR /app

COPY feed-service .
RUN chmod +x feed-service

CMD ["./feed-service"]