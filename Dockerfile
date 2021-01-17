# syntax = docker/dockerfile:experimental

FROM golang:latest AS base
WORKDIR /src
ENV CGO_ENABLED=0
COPY go.* .
COPY items/go.* .
RUN go mod download
COPY . .

FROM base AS build
ENV GOOS=linux
RUN --mount=type=cache,target=/root/.cache/go-build go build -o app .

FROM alpine:latest  AS bin
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build /src .
CMD ["./app"]  