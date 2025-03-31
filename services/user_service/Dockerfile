FROM ubuntu:latest

RUN apt-get update && \
    apt-get install -y golang ca-certificates && \
    update-ca-certificates

WORKDIR /app

COPY . .

RUN go mod tidy

CMD ["go", "run", "cmd/server/main.go"]