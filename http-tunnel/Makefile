build:
	docker build -t napnap75/rpi-http-tunnel:client -f Dockerfile.client .
	docker build -t napnap75/rpi-http-tunnel:server -f Dockerfile.server .

push:
	if [ "${DOCKER_USERNAME}" != "" ]; then docker login -u="${DOCKER_USERNAME}" -p="${DOCKER_PASSWORD}" ; fi
	docker push napnap75/rpi-http-tunnel:client
	version=$$(docker run --rm -it napnap75/rpi-http-tunnel:client -version | egrep -o "[a-z0-9]+") ; \
		docker tag napnap75/rpi-http-tunnel:client napnap75/rpi-http-tunnel:client-$$version ; \
		docker push napnap75/rpi-http-tunnel:client-$$version
	docker push napnap75/rpi-http-tunnel:server
	version=$$(docker run --rm -it napnap75/rpi-http-tunnel:server -version | egrep -o "[a-z0-9]+") ; \
		docker tag napnap75/rpi-http-tunnel:server napnap75/rpi-http-tunnel:server-$$version ; \
		docker push napnap75/rpi-http-tunnel:server-$$version
