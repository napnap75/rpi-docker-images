#!/bin/bash

set -e

# When used with S3 and docker secrets, get the credentials from files
if [[ "$AWS_ACCESS_KEY_ID" = /* && -f "$AWS_ACCESS_KEY_ID" ]] ; then
	AWS_ACCESS_KEY_ID=$(cat $AWS_ACCESS_KEY_ID)
fi
if [[ "$AWS_SECRET_ACCESS_KEY" = /* && -f "$AWS_SECRET_ACCESS_KEY" ]] ; then
	AWS_SECRET_ACCESS_KEY=$(cat $AWS_SECRET_ACCESS_KEY)
fi

# When used with SFTP set the SSH configuration file
if [[ "$RESTIC_REPOSITORY" = sftp:* ]] ; then
	# Copy the key and make it readable only by the current user to meet SSH security requirements
	cp $SFTP_KEY /tmp/foreign_host_key
	chmod 400 /tmp/foreign_host_key
	SFTP_KEY=/tmp/foreign_host_key

	# Initialize the SSH config file with the values provided in the environment
	mkdir -p /root/.ssh
	echo "Host $SFTP_HOST" > /root/.ssh/config
	if [[ "$SFTP_PORT" != "" ]] ; then echo "Port $SFTP_PORT" >> /root/.ssh/config ; fi
	echo "IdentityFile $SFTP_KEY" >> /root/.ssh/config
	echo "StrictHostKeyChecking no" >> /root/.ssh/config
fi

"$@"
