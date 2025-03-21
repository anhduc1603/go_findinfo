# syntax=docker/dockerfile:1
FROM golang:1.23-alpine

WORKDIR /app

# Add before RUN go mod download or go build
ENV GOPROXY=https://goproxy.cn,direct

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

# Copy .env file
COPY .env .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
