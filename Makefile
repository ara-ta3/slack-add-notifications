SLACK_TOKEN=
SLACK_NEW_CHANNEL_NOTIFICATION_CHANNEL_ID=

run:
	go run ./main.go ./config.go

help:
	@cat Makefile

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build
