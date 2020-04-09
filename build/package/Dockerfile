FROM golang:1.14-alpine as build

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /build

COPY . .

RUN apk --no-cache add ca-certificates && \
    go mod download && \
    go build /build/cmd/radarr-exporter/.

FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=build /build/radarr-exporter /

ENTRYPOINT ["/radarr-exporter"]