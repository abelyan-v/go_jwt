FROM golang:1.22-alpine as builder
RUN apk add --no-cache gcc musl-dev
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -o /app/main .
FROM alpine:3.19
RUN apk add --no-cache postgresql-client
COPY --from=builder /app/main /app/main
EXPOSE 8080
CMD ["/app/main"]