package sing

import (
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gvcgo/vpnparser/pkgs/parser"
	"github.com/gvcgo/vpnparser/pkgs/utils"
)

/*
http://sing-box.sagernet.org/zh/configuration/outbound/shadowsocks/

{
  "type": "shadowsocks",
  "tag": "ss-out",

  "server": "127.0.0.1",
  "server_port": 1080,
  "method": "2022-blake3-aes-128-gcm",
  "password": "8JCsPssfgS8tiRwiMlhARg==",
  "plugin": "",
  "plugin_opts": "",
  "network": "udp",
  "udp_over_tcp": false | {},
  "multiplex": {},

  ... // 拨号字段
}

Method:
2022-blake3-aes-128-gcm
2022-blake3-aes-256-gcm
2022-blake3-chacha20-poly1305
none
aes-128-gcm
aes-192-gcm
aes-256-gcm
chacha20-ietf-poly1305
xchacha20-ietf-poly1305
aes-128-ctr
aes-192-ctr
aes-256-ctr
aes-128-cfb
aes-192-cfb
aes-256-cfb
rc4-md5
chacha20-ietf
xchacha20
*/

var SingSS string = `{
	"type": "shadowsocks",
	"tag": "ss-out",
	"server": "127.0.0.1",
	"server_port": 1080,
	"method": "2022-blake3-aes-128-gcm",
	"password": "8JCsPssfgS8tiRwiMlhARg=="
}`

type SShadowSocksOut struct {
	RawUri   string
	Parser   *parser.ParserSS
	outbound string
}

func (that *SShadowSocksOut) Parse(rawUri string) {
	that.Parser = &parser.ParserSS{}
	that.Parser.Parse(rawUri)
}

func (that *SShadowSocksOut) Addr() string {
	if that.Parser == nil {
		return ""
	}
	return that.Parser.GetAddr()
}

func (that *SShadowSocksOut) Port() int {
	if that.Parser == nil {
		return 0
	}
	return that.Parser.GetPort()
}

func (that *SShadowSocksOut) Scheme() string {
	return parser.SchemeSS
}

func (that *SShadowSocksOut) GetRawUri() string {
	return that.RawUri
}

func (that *SShadowSocksOut) getSettings() string {
	if that.Parser.Address == "" || that.Parser.Port == 0 {
		return "{}"
	}
	j := gjson.New(SingSS)
	j.Set("type", "shadowsocks")
	j.Set("server", that.Parser.Address)
	j.Set("server_port", that.Parser.Port)
	j.Set("method", that.Parser.Method)
	j.Set("password", that.Parser.Password)
	j.Set("tag", utils.OutboundTag)

	if that.Parser.Plugin != "" {
		j.Set("plugin", that.Parser.Plugin)
	}

	if that.Parser.OBFS != "" && that.Parser.OBFSHost != "" {
		pluginOpts := []string{fmt.Sprintf("obfs=%s", that.Parser.OBFS), fmt.Sprintf("obfs-host=%s", that.Parser.OBFSHost)}
		j.Set("plugin_opts", strings.Join(pluginOpts, ";"))
	}

	vpluginOpts := []string{}
	if that.Parser.Mode != "" {
		vpluginOpts = append(vpluginOpts, fmt.Sprintf("mode=%s", that.Parser.Mode))
	}
	// if that.Parser.Mux != "" {
	// 	vpluginOpts = append(vpluginOpts, fmt.Sprintf("mux=%v", gconv.Bool(that.Parser.Mux)))
	// }
	if len(vpluginOpts) > 0 {
		j.Set("plugin_opts", strings.Join(vpluginOpts, ";"))
	}
	return j.MustToJsonString()
}

func (that *SShadowSocksOut) GetOutboundStr() string {
	if that.outbound == "" {
		settings := that.getSettings()
		if settings == "{}" {
			return ""
		}
		cnf := gjson.New(settings)
		// cnf = PrepareStreamStr(cnf, that.Parser.StreamField)
		that.outbound = cnf.MustToJsonString()
	}
	return that.outbound
}

func TestSS() {
	// rawUri := "ss://aes-256-gcm:bad5fba5-a7bc-4709-882b-e15edad16cef@ah-cmi-1000m.ikun666.club:18878#🇨🇳_CN_中国-\u003e🇸🇬_SG_新加坡"
	// rawUri := "ss://aes-128-gcm:g12sQi#ss#\u00261@183.232.170.32:20013?plugin=v2ray-plugin\u0026mode=websocket\u0026mux=undefined#🇨🇳_CN_中国-\u003e🇯🇵_JP_日本"
	rawUri := "ss://chacha20-ietf-poly1305:t0srmdxrm3xyjnvqz9ewlxb2myq7rjuv@4e168c3.h4.gladns.com:2377/?plugin=obfs-local\u0026obfs=tls\u0026obfs-host=(TG@WangCai_1)a83679f:53325#8DKJ|@Zyw_Channel"
	sso := &SShadowSocksOut{}
	sso.Parse(rawUri)
	o := sso.GetOutboundStr()
	j := gjson.New(o)
	fmt.Println(j.MustToJsonIndentString())
	fmt.Println(o)
}
