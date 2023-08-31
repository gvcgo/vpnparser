package main

import (
	"github.com/moqsien/vpnparser/pkgs/outbound/sing"
	_ "github.com/moqsien/vpnparser/pkgs/outbound/xray"
	_ "github.com/moqsien/vpnparser/pkgs/parser"
)

func main() {
	// parser.VlessTest()
	// parser.TrojanTest()
	// parser.SSRTest()
	// s := xray.GetPattern()
	// fmt.Println(s)
	// xray.TestVmess()
	// xray.TestTrojan()
	// xray.TestSS()

	// sing.TestVmess()
	// sing.TestVless()
	// sing.TestTrojan()
	// sing.TestSS()
	sing.TestSSR()
}
