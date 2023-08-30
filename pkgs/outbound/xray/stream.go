package xray

/*
https://xtls.github.io/config/transport.html#streamsettingsobject

{
  "network": "tcp",
  "security": "none",
  "tlsSettings": {},
  "tcpSettings": {},
  "kcpSettings": {},
  "wsSettings": {},
  "httpSettings": {},
  "quicSettings": {},
  "dsSettings": {},
  "grpcSettings": {},
  "sockopt": {
    "mark": 0,
    "tcpFastOpen": false,
    "tproxy": "off",
    "domainStrategy": "AsIs",
    "dialerProxy": "",
    "acceptProxyProtocol": false,
    "tcpKeepAliveInterval": 0
  }
}

TLSSettings:
{
  "serverName": "xray.com",
  "rejectUnknownSni": false,
  "allowInsecure": false,
  "alpn": ["h2", "http/1.1"],
  "minVersion": "1.2",
  "maxVersion": "1.3",
  "cipherSuites": "此处填写你需要的加密套件名称,每个套件名称之间用:进行分隔",
  "certificates": [],
  "disableSystemRoot": false,
  "enableSessionResumption": false,
  "fingerprint": "",
  "pinnedPeerCertificateChainSha256": [""]
}

TCPSettings:
https://xtls.github.io/config/transports/tcp.html#tcpobject
{
  "acceptProxyProtocol": false,
  "header": {
    "type": "none"
  }
}
HttpHeaderObject:
{
  "type": "http",
  "request": {},
  "response": {}
}

WSSettings:
{
  "acceptProxyProtocol": false,
  "path": "/",
  "headers": {
    "Host": "xray.com"
  }
}

GRPCSettings:
{
  "serviceName": "name",
  "multiMode": false,
  "user_agent": "custom user agent",
  "idle_timeout": 60,
  "health_check_timeout": 20,
  "permit_without_stream": false,
  "initial_windows_size": 0
}

HTTPSettings:
{
  "host": ["xray.com"],
  "path": "/random/path",
  "read_idle_timeout": 10,
  "health_check_timeout": 15,
  "method": "PUT",
  "headers": {
    "Header": ["value"]
  }
}

*/

var XrayStream string = `{
	"network": "tcp",
	"security": "none"
}`

var XrayStreamTLS string = `{
	"serverName": "xray.com",
	"allowInsecure": true
}`

var XrayStreamTCPNone string = `{
	"header": {
	  "type": "none"
	}
}`

var XrayStreamTCPHTTP string = `{
  "header": {
      "type": "http",
      "request": {
          "path": ["/"],
          "headers": {
              "Host": ["fast.com"]
          }
      }
  }
}`

var XrayStreamTCPHeader string = `{
	"type": "http",
	"request": {},
	"response": {}
}`

var XrayStreamWebSocket string = `{
	"path": "/",
	"headers": {
	  "Host": "xray.com"
	}
}`

var XrayStreamHTTP string = `{
	"host": [""],
	"path": ""
}`

var XrayStreamGRPC string = `{
	"serviceName": "",
	"multiMode": false
}`
