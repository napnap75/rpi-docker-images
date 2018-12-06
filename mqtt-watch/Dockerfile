FROM alpine:latest AS builder

ADD https://github.com/multiarch/qemu-user-static/releases/download/v2.6.0/qemu-arm-static.tar.gz /tmp
RUN tar xzf /tmp/qemu-arm-static.tar.gz -C /tmp 

FROM arm32v6/python:alpine3.7

COPY --from=builder /tmp/qemu-arm-static /usr/bin

RUN pip install paho-mqtt

ENV MQTT_HOST=localhost MQTT_PORT=1883

COPY mqtt-watch.py /usr/local/bin

CMD python /usr/local/bin/mqtt-watch.py
