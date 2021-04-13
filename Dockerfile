FROM golang:1.16.2-alpine as builder

ENV GOPATH=/

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./

# build go app
RUN go build -o gateway ./cmd/main.go

#Build destination container
FROM alpine:latest

ENV GOPATH=/go

# install psql
RUN apk --update add postgresql-client

# copy bin and pg-wait script
COPY --from=builder $GOPATH/gateway ./

# copy PG migrations
COPY --from=builder $GOPATH/pkg/repository/pg/migrations/*.sql ./migrations/

EXPOSE 8080

CMD ["./gateway", "run"]