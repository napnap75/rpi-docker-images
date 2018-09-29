build:
	docker build -f Dockerfile-3.7 -t napnap75/rpi-alpine-s6:3.7 .
	docker build -f Dockerfile-3.8 -t napnap75/rpi-alpine-s6:3.8 .
	docker build -f Dockerfile-edge -t napnap75/rpi-alpine-s6:edge .

push:
	if [ "${DOCKER_USERNAME}" != "" ]; then docker login -u="${DOCKER_USERNAME}" -p="${DOCKER_PASSWORD}" ; fi
	docker push napnap75/rpi-alpine-s6:3.7
	docker push napnap75/rpi-alpine-s6:3.8
	docker tag napnap75/rpi-alpine-s6:3.8 napnap75/rpi-alpine-s6:latest
	docker push napnap75/rpi-alpine-s6:latest
	docker push napnap75/rpi-alpine-s6:edge
