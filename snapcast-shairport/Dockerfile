FROM napnap75/rpi-alpine-base:3.8 AS builder

RUN apk update \
	&& apk add git build-base autoconf automake libtool alsa-lib-dev libdaemon-dev popt-dev libressl-dev soxr-dev avahi-dev libconfig-dev curl asio-dev flac-dev libvorbis-dev \
	&& mkdir /root/shairport-sync \
	&& cd /root/shairport-sync \
	&& curl -L -o ./shairport-sync.tar.gz https://github.com/mikebrady/shairport-sync/archive/3.2.2.tar.gz \
	&& tar -zxvf shairport-sync.tar.gz --strip-components=1 \
	&& autoreconf -i -f \
	&& ./configure --with-stdout --with-pipe --with-avahi --with-ssl=openssl --with-soxr --with-metadata \
	&& make \
	&& cd /root \
	&& git clone --recursive https://github.com/badaix/snapcast.git \
	&& cd snapcast \
	&& make server


FROM napnap75/rpi-alpine-base:3.8

RUN apk add --no-cache avahi-libs flac libvorbis dbus alsa-lib libdaemon popt libressl soxr avahi libconfig

COPY --from=builder /root/shairport-sync/shairport-sync /root/snapcast/server/snapserver /usr/local/bin/
COPY start.sh /usr/local/bin

ENV AIRPLAY_NAME=snapcast BUFFER_SIZE=1000
CMD /usr/local/bin/start.sh
