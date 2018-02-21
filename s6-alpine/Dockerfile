FROM alpine:3.7 AS builder

ADD https://github.com/just-containers/s6-overlay/releases/download/v1.21.2.2/s6-overlay-armhf.tar.gz /tmp
RUN mkdir /tmp/s6-overlay && tar xzf /tmp/s6-overlay-armhf.tar.gz -C /tmp/s6-overlay

ADD https://github.com/multiarch/qemu-user-static/releases/download/v2.6.0/qemu-arm-static.tar.gz /tmp
RUN tar xzf /tmp/qemu-arm-static.tar.gz -C /tmp

FROM arm32v6/alpine:3.7

COPY --from=builder /tmp/s6-overlay /
COPY --from=builder /tmp/qemu-arm-static /usr/bin 
ENTRYPOINT ["/init"]
