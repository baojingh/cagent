VERSION = 0.0.1
CURRENT_TIME = $(shell date +"%Y-%m-%d %H:%M:%S")

all: clean build

gen:
	rm -rf pb*.go
	protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative     pb/agent_service.proto

build:
	mkdir build
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -X 'main.buildVersion=v$(VERSION)' -X 'main.buildTime=$(CURRENT_TIME)'" -x -v -o build/agentserver

build-linux64: pre
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -X 'main.buildVersion=v$(VERSION)' -X 'main.buildTime=$(CURRENT_TIME)'" -x -v -o build/agent-server-linux64

cp:build-linux64
	cp setting/agent-server.yaml /opt/agent-server/
	cp -r cert/* /opt/agent-server/cert/

run:cp
	build/agent-server-linux64

clean:
	rm -rf build/
	rm -rf opt/agent-server/*
	rm -rf /var/log/agent-server/*

pre:
	mkdir -p /var/log/agent-server/
	mkdir -p build/
	mkdir -p /opt/agent-server/cert
	chmod u+x command/*.sh


flu:
	
	build/agentctl suricata    --ip  127.0.0.1 --port 50051  upload --file  /data/tmp/suricata.rules