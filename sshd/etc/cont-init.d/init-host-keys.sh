#!/bin/bash

if [ ! -f "/config/host_keys/ssh_host_rsa_key" ] ; then
	mkdir -p /config/host_keys
	ssh-keygen -t ed25519 -f /config/host_keys/ssh_host_ed25519_key -N "" < /dev/null
	ssh-keygen -t rsa -b 4096 -f /config/host_keys/ssh_host_rsa_key -N "" < /dev/null
fi
