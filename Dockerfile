FROM golang:1.21.0-alpine3.18 AS build

WORKDIR /build

RUN apk add --no-cache git gcc musl-dev

COPY . .

RUN go build -o ./bin/go-template .

FROM alpine:3.18

WORKDIR /app

COPY --from=build /build/bin/go-template /app/

RUN apk add --no-cache ca-certificates && \
    addgroup -S -g 5000 go-template && \
    adduser -S -u 5000 -G go-template go-template && \
    chown -R go-template:go-template .

USER go-template
EXPOSE 8080

CMD ["/app/go-template"]
