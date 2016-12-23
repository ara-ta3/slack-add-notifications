GOOS=
GOARCH=
goos_opt=GOOS=$(GOOS)
goarch_opt=GOARCH=$(GOARCH)
out=slack-new-channel
out_opt="-o $(out)"

help:
	@cat Makefile

install:
	go get golang.org/x/net/websocket

run:
	go run ./main.go ./config.go

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build

build: 
	$(goos_opt) $(goarch_opt) go build $(out_opt)

build_for_linux:
	$(MAKE) build GOOS=linux GOARCH=amd64 out_opt=""

build_for_local:
	$(MAKE) build goos_opt= goarch_opt= out_opt=
