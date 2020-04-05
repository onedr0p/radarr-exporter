FROM golang:1.14-alpine as build

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

RUN apk add --no-cache curl ca-certificates git alpine-sdk upx

WORKDIR /go/src/github.com/onedr0p/radarr-exporter
COPY . .

RUN export GOOS=$(echo ${TARGETPLATFORM} | cut -d / -f1) \
    && export GOARCH=$(echo ${TARGETPLATFORM} | cut -d / -f2) \
    && GOARM=$(echo ${TARGETPLATFORM} | cut -d / -f3); export GOARM=${GOARM:1} \
    && go mod vendor \
    && go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o radarr-exporter ./cmd/radarr-exporter/ \
    && upx -f --brute radarr-exporter \
    && chmod +x radarr-exporter

FROM alpine:3.11
RUN apk add --no-cache ca-certificates tini curl
COPY --from=build /go/src/github.com/onedr0p/radarr-exporter/radarr-exporter /usr/local/bin/radarr-exporter
ENTRYPOINT ["/sbin/tini", "--", "radarr-exporter"]
