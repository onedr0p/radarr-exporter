#!/bin/sh

apkArch="$(apk --print-arch)"
case "$apkArch" in
    armv7) export GOARCH='arm' GOARM=7 ;;
    aarch64) export GOARCH='arm64' ;;
    x86_64) export GOARCH='amd64' ;;
esac

go mod vendor
go build -o radarr-exporter ./cmd/radarr-exporter/