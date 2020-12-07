#!/usr/bin/python3

import board
import busio
from digitalio import DigitalInOut
from adafruit_pn532.spi import PN532_SPI
import paho.mqtt.client as mqtt
import logging
import time
import yaml

# Load configuration
with open('config.yml') as f:
	configuration = yaml.load(f, Loader=yaml.FullLoader)

# Initiate logging
logger = logging.getLogger(__name__)
logger.setLevel(int(configuration['log-level']))

# SPI connection:
spi = busio.SPI(board.SCK, board.MOSI, board.MISO)
cs_pin = DigitalInOut(board.D5)
pn532 = PN532_SPI(spi, cs_pin, debug=False)
ic, ver, rev, support = pn532.get_firmware_version()
logger.info('Found PN532 with firmware version: {0}.{1}'.format(ver, rev))
pn532.SAM_configuration()

# MQTT connection:
mqtt_client = mqtt.Client()
mqtt_client.enable_logger(logger)
mqtt_client.connect(configuration['mqtt']['host'], int(configuration['mqtt']['port']))
mqtt_client.loop_start()

# Wait for tag
last_uid = None
logger.info('Waiting for RFID/NFC card...')
while True:
	uid = pn532.read_passive_target()
	if uid is None:
		last_uid = None
		continue
	else:
		if uid == last_uid:
			continue
		else:
			last_uid = uid

	string_uid = ''.join(format(x, '02x') for x in uid)
	logger.info('Found card with UID: %s', string_uid)
	try:
		for messages in configuration['commands'][string_uid]:
			logger.info('Publishing: %s to topic %s', messages['message'], messages['topic'])
			mqtt_client.publish(messages['topic'], messages['message'])
			time.sleep(1)
	except KeyError:
		logger.warning('Unknown UID: %s', string_uid)
