VERSION = 0.0.1
CURRENT_TIME = $(shell date +"%Y-%m-%d %H:%M:%S")

all: clean build-linux64 run

build-linux64: pre
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -X 'main.buildVersion=v$(VERSION)' -X 'main.buildTime=$(CURRENT_TIME)'" -x -v -o build/agent-server-linux64

run:build-linux64
	build/agent-server-linux64

push:
	git add .
	git commit -m "update"
	git push origin main

clean:
	rm -rf build
	rm -rf var/log/agent-server/

pre:
	mkdir -p /var/log/agent-server/ build
