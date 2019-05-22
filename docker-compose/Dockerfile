FROM alpine:latest AS builder

RUN apk add --no-cache curl \
        && while [ "$QEMU_DOWNLOAD_URL" == "" ] ; do QEMU_DOWNLOAD_URL=$(curl -s https://api.github.com/repos/multiarch/qemu-user-static/releases/latest | grep "browser_download_url" | grep "\/qemu-arm-static.tar.gz" | cut -d\" -f4) ; done \
        && curl --retry 3 -L -s -o /tmp/qemu-arm-static.tar.gz $QEMU_DOWNLOAD_URL \
        && tar xzf /tmp/qemu-arm-static.tar.gz -C /tmp

FROM arm32v6/alpine:latest

COPY --from=builder /tmp/qemu-arm-static /usr/bin

RUN apk add --no-cache py2-pip py2-cffi \
	&& pip install docker-compose

ENTRYPOINT ["/usr/bin/docker-compose"]
