build: build-v4 build-v5

build-v4:
	docker build -t napnap75/rpi-gandi:v4 -f Dockerfile.v4 .

build-v5:
	docker build -t napnap75/rpi-gandi:v5 -f Dockerfile.v5 .

push: push-v4 push-v5

push-login:
	docker login -u="${DOCKER_USERNAME}" -p="${DOCKER_PASSWORD}"

push-v4: push-login
	docker push napnap75/rpi-gandi:v4
	
push-v5: push-login
	docker push napnap75/rpi-gandi:v5
	docker tag napnap75/rpi-gandi:v5 napnap75/rpi-gandi:latest
	docker push napnap75/rpi-gandi:latest
