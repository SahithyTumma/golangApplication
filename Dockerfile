FROM golang:1.19.0

WORKDIR /usr/src/app
COPY go.mod .
COPY go.sum .
COPY . .
RUN go mod tidy