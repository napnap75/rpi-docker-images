all: build push

build:
	docker build -t napnap75/rpi-mqtt-watch:latest .

push:
	docker login -u="${DOCKER_USERNAME}" -p="${DOCKER_PASSWORD}"
	docker push napnap75/rpi-mqtt-watch:latest
