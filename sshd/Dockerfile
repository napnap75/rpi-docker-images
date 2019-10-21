FROM napnap75/rpi-alpine-s6:latest

RUN apk add --no-cache bash curl openssh-server rsync rrsync borgbackup

ADD etc /etc/

CMD /usr/sbin/sshd -D -e
