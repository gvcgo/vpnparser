package sing

import (
	"fmt"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/moqsien/vpnparser/pkgs/parser"
	"github.com/moqsien/vpnparser/pkgs/utils"
)

var SingWireguard string = `{
	"type": "wireguard",
  	"tag": "wireguard-out",
	"server": "162.159.195.81", 
	"server_port": 928,
	"system_interface": false,
	"interface_name": "wg0",
	"local_address": [
		"172.16.0.2/32",
		"2606:4700:110:8bb9:68be:a130:cede:18bc/128" 
	],
	"private_key": "YNXtAzepDqRv9H52osJVDQnznT5AM11eCK3ESpwSt04=",
	"peer_public_key": "Z1XXLsKYkYxuiYjJIkRvtIKFepCYHTgON+GwPq7SOV4=",
	"mtu": 1280
}`

type SWireguardOut struct {
	RawUri   string
	Parser   *parser.ParserWirguard
	outbound string
}

func (that *SWireguardOut) Parse(rawUri string) {
	that.Parser = &parser.ParserWirguard{}
	that.Parser.Parse(rawUri)
}

func (that *SWireguardOut) Addr() string {
	if that.Parser == nil {
		return ""
	}
	return that.Parser.GetAddr()
}

func (that *SWireguardOut) Port() int {
	if that.Parser == nil {
		return 0
	}
	return that.Parser.GetPort()
}

func (that *SWireguardOut) Scheme() string {
	return parser.SchemeWireguard
}

func (that *SWireguardOut) GetRawUri() string {
	return that.RawUri
}

func (that *SWireguardOut) getSettings() string {
	if that.Parser.Address == "" || that.Parser.Port == 0 {
		return "{}"
	}
	j := gjson.New(SingWireguard)
	j.Set("tag", utils.OutboundTag)
	j.Set("server", that.Parser.Address)
	j.Set("server_port", that.Parser.Port)
	j.Set("interface_name", that.Parser.DeviceName)
	j.Set("local_address.0", fmt.Sprintf("%s/32", that.Parser.AddrV4))
	j.Set("local_address.1", fmt.Sprintf("%s/128", that.Parser.AddrV6))
	j.Set("private_key", that.Parser.PrivateKey)
	j.Set("peer_public_key", that.Parser.PublicKey)
	j.Set("mtu", that.Parser.MTU)
	return j.MustToJsonString()
}

func (that *SWireguardOut) GetOutboundStr() string {
	if that.outbound == "" {
		settings := that.getSettings()
		if settings == "{}" {
			return ""
		}
		that.outbound = settings
	}
	return that.outbound
}
