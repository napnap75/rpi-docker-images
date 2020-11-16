FROM alpine:latest AS builder

RUN apk add --no-cache curl \
  && while [ "$QEMU_DOWNLOAD_URL" == "" ] ; do QEMU_DOWNLOAD_URL=$(curl -s https://api.github.com/repos/multiarch/qemu-user-static/releases/latest | grep "browser_download_url" | grep "\/qemu-arm-static.tar.gz" | cut -d\" -f4) ; done \
  && curl --retry 3 -L -s -o /tmp/qemu-arm-static.tar.gz $QEMU_DOWNLOAD_URL \
  && tar xzf /tmp/qemu-arm-static.tar.gz \
	&& while [ "$DOWNLOAD_URL" == "" ] ; do DOWNLOAD_URL=$(curl -s https://api.github.com/repos/42wim/matterbridge/releases/latest | grep "browser_download_url" | grep "linux-armv6" | cut -d\" -f4) ; done \
	&& curl --retry 3 -L -s -o matterbridge ${DOWNLOAD_URL} \
	&& chmod +x matterbridge

FROM arm32v6/alpine:latest

COPY --from=builder qemu-arm-static matterbridge /usr/bin/

RUN apk --no-cache add ca-certificates mailcap

ENTRYPOINT ["/usr/bin/matterbridge", "-conf", "/etc/matterbridge/matterbridge.toml"]
