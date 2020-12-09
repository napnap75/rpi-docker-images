build:
	docker build -t napnap75/rpi-restic:rest-server .

push:
	if [ "${DOCKER_USERNAME}" != "" ]; then docker login -u="${DOCKER_USERNAME}" -p="${DOCKER_PASSWORD}" ; fi
	docker push napnap75/rpi-restic:rest-server
	version=$$(docker run napnap75/rpi-restic:rest-server --version | egrep -o "\((.+)\)" | egrep -o "[-.a-z0-9]+") ; \
		docker tag napnap75/rpi-restic:rest-server napnap75/rpi-restic:rest-server-$$version ; \
		docker push napnap75/rpi-restic:rest-server-$$version

