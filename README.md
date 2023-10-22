# cagent

### Components Client


### Components Server


### Components ctl



# Setup ENV
### Install protoc cmd
```
apt install -y protobuf-compiler
protoc --version  # Ensure compiler version is 3+
```
### Install go plugins
Install the protocol compiler plugins for Go using the following commands:

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

Update your PATH so that the protoc compiler can find the plugins:
export PATH="$PATH:$(go env GOPATH)/bin"

```
protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative     helloworld/helloworld.proto
```

# Roadmap
### agentctl
It is a linux tool for the client host to use.
It is not a daemon process.
ctl has a dedicated feature to do the biz.

### Features
1. Update conf files
2. Restart process when conf files updated success
3. Upload files with big size, for example 50MB





