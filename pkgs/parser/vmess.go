package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gogf/gf/encoding/gjson"
)

/*
vmess: ['v', 'ps', 'add', 'port', 'aid', 'scy', 'net', 'type', 'tls', 'id', 'sni', 'host', 'path', 'alpn', 'security', 'skip-cert-verify', 'fp', 'test_name', 'serverPort', 'nation']
*/

type ParserVmess struct {
	Address  string
	Port     int
	UUID     string
	Security string
	AID      string

	Nation         string
	Path           string
	PS             string
	ServerPort     string
	SkipCertVerify bool
	TestName       string
	V              string

	*StreamField
}

func (that *ParserVmess) Parse(rawUri string) {
	r := strings.ReplaceAll(rawUri, SchemeVmess, "")
	j := gjson.New(r)
	if j == nil {
		return
	}
	that.Address = j.GetString("add")
	if !strings.Contains(that.Address, ".") {
		that.Address = ""
		return
	}
	that.Port = j.GetInt("port")
	that.UUID = j.GetString("id")
	that.AID = j.GetString("aid")
	that.Security = j.GetString("security")
	that.Security = j.GetString("security")
	if that.Security == "" {
		that.Security = j.GetString("scy")
	}

	that.Nation = j.GetString("nation")
	that.PS = j.GetString("ps")
	that.ServerPort = j.GetString("serverPort")
	that.SkipCertVerify = j.GetBool("skip-cert-verify")
	that.TestName = j.GetString("test_name")
	that.V = j.GetString("v")

	that.StreamField = &StreamField{}
	that.StreamField.Network = j.GetString("net")
	that.StreamField.StreamSecurity = j.GetString("tls")
	that.StreamField.Path = j.GetString("path")
	that.StreamField.Host = j.GetString("host")
	// that.StreamField.GRPCServiceName = j.GetString("serviceName")
	// that.StreamField.GRPCMultiMode = j.GetString("mode")
	that.StreamField.ServerName = j.GetString("sni")
	that.StreamField.TCPHeaderType = j.GetString("type")
	that.StreamField.TLSALPN = j.GetString("alpn")
	that.StreamField.Fingerprint = j.GetString("fp")

	// that.StreamField.RealityShortId = j.GetString("sid")
	// that.StreamField.RealitySpiderX = j.GetString("spx")
	// that.StreamField.RealityPublicKey = j.GetString("pbk")
}

func (that *ParserVmess) GetAddr() string {
	return that.Address
}

func (that *ParserVmess) GetPort() int {
	return that.Port
}

func (that *ParserVmess) Show() {
	fmt.Printf("addr: %s, port: %v, uuid: %s, net: %s", that.Address, that.Port, that.UUID, that.Network)
}

func VmessTest() {
	type Vmess struct {
		Vmess []string `json:"Vmess"`
	}

	v := &Vmess{}
	content, _ := os.ReadFile(`C:\Users\moqsien\data\projects\go\src\vpnparser\misc\vmess.json`)
	json.Unmarshal(content, v)
	for _, rawUri := range v.Vmess {
		p := &ParserVmess{}
		p.Parse(rawUri)
		if p.TLSALPN != "" {
			fmt.Println(p.TLSALPN)
		}
	}
}
