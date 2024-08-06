FROM golang:latest

WORKDIR /usr/src/app
COPY go.mod .
COPY go.sum .
COPY . .
RUN go mod tidy