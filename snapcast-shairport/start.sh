#!/bin/sh
rm -rf /var/run
mkdir -p /var/run/dbus
dbus-uuidgen --ensure
dbus-daemon --system
avahi-daemon --daemonize --no-chroot
if [ -p /fifo/snapfifo ] ; then
	/usr/bin/snapserver -b $BUFFER_SIZE -s "airplay:///shairport-sync?name=Airplay&devicename=$AIRPLAY_NAME" -s "pipe:///fifo/snapfifo?name=$STREAM_NAME"
else
	/usr/bin/snapserver -b $BUFFER_SIZE -s "airplay:///shairport-sync?name=Airplay&devicename=$AIRPLAY_NAME"
fi
