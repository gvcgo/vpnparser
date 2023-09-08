package parser

import (
	"fmt"

	json "github.com/bytedance/sonic"
	"github.com/moqsien/goutils/pkgs/gtui"
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
