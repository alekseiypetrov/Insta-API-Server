FROM alpine:latest

WORKDIR /app

COPY user-service .
RUN chmod +x user-service

CMD ["./user-service"]