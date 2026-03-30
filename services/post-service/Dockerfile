FROM alpine:latest

WORKDIR /app

COPY post-service .
COPY .env .
RUN chmod +x post-service

CMD ["./post-service"]