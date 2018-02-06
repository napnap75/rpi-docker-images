all: build push

build:
	docker build -t napnap75/rpi-snips-handler:display .

push:
	docker login -u="${DOCKER_USERNAME}" -p="${DOCKER_PASSWORD}"
	docker push napnap75/rpi-snips-handler:display
