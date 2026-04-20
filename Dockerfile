FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY . .
RUN go build -mod=vendor -o app .

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/app .
CMD [ "/app/app" ]