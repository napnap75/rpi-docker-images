FROM napnap75/rpi-alpine-s6:latest

RUN apk add --no-cache mopidy \
	&& echo "30 7 * * * mopidy local scan" | crontab -

COPY etc /etc/
