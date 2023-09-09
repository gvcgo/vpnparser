package main

import (
	"fmt"

	"github.com/moqsien/vpnparser/pkgs/outbound"
	_ "github.com/moqsien/vpnparser/pkgs/outbound/sing"
	_ "github.com/moqsien/vpnparser/pkgs/outbound/xray"
	_ "github.com/moqsien/vpnparser/pkgs/parser"
)

func main() {
	// parser.VlessTest()
	// parser.TrojanTest()
	// parser.SSRTest()
	// parser.TestWireguard()

	// s := xray.GetPattern()
	// fmt.Println(s)
	// xray.TestVmess()
	// xray.TestTrojan()
	// xray.TestSS()

	// sing.TestVmess()
	// sing.TestVless()
	// sing.TestTrojan()
	// sing.TestSS()

	// cmd.StartApp()

	rawUri := `vless://15f430e8-a55a-48ca-92de-305fd4305767@xxx.dev:443?security=tls&type=ws&sni=xxx.dev&path=/&encryption=none&headerType=none&host=xxx.dev&fp=random&alpn=h2&allowInsecure=1`
	p := outbound.ParseRawUriToProxyItem(rawUri, outbound.SingBox)
	fmt.Println(p)
}
