all: build push

build-pre:
	curl -L -o qemu-arm-static.tar.gz https://github.com/multiarch/qemu-user-static/releases/download/v2.6.0/qemu-arm-static.tar.gz && tar xzf qemu-arm-static.tar.gz

build: build-pre
	docker build -t napnap75/rpi-snips-handler:display .

push:
	docker login -u="${DOCKER_USERNAME}" -p="${DOCKER_PASSWORD}"
	docker push napnap75/rpi-snips-handler:display

clean:
	rm qemu-arm-static*
