FROM napnap75/rpi-alpine-base:edge

RUN cat /etc/apk/repositories | grep main | sed "s/main/testing/g" >> /etc/apk/repositories \
	&& apk update \
	&& apk add --no-cache shairport-sync snapcast-server

COPY start.sh /usr/local/bin

ENV AIRPLAY_NAME=snapcast BUFFER_SIZE=1000
CMD /usr/local/bin/start.sh
