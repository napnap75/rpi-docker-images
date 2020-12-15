build:
	docker build -t napnap75/rpi-gandi:v5 -f Dockerfile.v5 .

push:
	if [ "${DOCKER_USERNAME}" != "" ]; then docker login -u="${DOCKER_USERNAME}" -p="${DOCKER_PASSWORD}" ; fi
	docker push napnap75/rpi-gandi:v5
	docker tag napnap75/rpi-gandi:v5 napnap75/rpi-gandi:latest
	docker push napnap75/rpi-gandi:latest
