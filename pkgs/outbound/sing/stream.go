package sing

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gvcgo/vpnparser/pkgs/parser"
	"github.com/gvcgo/vpnparser/pkgs/utils"
)

/*
Transport:
http://sing-box.sagernet.org/zh/configuration/shared/v2ray-transport/

HTTP:
{
  "type": "http",
  "host": [],
  "path": "",
  "method": "",
  "headers": {},
  "idle_timeout": "15s",
  "ping_timeout": "15s"
}

WebSocket:
{
  "type": "ws",
  "path": "",
  "headers": {},
  "max_early_data": 0,
  "early_data_header_name": ""
}

GRPC:
{
  "type": "grpc",
  "service_name": "TunService",
  "idle_timeout": "15s",
  "ping_timeout": "15s",
  "permit_without_stream": false
}

QUIC:
{
  "type": "quic"
}

TLS:
http://sing-box.sagernet.org/zh/configuration/shared/tls/

{
  "enabled": true,
  "disable_sni": false,
  "server_name": "",
  "insecure": false,
  "alpn": [],
  "min_version": "",
  "max_version": "",
  "cipher_suites": [],
  "certificate": "",
  "certificate_path": "",
  "ech": {
    "enabled": false,
    "pq_signature_schemes_enabled": false,
    "dynamic_record_sizing_disabled": false,
    "config": ""
  },
  "utls": {
    "enabled": false,
    "fingerprint": ""
  },
  "reality": {
    "enabled": false,
    "public_key": "jNXHt1yRo0vDuchQlIP6Z0ZvjT3KtzVI-T4E7RoLJS0",
    "short_id": "0123456789abcdef"
  }
}
*/

var SingHTTPandTCP string = `{
	"type": "http",
	"host": [],
	"path": ""
}`

var SingHTTPHeaders string = `{
	"Host": []
}`

var SingWebSocket string = `{
	"type": "ws",
	"path": ""
}`

var SingWebsocketHeaders string = `{
	"Host": ""
}`

var SingGRPC string = `{
	"type": "grpc",
	"service_name": ""
}`

var SingTLS string = `{
	"enabled": true,
	"disable_sni": false,
	"server_name": "",
	"insecure": false,
  }`

var SinguTLS string = `{
	"enabled": false,
	"fingerprint": ""
}`

var SingReality string = `{
	"enabled": false,
	"public_key": "",
	"short_id": ""
}`

func PrepareStreamStr(cnf *gjson.Json, sf *parser.StreamField) (result *gjson.Json) {
	var tp string
	switch sf.Network {
	case "tcp", "http":
		j := gjson.New(SingHTTPandTCP)
		host := sf.Host
		if host == "" {
			host = sf.ServerName
		}
		if host != "" {
			j.Set("host.0", host)
			h := gjson.New(SingHTTPHeaders)
			h.Set("Host.0", host)
			j = utils.SetJsonObjectByString("headers", h.MustToJsonString(), j)
		}
		SetPathForSingBoxTransport(sf.Path, j)
		// j.Set("path", sf.Path)
		tp = j.MustToJsonString()
	case "ws":
		j := gjson.New(SingWebSocket)
		host := sf.Host
		if host == "" {
			host = sf.ServerName
		}
		if host != "" {
			h := gjson.New(SingHTTPHeaders)
			h.Set("Host", host)
			j = utils.SetJsonObjectByString("headers", h.MustToJsonString(), j)
		}
		if sf.Path == "" {
			sf.Path = "/"
		}
		SetPathForSingBoxTransport(sf.Path, j)
		// j.Set("path", sf.Path)
		tp = j.MustToJsonString()
	case "grpc":
		j := gjson.New(SingGRPC)
		j.Set("service_name", sf.GRPCServiceName)
		tp = j.MustToJsonString()
	default:
		tp = "{}"
	}
	cnf = utils.SetJsonObjectByString("transport", tp, cnf)

	var tlsStr string
	switch sf.StreamSecurity {
	case "tls", "reality":
		j := gjson.New(SingTLS)
		if sf.ServerName == "" {
			sf.ServerName = sf.Host
		}
		j.Set("server_name", sf.ServerName)
		allowInsecure := false
		if sf.TLSAllowInsecure != "" {
			allowInsecure = gconv.Bool(sf.TLSAllowInsecure)
		}
		j.Set("enabled", true)
		j.Set("insecure", allowInsecure)
		if sf.TLSALPN != "" {
			j.Set("alpn", strings.Split(sf.TLSALPN, ","))
		}
		if sf.Fingerprint != "" {
			utls := gjson.New(SinguTLS)
			utls.Set("enabled", true)
			utls.Set("fingerprint", sf.Fingerprint)
			j = utils.SetJsonObjectByString("utls", utls.MustToJsonString(), j)
		}

		if sf.RealityShortId != "" && sf.RealityPublicKey != "" {
			reality := gjson.New(SingReality)
			reality.Set("short_id", sf.RealityShortId)
			reality.Set("public_key", sf.RealityPublicKey)
			reality.Set("enabled", true)
			j = utils.SetJsonObjectByString("reality", reality.MustToJsonString(), j)
		}
		tlsStr = j.MustToJsonString()
		// fmt.Println(tlsStr)
	default:
		tlsStr = `{"enabled": false}`
	}
	result = utils.SetJsonObjectByString("tls", tlsStr, cnf)
	return
}

func SetPathForSingBoxTransport(pathStr string, j *gjson.Json) {
	if u := ParseSingBoxPathToURL(pathStr); u != nil {
		if uPath := u.Path; uPath != "" {
			j.Set("path", uPath)
		}
		if ed, err := strconv.Atoi(u.Query().Get("ed")); err == nil && ed > 0 {
			j.Set("max_early_data", ed)
			j.Set("early_data_header_name", "Sec-WebSocket-Protocol")
		}
	}
}

func ParseSingBoxPathToURL(pathStr string) (result *url.URL) {
	if pathStr == "" {
		return
	}
	if strings.HasPrefix(pathStr, "/") {
		pathStr = "http://www.test.com" + pathStr
	}
	result, _ = url.Parse(pathStr)
	return
}
