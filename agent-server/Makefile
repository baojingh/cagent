VERSION = 1.0.10
CURRENT_TIME = $(shell date +"%Y-%m-%d %H:%M:%S")

build-all: build build-linux64 build-linux32
build: pre
	go build -ldflags="-s -w -X 'main.buildVersion=v$(VERSION)' -X 'main.buildTime=$(CURRENT_TIME)'" -x -v -o build/ssmctl-grpc-client
build-linux64: pre
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -X 'main.buildVersion=v$(VERSION)' -X 'main.buildTime=$(CURRENT_TIME)'" -x -v -o build/ssmctl-grpc-client-linux64
build-linux32: pre
	GOOS=linux GOARCH=386 go build -ldflags="-s -w -X 'main.buildVersion=v$(VERSION)' -X 'main.buildTime=$(CURRENT_TIME)'" -x -v -o build/ssmctl-grpc-client-linux32

clean:
	rm -rf build
pre:
	mkdir -p build
	mkdir -p build/logs
	cp -r ./cer build/