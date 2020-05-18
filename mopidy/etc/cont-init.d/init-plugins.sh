#!/usr/bin/with-contenv sh

if [ "$MOPIDY_PLUGINS" != "" ]; then
	echo "Installing plugins $MOPIDY_PLUGINS"
	pip3 install $MOPIDY_PLUGINS
fi
