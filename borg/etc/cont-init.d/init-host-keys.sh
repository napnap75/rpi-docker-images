#!/bin/bash

if [ -f "/config/host_keys/ssh_host_rsa_key" ] ; then
	for keyfile in /config/host_keys/ssh_host_* ; do
		cp $keyfile /etc/ssh/
	done
elif [ ! -f "/etc/ssh/ssh_host_rsa_key" ] ; then
	for keytype in dsa ecdsa ed25519 rsa ; do
		ssh-keygen -q -f "/etc/ssh/ssh_host_${keytype}_key" -N '' -t $keytype
	done
fi
