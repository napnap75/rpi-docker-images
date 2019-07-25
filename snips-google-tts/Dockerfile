FROM golang:alpine AS builder

WORKDIR $GOPATH/src/napnap75/snips-google-tts/
COPY snips-google-tts.go .
RUN apk add --no-cache git \
	&& go get -d -v
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/snips-google-tts

FROM alpine

COPY --from=builder /go/bin/snips-google-tts /go/bin/snips-google-tts
ENTRYPOINT ["/go/bin/snips-google-tts"]
