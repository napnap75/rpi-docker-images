all: build push

build:
	docker build -t napnap75/rpi-snips-google-tts .

login:
	if [ "${DOCKER_USERNAME}" != "" ]; then docker login -u="${DOCKER_USERNAME}" -p="${DOCKER_PASSWORD}" ; fi

push: login
	docker push napnap75/rpi-snips-google-tts
