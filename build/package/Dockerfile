FROM golang:1.14-alpine as build

ENV GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /build

COPY . .

RUN export GOOS=${TARGETPLATFORM%%/*} && \
    export GOARCH=$(echo ${TARGETPLATFORM} | cut -d/ -f2) && \
    export GOARM=${TARGETPLATFORM##*/} &&
    apk --no-cache add ca-certificates && \
    go mod download && \
    go build -a -tags netgo -ldflags '-w -extldflags "-static"' /build/cmd/radarr-exporter/.

FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=build /build/radarr-exporter /

ENTRYPOINT ["/radarr-exporter"]