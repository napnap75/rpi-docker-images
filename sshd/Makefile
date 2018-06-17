build:
	docker build -t napnap75/rpi-sshd -f Dockerfile .


test: build
	docker run --rm -it -v test-sshd:/config -p 2200:22 napnap75/rpi-sshd

test-bash: build
	docker run --rm -it -v test-sshd:/config -p 2200:22 napnap75/rpi-sshd /bin/bash


push-init:
	if [ "${DOCKER_USERNAME}" != "" ]; then docker login -u="${DOCKER_USERNAME}" -p="${DOCKER_PASSWORD}" ; fi

push: push-init
	docker push napnap75/rpi-sshd
