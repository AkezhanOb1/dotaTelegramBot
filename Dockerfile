FROM golang:1.13-alpine AS build_base
RUN apk add bash ca-certificates git gcc g++ libc-dev
WORKDIR /go/src/dotaPollBot
ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .
RUN go mod download
FROM build_base AS server_builder
COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go install -a -tags netgo -ldflags '-w -extldflags "-static"' ./main.go
FROM alpine AS weaviate
RUN apk add ca-certificates
COPY --from=server_builder /go/bin/main /bin/main
ENTRYPOINT ["/bin/main"]