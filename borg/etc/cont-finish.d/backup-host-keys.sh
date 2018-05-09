#!/bin/bash

mkdir -p /config/host_keys
for keyfile in /etc/ssh/ssh_host_* ; do
	cp $keyfile /config/host_keys/
done
