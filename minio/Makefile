all: build push clean

build: build-server build-client

build-server:
	if [ ! -f qemu-arm-static ]; then curl -L -o qemu-arm-static.tar.gz https://github.com/multiarch/qemu-user-static/releases/download/v2.6.0/qemu-arm-static.tar.gz && tar xzf qemu-arm-static.tar.gz ; fi
	docker build -t napnap75/rpi-minio:server -f Dockerfile.server .

build-client:
	if [ ! -f qemu-arm-static ]; then curl -L -o qemu-arm-static.tar.gz https://github.com/multiarch/qemu-user-static/releases/download/v2.6.0/qemu-arm-static.tar.gz && tar xzf qemu-arm-static.tar.gz ; fi
	docker build -t napnap75/rpi-minio:client -f Dockerfile.client .

test:
	docker run --rm -e "MINIO_ACCESS_KEY=access_key" -e "MINIO_SECRET_KEY=secret_key" --name minio -d napnap75/rpi-minio:server server /data
	mkdir mc_config
	docker run --rm -v ${PWD}/mc_config:/root/.mc --link minio napnap75/rpi-minio:client config host add test http://minio:9000 access_key secret_key
	docker run --rm -v ${PWD}/mc_config:/root/.mc --link minio napnap75/rpi-minio:client mb test/test
	docker run --rm -v ${PWD}/mc_config:/root/.mc --link minio napnap75/rpi-minio:client cp /usr/bin/mc test/test/
	docker run --rm -v ${PWD}/mc_config:/root/.mc --link minio napnap75/rpi-minio:client mirror test/test/ /root/.mc/test/
	docker run --rm -v ${PWD}/mc_config:/root/.mc --link minio napnap75/rpi-minio:client cp /usr/bin/mc /root/.mc/test/mc.origin
	diff mc_config/test/mc.origin mc_config/test/mc
	rm -fr mc_config mc_config
	docker stop minio

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
