package xray

import (
	"fmt"

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
	return ""
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
