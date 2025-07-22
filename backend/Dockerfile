
FROM golang:1.21-alpine as builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

EXPOSE 8000
CMD ["./main"]