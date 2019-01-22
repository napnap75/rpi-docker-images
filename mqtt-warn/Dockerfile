FROM alpine:latest AS builder

ADD https://github.com/multiarch/qemu-user-static/releases/download/v2.6.0/qemu-arm-static.tar.gz /tmp
RUN tar xzf /tmp/qemu-arm-static.tar.gz -C /tmp \
	&& apk add --no-cache git \
	&& cd /tmp \
	&& git clone https://github.com/jpmens/mqttwarn

#FROM arm32v6/python:alpine3.7
FROM python:2-alpine3.8

COPY --from=builder /tmp/qemu-arm-static /usr/bin
COPY --from=builder /tmp/mqttwarn /mqttwarn

RUN pip install paho-mqtt requests jinja2 slacker

WORKDIR /mqttwarn

ENV MQTTWARNINI="/mqttwarn/conf/mqttwarn.ini"

CMD python mqttwarn.py
