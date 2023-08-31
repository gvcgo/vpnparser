package sing

import (
	"fmt"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/util/gconv"
	"github.com/moqsien/vpnparser/pkgs/parser"
	"github.com/moqsien/vpnparser/pkgs/utils"
)

/*
http://sing-box.sagernet.org/zh/configuration/outbound/vmess/
{
  "type": "vmess",
  "tag": "vmess-out",

  "server": "127.0.0.1", # 必填
  "server_port": 1080, # 必填
  "uuid": "bf000d23-0752-40b4-affe-68f7707a9661", # 必填
  "security": "auto",
  "alter_id": 0,
  "global_padding": false,
  "authenticated_length": true,
  "network": "tcp",
  "tls": {},
  "packet_encoding": "",
  "multiplex": {},
  "transport": {},

  ... // 拨号字段
}

Security:
auto
none
zero
aes-128-gcm
chacha20-poly1305
aes-128-ctr

alter_id:
0	禁用旧协议
1	启用旧协议
>1	未使用, 行为同 1

packet_encoding:
(空)		禁用
packetaddr	由 v2ray 5+ 支持
xudp		由 xray 支持

network:
启用的网络协议, tcp 或 udp。默认所有。
*/

var SingVmess string = `{
	"type": "vmess",
	"tag": "vmess-out",
	"server": "127.0.0.1",
	"server_port": 1080,
	"uuid": "bf000d23-0752-40b4-affe-68f7707a9661",
	"security": "auto",
	"alter_id": 0
`

type SVmessOut struct {
	RawUri   string
	Parser   *parser.ParserVmess
	outbound string
}

func (that *SVmessOut) Parse(rawUri string) {
	that.Parser = &parser.ParserVmess{}
	that.Parser.Parse(rawUri)
}

func (that *SVmessOut) Addr() string {
	if that.Parser == nil {
		return ""
	}
	return that.Parser.GetAddr()
}

func (that *SVmessOut) Port() int {
	if that.Parser == nil {
		return 0
	}
	return that.Parser.GetPort()
}

func (that *SVmessOut) Scheme() string {
	return parser.SchemeVmess
}

func (that *SVmessOut) getSettings() string {
	if that.Parser.Address == "" || that.Parser.Port == 0 {
		return "{}"
	}
	j := gjson.New(SingVmess)
	j.Set("type", "vmess")
	j.Set("server", that.Parser.Address)
	j.Set("server_port", that.Parser.Port)
	j.Set("uuid", that.Parser.UUID)
	j.Set("alter_id", gconv.Int(that.Parser.AID))
	if that.Parser.Security == "" {
		that.Parser.Security = "none"
	}
	j.Set("security", that.Parser.Security)
	if that.Parser.PacketEncoding != "" {
		j.Set("packet_encoding", that.Parser.PacketEncoding)
	}
	j.Set("tag", utils.OutboundTag)
	return j.MustToJsonString()
}

func (that *SVmessOut) GetOutboundStr() string {
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

func TestVmess() {
	// rawUri := "vmess://{\"v\": \"2\", \"ps\": \"13|西班牙 02 | 1x ES\", \"add\": \"2d3e6s01.mcfront.xyz\", \"port\": \"31884\", \"aid\": 0, \"scy\": \"auto\", \"net\": \"tcp\", \"type\": \"none\", \"tls\": \"tls\", \"id\": \"82a934c7-d98d-4e08-b63f-827b132d42ac\", \"sni\": \"es04.lovemc.xyz\"}"
	rawUri := "vmess://{\"add\":\"bobbykotick.rip\",\"host\":\"Kansas.bobbykotick.rip\",\"sni\":\"Kansas.bobbykotick.rip\",\"id\":\"D213ED80-199B-4A01-9D62-BBCBA9C16226\",\"net\":\"ws\",\"path\":\"\\/speedtest\",\"port\":\"443\",\"ps\":\"GetAFreeNode.com-Kansas\",\"tls\":\"tls\",\"fp\":\"android\",\"alpn\":\"h2,http\\/1.1\",\"v\":2,\"aid\":0,\"type\":\"none\"}"
	vo := &SVmessOut{}
	vo.Parse(rawUri)
	o := vo.GetOutboundStr()
	j := gjson.New(o)
	fmt.Println(j.MustToJsonIndentString())
}
