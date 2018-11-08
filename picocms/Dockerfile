FROM alpine:latest AS builder

RUN apk add --no-cache curl \
        && while [ "$QEMU_DOWNLOAD_URL" == "" ] ; do QEMU_DOWNLOAD_URL=$(curl -s https://api.github.com/repos/multiarch/qemu-user-static/releases/latest | grep "browser_download_url" | grep "\/qemu-arm-static.tar.gz" | cut -d\" -f4) ; done \
        && curl --retry 3 -L -s -o /tmp/qemu-arm-static.tar.gz $QEMU_DOWNLOAD_URL \
        && tar xzf /tmp/qemu-arm-static.tar.gz -C /tmp

FROM arm32v7/php:apache

COPY --from=builder /tmp/qemu-arm-static /usr/bin

RUN a2enmod rewrite \
	&& DOWNLOAD_URL=$(curl -s --retry 3 https://api.github.com/repos/picocms/Pico/releases/latest | grep "browser_download_url" | cut -d\" -f4) \
	&& curl -Ls --retry 3 $DOWNLOAD_URL | tar zx \
	&& mkdir -p plugins/pico_edit \
	&& DOWNLOAD_URL=$(curl -s --retry 3 https://api.github.com/repos/blocknotes/pico_edit/releases/latest | grep "tarball_url" | cut -d\" -f4) \
	&& curl -Ls --retry 3 $DOWNLOAD_URL | tar -zx -C plugins/pico_edit --strip-components 1 \
	&& chown -R www-data:www-data .

