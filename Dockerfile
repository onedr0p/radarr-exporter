FROM golang:1.13-alpine as build

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

# hadolint ignore=DL3018
RUN apk add --no-cache curl ca-certificates git

WORKDIR /go/src/github.com/onedr0p/radarr-exporter
COPY . .
RUN chmod +x build.sh \
    && sh build.sh

FROM alpine:3.11

# hadolint ignore=DL3018
RUN apk add --no-cache ca-certificates tini curl

COPY --from=build /go/src/github.com/onedr0p/radarr-exporter/radarr-exporter /usr/local/bin/radarr-exporter
RUN chmod +x /usr/local/bin/radarr-exporter

ENTRYPOINT ["/sbin/tini", "--", "radarr-exporter"]