FROM napnap75/rpi-alpine-s6:latest

RUN apk add --no-cache dcron py2-pip gstreamer gst-plugins-base gst-plugins-good gst-plugins-bad py2-gst  \
	&& pip install mopidy Mopidy-Local-SQLite

COPY etc /etc/
