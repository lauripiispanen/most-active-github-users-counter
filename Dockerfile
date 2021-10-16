FROM alpine:latest AS ca-certs
RUN apk add -U --no-cache ca-certificates

FROM golang:1.17 AS go-builder

WORKDIR /go/src/main
COPY . .

RUN go get -d -v .
RUN CGO_ENABLED=0 go install -v .

FROM scratch
COPY --from=go-builder /go/bin/most-active-github-users-counter /go/bin/most-active-github-users-counter
COPY --from=ca-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/go/bin/most-active-github-users-counter"]
