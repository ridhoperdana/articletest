## Builder
FROM golang:1.14-alpine3.12 as builder

RUN apk update && apk upgrade && \
    apk --no-cache --update add git make gcc musl-dev ca-certificates openssh && \
    rm -rf /var/cache/apk/*

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o articletest github.com/ridhoperdana/articletest/cmd/articletest

## Distribution
FROM alpine:3.12

RUN apk update && apk upgrade \
    && apk --no-cache --update add ca-certificates tzdata \
    && rm -f /var/cache/apk/*

ENV TZ Asia/Jakarta

WORKDIR /app

COPY --from=builder /app/articletest .

EXPOSE 6969

CMD ["./articletest", "http"]
