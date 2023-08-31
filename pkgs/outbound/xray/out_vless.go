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

func (that *VlessOut) Parse(rawUri string) {
	that.Parser = &parser.ParserVless{}
	that.Parser.Parse(rawUri)
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
	// if that.Parser.PacketEncoding != "" {
	// 	j.Set("vnext.0.packetEncoding", that.Parser.PacketEncoding)
	// }
	return j.MustToJsonString()
}

// func (that *VlessOut) getStreamString() string {
// 	stream := gjson.New(XrayStream)
// 	stream.Set("network", that.Parser.Type)
// 	stream.Set("security", that.Parser.Security)
// 	host_ := that.Parser.Host
// 	if host_ == "" {
// 		host_ = that.Parser.SNI
// 	}
// 	switch that.Parser.Type {
// 	case "tcp":
// 		if that.Parser.HeaderType != "http" {
// 			stream = utils.SetJsonObjectByString("tcpSetting", XrayStreamTCPNone, stream)
// 		} else {
// 			j := gjson.New(XrayStreamTCPHTTP)
// 			j.Set("header.request.path.0", that.Parser.Path)
// 			j.Set("header.request.headers.Host.0", host_)
// 			stream = utils.SetJsonObjectByString("tcpSetting", j.MustToJsonString(), stream)
// 		}
// 	case "ws":
// 		j := gjson.New(XrayStreamWebSocket)
// 		j.Set("path", that.Parser.Path)
// 		j.Set("headers.Host", host_)
// 		stream = utils.SetJsonObjectByString("wsSettings", j.MustToJsonString(), stream)
// 	case "grpc":
// 		j := gjson.New(XrayStreamGRPC)
// 		j.Set("serviceName", that.Parser.ServiceName)
// 		multiMode := false
// 		if that.Parser.Mode == "multi" {
// 			multiMode = true
// 		}
// 		j.Set("multiMode", multiMode)
// 		stream = utils.SetJsonObjectByString("grpcSettings", j.MustToJsonString(), stream)
// 	default:
// 		return "{}"
// 	}

// 	switch that.Parser.Security {
// 	case "tls":
// 		j := gjson.New(XrayStreamTLS)
// 		serverName := that.Parser.SNI
// 		if serverName == "" {
// 			serverName = that.Parser.Host
// 		}
// 		j.Set("serverName", serverName)
// 		if that.Parser.ALPN != "" {
// 			aList := strings.Split(that.Parser.ALPN, ",")
// 			j.Set("alpn", aList)
// 		}
// 		if that.Parser.FP != "" {
// 			j.Set("fingerprint", that.Parser.FP)
// 		}
// 		stream = utils.SetJsonObjectByString("tlsSettings", j.MustToJsonString(), stream)
// 	case "reality":
// 		j := gjson.New(XrayStreamReality)
// 		serverName := that.Parser.SNI
// 		if serverName == "" {
// 			serverName = that.Parser.Host
// 		}
// 		j.Set("serverName", serverName)
// 		j.Set("shortId", that.Parser.SID)
// 		j.Set("fingerprint", that.Parser.FP)
// 		j.Set("spiderX", that.Parser.SPX)
// 		j.Set("publicKey", that.Parser.PBK)
// 		stream = utils.SetJsonObjectByString("realitySettings", j.MustToJsonString(), stream)
// 	default:
// 	}
// 	return stream.MustToJsonString()
// }

func (that *VlessOut) setProtocolAndTag(outStr string) string {
	j := gjson.New(outStr)
	j.Set("protocol", "vless")
	j.Set("tag", utils.OutboundTag)
	return j.MustToJsonString()
}

func (that *VlessOut) GetOutboundStr() string {
	if that.outbound == "" {
		settings := that.getSettings()
		stream := PrepareStreamString(that.Parser.StreamField)
		outStr := fmt.Sprintf(XrayOut, settings, stream)
		that.outbound = that.setProtocolAndTag(outStr)
	}
	return that.outbound
}

func TestVless() {
	rawUri := "vless://f0f4eabc-2747-4656-99b5-81ab6d32a8ab@172.67.33.254:443?alpn=http/1.1\u0026headerType=ws\u0026host=hct2.jensenk.cf\u0026path=/f0f4eabc-2747-4656-99b5-81ab6d32a8ab-vless\u0026security=tls\u0026sni=hct2.jensenk.cf\u0026type=ws#美国_08281722"
	// rawUri := "vless://882b8757-9244-404b-fee6-9ec7c3d8fd82@b2.v2parsin.site:17407?encryption=none\u0026security=none\u0026type=tcp\u0026headerType=http\u0026host=telewebion.com#德国_0828093"
	// rawUri := "vless://9890111b-a139-4a87-89d5-b9b18dd05e46@mci-shhproxy.ddns.net:8443?encryption=none\u0026security=tls\u0026sni=dl.SpV2ray.cfd\u0026type=grpc\u0026serviceName=@Shh_Proxy\u0026mode=gun#中国_0828245"
	// rawUri := "vless://d572752d-b079-4169-a1a1-3da5721a8ab4@m2rel.siasepid.sbs:80?encryption=none\u0026security=reality\u0026sni=tgju.org\u0026fp=firefox\u0026pbk=HgrpXJzQo2liQMY9YAPq1_PuiDXNNBLx8hRyVVfUZko\u0026sid=af41f983\u0026spx=/\u0026type=grpc\u0026serviceName=@V2rayNGmat\u0026mode=multi#德国_0828096"
	vo := &VlessOut{}
	vo.Parse(rawUri)
	o := vo.GetOutboundStr()
	j := gjson.New(o)
	fmt.Println(j.MustToJsonIndentString())
	fmt.Println(o)
}
