FROM golang:1.19-alpine

WORKDIR /app

COPY ./ ./

RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -o bin/cmd cmd/main.go

ENTRYPOINT ["./bin/cmd"]