build:
	cat Dockerfile.template | sed "s/%VERSION%/${VERSION}/g" | sed "s/%VARIANT%/${VARIANT}/g" > Dockerfile.${VERSION}-${VARIANT}
	docker build -t napnap75/rpi-nextcloud-smb:${VERSION}-${VARIANT} -f Dockerfile.${VERSION}-${VARIANT} .

push:
	if [ "${DOCKER_USERNAME}" != "" ]; then docker login -u="${DOCKER_USERNAME}" -p="${DOCKER_PASSWORD}" ; fi
	docker push napnap75/rpi-nextcloud-smb:${VERSION}-${VARIANT}
