package sing

import (
	"fmt"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/moqsien/vpnparser/pkgs/parser"
	"github.com/moqsien/vpnparser/pkgs/utils"
)

/*
http://sing-box.sagernet.org/zh/configuration/outbound/vless/

{
  "type": "vless",
  "tag": "vless-out",

  "server": "127.0.0.1", # 必填
  "server_port": 1080, # 必填
  "uuid": "bf000d23-0752-40b4-affe-68f7707a9661", # 必填
  "flow": "xtls-rprx-vision",
  "network": "tcp",
  "tls": {},
  "packet_encoding": "",
  "transport": {},

  ... // 拨号字段
}

*/

var SingVless string = `{
	"type": "vless",
	"tag": "vless-out",
	"server": "127.0.0.1",
	"server_port": 1080,
	"uuid": "bf000d23-0752-40b4-affe-68f7707a9661",
	"flow": ""
}`

type SVlessOut struct {
	RawUri   string
	Parser   *parser.ParserVless
	outbound string
}

func (that *SVlessOut) Parse(rawUri string) {
	that.Parser = &parser.ParserVless{}
	that.Parser.Parse(rawUri)
}

func (that *SVlessOut) Addr() string {
	if that.Parser == nil {
		return ""
	}
	return that.Parser.GetAddr()
}

func (that *SVlessOut) Port() int {
	if that.Parser == nil {
		return 0
	}
	return that.Parser.GetPort()
}

func (that *SVlessOut) Scheme() string {
	return parser.SchemeVless
}

func (that *SVlessOut) GetRawUri() string {
	return that.RawUri
}

func (that *SVlessOut) getSettings() string {
	if that.Parser.Address == "" || that.Parser.Port == 0 {
		return "{}"
	}
	j := gjson.New(SingVless)
	j.Set("type", "vless")
	j.Set("server", that.Parser.Address)
	j.Set("server_port", that.Parser.Port)
	j.Set("uuid", that.Parser.UUID)
	j.Set("flow", that.Parser.Flow)
	if that.Parser.PacketEncoding != "" {
		j.Set("packet_encoding", that.Parser.PacketEncoding)
	}
	j.Set("tag", utils.OutboundTag)
	return j.MustToJsonString()
}

func (that *SVlessOut) GetOutboundStr() string {
	if that.outbound == "" {
		settings := that.getSettings()
		if settings == "{}" {
			return ""
		}
		cnf := gjson.New(settings)
		cnf = PrepareStreamStr(cnf, that.Parser.StreamField)
		that.outbound = cnf.MustToJsonString()
	}
	return that.outbound
}

func TestVless() {
	// rawUri := "vless://f0f4eabc-2747-4656-99b5-81ab6d32a8ab@172.67.33.254:443?alpn=http/1.1\u0026headerType=ws\u0026host=hct2.jensenk.cf\u0026path=/f0f4eabc-2747-4656-99b5-81ab6d32a8ab-vless\u0026security=tls\u0026sni=hct2.jensenk.cf\u0026type=ws#美国_08281722"
	// rawUri := "vless://882b8757-9244-404b-fee6-9ec7c3d8fd82@b2.v2parsin.site:17407?encryption=none\u0026security=none\u0026type=tcp\u0026headerType=http\u0026host=telewebion.com#德国_0828093"
	// rawUri := "vless://9890111b-a139-4a87-89d5-b9b18dd05e46@mci-shhproxy.ddns.net:8443?encryption=none\u0026security=tls\u0026sni=dl.SpV2ray.cfd\u0026type=grpc\u0026serviceName=@Shh_Proxy\u0026mode=gun#中国_0828245"
	rawUri := "vless://d572752d-b079-4169-a1a1-3da5721a8ab4@m2rel.siasepid.sbs:80?encryption=none\u0026security=reality\u0026sni=tgju.org\u0026fp=firefox\u0026pbk=HgrpXJzQo2liQMY9YAPq1_PuiDXNNBLx8hRyVVfUZko\u0026sid=af41f983\u0026spx=/\u0026type=grpc\u0026serviceName=@V2rayNGmat\u0026mode=multi#德国_0828096"
	// rawUri := "vless://c07fef7d-e8d2-42fe-b977-50e368f18293@104.17.36.178:443?flow=xtls-rprx-origin\u0026encryption=none\u0026security=tls\u0026sni=vincent-jackson2021.ga\u0026type=ws\u0026host=vincent-jackson2021.ga\u0026path=/The-Great-Awakening_ws#未知_0828876"
	vo := &SVlessOut{}
	vo.Parse(rawUri)
	o := vo.GetOutboundStr()
	j := gjson.New(o)
	fmt.Println(j.MustToJsonIndentString())
	fmt.Println(o)
}
