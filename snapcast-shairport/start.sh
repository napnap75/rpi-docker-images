#!/bin/sh
rm -rf /var/run
mkdir -p /var/run/dbus
dbus-uuidgen --ensure
dbus-daemon --system
avahi-daemon --daemonize --no-chroot
/usr/local/bin/snapserver -b $BUFFER_SIZE -s "airplay:///shairport-sync?name=Airplay&devicename=$AIRPLAY_NAME"

