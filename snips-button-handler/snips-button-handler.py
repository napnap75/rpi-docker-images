import datetime
import os
import paho.mqtt.client as mqtt
import RPi.GPIO as GPIO


# MQTT client to connect to the bus
mqtt_client = mqtt.Client()

# Initialize the GPIO ports
button_channel = int(os.environ.get('GPIO_BUTTON'))
led_channel = int(os.environ.get('GPIO_LED'))
GPIO.setmode(GPIO.BCM)
GPIO.setup(button_channel, GPIO.IN, pull_up_down=GPIO.PUD_DOWN)
GPIO.setup(led_channel, GPIO.OUT)

try:
	# Simulates the hotword on button pressed
	def on_button(channel):
		if GPIO.input(channel):
			mqtt_client.publish("hermes/hotword/default/detected", '{"siteId":"default","modelId":"button","modelVersion":null,"modelType":"personal"}')

	# Subscribe to the messages
	def on_connect(client, userdata, flags, rc):
		mqtt_client.subscribe('hermes/asr/#')

	# Process a message as it arrives
	def on_message(client, userdata, msg):
		if msg.topic == 'hermes/asr/startListening':
			GPIO.output(led_channel, True)
		elif msg.topic == 'hermes/asr/stopListening':
			GPIO.output(led_channel, False)

	GPIO.add_event_detect(button_channel, GPIO.RISING, callback=on_button, bouncetime=200)

	mqtt_client.on_connect = on_connect
	mqtt_client.on_message = on_message
	mqtt_client.connect(os.environ.get('MQTT_HOST'), int(os.environ.get('MQTT_PORT')))
	mqtt_client.loop_forever()

except KeyboardInterrupt as e:
	GPIO.cleanup()
