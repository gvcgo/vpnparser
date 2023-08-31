package xray

import (
	"fmt"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/moqsien/vpnparser/pkgs/parser"
	"github.com/moqsien/vpnparser/pkgs/utils"
)

/*
https://xtls.github.io/config/outbounds/trojan.html#outboundconfigurationobject

{
	"servers": [
	  {
		"address": "127.0.0.1",
		"port": 1234,
		"password": "password",
		"email": "love@xray.com",
		"level": 0
	  }
	]
}
*/

var XrayTrojan string = `{
	"servers": [
	  {
		"address": "127.0.0.1",
		"port": 1234,
		"password": "password"
	  }
	]
}`

type TrojanOut struct {
	RawUri   string
	Parser   *parser.ParserTrojan
	outbound string
}

func (that *TrojanOut) Parse(rawUri string) {
	that.Parser = &parser.ParserTrojan{}
	that.Parser.Parse(rawUri)
}

func (that *TrojanOut) Addr() string {
	if that.Parser == nil {
		return ""
	}
	return that.Parser.GetAddr()
}

func (that *TrojanOut) Port() int {
	if that.Parser == nil {
		return 0
	}
	return that.Parser.GetPort()
}

func (that *TrojanOut) Scheme() string {
	return parser.SchemeTrojan
}

func (that *TrojanOut) GetRawUri() string {
	return that.RawUri
}

func (that *TrojanOut) getSettings() string {
	j := gjson.New(XrayTrojan)
	j.Set("servers.0.address", that.Parser.Address)
	j.Set("servers.0.port", that.Parser.Port)
	j.Set("servers.0.password", that.Parser.Password)
	return j.MustToJsonString()
}

func (that *TrojanOut) setProtocolAndTag(outStr string) string {
	j := gjson.New(outStr)
	j.Set("protocol", "trojan")
	j.Set("tag", utils.OutboundTag)
	return j.MustToJsonString()
}

func (that *TrojanOut) GetOutboundStr() string {
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

func TestTrojan() {
	// rawUri := "trojan://2a898bd8-c0d1-4f7d-a88e-831d5682a9b9@hk02.isddns.tk:65527?allowInsecure=0\u0026peer=hk02.isddns.tk\u0026sni=hk02.isddns.tk#RPD|www.zyw.asia ZYW免费节点"
	// rawUri := "trojan://da88864b-6aa5-4d18-8e36-ac809a24c571@uk.stablize.top:443?allowInsecure=1#8DKJ|@Zyw_Channel"
	rawUri := "trojan://4d706727-996f-4427-930d-60f3bd414cf9@cnamemk.ciscocdn1.live:443?type=ws\u0026sni=c2mk.ciscocdn1.live\u0026allowInsecure=1\u0026path=/rDCYQta83d0oPABKBhcIX#🇺🇸_US_美国-\u003e🇵🇱_PL_波兰"
	to := &TrojanOut{}
	to.Parse(rawUri)
	o := to.GetOutboundStr()
	j := gjson.New(o)
	fmt.Println(j.MustToJsonIndentString())
	fmt.Println(o)
}
