package outbound

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/moqsien/vpnparser/pkgs/parser"
	"github.com/moqsien/vpnparser/pkgs/utils"
)

type ProxyItem struct {
	Address      string     `json:"address"`
	Port         int        `json:"port"`
	RTT          int        `json:"rtt"`
	RawUri       string     `json:"raw_uri"`
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
	scheme := utils.ParseScheme(that.RawUri)
	var ob IOutbound
	if scheme == parser.SchemeSSR || (scheme == parser.SchemeSS && strings.Contains(that.RawUri, "plugin=")) {
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
