FROM alpine:latest AS builder


RUN apk add --no-cache curl \
	&& while [ "$QEMU_DOWNLOAD_URL" == "" ] ; do QEMU_DOWNLOAD_URL=$(curl -s https://api.github.com/repos/multiarch/qemu-user-static/releases/latest | grep "browser_download_url" | grep "\/qemu-arm-static.tar.gz" | cut -d\" -f4) ; done \
	&& curl --retry 3 -L -s -o /tmp/qemu-arm-static.tar.gz $QEMU_DOWNLOAD_URL \
	&& tar xzf /tmp/qemu-arm-static.tar.gz -C /tmp

FROM arm32v7/python:slim AS python-builder

COPY --from=builder /tmp/qemu-arm-static /usr/bin

RUN apt update \
	&& apt install -y git gcc \
	&& pip install RPi.GPIO spidev gpiozero paho-mqtt pytoml \
	&& git clone https://github.com/Psychokiller1888/snipsLedControl \
	&& mkdir /snipsLedControl/logs

FROM arm32v7/python:slim

COPY --from=builder /tmp/qemu-arm-static /usr/bin
COPY --from=python-builder /root/.cache /root/.cache
COPY --from=python-builder /snipsLedControl /snipsLedControl

RUN pip install RPi.GPIO spidev gpiozero paho-mqtt pytoml

WORKDIR /snipsLedControl

CMD python /snipsLedControl/main.py
