package parser

import (
	"fmt"
	"strings"

	"encoding/json"

	"github.com/gvcgo/goutils/pkgs/gtui"
)

/*
PrivateKey string   `koanf,json:"private_key"`
AddrV4     string   `koanf,json:"addr_v4"`
AddrV6     string   `koanf,json:"addr_v6"`
DNS        string   `koanf,json:"dns"`
MTU        int      `koanf,json:"mtu"`
PublicKey  string   `koanf,json:"public_key"`
AllowedIPs []string `koanf,json:"allowed_ips"`
Endpoint   string   `koanf,json:"endpoint"`
ClientID   string   `koanf,json:"client_id"`
DeviceName string   `koanf,json:"device_name"`
Reserved   []int    `koanf,json:"reserved"`
*/

type ParserWirguard struct {
	PrivateKey string   `koanf,json:"private_key"`
	AddrV4     string   `koanf,json:"addr_v4"`
	AddrV6     string   `koanf,json:"addr_v6"`
	DNS        string   `koanf,json:"dns"`
	MTU        int      `koanf,json:"mtu"`
	PublicKey  string   `koanf,json:"public_key"`
	AllowedIPs []string `koanf,json:"allowed_ips"`
	Endpoint   string   `koanf,json:"endpoint"`
	ClientID   string   `koanf,json:"client_id"`
	DeviceName string   `koanf,json:"device_name"`
	Reserved   []int    `koanf,json:"reserved"`
	Address    string   `koanf,json:"address"`
	Port       int      `koanf,json:"port"`
}

func (that *ParserWirguard) Parse(rawUri string) {
	if strings.Contains(rawUri, SchemeWireguard) {
		rawUri = strings.ReplaceAll(rawUri, SchemeWireguard, "")
	}
	if err := json.Unmarshal([]byte(rawUri), that); err != nil {
		gtui.PrintError(err)
	}
}

func (that *ParserWirguard) GetAddr() string {
	return that.Address
}

func (that *ParserWirguard) GetPort() int {
	return that.Port
}

func (that *ParserWirguard) Show() {
	fmt.Printf("addr: %s, port: %d, privateKey: %s, publicKey: %s\n",
		that.Address,
		that.Port,
		that.PrivateKey,
		that.PublicKey,
	)
}

func TestWireguard() {
	rawUri := `wireguard://{"PrivateKey":"2B8LLjlXkJ608ct0LD0UnuuR9A2GuZUFBMBQJ9GFn1I=","AddrV4":"172.16.0.2","AddrV6":"2606:4700:110:8dad:87b4:b141:584d:e9dc","DNS":"1.1.1.1","MTU":1280,"PublicKey":"bmXOC+F1FxEMF9dyiK2H5/1SUtzH0JuVo51h2wPfgyo=","AllowedIPs":["0.0.0.0/0","::/0"],"Endpoint":"198.41.222.233:2087","ClientID":"GpxH","DeviceName":"D9D669","Reserved":null,"Address":"198.41.222.233","Port":2087}`
	p := &ParserWirguard{}
	p.Parse(rawUri)
	p.Show()
}
