FROM alpine:latest

WORKDIR /app

COPY user-service .
COPY .env .
RUN chmod +x user-service

CMD ["./user-service"]