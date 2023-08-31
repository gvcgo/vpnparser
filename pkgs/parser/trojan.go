package parser

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strconv"

	"github.com/moqsien/goutils/pkgs/gtui"
)

/*
trojan: ['allowInsecure', 'peer', 'sni', 'type', 'path', 'security', 'headerType']
*/

type ParserTrojan struct {
	Address  string
	Port     int
	Password string

	*StreamField
}

func (that *ParserTrojan) Parse(rawUri string) {
	if u, err := url.Parse(rawUri); err == nil {
		that.Address = u.Hostname()
		that.Port, _ = strconv.Atoi(u.Port())
		that.Password = u.User.Username()

		query := u.Query()

		that.StreamField = &StreamField{
			Network:          query.Get("type"),
			Host:             query.Get("peer"),
			Path:             query.Get("path"),
			StreamSecurity:   query.Get("security"),
			ServerName:       query.Get("sni"),
			TCPHeaderType:    query.Get("headerType"),
			TLSAllowInsecure: query.Get("allowInsecure"),
		}
	} else {
		gtui.PrintError(err)
		fmt.Println(rawUri)
		return
	}

	if that.StreamField.TLSAllowInsecure != "" && that.StreamField.ServerName != "" {
		if that.Network == "" {
			that.Network = "tcp"
		}
		if that.StreamSecurity == "" {
			that.StreamSecurity = "tls"
		}
		if that.Host == "" {
			that.Host = that.ServerName
		}
		if that.ServerName == "" {
			that.ServerName = that.Address
		}
	}
}

func (that *ParserTrojan) GetAddr() string {
	return that.Address
}

func (that *ParserTrojan) GetPort() int {
	return that.Port
}

func (that *ParserTrojan) Show() {
	fmt.Printf("addr: %s, port: %v, password: %s\n",
		that.Address,
		that.Port,
		that.Password)
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
		if p.StreamSecurity != "" {
			fmt.Println(p.Network, p.StreamSecurity, p.Path, p.ServerName, p.TCPHeaderType, p.TLSAllowInsecure)
		}
	}
	fmt.Println("total: ", i, len(t.Trojan))
}
