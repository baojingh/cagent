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

echo "export PATH="$PATH:$(go env GOPATH)/bin"" >> /etc/profile
. /etc/profile

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


# Detailed Design
### Suricata
1. Upload files with large size
2. update net in suricata.yaml
3. Update logrotate

### Fluentbit
1. Update host and port in fluent-bit.conf
2. Update logrotate



# Scripts
```
git config   user.email  "baojingh@163.com"
git config   user.name  "baojingh"


sudo mkdir -p /var/log/agentctl /opt/agentctl/cert
sudo chown hadoop:hadoop -R /var/log/agentctl /opt/agentctl/cert

sudo mkdir -p /var/log/agent-server /opt/agent-server/cert
sudo chown hadoop:hadoop -R /var/log/agent-server /opt/agent-server/cert
```



# Ref
https://www.cnblogs.com/borey/p/5715641.html
https://github.com/dimk00z/grpc-filetransfer
https://zhuanlan.zhihu.com/p/627848739
