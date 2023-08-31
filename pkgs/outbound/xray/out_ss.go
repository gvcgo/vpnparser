package xray

import (
	"fmt"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/moqsien/vpnparser/pkgs/parser"
	"github.com/moqsien/vpnparser/pkgs/utils"
)

/*
https://xtls.github.io/config/outbounds/shadowsocks.html#serverobject

{
	"servers": [
	  {
		"email": "love@xray.com",
		"address": "127.0.0.1",
		"port": 1234,
		"method": "加密方式",
		"password": "密码",
		"uot": true,
		"UoTVersion": 2,
		"level": 0
	  }
	]
}

Method:
2022-blake3-aes-128-gcm
2022-blake3-aes-256-gcm
2022-blake3-chacha20-poly1305
aes-256-gcm
aes-128-gcm
chacha20-poly1305 或称 chacha20-ietf-poly1305
xchacha20-poly1305 或称 xchacha20-ietf-poly1305
none 或 plain

UoTVersion:
UDP over TCP 的实现版本。
当前可选值：1, 2

*/

var XraySS string = `{
	"servers": [
	  {
		"address": "127.0.0.1",
		"port": 1234,
		"method": "加密方式",
		"password": "密码"
	  }
	]
}`

type ShadowSocksOut struct {
	RawUri   string
	Parser   *parser.ParserSS
	outbound string
}

func (that *ShadowSocksOut) Parse(rawUri string) {
	that.Parser = &parser.ParserSS{}
	that.Parser.Parse(rawUri)
}

func (that *ShadowSocksOut) Addr() string {
	if that.Parser == nil {
		return ""
	}
	return that.Parser.GetAddr()
}

func (that *ShadowSocksOut) Port() int {
	if that.Parser == nil {
		return 0
	}
	return that.Parser.GetPort()
}

func (that *ShadowSocksOut) Scheme() string {
	return parser.SchemeSS
}

func (that *ShadowSocksOut) GetRawUri() string {
	return that.RawUri
}

func (that *ShadowSocksOut) getSettings() string {
	j := gjson.New(XraySS)
	j.Set("servers.0.address", that.Parser.Address)
	j.Set("servers.0.port", that.Parser.Port)
	j.Set("servers.0.method", that.Parser.Method)
	j.Set("servers.0.password", that.Parser.Password)
	return j.MustToJsonString()
}

func (that *ShadowSocksOut) setProtocolAndTag(outStr string) string {
	j := gjson.New(outStr)
	j.Set("protocol", "shadowsocks")
	j.Set("tag", utils.OutboundTag)
	return j.MustToJsonString()
}

func (that *ShadowSocksOut) GetOutboundStr() string {
	if that.Parser.Address == "" && that.Parser.Port == 0 {
		return ""
	}
	if that.outbound == "" {
		settings := that.getSettings()
		stream := PrepareStreamString(that.Parser.StreamField)
		outStr := fmt.Sprintf(XrayOut, settings, stream)
		that.outbound = that.setProtocolAndTag(outStr)
	}
	return that.outbound
}

func TestSS() {
	rawUri := "ss://aes-256-gcm:bad5fba5-a7bc-4709-882b-e15edad16cef@ah-cmi-1000m.ikun666.club:18878#🇨🇳_CN_中国-\u003e🇸🇬_SG_新加坡"
	// rawUri := "ss://aes-128-gcm:g12sQi#ss#\u00261@183.232.170.32:20013?plugin=v2ray-plugin\u0026mode=websocket\u0026mux=undefined#🇨🇳_CN_中国-\u003e🇯🇵_JP_日本"
	// rawUri := "ss://chacha20-ietf-poly1305:t0srmdxrm3xyjnvqz9ewlxb2myq7rjuv@4e168c3.h4.gladns.com:2377/?plugin=obfs-local\u0026obfs=tls\u0026obfs-host=(TG@WangCai_1)a83679f:53325#8DKJ|@Zyw_Channel"
	sso := &ShadowSocksOut{}
	sso.Parse(rawUri)
	o := sso.GetOutboundStr()
	j := gjson.New(o)
	fmt.Println(j.MustToJsonIndentString())
	fmt.Println(o)
}
