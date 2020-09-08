FROM golang:1.15-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app

FROM alpine:3.12
RUN apk --no-cache add ca-certificates
# Todo: Use non-root user
WORKDIR /root/ 
COPY --from=builder /app/app .
CMD ["./app"]
