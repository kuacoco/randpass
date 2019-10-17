FROM golang:1.12-alpine as builder

RUN apk add --no-cache git

RUN go install github.com/kuacoco/randpass


FROM alpine:3.10

RUN apk add --no-cache --no-progress curl tini ca-certificates

COPY --from=builder /go/bin/randpass /usr/bin/randpass

ENTRYPOINT ["/sbin/tini", "--"]
CMD ["randpass", "-n", "8"]