FROM napnap75/rpi-alpine-base:latest AS builder

RUN apk update \
	&& apk add git build-base autoconf automake libtool alsa-lib-dev libdaemon-dev popt-dev libressl-dev soxr-dev avahi-dev libconfig-dev curl asio-dev flac-dev libvorbis-dev mosquitto-dev \
	&& mkdir /root/shairport-sync \
	&& cd /root/shairport-sync \
	&& while [ "$DOWNLOAD_URL" == "" ] ; do DOWNLOAD_URL=$(curl -s https://api.github.com/repos/mikebrady/shairport-sync/releases/latest | grep "tarball_url" | cut -d\" -f4) ; done \
	&& curl --retry 3 -L -s -o /tmp/shairport-sync.tar.gz $DOWNLOAD_URL \
	&& tar xzf /tmp/shairport-sync.tar.gz --strip-components=1 \
	&& autoreconf -i -f \
	&& ./configure --with-stdout --with-pipe --with-avahi --with-ssl=openssl --with-soxr --with-metadata --with-mqtt-client  \
	&& make

FROM napnap75/rpi-alpine-base:latest

RUN apk add --no-cache snapcast-server avahi-libs flac libvorbis dbus alsa-lib libdaemon popt libressl soxr avahi libconfig

COPY --from=builder /root/shairport-sync/shairport-sync /usr/local/bin/
COPY start.sh /usr/local/bin

ENV AIRPLAY_NAME=snapcast BUFFER_SIZE=1000
CMD /usr/local/bin/start.sh
