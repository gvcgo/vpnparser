package main

import (
	"github.com/moqsien/vpnparser/pkgs/outbound/xray"
	_ "github.com/moqsien/vpnparser/pkgs/parser"
)

func main() {
	// parser.VlessTest()
	// s := xray.GetPattern()
	// fmt.Println(s)
	// xray.TestVmess()
	xray.TestVless()
}