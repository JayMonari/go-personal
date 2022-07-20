FROM golang:stretch AS builder

ENV GOPATH=

WORKDIR .

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/main server/server.go

FROM gcr.io/distroless/base

COPY --from=builder /go/bin/main /go/bin/main

ENTRYPOINT ["/go/bin/main"]
