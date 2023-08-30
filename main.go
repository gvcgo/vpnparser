package main

import (
	_ "github.com/moqsien/vpnparser/pkgs/outbound/xray"
	"github.com/moqsien/vpnparser/pkgs/parser"
)

func main() {
	// parser.VlessTest()
	parser.TrojanTest()
	// s := xray.GetPattern()
	// fmt.Println(s)
	// xray.TestVmess()
	// xray.TestVless()
}
