FROM alpine:3.7 AS builder

RUN apk add --no-cache curl \
  && while [ "$QEMU_DOWNLOAD_URL" == "" ] ; do QEMU_DOWNLOAD_URL=$(curl -s https://api.github.com/repos/multiarch/qemu-user-static/releases/latest | grep "browser_download_url" | grep "\/qemu-arm-static.tar.gz" | cut -d\" -f4) ; done \
  && curl --retry 3 -L -s -o /tmp/qemu-arm-static.tar.gz $QEMU_DOWNLOAD_URL \
  && tar xzf /tmp/qemu-arm-static.tar.gz -C /tmp \
  && while [ "$GITEA_DOWNLOAD_URL" == "" ] ; do GITEA_DOWNLOAD_URL=$(curl -s https://api.github.com/repos/go-gitea/gitea/releases/latest | grep "browser_download_url" | grep "arm-6.xz\"" | cut -d\" -f4) ; done \
  && curl --retry 3 -L -s -o /tmp/gitea-linux-arm-6.xz $GITEA_DOWNLOAD_URL \
  && unxz -c /tmp/gitea-linux-arm-6.xz > /tmp/gitea
  
FROM arm32v6/alpine:3.7
COPY --from=builder /tmp/qemu-arm-static /usr/bin

EXPOSE 22 3000

RUN apk --no-cache add bash ca-certificates curl gettext git linux-pam openssh s6 sqlite su-exec tzdata

RUN addgroup -S -g 1000 git \
    && adduser -S -H -D -h /data/git -s /bin/bash -u 1000 -G git git \
    && echo "git:$(dd if=/dev/urandom bs=24 count=1 status=none | base64)" | chpasswd

ENV USER git
ENV GITEA_CUSTOM /data/gitea

VOLUME ["/data"]

ENTRYPOINT ["/usr/bin/entrypoint"]
CMD ["/bin/s6-svscan", "/etc/s6"]

COPY docker /
COPY --from=builder /tmp/gitea /app/gitea/gitea
RUN ln -s /app/gitea/gitea /usr/local/bin/gitea
