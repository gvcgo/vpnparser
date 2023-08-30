package xray

import (
	"fmt"
	"strings"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/moqsien/vpnparser/pkgs/parser"
	"github.com/moqsien/vpnparser/pkgs/utils"
)

/*
https://xtls.github.io/config/outbounds/vless.html#outboundconfigurationobject

{
	"vnext": [
	  {
		"address": "example.com",
		"port": 443,
		"users": [
		  {
			"id": "5783a3e7-e373-51cd-8642-c83782b807c5",
			"encryption": "none",
			"flow": "xtls-rprx-vision",
			"level": 0
		  }
		]
	  }
	]
}

Flow:
流控模式，用于选择 XTLS 的算法。
目前出站协议中有以下流控模式可选：
无 flow，空字符或者 none：使用普通 TLS 代理
xtls-rprx-vision：使用新 XTLS 模式 包含内层握手随机填充 支持 uTLS 模拟客户端指纹
xtls-rprx-vision-udp443：同 xtls-rprx-vision, 但是放行了目标为 443 端口的 UDP 流量
此外，目前 XTLS 仅支持 TCP、mKCP、DomainSocket 这三种传输方式。

*/

var XrayVless string = `{
	"vnext": [
	  {
		"address": "example.com",
		"port": 443,
		"users": [
		  {
			"id": "5783a3e7-e373-51cd-8642-c83782b807c5",
			"encryption": "none",
			"flow": "xtls-rprx-vision"
		  }
		]
	  }
	]
}`

type VlessOut struct {
	RawUri   string
	Parser   *parser.ParserVless
	outbound string
}

func (that *VlessOut) Addr() string {
	if that.Parser == nil {
		return ""
	}
	return that.Parser.GetAddr()
}

func (that *VlessOut) Port() int {
	if that.Parser == nil {
		return 0
	}
	return that.Parser.GetPort()
}

func (that *VlessOut) Scheme() string {
	return parser.SchemeVless
}

func (that *VlessOut) getSettings() string {
	j := gjson.New(XrayVless)
	j.Set("vnext.0.address", that.Parser.Address)
	j.Set("vnext.0.port", that.Parser.Port)
	j.Set("vnext.0.users.0.id", that.Parser.UUID)
	j.Set("vnext.0.users.0.encryption", that.Parser.Encryption)
	j.Set("vnext.0.users.0.flow", that.Parser.Flow)
	return j.MustToJsonIndentString()
}

func (that *VlessOut) getStreamString() string {
	stream := gjson.New(XrayStream)
	stream.Set("network", that.Parser.Type)
	stream.Set("security", that.Parser.Security)
	host_ := that.Parser.Host
	if host_ == "" {
		host_ = that.Parser.SNI
	}
	switch that.Parser.Type {
	case "tcp":
		if that.Parser.HeaderType != "http" {
			stream = utils.SetJsonObjectByString("tcpSetting", XrayStreamTCPNone, stream)
		} else {
			j := gjson.New(XrayStreamTCPHTTP)
			j.Set("header.request.path.0", that.Parser.Path)
			j.Set("header.request.headers.Host.0", host_)
			stream = utils.SetJsonObjectByString("tcpSetting", j.MustToJsonIndentString(), stream)
		}
	case "ws":
		j := gjson.New(XrayStreamWebSocket)
		j.Set("path", that.Parser.Path)
		j.Set("headers.Host", host_)
		stream = utils.SetJsonObjectByString("wsSettings", j.MustToJsonIndentString(), stream)
	case "grpc":
		j := gjson.New(XrayStreamGRPC)
		j.Set("serviceName", that.Parser.ServiceName)
		multiMode := false
		if that.Parser.Mode == "multi" {
			multiMode = true
		}
		j.Set("multiMode", multiMode)
		stream = utils.SetJsonObjectByString("grpcSettings", j.MustToJsonIndentString(), stream)
	default:
		return "{}"
	}

	switch that.Parser.Security {
	case "tls":
		j := gjson.New(XrayStreamTLS)
		serverName := that.Parser.SNI
		if serverName == "" {
			serverName = that.Parser.Host
		}
		j.Set("serverName", serverName)
		if that.Parser.ALPN != "" {
			aList := strings.Split(that.Parser.ALPN, ",")
			j.Set("alpn", aList)
		}
		if that.Parser.FP != "" {
			j.Set("fingerprint", that.Parser.FP)
		}
		stream = utils.SetJsonObjectByString("tlsSettings", j.MustToJsonIndentString(), stream)
	case "reality":
		j := gjson.New(XrayStreamReality)
		serverName := that.Parser.SNI
		if serverName == "" {
			serverName = that.Parser.Host
		}
		j.Set("serverName", serverName)
		j.Set("shortId", that.Parser.SID)
		j.Set("fingerprint", that.Parser.FP)
		j.Set("spiderX", that.Parser.SPX)
		j.Set("publicKey", that.Parser.PBK)
		stream = utils.SetJsonObjectByString("tlsSettings", j.MustToJsonIndentString(), stream)
	default:
	}
	return stream.MustToJsonIndentString()
}

func (that *VlessOut) setProtocolAndTag(outStr string) string {
	j := gjson.New(outStr)
	j.Set("protocol", "vless")
	j.Set("tag", utils.OutboundTag)
	return j.MustToJsonIndentString()
}

func (that *VlessOut) GetOutboundStr() string {
	if that.outbound == "" {
		settings := that.getSettings()
		stream := that.getStreamString()
		outStr := fmt.Sprintf(XrayOut, settings, stream)
		that.outbound = that.setProtocolAndTag(outStr)
	}
	return that.outbound
}
