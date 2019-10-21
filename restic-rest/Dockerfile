FROM golang:alpine AS builder

RUN apk add --no-cache curl git \
        && while [ "$QEMU_DOWNLOAD_URL" == "" ] ; do QEMU_DOWNLOAD_URL=$(curl -s https://api.github.com/repos/multiarch/qemu-user-static/releases/latest | grep "browser_download_url" | grep "\/qemu-arm-static.tar.gz" | cut -d\" -f4) ; done \
        && curl --retry 3 -L -s -o /tmp/qemu-arm-static.tar.gz $QEMU_DOWNLOAD_URL \
        && tar xzf /tmp/qemu-arm-static.tar.gz \
	&& mkdir -p /go/src/github.com/restic \
	&& cd /go/src/github.com/restic \
	&& git clone https://github.com/restic/rest-server \
	&& cd rest-server \
	&& go run build.go -v --goos linux --goarch arm

FROM arm32v6/alpine:latest

COPY --from=builder /go/qemu-arm-static /go/src/github.com/restic/rest-server/rest-server /usr/bin/
RUN adduser -DH rest-server
USER rest-server
ENTRYPOINT ["/usr/bin/rest-server"]
