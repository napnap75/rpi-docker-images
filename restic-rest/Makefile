all: build push

build:
	docker build -t napnap75/rpi-restic:rest-server .

login:
	if [ "${DOCKER_USERNAME}" != "" ]; then docker login -u="${DOCKER_USERNAME}" -p="${DOCKER_PASSWORD}" ; fi

push: login
	docker push napnap75/rpi-restic:rest-server
	version=$$(docker run napnap75/rpi-restic:rest-server --version | egrep -o "\((.+)\)" | egrep -o "[-.a-z0-9]+") ; \
		docker tag napnap75/rpi-restic:rest-server napnap75/rpi-restic:rest-server-$$version ; \
		docker push napnap75/rpi-restic:rest-server-$$version

