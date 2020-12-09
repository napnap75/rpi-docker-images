build:
	docker build -t napnap75/rpi-snips:common -f Dockerfile-common .
	set -e ; for template in analytics asr audio-server dialogue nlu tts watch ; do \
		sed "s/template/$$template/g" Dockerfile-template > Dockerfile.$$template ; \
		docker build -t napnap75/rpi-snips:$$template -f Dockerfile.$$template . ; \
	done
	set -e ; for template in hotword sam skill-server ; do \
		docker build -t napnap75/rpi-snips:$$template -f Dockerfile-$$template . ; \
	done

push:
	if [ "${DOCKER_USERNAME}" != "" ]; then docker login -u="${DOCKER_USERNAME}" -p="${DOCKER_PASSWORD}" ; fi
	set -e ; for template in analytics asr audio-server hotword dialogue nlu skill-server tts watch ; do \
		docker push napnap75/rpi-snips:$$template ; \
		version=$$(docker run --rm -it napnap75/rpi-snips:$$template /usr/bin/snips-$$template --version | egrep -o "\([0-9\.]+\)" | egrep -o "[0-9]+\.[0-9]+\.[0-9]+") ; \
		docker tag napnap75/rpi-snips:$$template napnap75/rpi-snips:$$template-$$version ; \
		docker push napnap75/rpi-snips:$$template-$$version ; \
	done
	set -e ; for template in sam ; do \
		docker push napnap75/rpi-snips:$$template ; \
		version=$$(docker run --rm -it napnap75/rpi-snips:$$template --version | egrep -o "[0-9]+\.[0-9]+\.[0-9]+") ; \
		docker tag napnap75/rpi-snips:$$template napnap75/rpi-snips:$$template-$$version ; \
		docker push napnap75/rpi-snips:$$template-$$version ; \
	done

clean:
	for template in analytics asr audio-server dialogue nlu tts watch ; do rm Dockerfile.$$template ; done
