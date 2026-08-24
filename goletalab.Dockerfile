# syntax=docker/dockerfile:1.7
FROM golang:1.26
WORKDIR /workspace
ENV GOPROXY=https://goproxy.cn,direct
COPY go.mod go.sum* ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/root/.cache/go-build go build ./...
