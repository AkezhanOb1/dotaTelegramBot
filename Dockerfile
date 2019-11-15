FROM golang:1.13-alpine AS build_base
LABEL maintainer="Yesbolatov Akezhan <esbolatovakezhan@gmail.com>"
WORKDIR /go/src/dotaPollBot
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .
CMD ["./main"]