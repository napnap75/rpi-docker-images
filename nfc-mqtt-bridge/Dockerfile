FROM alpine:latest AS loader

RUN apk add --no-cache curl \
        && while [ "$QEMU_DOWNLOAD_URL" == "" ] ; do QEMU_DOWNLOAD_URL=$(curl -s https://api.github.com/repos/multiarch/qemu-user-static/releases/latest | grep "browser_download_url" | grep "\/qemu-arm-static.tar.gz" | cut -d\" -f4) ; done \
        && curl --retry 3 -L -s -o /tmp/qemu-arm-static.tar.gz $QEMU_DOWNLOAD_URL \
	&& tar xzf /tmp/qemu-arm-static.tar.gz -C /tmp

FROM arm32v6/python:alpine AS builder

COPY --from=loader /tmp/qemu-arm-static /usr/bin 

ENV PATH="/opt/venv/bin:$PATH"

RUN apk add --no-cache alpine-sdk \
	&& python -m venv /opt/venv \
	&& pip install adafruit-circuitpython-pn532 paho-mqtt PyYAML RPi.GPIO

FROM arm32v6/python:alpine

COPY --from=loader /tmp/qemu-arm-static /usr/bin 

COPY --from=builder /opt/venv /opt/venv

ENV PATH="/opt/venv/bin:$PATH"

ADD nfc-mqtt-bridge.py /

CMD ["python3", "/nfc-mqtt-bridge.py"]
