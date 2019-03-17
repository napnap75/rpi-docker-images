FROM linuxserver/daapd:arm32v6-latest AS builder

RUN apk update \
	&& apk add git build-base autoconf automake libtool alsa-lib-dev libdaemon-dev popt-dev libressl-dev soxr-dev avahi-dev libconfig-dev curl asio-dev flac-dev libvorbis-dev \
	&& mkdir /root/shairport-sync \
	&& cd /root/shairport-sync \
	&& while [ "$DOWNLOAD_URL" == "" ] ; do DOWNLOAD_URL=$(curl -s https://api.github.com/repos/mikebrady/shairport-sync/releases/latest | grep "tarball_url" | cut -d\" -f4) ; done \
	&& curl --retry 3 -L -s -o /tmp/shairport-sync.tar.gz $DOWNLOAD_URL \
	&& tar xzf /tmp/shairport-sync.tar.gz --strip-components=1 \
	&& autoreconf -i -f \
	&& ./configure --with-stdout --with-pipe --with-avahi --with-ssl=openssl --with-soxr --with-metadata \
	&& make

FROM linuxserver/daapd:arm32v6-latest

COPY --from=builder /root/shairport-sync/shairport-sync /usr/local/bin/

RUN apk add --no-cache flac libvorbis alsa-lib libdaemon popt libressl soxr libconfig \
	&& rm /etc/avahi/services/*

COPY root /
