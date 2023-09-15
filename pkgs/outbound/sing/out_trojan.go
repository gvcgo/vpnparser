package sing

import (
	"fmt"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/moqsien/vpnparser/pkgs/parser"
	"github.com/moqsien/vpnparser/pkgs/utils"
)

/*
http://sing-box.sagernet.org/zh/configuration/outbound/trojan/
{
  "type": "trojan",
  "tag": "trojan-out",

  "server": "127.0.0.1", # 必填
  "server_port": 1080, # 必填
  "password": "8JCsPssfgS8tiRwiMlhARg==", # 必填
  "network": "tcp",
  "tls": {},
  "multiplex": {},
  "transport": {},

  ... // 拨号字段
}

*/

var SingTrojan string = `{
	"type": "trojan",
	"tag": "trojan-out",
	"server": "127.0.0.1",
	"server_port": 1080,
	"password": "8JCsPssfgS8tiRwiMlhARg=="
}`

type STrojanOut struct {
	RawUri   string
	Parser   *parser.ParserTrojan
	outbound string
}

func (that *STrojanOut) Parse(rawUri string) {
	that.Parser = &parser.ParserTrojan{}
	that.Parser.Parse(rawUri)
}

func (that *STrojanOut) Addr() string {
	if that.Parser == nil {
		return ""
	}
	return that.Parser.GetAddr()
}

func (that *STrojanOut) Port() int {
	if that.Parser == nil {
		return 0
	}
	return that.Parser.GetPort()
}

func (that *STrojanOut) Scheme() string {
	return parser.SchemeTrojan
}

func (that *STrojanOut) GetRawUri() string {
	return that.RawUri
}

func (that *STrojanOut) getSettings() string {
	if that.Parser.Address == "" || that.Parser.Port == 0 {
		return "{}"
	}
	j := gjson.New(SingTrojan)
	j.Set("type", "trojan")
	j.Set("server", that.Parser.Address)
	j.Set("server_port", that.Parser.Port)
	j.Set("password", that.Parser.Password)
	j.Set("tag", utils.OutboundTag)
	return j.MustToJsonString()
}

func (that *STrojanOut) GetOutboundStr() string {
	if that.outbound == "" {
		settings := that.getSettings()
		if settings == "{}" {
			return ""
		}
		cnf := gjson.New(settings)
		// trojan doesn't need to prepare stream settings
		// cnf = PrepareStreamStr(cnf, that.Parser.StreamField)
		that.outbound = cnf.MustToJsonString()
	}
	return that.outbound
}

func TestTrojan() {
	// rawUri := "trojan://2a898bd8-c0d1-4f7d-a88e-831d5682a9b9@hk02.isddns.tk:65527?allowInsecure=0\u0026peer=hk02.isddns.tk\u0026sni=hk02.isddns.tk#RPD|www.zyw.asia ZYW免费节点"
	// rawUri := "trojan://da88864b-6aa5-4d18-8e36-ac809a24c571@uk.stablize.top:443?allowInsecure=1#8DKJ|@Zyw_Channel"
	rawUri := "trojan://4d706727-996f-4427-930d-60f3bd414cf9@cnamemk.ciscocdn1.live:443?type=ws\u0026sni=c2mk.ciscocdn1.live\u0026allowInsecure=1\u0026path=/rDCYQta83d0oPABKBhcIX#🇺🇸_US_美国-\u003e🇵🇱_PL_波兰"
	to := &STrojanOut{}
	to.Parse(rawUri)
	o := to.GetOutboundStr()
	j := gjson.New(o)
	fmt.Println(j.MustToJsonIndentString())
	fmt.Println(o)
}
