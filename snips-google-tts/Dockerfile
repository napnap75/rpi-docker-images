FROM golang:alpine AS builder
WORKDIR $GOPATH/src/napnap75/snips-google-tts/
COPY snips-google-tts.go .
RUN apk add --no-cache git gcc musl-dev \
	&& go get -d -v \
	&& GOOS=linux GOARCH=arm GOARM=6 go build -ldflags="-w -s" -o /go/bin/snips-google-tts

FROM arm32v6/alpine:latest
COPY --from=builder /go/bin/snips-google-tts /go/bin/snips-google-tts
ENV GOOGLE_APPLICATION_CREDENTIALS=/google-credentials.json
RUN apk add --no-cache ca-certificates
VOLUME /tmp/messages
ENTRYPOINT ["/go/bin/snips-google-tts"]
