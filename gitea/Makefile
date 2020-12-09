build:
	if [ ! -d gitea ]; then git clone https://github.com/go-gitea/gitea ; fi
	cp Dockerfile gitea/
	docker build -t napnap75/rpi-gitea:latest gitea

push:
	if [ "${DOCKER_USERNAME}" != "" ]; then docker login -u="${DOCKER_USERNAME}" -p="${DOCKER_PASSWORD}" ; fi
	docker push napnap75/rpi-gitea:latest
