package xray

import (
	"fmt"
	"strings"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/util/gconv"
	"github.com/moqsien/vpnparser/pkgs/parser"
	"github.com/moqsien/vpnparser/pkgs/utils"
)

/*
https://xtls.github.io/config/outbounds/vmess.html#outboundconfigurationobject

{
	"vnext": [
	  {
		"address": "127.0.0.1",
		"port": 37192,
		"users": [
		  {
			"id": "5783a3e7-e373-51cd-8642-c83782b807c5",
			"security": "auto",
			"level": 0
		  }
		]
	  }
	]
}

Security:
"aes-128-gcm" | "chacha20-poly1305" | "auto" | "none" | "zero"
加密方式，客户端将使用配置的加密方式发送数据，服务器端自动识别，无需配置。

"aes-128-gcm"：推荐在 PC 上使用
"chacha20-poly1305"：推荐在手机端使用
"auto"：默认值，自动选择（运行框架为 AMD64、ARM64 或 s390x 时为 aes-128-gcm 加密方式，其他情况则为 Chacha20-Poly1305 加密方式）
"none"：不加密
"zero"：不加密，也不进行消息认证 (v1.4.0+)
提示:
推荐使用"auto"加密方式，这样可以永久保证安全性和兼容性。
"none" 伪加密方式会计算并验证数据包的校验数据，由于认证算法没有硬件支持，在部分平台可能速度比有硬件加速的 "aes-128-gcm" 还慢。
"zero" 伪加密方式不会加密消息也不会计算数据的校验数据，因此理论上速度会高于其他任何加密方式。实际速度可能受到其他因素影响。
不推荐在未开启 TLS 加密并强制校验证书的情况下使用 "none" "zero" 伪加密方式。 如果使用 CDN 或其他会解密 TLS 的中转平台或网络环境建立连接，不建议使用 "none" "zero" 伪加密方式。
无论使用哪种加密方式， VMess 的包头都会受到加密和认证的保护。

*/

var XrayVmessSettings string = `{
	"vnext": [
	  {
		"address": "127.0.0.1",
		"port": 37192,
		"users": [
		  {
			"id": "5783a3e7-e373-51cd-8642-c83782b807c5",
			"alterId": 0,
			"security": "auto"
		  }
		]
	  }
	]
}`

type VmessOut struct {
	RawUri   string
	Parser   *parser.ParserVmess
	outbound string
}

func (that *VmessOut) Parse(rawUri string) {
	that.Parser = &parser.ParserVmess{}
	that.Parser.Parse(rawUri)
}

func (that *VmessOut) Addr() string {
	if that.Parser == nil {
		return ""
	}
	return that.Parser.GetAddr()
}

func (that *VmessOut) Port() int {
	if that.Parser == nil {
		return 0
	}
	return that.Parser.GetPort()
}

func (that *VmessOut) Scheme() string {
	return parser.SchemeVmess
}

func (that *VmessOut) getSettings() string {
	if that.Parser.Address == "" || that.Parser.Port == 0 {
		return "{}"
	}
	j := gjson.New(XrayVmessSettings)
	j.Set("vnext.0.address", that.Parser.Address)
	j.Set("vnext.0.port", that.Parser.Port)
	j.Set("vnext.0.users.0.id", that.Parser.UUID)
	j.Set("vnext.0.users.0.alterId", gconv.Int(that.Parser.AID))
	security := that.Parser.Security
	if security == "" && that.Parser.SCY != "" {
		security = that.Parser.SCY
	}
	j.Set("vnext.0.users.0.security", security)
	return j.MustToJsonString()
}

func (that *VmessOut) getStreamString() string {
	stream := gjson.New(XrayStream)
	stream.Set("network", that.Parser.Net)
	stream.Set("security", that.Parser.TLS)
	switch that.Parser.Net {
	case "tcp":
		if that.Parser.Type != "http" {
			stream = utils.SetJsonObjectByString("tcpSetting", XrayStreamTCPNone, stream)
		} else {
			j := gjson.New(XrayStreamTCPHTTP)
			j.Set("header.request.path.0", that.Parser.Path)
			j.Set("header.request.headers.Host.0", that.Parser.Host)
			stream = utils.SetJsonObjectByString("tcpSetting", j.MustToJsonString(), stream)
		}
	case "ws":
		j := gjson.New(XrayStreamWebSocket)
		j.Set("path", that.Parser.Path)
		j.Set("headers.Host", that.Parser.Host)
		stream = utils.SetJsonObjectByString("wsSettings", j.MustToJsonString(), stream)
	case "http":
		j := gjson.New(XrayStreamHTTP)
		j.Set("host.0", that.Parser.Host)
		j.Set("path", that.Parser.Path)
		stream = utils.SetJsonObjectByString("httpSettings", j.MustToJsonString(), stream)
	case "grpc":
		j := gjson.New(XrayStreamGRPC)
		j.Set("serviceName", that.Parser.Host)
		stream = utils.SetJsonObjectByString("grpcSettings", j.MustToJsonString(), stream)
	default:
		return "{}"
	}
	if that.Parser.TLS != "" {
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
		stream = utils.SetJsonObjectByString("tlsSettings", j.MustToJsonString(), stream)
	}
	return stream.MustToJsonString()
}

func (that *VmessOut) setProtocolAndTag(outStr string) string {
	j := gjson.New(outStr)
	j.Set("protocol", "vmess")
	j.Set("tag", utils.OutboundTag)
	return j.MustToJsonIndentString()
}

func (that *VmessOut) GetOutboundStr() string {
	if that.outbound == "" {
		settings := that.getSettings()
		stream := that.getStreamString()
		outStr := fmt.Sprintf(XrayOut, settings, stream)
		that.outbound = that.setProtocolAndTag(outStr)
	}
	return that.outbound
}

func TestVmess() {
	// rawUri := "vmess://{\"v\": \"2\", \"ps\": \"13|西班牙 02 | 1x ES\", \"add\": \"2d3e6s01.mcfront.xyz\", \"port\": \"31884\", \"aid\": 0, \"scy\": \"auto\", \"net\": \"tcp\", \"type\": \"none\", \"tls\": \"tls\", \"id\": \"82a934c7-d98d-4e08-b63f-827b132d42ac\", \"sni\": \"es04.lovemc.xyz\"}"
	rawUri := "vmess://{\"add\":\"bobbykotick.rip\",\"host\":\"Kansas.bobbykotick.rip\",\"sni\":\"Kansas.bobbykotick.rip\",\"id\":\"D213ED80-199B-4A01-9D62-BBCBA9C16226\",\"net\":\"ws\",\"path\":\"\\/speedtest\",\"port\":\"443\",\"ps\":\"GetAFreeNode.com-Kansas\",\"tls\":\"tls\",\"fp\":\"android\",\"alpn\":\"h2,http\\/1.1\",\"v\":2,\"aid\":0,\"type\":\"none\"}"
	vo := &VmessOut{}
	vo.Parse(rawUri)
	o := vo.GetOutboundStr()
	j := gjson.New(o)
	fmt.Println(j.MustToJsonIndentString())
}
