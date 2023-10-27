VERSION = 0.0.1
CURRENT_TIME = $(shell date +"%Y-%m-%d %H:%M:%S")

all: clean build-linux64 run


gen:
	rm -rf pb/*.go
	protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative     pb/agent_service.proto

build-linux64: pre
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -X 'main.buildVersion=v$(VERSION)' -X 'main.buildTime=$(CURRENT_TIME)'" -x -v -o build/agentctl-linux64

cp:build-linux64
	sudo cp build/agentctl-linux64 /usr/local/bin/agentctl-linux64
	cp -r cert/* /opt/agentctl/cert/

run:cp
	agentctl-linux64

clean:
	rm -rf build/
	rm -rf /opt/agentctl/cert/*
	rm -rf /var/log/agentctl/*
	sudo rm -rf /usr/local/bin/agentctl-linux64

pre:
	mkdir -p /var/log/agentctl/
	mkdir -p build/
	mkdir -p /opt/agentctl/cert
