# build
FROM golang:1.14-alpine as builder

RUN apk update && apk add tzdata \
    && cp /usr/share/zoneinfo/Asia/Bangkok /etc/localtime \
    && echo "Asia/Bangkok" >  /etc/timezone \
    && apk del tzdata

WORKDIR /app

COPY go.mod go.sum .
RUN go mod download

COPY . .

RUN go clean --cache

RUN make build

# ---------------------------------------------------------

# run
FROM alpine:latest

RUN apk update && apk add tzdata \
    && cp /usr/share/zoneinfo/Asia/Bangkok /etc/localtime \
    && echo "Asia/Bangkok" >  /etc/timezone \
    && apk del tzdata

WORKDIR /app

COPY --from=builder /app/build/goapp .
COPY --from=builder /app/config.json .

CMD ["./goapp"]
