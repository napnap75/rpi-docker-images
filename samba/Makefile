build:
	docker build -t napnap75/rpi-samba:latest .

push-login:
	docker login -u="${DOCKER_USERNAME}" -p="${DOCKER_PASSWORD}"
	
push: push-login
	docker push napnap75/rpi-samba:latest
