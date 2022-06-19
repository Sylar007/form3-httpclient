FROM golang:1.16-alpine

ENV CGO_ENABLED 0

RUN set -xe \
    && mkdir /app

COPY go.* /app/

WORKDIR /app/accounts

RUN set -xe \
    && go mod download

CMD ["go", "test", "-v", "./..."]