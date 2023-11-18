
# golang grpc tls通信 新手 执行步骤s
没有启用证书的gRPC服务和客户端进行的是明文通信，信息面临被任何第三方监听的风险。为了保证gRPC通信不被第三方监听、篡改或伪造，可以对服务器启动TLS加密特性。

CA 体系
什么是 CA

CA是Certificate Authority的缩写，也叫“证书授权中心”。它是负责管理和签发证书的第三方机构，作用是检查证书持有者身份的合法性，并签发证书，以防证书被伪造或篡改。
什么是 CA 证书

CA 证书就是CA颁发的证书。我们常听到的数字证书就是CA证书,CA证书包含信息有:证书拥有者的身份信息，CA 机构的签名，公钥和私钥。CA 一般用其根证书的私钥进行签名生成证书公钥文件，爷就是颁发给用户的证书。
什么是根证书

根证书（root certificate）是属于根证书颁发机构（CA）的公钥证书。我们可以通过验证 CA 的签名从而信任 CA ，任何人都可以得到 CA 的证书（含公钥），用以验证它所签发的证书（客户端、服务端）。
什么是 CSR

CSR（Certificate Signing Request），即证书签名请求文件，是证书申请者在申请数字证书时由 CSP(加密服务提供者)在生成私钥的同时也生成证书请求文件，证书申请者只要把 CSR 文件提交给证书颁发机构后，证书颁发机构使用其根证书私钥签名就生成了证书公钥文件，也就是颁发给用户的证书。

为什么要基于 CA

如果采用自签的方式，即服务端通过私钥自签公钥，客户端端通过证书文件和服务名称构造 TLS 凭证连接，有一个很大的问题。那就是服务端怎么安全的把证书传给客户端，必定要通过公共的网络环境，且任何攻击者都可以伪造公钥证书，并在服务端发现有人伪造之前，提前获取/发送信息于/给客户端。使用 CA 可以避免这一步，也是常见的公钥分发方式之一。

完整流程
生成 CA 根证书
1.新建 ca.conf 来放相关信息

```
[ req ]
default_bits       = 4096
distinguished_name = req_distinguished_name

[ req_distinguished_name ]
countryName                 = GB
countryName_default         = CN
stateOrProvinceName         = State or Province Name (full name)
stateOrProvinceName_default = Zhejiang
localityName                = Locality Name (eg, city)
localityName_default        = Hangzhou
organizationName            = Organization Name (eg, company)
organizationName_default    = ATech
commonName                  = CommonName (e.g. server FQDN or YOUR name)
commonName_max              = 64
commonName_default          = an1ex.top
```
2.生成 CA 私钥
openssl genrsa -out ca.key 2048

3.生成根证书
这里是用了-subjflag 来传入非 ca.conf 中默认的信息，当然也可以不用 conf，只用-subj传，未传的参数默认为空
openssl req -new -x509 -days 365 -subj "/C=GB/L=Zhejiang/O=github/CN=an1ex.top" -key ca.key -out ca.crt -config ca.conf


生成服务端证书
1.新建 server.conf

[ req ]
default_bits       = 2048
distinguished_name = req_distinguished_name

[ req_distinguished_name ]
countryName                 = Country Name (2 letter code)
countryName_default         = CN
stateOrProvinceName         = State or Province Name (full name)
stateOrProvinceName_default = Zhejiang
localityName                = Locality Name (eg, city)
localityName_default        = Hangzhou
organizationName            = Organization Name (eg, company)
organizationName_default    = ATech
commonName                  = CommonName (e.g. server FQDN or YOUR name)
commonName_max              = 64
commonName_default          = an1ex.top
[ req_ext ]
subjectAltName = @alt_names
[alt_names]
DNS.1   = an1ex.top
IP      = 127.0.0.1

2.生成服务端私钥
$ openssl genrsa -out server.key 2048

3. 生成CSR
当然也可以只通过-subj传证书参数
openssl req -new  -subj "/C=GB/L=Zhejiang/O=github/CN=an1ex.top" -key server.key -out server.csr -config server.conf

4. 基于 CA 签发证书

签发注意下用的 flag，如果用配置文件传：
openssl x509 -req -sha256 -CA ca.crt -CAkey ca.key -CAcreateserial -days 365 -in server.csr -out server.crt -extensions req_ext -extfile server.conf

如果命令行参数传，太长了太长了，还是少用：
openssl x509 -req -subj "/C=GB/L=Zhejiang/O=github/CN=an1ex.top" -extfile <(printf "subjectAltNa


生成客户端证书
1.生成客户端私钥
openssl genrsa -out client.key 2048

2.生成 CSR

服务端证书参数少，没了 SAN 就直接命令行敲吧

openssl req -new -subj "/C=GB/L=Zhejiang/O=github/CN=an1ex.top" -key client.key -out client.csr

3.基于 CA 签发证书


openssl x509 -req -sha256 -CA ca.crt -CAkey ca.key -CAcreateserial -days 365 -in client.csr -out client.crt


证书校验，用下面命令来校验是否签发成功
openssl verify -CAfile ca.crt server.crt
openssl verify -CAfile ca.crt client.crt


# REF
https://www.an1ex.top/2022/08/08/gRPC%E4%B8%AD%E5%9F%BA%E4%BA%8ECA%E7%9A%84TLS%E8%AE%A4%E8%AF%81/
https://blog.csdn.net/Mr_XiMu/article/details/124937025