FROM golang:alpine AS builder

WORKDIR $GOPATH/src/napnap75/snips-google-tts/

RUN apk add --no-cache git gcc musl-dev \
	&& go get -d -v  github.com/eclipse/paho.mqtt.golang cloud.google.com/go/texttospeech/apiv1 google.golang.org/genproto/googleapis/cloud/texttospeech/v1 google.golang.org/api/option

COPY snips-google-tts.go .

RUN go get -d -v \
	&& go build -ldflags="-w -s" -o /go/bin/snips-google-tts
#	&& GOOS=linux GOARCH=arm GOARM=6 go build -ldflags="-w -s" -o /go/bin/snips-google-tts

FROM alpine

COPY --from=builder /go/bin/snips-google-tts /go/bin/snips-google-tts
COPY google-credentials.json /
ENV GOOGLE_APPLICATION_CREDENTIALS=/google-credentials.json
RUN apk add --no-cache ca-certificates

VOLUME /tmp/messages

ENTRYPOINT ["/go/bin/snips-google-tts"]
