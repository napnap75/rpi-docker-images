build:
	docker build -t napnap75/rpi-restic:latest .

push:
	if [ "${DOCKER_USERNAME}" != "" ]; then docker login -u="${DOCKER_USERNAME}" -p="${DOCKER_PASSWORD}" ; fi
	docker push napnap75/rpi-restic:latest
	version=$$(docker run --rm -it napnap75/rpi-restic:latest restic version | egrep -o "restic [.0-9]+ compiled" | egrep -o "[.0-9]+") ; \
		docker tag napnap75/rpi-restic:latest napnap75/rpi-restic:$$version ; \
		docker push napnap75/rpi-restic:$$version
