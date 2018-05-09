#!/bin/bash

if [ -f /home/borg/.ssh/authorized_keys ] ; then
	rm /home/borg/.ssh/authorized_keys
fi

for keyfile in $(find "/config/clients" -type f); do
	client_name=$(basename $keyfile)

	echo -n "command=\"cd /backup/${client_name}; borg serve --restrict-to-path /backup/${client_name}\" " >> /home/borg/.ssh/authorized_keys
	cat $keyfile >> /home/borg/.ssh/authorized_keys

	mkdir -p /backup/$client_name
	chown -R borg:borg /backup/$client_name
done
