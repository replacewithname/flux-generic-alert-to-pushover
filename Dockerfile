# syntax=docker/dockerfile:1
FROM golang:1.18 AS builder
WORKDIR /go/src/flux-generic-alert-to-pushover/
COPY . ./
# Use CGO_ENABLED=0 and GOOS=linux to run on alpine
RUN CGO_ENABLED=0 GOOS=linux go build -a -o server .

# FROM debian:11.3-slim
# RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/flux-generic-alert-to-pushover/server ./
CMD ["./server"]