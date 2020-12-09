build:
	docker build -t napnap75/rpi-piwigo-souvenirs .

push:
	if [ "${DOCKER_USERNAME}" != "" ]; then docker login -u="${DOCKER_USERNAME}" -p="${DOCKER_PASSWORD}" ; fi
	docker push napnap75/rpi-piwigo-souvenirs
