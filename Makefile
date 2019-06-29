GO=go
GOOS=
GOARCH=
goos_opt=GOOS=$(GOOS)
goarch_opt=GOARCH=$(GOARCH)
out=slack-add-notifications
out_opt="-o $(out)"

help:
	@cat Makefile

install:
	$(GO) mod vendor

run: config.json
	$(GO) run ./main.go ./config.go

config.json: config.sample.json
	cp -f $< $@

build: 
	$(goos_opt) $(goarch_opt) $(GO) build $(out_opt)

build_for_linux:
	$(MAKE) build GOOS=linux GOARCH=amd64 out_opt=""

build_for_local:
	$(MAKE) build goos_opt= goarch_opt= out_opt=

test/new_channels:
	curl -i -X POST localhost:8080 -d '{"type": "channel_created", "channel": {"name": "test"}}'

test/new_emojis:
	curl -i -X POST localhost:8080 -d '{"type": "emoji_changed", "subtype": "add", "name": "test"}'

test/team_joined: UserID=XXX
test/team_joined:
	curl -i -X POST localhost:8080 -d '{"type": "team_join", "user": {"id": "$(UserID)", "is_bot": false}}'

