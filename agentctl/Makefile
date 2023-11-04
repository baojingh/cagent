VERSION = 0.0.1
CURRENT_TIME = $(shell date +"%Y-%m-%d %H:%M:%S")
GIT_HASH = $(git show -s --format=%H)
GO_VERSION = $(go version)

all: clean build-linux64


gen:
	rm -rf pb/*.go
	protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative     pb/agent_service.proto

build-linux64: pre
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -X 'agentctl/cmd.buildVersion=$(VERSION)' -X 'agentctl/cmd.buildTime=$(CURRENT_TIME)'" -x -v -o build/agentctl
	chmod u+x build/agentctl

cp:build-linux64
	cp -r cert/* /opt/agentctl/cert/

run:cp
	build/agentctl

clean:
	rm -rf build/
	rm -rf /opt/agentctl/cert/*
	rm -rf /var/log/agentctl/*

pre:
	mkdir -p /var/log/agentctl/
	mkdir -p build/
	mkdir -p /opt/agentctl/cert

suri:
	sudo mkdir -p /data/tmp/rules
	sudo chown -R hadoop:hadoop  /data/tmp/rules
	dd if=/dev/urandom  of=/data/tmp/rules/hello.txt  count=500 bs=1M
	build/agentctl suricata    --ip  127.0.0.1 --port 50051  upload --file  /data/tmp/rules/hello.txt
	