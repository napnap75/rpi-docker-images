build:
	docker build -t napnap75/rpi-docker-compose:latest .

push:
	if [ "${DOCKER_USERNAME}" != "" ]; then docker login -u="${DOCKER_USERNAME}" -p="${DOCKER_PASSWORD}" ; fi
	docker push napnap75/rpi-docker-compose:latest
	version=$$(docker run --rm -it napnap75/rpi-docker-compose version | grep docker-compose | egrep -o "[0-9]+\.[0-9]+\.[0-9]+") ; \
	docker tag napnap75/rpi-docker-compose:latest napnap75/rpi-docker-compose:$$version ; \
	docker push napnap75/rpi-docker-compose:$$version
