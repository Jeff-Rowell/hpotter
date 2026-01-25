ARG GO_VERSION="1.25.3"
FROM golang:${GO_VERSION}-alpine3.22 AS builder
COPY . /app
WORKDIR /app
RUN go build -o main main.go

FROM alpine:3.22
WORKDIR /app
COPY --from=builder /app/public /app/public
COPY --from=builder /app/main /app/main
CMD ["/app/main"]
EXPOSE 80
