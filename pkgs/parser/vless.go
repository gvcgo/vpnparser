package parser

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strconv"
)

/*
vless: ['security', 'type', 'sni', 'path', 'encryption', 'headerType', 'packetEncoding', 'serviceName', 'mode', 'flow', 'alpn', 'host', 'fp', 'pbk', 'sid', 'spx']
*/

type ParserVless struct {
	Address    string
	Port       int
	UUID       string
	Encryption string
	Flow       string

	ALPN           string
	FP             string
	HeaderType     string
	Host           string
	Mode           string
	PacketEncoding string
	Path           string
	PBK            string
	Security       string
	ServiceName    string
	SID            string
	SNI            string
	SPX            string
	Type           string
}

func (that *ParserVless) Parse(rawUri string) {
	if r, err := url.Parse(rawUri); err == nil {
		that.Address = r.Hostname()
		that.Port, _ = strconv.Atoi(r.Port())
		that.UUID = r.User.Username()
		query := r.Query()
		that.Encryption = query.Get("encryption")
		that.Flow = query.Get("flow")

		that.ALPN = query.Get("alpn")
		that.FP = query.Get("fp")
		that.HeaderType = query.Get("headerType")
		that.Host = query.Get("host")
		that.Mode = query.Get("mode")
		that.PacketEncoding = query.Get("packetEncoding")
		that.Path = query.Get("path")
		that.PBK = query.Get("pbk")
		that.Security = query.Get("security")
		that.ServiceName = query.Get("serviceName")
		that.SID = query.Get("sid")
		that.SNI = query.Get("sni")
		that.SPX = query.Get("spx")
		that.Type = query.Get("type")
	}
}

func (that *ParserVless) GetAddr() string {
	return that.Address
}

func (that *ParserVless) GetPort() int {
	return that.Port
}

func (that *ParserVless) Show() {
	fmt.Printf("addr: %s, port: %v, uuid: %s, serviceName: %s", that.Address, that.Port, that.UUID, that.ServiceName)
}

func VlessTest() {
	type Vless struct {
		Vless []string `json:"Vless"`
	}

	v := &Vless{}
	content, _ := os.ReadFile(`C:\Users\moqsien\data\projects\go\src\vpnparser\misc\vless.json`)
	json.Unmarshal(content, v)
	i := 0
	for _, rawUri := range v.Vless {
		p := &ParserVless{}
		p.Parse(rawUri)
		if p.Address != "" {
			fmt.Println(p.Flow, p.Type)
			i++
		}
	}
	fmt.Println("total: ", i, len(v.Vless))
}
