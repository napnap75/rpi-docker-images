build:
	if [ ! -d gitea ]; then git clone https://github.com/go-gitea/gitea ; fi
	if [ ! -f gitea/qemu-arm-static ]; then curl -L https://github.com/resin-io/qemu/releases/download/v2.9.0%2Bresin1/qemu-2.9.0.resin1-arm.tar.gz | tar xz --strip-components=1 ; cp qemu-arm-static gitea/qemu-arm-static ; fi
	patch gitea/Dockerfile Dockerfile.patch
	docker build -t napnap75/rpi-gitea:latest gitea

push-init:
	if [ "${DOCKER_USERNAME}" != "" ]; then docker login -u="${DOCKER_USERNAME}" -p="${DOCKER_PASSWORD}" ; fi

push: push-init
	docker push napnap75/rpi-gitea:latest
