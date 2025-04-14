FROM golang:alpine

# Install necessary dependencies
RUN apk add --no-cache ca-certificates && \
    update-ca-certificates

WORKDIR /app

COPY . .

RUN go mod tidy

CMD ["go", "run", "cmd/main.go", "deployment"]
