build: build-borg build-borgmatic build-server

build-borg:
	docker build -t napnap75/rpi-borg:borg -f Dockerfile.borg .

build-borgmatic:
	docker build -t napnap75/rpi-borg:borgmatic -f Dockerfile.borgmatic .

build-server:
	docker build -t napnap75/rpi-borg:server -f Dockerfile.server .

push:
	push-borg push-borgmatic push-server

push-init:
	if [ "${DOCKER_USERNAME}" != "" ]; then docker login -u="${DOCKER_USERNAME}" -p="${DOCKER_PASSWORD}" ; fi

push-borg: push-init
	docker push napnap75/rpi-borg:borg

push-borgmatic: push-init
	docker push napnap75/rpi-borg:borgmatic

push-server: push-init
	docker push napnap75/rpi-borg:server
