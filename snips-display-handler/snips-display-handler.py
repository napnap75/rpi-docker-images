import paho.mqtt.client as mqtt
import datetime
import os


def time_now():
    return datetime.datetime.now().strftime('%H:%M:%S.%f')

# MQTT client to connect to the bus
mqtt_client = mqtt.Client()


def on_connect(client, userdata, flags, rc):
    # subscribe to all messages
    mqtt_client.subscribe('#')


# Process a message as it arrives
def on_message(client, userdata, msg):
        if len(msg.payload) > 0:
            print('[{}] - {}: {}'.format(time_now(), msg.topic, msg.payload))
        else:
            print('[{}] - {}'.format(time_now(), msg.topic))

mqtt_client.on_connect = on_connect
mqtt_client.on_message = on_message
mqtt_client.connect(os.environ.get('MQTT_HOST'), int(os.environ.get('MQTT_PORT')))
mqtt_client.loop_forever()
