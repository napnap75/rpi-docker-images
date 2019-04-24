FROM napnap75/rpi-alpine-base:latest

RUN apk add --no-cache bash samba-common-tools samba

COPY run.sh /run.sh
ENTRYPOINT ["/run.sh"]

CMD ["-h"]
