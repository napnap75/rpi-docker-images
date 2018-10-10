all: build push

build:
	docker build -t napnap75/rpi-snips-button-handler:latest .

push:
	docker login -u="${DOCKER_USERNAME}" -p="${DOCKER_PASSWORD}"
	docker push napnap75/rpi-snips-button-handler:latest

test:
	docker run --rm -it --device /dev/mem --privileged -e "MQTT_HOST=${MQTT_HOST}" napnap75/rpi-snips-button-handler:latest
