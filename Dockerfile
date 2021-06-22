FROM golang:1.16-alpine AS base
WORKDIR /app

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0

# System dependencies
RUN apk update \
    && apk add --no-cache \
        ca-certificates \
        git \
        wget \
        curl \
    && update-ca-certificates

# Application dependencies
COPY . /app
RUN go mod download \
    && go mod verify

### Development
FROM base AS dev
WORKDIR /app

ARG AIR_VERSION
ARG AIR_OS
ARG AIR_ARCH

# Hot reloading mod
RUN wget -O air https://github.com/cosmtrek/air/releases/download/v${AIR_VERSION}/air_${AIR_VERSION}_${AIR_OS}_${AIR_ARCH} \
    && mv ./air /usr/local/bin/air \
    && chmod +x /usr/local/bin/air \
    && git clone https://github.com/go-delve/delve.git \
    && cd delve \
    && go install github.com/go-delve/delve/cmd/dlv 
     
EXPOSE 8080
EXPOSE 2345

ENTRYPOINT ["air"]

### Executable builder
FROM base AS builder
WORKDIR /app

RUN go build -o shortener -a .

### Production
FROM alpine:latest

ARG APP_VERSION="v0.0.1"
ENV APP_VERSION="${APP_VERSION}"

RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    curl \
    tzdata \
    && update-ca-certificates

# Copy executable and use an unprivileged user
COPY --from=builder /app/shortener /usr/local/bin/shortener
EXPOSE 8080

ENTRYPOINT ["shortener"]
