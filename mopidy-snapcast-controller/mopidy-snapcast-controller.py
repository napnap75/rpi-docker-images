import paho.mqtt.client as mqtt
import os
import asyncio
import logging
import snapcast.control

# Initiate logging
logger = logging.getLogger()
logger.setLevel(logging.INFO)
stream_handler = logging.StreamHandler()
stream_handler.setLevel(logging.INFO)
logger.addHandler(stream_handler)

# Connect to the Snapcast JSON-RPC API
loop = asyncio.get_event_loop()
server = loop.run_until_complete(snapcast.control.create_server(loop, os.environ.get('SNAPCAST_HOST')))

# Connect to the MQTT Broker
mqtt_client = mqtt.Client()

def on_connect(client, userdata, flags, rc):
	logger.info('connected to MQTT on %s:%s', os.environ.get('MQTT_HOST'), os.environ.get('MQTT_PORT'))
	mqtt_client.subscribe(os.environ.get('MQTT_TOPIC'))

def on_message(client, userdata, msg):
	state = msg.payload.decode('utf-8')
	logger.debug('received status %s from Mopidy', state)
	if state == 'playing':
		loop.run_until_complete(server.groups[0].set_stream(os.environ.get('SNAPCAST_MOPIDY_STREAM')))
	elif state == 'stopped' or state == 'paused':
		loop.run_until_complete(server.groups[0].set_stream(os.environ.get('SNAPCAST_ALTERNATE_STREAM')))

mqtt_client.on_connect = on_connect
mqtt_client.on_message = on_message
mqtt_client.connect(os.environ.get('MQTT_HOST'), int(os.environ.get('MQTT_PORT')))
mqtt_client.loop_forever()
