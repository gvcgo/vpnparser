package parser

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strconv"
)

/*
trojan: ['allowInsecure', 'peer', 'sni', 'type', 'path', 'security', 'headerType']
*/

type ParserTrojan struct {
	Address  string
	Port     int
	Password string

	AllowInsecure string
	HeaderType    string
	Path          string
	Peer          string
	Security      string
	SNI           string
	Type          string
}

func (that *ParserTrojan) Parse(rawUri string) {
	if u, err := url.Parse(rawUri); err == nil {
		that.Address = u.Hostname()
		that.Port, _ = strconv.Atoi(u.Port())
		that.Password = u.User.Username()

		query := u.Query()
		that.AllowInsecure = query.Get("allowInsecure")
		that.HeaderType = query.Get("headerType")
		that.Path = query.Get("path")
		that.Peer = query.Get("peer")
		that.Security = query.Get("security")
		that.SNI = query.Get("sni")
		that.Type = query.Get("type")
	}
}

func (that *ParserTrojan) GetAddr() string {
	return that.Address
}

func (that *ParserTrojan) GetPort() int {
	return that.Port
}

func (that *ParserTrojan) Show() {
	fmt.Printf("addr: %s, port: %v, password: %s", that.Address, that.Port, that.Password)
}

func TrojanTest() {
	type Trojan struct {
		Trojan []string `json:"Trojan"`
	}

	t := &Trojan{}
	content, _ := os.ReadFile(`C:\Users\moqsien\data\projects\go\src\vpnparser\misc\trojan.json`)
	json.Unmarshal(content, t)
	i := 0
	for _, rawUri := range t.Trojan {
		p := &ParserTrojan{}
		p.Parse(rawUri)
		if p.Address != "" {
			i++
		}
		if p.Security != "" {
			fmt.Println(p.Type, p.Security, p.Path, p.SNI, p.HeaderType, p.AllowInsecure)
		}
	}
	fmt.Println("total: ", i, len(t.Trojan))
}
