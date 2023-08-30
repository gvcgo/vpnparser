package main

import (
	"github.com/moqsien/vpnparser/pkgs/outbound/xray"
	"github.com/moqsien/vpnparser/pkgs/parser"
)

func main() {
	parser.VmessTest()
	// s := xray.GetPattern()
	// fmt.Println(s)
	xray.TestVmess()
}
