# syntax=docker/dockerfile:1
ARG GO_VERSION=1.23.2
FROM golang:${GO_VERSION} AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download



COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /router-os


FROM gcr.io/distroless/static-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /router-os router-os
COPY .env .env

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT [ "/router-os" ]