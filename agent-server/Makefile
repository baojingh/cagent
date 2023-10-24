VERSION = 0.0.1
CURRENT_TIME = $(shell date +"%Y-%m-%d %H:%M:%S")

all: clean build-linux64 run

gen:
	protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative     pb/agent_service.proto

build-linux64: pre
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -X 'main.buildVersion=v$(VERSION)' -X 'main.buildTime=$(CURRENT_TIME)'" -x -v -o build/agent-server-linux64

cp:build-linux64
	sudo cp build/agent-server-linux64 /usr/local/bin/agent-server-linux64
	cp setting/agent-server.yaml /opt/agent-server/
	cp -r cert/* /opt/agent-server/cert/

run:cp
	agent-server-linux64

clean:
	rm -rf build/
	rm -rf opt/agent-server/*
	rm -rf /var/log/agent-server/*
	rm -rf pb/*.go

pre:
	mkdir -p /var/log/agent-server/
	mkdir -p build/
	mkdir -p /opt/agent-server/cert
	chmod u+x command/*.sh
