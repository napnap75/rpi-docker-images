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
//	"strconv"
	"syscall"
//	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
	"google.golang.org/api/option"
)

var googleApiKey string
var googleVoice string

func getAudioFileFromGoogle(apiKey string, textInput string, voiceName string, filename string) {
	// Instantiates a client.
	ctx := context.Background()
	client, err := texttospeech.NewClient(ctx, option.WithAPIKey(apiKey))
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
			Name: voiceName,
		},
		// Select the type of audio file you want returned.
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
			EffectsProfileId: []string{"small-bluetooth-speaker-class-device"},
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

	log.Print("Audio content written to file: %v\n", filename)
}

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	log.Print("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())
	
	type Payload struct {
		text []byte `json:"text"`
	}
	var messagePayload Payload
	err := json.Unmarshal(message.Payload(), &messagePayload)
	if err != nil {
		log.Fatal(err)
	}

	if messagePayload.text != "" {
		hash := md5.Sum(messagePayload.text)

	}
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
	flag.StringVar(&googleApiKey, "google-api-key", "", "Google Cloud API Key")
	flag.StringVar(&googleVoice, "google-voice", "fr-FR-Wavenet-C", "Google TTS voice identifier")
	flag.Parse()

	// Connect to MQTT
	mqttConnectAndSubscribe(*mqttServer, *mqttClientid, *mqttUsername, *mqttPassword)

	getAudioFileFromGoogle(googleApiKey, "J'ai ajoute du pain", googleVoice, "test.mp3")

	<-c
}
