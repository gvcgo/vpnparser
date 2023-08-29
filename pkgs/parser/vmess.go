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
	Address string
	Port    int
	UUID    string

	AID            string
	ALPN           string
	FP             string
	Host           string
	Net            string
	Nation         string
	Path           string
	PS             string
	SCY            string
	Security       string
	ServerPort     string
	SkipCertVerify bool
	SNI            string
	TLS            string
	Type           string
	TestName       string
	V              string
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
	that.ALPN = j.GetString("alpn")
	that.FP = j.GetString("fp")
	that.Host = j.GetString("host")
	that.Net = j.GetString("net")
	that.Nation = j.GetString("nation")
	that.Path = j.GetString("path")
	that.PS = j.GetString("ps")
	that.SCY = j.GetString("scy")
	that.Security = j.GetString("security")
	that.ServerPort = j.GetString("serverPort")
	that.SkipCertVerify = j.GetBool("skip-cert-verify")
	that.SNI = j.GetString("sni")
	that.TLS = j.GetString("tls")
	that.Type = j.GetString("type")
	that.TestName = j.GetString("test_name")
	that.V = j.GetString("v")
}

func (that *ParserVmess) GetAddr() string {
	if that.Address == "" {
		return ""
	}
	return that.Address
}

func (that *ParserVmess) GetHost() string {
	if that.Address == "" || that.Port == 0 {
		return ""
	}
	return fmt.Sprintf("%s:%d", that.Address, that.Port)
}

func (that *ParserVmess) Show() {
	fmt.Printf("addr: %s, port: %v, uuid: %s, net: %s", that.Address, that.Port, that.UUID, that.Net)
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
		if p.ALPN != "" {
			fmt.Println(p.ALPN)
		}
	}
}
