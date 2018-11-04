#!/bin/bash
set -e

# Build the templates
if [ -d "/usr/share/snips/assistant/snippets" ] ; then
	/usr/bin/snips-template
fi

# Load the skills from git
cd /var/lib/snips/skills
for repository in $(cat /usr/share/snips/assistant/Snipsfile.yaml | grep "url:" | sed "s/\ *url:\ *//g") ; do
	echo "Loading repository $repository from Git"
	dirname=$(echo "$repository" | sed "s#.*/##g")
	if [ -d "$dirname" ] ; then
		cd "$dirname"
		git pull
		cd ..
	else
		git clone $repository
	fi
done

# Install the dependencies
for dirname in $(ls * -d) ; do
	cd "$dirname"
	if [ -f "requirements.txt" ] ; then
		pip install -r requirements.txt
	fi
done

# Run the server
/usr/bin/snips-skill-server
