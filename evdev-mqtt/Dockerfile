FROM alpine:latest AS builder

ADD https://github.com/multiarch/qemu-user-static/releases/download/v2.6.0/qemu-arm-static.tar.gz /tmp
RUN tar xzf /tmp/qemu-arm-static.tar.gz -C /tmp 

FROM arm32v7/debian:10-slim

COPY --from=builder /tmp/qemu-arm-static /usr/bin

RUN apt update \
	&& apt-get install --no-install-recommends -y python3-evdev python3-paho-mqtt \
	&& rm -rf /var/lib/apt/lists/* 

ENV MQTT_HOST=localhost MQTT_PORT=1883

COPY evdev-mqtt.py /usr/local/bin

CMD python3 /usr/local/bin/evdev-mqtt.py
