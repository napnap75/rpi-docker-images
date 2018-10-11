FROM alpine:latest AS builder

ADD https://github.com/multiarch/qemu-user-static/releases/download/v2.6.0/qemu-arm-static.tar.gz /tmp
RUN tar xzf /tmp/qemu-arm-static.tar.gz -C /tmp

FROM arm32v6/alpine:latest

COPY --from=builder /tmp/qemu-arm-static /usr/bin

RUN apk add --no-cache py3-paho-mqtt py3-rpigpio

COPY snips-button-handler.py /usr/local/bin

ENV GPIO_BUTTON=23 GPIO_LED=25 MQTT_HOST=localhost MQTT_PORT=1883

CMD python3 /usr/local/bin/snips-button-handler.py
