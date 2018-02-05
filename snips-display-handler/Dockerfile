FROM arm32v6/python:alpine3.7

COPY qemu-arm-static /usr/bin/qemu-arm-static

RUN pip install paho-mqtt

COPY snips-display-handler.py /usr/local/bin

ENV MQTT_HOST=localhost MQTT_PORT=1883

CMD python /usr/local/bin/snips-display-handler.py
