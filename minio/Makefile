build:
	curl -L -o qemu-arm-static.tar.gz https://github.com/multiarch/qemu-user-static/releases/download/v2.6.0/qemu-arm-static.tar.gz && tar xzf qemu-arm-static.tar.gz
	git clone https://github.com/minio/minio && mv minio/dockerscripts . && rm -fr minio
	docker build -t napnap75/rpi-minio:server -f Dockerfile.server .

push:
	docker login -u="${DOCKER_USERNAME}" -p="${DOCKER_PASSWORD}"
	docker push napnap75/rpi-minio:server
