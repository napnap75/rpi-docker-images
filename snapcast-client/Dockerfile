FROM alpine:latest AS loader

RUN apk add --no-cache curl \
        && while [ "$QEMU_DOWNLOAD_URL" == "" ] ; do QEMU_DOWNLOAD_URL=$(curl -s https://api.github.com/repos/multiarch/qemu-user-static/releases/latest | grep "browser_download_url" | grep "\/qemu-arm-static.tar.gz" | cut -d\" -f4) ; done \
        && curl --retry 3 -L -s -o /tmp/qemu-arm-static.tar.gz $QEMU_DOWNLOAD_URL \
	&& tar xzf /tmp/qemu-arm-static.tar.gz -C /tmp

FROM arm32v6/alpine:latest

COPY --from=loader /tmp/qemu-arm-static /usr/bin 

RUN apk add --no-cache dbus avahi snapcast-client

COPY start.sh /usr/local/bin

CMD /usr/local/bin/start.sh
