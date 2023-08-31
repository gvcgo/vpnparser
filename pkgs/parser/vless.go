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

	*StreamField
}

func (that *ParserVless) Parse(rawUri string) {
	if r, err := url.Parse(rawUri); err == nil {
		that.Address = r.Hostname()
		that.Port, _ = strconv.Atoi(r.Port())
		that.UUID = r.User.Username()
		query := r.Query()
		that.Encryption = query.Get("encryption")
		that.Flow = query.Get("flow")

		that.StreamField = &StreamField{
			Network:          query.Get("type"),
			StreamSecurity:   query.Get("security"),
			Path:             query.Get("path"),
			Host:             query.Get("host"),
			GRPCServiceName:  query.Get("serviceName"),
			GRPCMultiMode:    query.Get("mode"),
			ServerName:       query.Get("sni"),
			TLSALPN:          query.Get("alpn"),
			Fingerprint:      query.Get("fp"),
			RealityShortId:   query.Get("sid"),
			RealitySpiderX:   query.Get("spx"),
			RealityPublicKey: query.Get("pbk"),
			PacketEncoding:   query.Get("packetEncoding"),
			TCPHeaderType:    query.Get("headerType"),
		}
	}
}

func (that *ParserVless) GetAddr() string {
	return that.Address
}

func (that *ParserVless) GetPort() int {
	return that.Port
}

func (that *ParserVless) Show() {
	fmt.Printf("addr: %s, port: %v, uuid: %s\n",
		that.Address,
		that.Port,
		that.UUID)
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
			fmt.Println(p.Flow, p.Network)
			i++
		}
	}
	fmt.Println("total: ", i, len(v.Vless))
}
