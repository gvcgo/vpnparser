package outbound

import (
	"fmt"
	"strings"

	json "github.com/bytedance/sonic"

	"github.com/moqsien/vpnparser/pkgs/parser"
	"github.com/moqsien/vpnparser/pkgs/utils"
)

var ShadowSocksMethodOnlyBySing = []string{
	"aes-256-cfb",
	"aes-128-ctr",
	"aes-192-ctr",
	"aes-256-ctr",
	"aes-128-cfb",
	"aes-192-cfb",
	"rc4-md5",
	"rc4",
	"chacha20-ietf",
	"xchacha20",
}

func EnableSingBox(rawUri string) bool {
	for _, m := range ShadowSocksMethodOnlyBySing {
		if strings.Contains(rawUri, m) {
			return true
		}
	}
	return false
}

type ProxyItem struct {
	Scheme       string     `json:"scheme"`
	Address      string     `json:"address"`
	Port         int        `json:"port"`
	RTT          int64      `json:"rtt"`
	RawUri       string     `json:"raw_uri"`
	Location     string     `json:"location"`
	Outbound     string     `json:"outbound"`
	OutboundType ClientType `json:"outbound_type"`
}

func NewItem(rawUri string) *ProxyItem {
	return &ProxyItem{RawUri: rawUri}
}

func NewItemByEncryptedRawUri(enRawUri string) (item *ProxyItem) {
	rawUri := parser.ParseRawUri(enRawUri)
	if rawUri == "" {
		return
	}
	return &ProxyItem{RawUri: rawUri}
}

func (that *ProxyItem) parse() bool {
	that.Scheme = utils.ParseScheme(that.RawUri)
	var ob IOutbound
	if that.Scheme == parser.SchemeSSR || (that.Scheme == parser.SchemeSS && strings.Contains(that.RawUri, "plugin=")) {
		that.OutboundType = SingBox
		ob = GetOutbound(SingBox, that.RawUri)
	} else if that.Scheme == parser.SchemeSS && EnableSingBox(that.RawUri) {
		that.OutboundType = SingBox
		ob = GetOutbound(SingBox, that.RawUri)
	} else if that.Scheme == parser.SchemeWireguard {
		that.OutboundType = SingBox
		ob = GetOutbound(SingBox, that.RawUri)
	} else {
		that.OutboundType = XrayCore
		ob = GetOutbound(XrayCore, that.RawUri)
	}

	if ob == nil {
		return false
	}
	ob.Parse(that.RawUri)
	that.Outbound = ob.GetOutboundStr()
	that.Address = ob.Addr()
	that.Port = ob.Port()
	return true
}

// Item string for conf.txt
func (that *ProxyItem) String() string {
	if ok := that.parse(); !ok {
		return ""
	}
	if r, err := json.Marshal(that); err == nil {
		return string(r)
	}
	return ""
}

func (that *ProxyItem) GetHost() string {
	if that.Address == "" && that.Port == 0 {
		return ""
	}
	return fmt.Sprintf("%s:%d", that.Address, that.Port)
}

func (that *ProxyItem) GetOutbound() string {
	if that.Outbound == "" {
		that.parse()
	}
	return that.Outbound
}

func (that *ProxyItem) GetOutboundType() ClientType {
	return that.OutboundType
}
