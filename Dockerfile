# Build stage
FROM golang:1.21.7-alpine3.18 AS BuildStage

WORKDIR /bot

COPY ./bot .

RUN go mod download

RUN go build -o /build main.go

# Deploiement stage
FROM alpine:latest

WORKDIR /

COPY --from=BuildStage /build /build

COPY .env /

ENTRYPOINT [ "/build" ]