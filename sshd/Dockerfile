FROM napnap75/rpi-alpine-s6:3.7

RUN apk add --no-cache bash curl openssh-server rsync rrsync borgbackup

ADD etc /etc/

CMD /usr/sbin/sshd -D -e
