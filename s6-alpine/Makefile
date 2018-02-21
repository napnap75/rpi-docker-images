build:
	docker build -t napnap75/rpi-alpine-s6:3.7 .

push:
	docker login -u="${DOCKER_USERNAME}" -p="${DOCKER_PASSWORD}"
	docker push napnap75/rpi-alpine-s6:3.7
