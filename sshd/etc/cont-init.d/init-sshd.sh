#! /bin/bash

cp /etc/ssh/sshd_config.default /etc/ssh/sshd_config

function init_user {
	# First create the user
	options=
	if [[ "$3" != "" ]]; then
		options+="-u $3 "
	fi
	if [[ "$4" != "" ]]; then
		grep ":$4:" /etc/group || addgroup -g $4 "group-$4"
		options+="-G `getent group $4 | sed 's/:.*//'` "
	fi
	if [[ "$5" != "" ]]; then
		options+="-h $5 "
	else
		options+="-h /home/$1 "
	fi
	if [[ "$2" = "ssh" || "$2" = "borg" || "$2" = "rsync" ]]; then
		adduser -D $options -s /bin/bash $1
	else
		adduser -D $options -s /bin/false $1
	fi
	passwd -u $1

	# Adjust the keys permissions
	chown $1 /config/users_keys/$1
	chmod 400 /config/users_keys/$1

	# Update sshd-config
	sed -i "/^AllowUsers/ s/$/ $1/" /etc/ssh/sshd_config
	if [[ $2 = "sftp" ]]; then
		echo "Match User $1" >> /etc/ssh/sshd_config
		echo "  ForceCommand internal-sftp" >> /etc/ssh/sshd_config
		if [[ $6 != "" ]]; then
			echo "  ChrootDirectory $6" >> /etc/ssh/sshd_config
			chown root:root $6
		fi
	elif [[ $2 = "borg" ]]; then
		echo "Match User $1" >> /etc/ssh/sshd_config
		if [[ $6 != "" ]]; then
			echo "  ForceCommand borg serve --restrict-to-path $6" >> /etc/ssh/sshd_config
		else
			echo "  ForceCommand borg serve" >> /etc/ssh/sshd_config
		fi
	elif [[ $2 = "rsync" ]]; then
		echo "Match User $1" >> /etc/ssh/sshd_config
		if [[ $6 != "" ]]; then
			echo "  ForceCommand rrsync $6" >> /etc/ssh/sshd_config
		else
			echo "  ForceCommand rrsync ." >> /etc/ssh/sshd_config
		fi
	fi
}

while read line; do
	if [[ "$line" =~ ^\[ ]]; then
		if [[ "$user" != "" ]]; then
			init_user "$user" "$type" "$uid" "$gid" "$home" "$chroot"
		fi

		name=${line#*\[}
		name=${name%%\]}

		user=$name
		type=
		uid=
		gid=
		home=
		chroot=
	elif [[ "$line" =~ ^[^#]*= ]]; then
		name=${line%% =*}
		value=${line#*= }
		if [[ $name = "Type" ]]; then
			type=$value
		elif [[ $name = "UID" ]]; then
			uid=$value
		elif [[ $name = "GID" ]]; then
			gid=$value
		elif [[ $name = "Home" ]]; then
			home=$value
		elif [[ $name = "Chroot" ]]; then
			chroot=$value
		fi
	fi
done < /config/config.ini
init_user "$user" "$type" "$uid" "$gid" "$home" "$chroot"
