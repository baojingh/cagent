VERSION = 0.0.1
CURRENT_TIME = $(shell date +"%Y-%m-%d %H:%M:%S")

all: clean build-linux64 run

build-linux64: pre
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -X 'main.buildVersion=v$(VERSION)' -X 'main.buildTime=$(CURRENT_TIME)'" -x -v -o build/agent-client-linux64

run:build-linux64
	build/agent-client-linux64

clean:
	rm -rf build
	rm -rf var/log/agent-client/

pre:
	mkdir -p /var/log/agent-client/ build
