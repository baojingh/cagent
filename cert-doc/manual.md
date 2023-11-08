
golang grpc tls通信 新手 执行步骤s

# Steps
```bash
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout server.key -out server.crt

openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout client.key -out client.crt
```




# REF
http://liuqh.icu/2022/02/23/go/rpc/06-tls/