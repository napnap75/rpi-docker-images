all: build push

build:
	docker build -t napnap75/rpi-mopidy-snapcast-controller:latest .

push:
	if [ "${DOCKER_USERNAME}" != "" ]; then docker login -u="${DOCKER_USERNAME}" -p="${DOCKER_PASSWORD}" ; fi
	docker push napnap75/rpi-mopidy-snapcast-controller:latest
