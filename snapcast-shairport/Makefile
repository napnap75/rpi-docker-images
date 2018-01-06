build:
	docker build -t napnap75/rpi-snapcast-shairport .

push:
	docker login -u="$DOCKER_USERNAME" -p="$DOCKER_PASSWORD"
	docker push napnap75/rpi-snapcast-shairport:latest
