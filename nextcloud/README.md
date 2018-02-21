Nextcloud with Samba client installed for the Raspberry Pi.

# Status
[![Github link](https://assets-cdn.github.com/favicon.ico)](https://github.com/napnap75/rpi-docker-images/) [![Docker hub link](https://www.docker.com/favicon.ico)](https://hub.docker.com/r/napnap75/rpi-nextcloud-smb/)


# Content
This image is based on [the official Nextcloud image](https://hub.docker.com/r/arm32v7/nextcloud/).

The changes from the original images are :
- added QEmu in the image to allow it to be build on x86 systems (Travis CI)
- installed the smbclient package
