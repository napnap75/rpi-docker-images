build:
	docker build -t napnap75/rpi-mopidy:latest .

push:
	docker login -u="${DOCKER_USERNAME}" -p="${DOCKER_PASSWORD}"
	docker push napnap75/rpi-mopidy:latest

