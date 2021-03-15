FROM golang:alpine AS builder
RUN apk add --no-cache curl \
        && while [ "$QEMU_DOWNLOAD_URL" == "" ] ; do QEMU_DOWNLOAD_URL=$(curl -s https://api.github.com/repos/multiarch/qemu-user-static/releases/latest | grep "browser_download_url" | grep "\/qemu-arm-static.tar.gz" | cut -d\" -f4) ; done \
        && curl --retry 3 -L -s -o /tmp/qemu-arm-static.tar.gz $QEMU_DOWNLOAD_URL \
        && tar xzf /tmp/qemu-arm-static.tar.gz
WORKDIR $GOPATH/src/napnap75/piwigo-souvenirs/
COPY piwigo-souvenirs.go .
RUN apk add --no-cache git gcc musl-dev \
	&& go mod init github.com/napnap75/rpi-docker-files/piwigo-souvenirs \
	&& go get -d -v \
	&& GOOS=linux GOARCH=arm GOARM=6 go build -ldflags="-w -s" -o /go/bin/piwigo-souvenirs

FROM arm32v6/alpine:latest
COPY --from=builder /go/qemu-arm-static /go/bin/piwigo-souvenirs /usr/bin/
VOLUME /tmp/messages
ENTRYPOINT ["/usr/bin/piwigo-souvenirs"]
