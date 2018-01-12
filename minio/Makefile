build:
	if [ ! -f qemu-arm-static ]; then curl -L -o qemu-arm-static.tar.gz https://github.com/multiarch/qemu-user-static/releases/download/v2.6.0/qemu-arm-static.tar.gz && tar xzf qemu-arm-static.tar.gz ; fi
	if [ ! -d dockerscripts ]; then git clone https://github.com/minio/minio && mv minio/dockerscripts . && rm -fr minio ; fi
	docker build -t napnap75/rpi-minio:server -f Dockerfile.server .

push:
	if [ "${DOCKER_USERNAME}" == "" || "${DOCKER_PASSWORD}" ]; then docker login -u="${DOCKER_USERNAME}" -p="${DOCKER_PASSWORD}" ; fi
	docker push napnap75/rpi-minio:server
