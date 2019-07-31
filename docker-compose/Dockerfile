FROM alpine:latest AS builder

RUN apk add --no-cache curl \
#        && while [ "$QEMU_DOWNLOAD_URL" == "" ] ; do QEMU_DOWNLOAD_URL=$(curl -s https://api.github.com/repos/multiarch/qemu-user-static/releases/latest | grep "browser_download_url" | grep "\/qemu-arm-static.tar.gz" | cut -d\" -f4) ; done \
        && QEMU_DOWNLOAD_URL="https://github.com/multiarch/qemu-user-static/releases/download/v3.1.0-3/qemu-arm-static.tar.gz" \
	&& curl --retry 3 -L -s -o /tmp/qemu-arm-static.tar.gz $QEMU_DOWNLOAD_URL \
        && tar xzf /tmp/qemu-arm-static.tar.gz -C /tmp

FROM arm32v6/alpine:latest

COPY --from=builder /tmp/qemu-arm-static /usr/bin

RUN apk add --no-cache python2 \
	&& apk add --no-cache --virtual .build-dependencies py2-pip gcc make python2-dev musl-dev libffi-dev libressl-dev \
	&& pip install --no-cache-dir docker-compose \
	&& apk del .build-dependencies

ENTRYPOINT ["/usr/bin/docker-compose"]
