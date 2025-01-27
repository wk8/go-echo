## docker buildx create --use
## docker buildx build --platform linux/amd64,linux/arm64 -t wk88/httpecho . --push
## docker run --rm -p 8282:8282 wk88:httpecho

FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.* .

RUN go mod download

COPY *.go .

RUN go build \
    -o echo \
    *.go

###

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/echo .

CMD ["/app/echo"]
