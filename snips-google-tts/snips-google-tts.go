package main

import (
	"context"
	"crypto/md5"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
)

var googleVoice string

func getAudioFileFromGoogle(textInput string, filename string) {
	// Instantiates a client.
	ctx := context.Background()
	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Perform the text-to-speech request on the text input with the selected voice parameters and audio file type.
	req := texttospeechpb.SynthesizeSpeechRequest{
		// Set the text input to be synthesized.
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: textInput},
		},
		// Build the voice request, select the voice
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: "fr-FR",
			Name: googleVoice,
		},
		// Select the type of audio file you want returned.
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_LINEAR16,
			EffectsProfileId: []string{"small-bluetooth-speaker-class-device"},
			Pitch: -2,
		},
	}

	// Send the request
	resp, err := client.SynthesizeSpeech(ctx, &req)
	if err != nil {
		log.Fatal(err)
	}

	// Write the response to the file
	err = ioutil.WriteFile(filename, resp.AudioContent, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Close the connection
	err = client.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("The message %v was written to file: %v\n", textInput, filename)
}

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	log.Printf("Received message on topic: %+v\nMessage: %+v\n", message.Topic(), message.Payload())
	log.Printf("Received message on topic: %s\nMessage: %s\n", string(message.Topic()), string(message.Payload()))

	// Decode the message
	type Payload struct {
		Text string `json:"text"`
		RequestId string `json:"id"`
		SiteId string `json:"siteId"`
		SessionId string `json:"sessionId"`
	}
	var messagePayload Payload
	err := json.Unmarshal(message.Payload(), &messagePayload)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Full message: %+v\n", messagePayload.Text)
	log.Printf("Message text: %s\n", messagePayload.Text)

	// Get the audio file from Google TTS if necessary
	hash := fmt.Sprintf("%x", md5.Sum([]byte(messagePayload.Text)))

	log.Printf("Message hash: %s\n", hash)

	if _, err := os.Stat("/tmp/messages/" + hash); os.IsNotExist(err) {
		getAudioFileFromGoogle(string(messagePayload.Text), "/tmp/messages/" + hash)
	}

	// Sent the audio file
	audio, err := ioutil.ReadFile("/tmp/messages/" + hash)
	if err != nil {
		log.Fatal(err)
	}
	client.Publish(fmt.Sprintf("hermes/audioServer/default/playBytes/%s", hash), 0, false, audio)

	// Answer back
	client.Publish("hermes/tts/sayFinished", 0, false, []byte(fmt.Sprintf("{\"id\": \"%x\", \"sessionId\": \"%x\"}", messagePayload.RequestId, messagePayload.SessionId)))
}

func mqttConnectAndSubscribe(server string, clientid string, username string, password string) {
	connOpts := mqtt.NewClientOptions().AddBroker(server).SetClientID(clientid).SetCleanSession(true)
	if username != "" {
		connOpts.SetUsername(username)
		if password != "" {
			connOpts.SetPassword(password)
		}
	}
	tlsConfig := &tls.Config{InsecureSkipVerify: true, ClientAuth: tls.NoClientCert}
	connOpts.SetTLSConfig(tlsConfig)
	connOpts.OnConnect = func(c mqtt.Client) {
		if token := c.Subscribe("hermes/tts/say", 0, onMessageReceived); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}
	client := mqtt.NewClient(connOpts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		fmt.Printf("Connected to %s\n", server)
	}

}

func main() {
	// Logs
	mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)

	// Handle interrupts to clean properly
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Load the parameters
	mqttServer := flag.String("mqtt-server", "tcp://127.0.0.1:1883", "The full url of the MQTT server to connect to")
	mqttClientid := flag.String("mqtt-clientid", "snips-google-tts", "A clientid for the connection")
	mqttUsername := flag.String("mqtt-username", "", "A username to authenticate to the MQTT server")
	mqttPassword := flag.String("mqtt-password", "", "Password to match username")
	flag.StringVar(&googleVoice, "google-voice", "fr-FR-Wavenet-A", "Google TTS voice identifier")
	flag.Parse()

	// Connect to MQTT
	mqttConnectAndSubscribe(*mqttServer, *mqttClientid, *mqttUsername, *mqttPassword)

	<-c
}
