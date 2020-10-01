import evdev
import selectors
import paho.mqtt.client as mqtt
import os
import time

#print(os.environ)

while True:
	try:
		keyboard = evdev.InputDevice('/dev/input/event0')
		remote = evdev.InputDevice('/dev/input/event1')
		break
	except FileNotFoundError as e:
		print(".")
		time.sleep(10)

selector = selectors.DefaultSelector()
selector.register(keyboard, selectors.EVENT_READ)
selector.register(remote, selectors.EVENT_READ)
print("Remote found")

mqtt_client = mqtt.Client()
mqtt_client.connect(os.environ.get('MQTT_HOST'), int(os.environ.get('MQTT_PORT')))
mqtt_client.loop_start()
print("Connected to MQTT")

try:
	while True:
		for key, mask in selector.select():
			device = key.fileobj
			for event in device.read():
				if event.type == evdev.ecodes.EV_KEY:
					if event.value == 1:
						print("Received code: " + evdev.ecodes.KEY[event.code])
						mqtt_client.publish(os.environ.get('MQTT_TOPIC'), evdev.ecodes.KEY[event.code])

except KeyboardInterrupt as e:
	mqtt_client.loop_stop()
