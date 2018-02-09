Docker image to automatically update Gandi DNS with the current IP adress

# Status
[![Github link](https://assets-cdn.github.com/favicon.ico)](https://github.com/napnap75/rpi-docker-images/) [![Docker hub link](https://www.docker.com/favicon.ico)](https://hub.docker.com/r/napnap75/rpi-gandi/)


# Content
This image is based [my own Alpine Linux base image](https://hub.docker.com/r/napnap75/rpi-alpine-base/).

This image contains :
- the python Gandi CLI for the v4 version
- curl and jq to access Gandi REST API for the v5 version

# Usage
Use this image to update a DNS record to the current IP of the host: `docker run -e GANDI_API_KEY="YOUR GANDI API KEY" -e GANDI_HOST="THE HOST NAME IN THE GANDI DOMAIN" -e GANDI_DOMAIN="YOUR GANDI DOMAIN" --name gandi napnap75/rpi-gandi:latest`
Use the following images :
- `napnap75/rpi-gandi:v4` if you use the v4 version of Gandi
- `napnap75/rpi-gandi:v5` or `napnap75/rpi-gandi:latest` if you use the v5 version of Gandi

Every 5 minutes, the image will automatically check the current IP address of the host and, if necessary, update the DNS.
