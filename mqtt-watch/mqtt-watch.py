import paho.mqtt.client as mqtt
import datetime
import os
import re

def time_now():
    return datetime.datetime.now().strftime('%H:%M:%S.%f')

mqtt_client = mqtt.Client()
prog = re.compile("hermes/audioServer/\w+/audioFrame")

def on_connect(client, userdata, flags, rc):
	mqtt_client.subscribe('#')

def on_message(client, userdata, msg):
	if prog.match(msg.topic) is None:
		if len(msg.payload) > 0:
			print('[{}] - {}: {}'.format(time_now(), msg.topic, msg.payload))
		else:
			print('[{}] - {}'.format(time_now(), msg.topic))

mqtt_client.on_connect = on_connect
mqtt_client.on_message = on_message
mqtt_client.connect(os.environ.get('MQTT_HOST'), int(os.environ.get('MQTT_PORT')))
mqtt_client.loop_forever()
