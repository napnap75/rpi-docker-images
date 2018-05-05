FROM alpine:3.7 AS builder

RUN apk add --no-cache curl \
	&& QEMU_DOWNLOAD_URL=$(curl -s https://api.github.com/repos/multiarch/qemu-user-static/releases/latest | grep "browser_download_url" | grep "\/qemu-arm-static.tar.gz" | cut -d\" -f4) \
	&& curl -L -s -o /tmp/qemu-arm-static.tar.gz $QEMU_DOWNLOAD_URL \
	&& tar xzf /tmp/qemu-arm-static.tar.gz -C /tmp

FROM arm32v6/python:alpine3.7

COPY --from=builder /tmp/qemu-arm-static /usr/bin 

RUN pip install slacker

ADD slack-eraser.py /

ENTRYPOINT ["python", "/slack-eraser.py"]
