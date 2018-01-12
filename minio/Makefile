all: build push clean

build: build-server build-client

build-server:
	if [ ! -f qemu-arm-static ]; then curl -L -o qemu-arm-static.tar.gz https://github.com/multiarch/qemu-user-static/releases/download/v2.6.0/qemu-arm-static.tar.gz && tar xzf qemu-arm-static.tar.gz ; fi
	docker build -t napnap75/rpi-minio:server -f Dockerfile.server .

build-client:
	if [ ! -f qemu-arm-static ]; then curl -L -o qemu-arm-static.tar.gz https://github.com/multiarch/qemu-user-static/releases/download/v2.6.0/qemu-arm-static.tar.gz && tar xzf qemu-arm-static.tar.gz ; fi
	docker build -t napnap75/rpi-minio:client -f Dockerfile.client .

push: push-server push-client

push-server:
	if [ "${DOCKER_USERNAME}" != "" ] && [ "${DOCKER_PASSWORD}" != "" ]; then docker login -u="${DOCKER_USERNAME}" -p="${DOCKER_PASSWORD}" ; fi
	docker push napnap75/rpi-minio:server

push-client:
	if [ "${DOCKER_USERNAME}" != "" ] && [ "${DOCKER_PASSWORD}" != "" ]; then docker login -u="${DOCKER_USERNAME}" -p="${DOCKER_PASSWORD}" ; fi
	docker push napnap75/rpi-minio:client

clean:
	rm -fr qemu-arm-static*
	docker image prune -f
