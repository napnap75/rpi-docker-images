all: build push

build:
	docker build -t napnap75/rpi-snips:common -f Dockerfile.common .
	set -e ; for template in analytics asr audio-server hotword dialogue nlu skill-server template tts watch ; do sed "s/template/$$template/g" Dockerfile-template > Dockerfile.$$template ; docker build -t napnap75/rpi-snips:$$template -f Dockerfile.$$template . ; done

push:
	if [ "${DOCKER_USERNAME}" != "" ]; then docker login -u="${DOCKER_USERNAME}" -p="${DOCKER_PASSWORD}" ; fi
	set -e ; for template in analytics asr audio-server hotword dialogue nlu skill-server template tts watch ; do docker push napnap75/rpi-snips:$$template ; done
	docker push napnap75/rpi-snips:hotword

clean:
	for template in analytics asr audio-server hotword dialogue nlu skill-server template tts watch ; do rm Dockerfile.$$template ; done
