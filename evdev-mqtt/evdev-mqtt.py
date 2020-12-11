import evdev
import logging
import os
import paho.mqtt.client as mqtt
import selectors
import sys
import time

logging.basicConfig(stream=sys.stdout, level=int(os.environ.get('LOG_LEVEL', "20")))
logger = logging.getLogger()

while True:
	try:
		keyboard = evdev.InputDevice('/dev/input/event0')
		remote = evdev.InputDevice('/dev/input/event1')
		break
	except FileNotFoundError as e:
		logger.debug("Remote not ready, waiting...")
		time.sleep(10)

selector = selectors.DefaultSelector()
selector.register(keyboard, selectors.EVENT_READ)
selector.register(remote, selectors.EVENT_READ)
logger.info("Remote found")

mqtt_client = mqtt.Client()
mqtt_client.enable_logger(logger)
mqtt_client.connect(os.environ.get('MQTT_HOST'), int(os.environ.get('MQTT_PORT')))
mqtt_client.loop_start()
logger.info("Connected to MQTT")

try:
	while True:
		for key, mask in selector.select():
			device = key.fileobj
			for event in device.read():
				if event.type == evdev.ecodes.EV_KEY:
					if event.value == 1:
						logger.debug("Received code: " + evdev.ecodes.KEY[event.code])
						mqtt_client.publish(os.environ.get('MQTT_TOPIC'), evdev.ecodes.KEY[event.code])

except KeyboardInterrupt as e:
	mqtt_client.loop_stop()
