FROM golang:alpine3.15 AS builder
WORKDIR /app
COPY vendor vendor
COPY . .
ENV CGO_ENABLED=0
RUN --mount=type=cache,target=/root/.cache/go-build go build -ldflags="-w -s" -o main -mod=vendor main.go
# RUN --mount=type=cache,target=/root/.cache/go-build go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

FROM alpine:3.15
COPY --from=builder /app/main .
# COPY --from=builder /go/bin/migrate .
COPY app.env .
COPY start.sh .
copy wait-for .
COPY db/migration ./migration
EXPOSE 9001
CMD [ "./main" ]
ENTRYPOINT [ "./start.sh" ]
