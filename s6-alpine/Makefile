build:
	docker build -f Dockerfile-latest -t napnap75/rpi-alpine-s6:latest .
	docker build -f Dockerfile-edge -t napnap75/rpi-alpine-s6:edge .

push:
	if [ "${DOCKER_USERNAME}" != "" ]; then docker login -u="${DOCKER_USERNAME}" -p="${DOCKER_PASSWORD}" ; fi
	docker push napnap75/rpi-alpine-s6:latest
	VERSION=$(shell docker run --rm alpine cat /etc/os-release | grep PRETTY_NAME | egrep -o "[0-9]+\.[0-9]+") ; \
		docker tag napnap75/rpi-alpine-s6:latest napnap75/rpi-alpine-s6:$$VERSION ; \
		docker push napnap75/rpi-alpine-s6:$$VERSION
	docker push napnap75/rpi-alpine-s6:edge
