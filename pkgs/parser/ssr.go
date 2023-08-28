package parser

/*
shadowsocksr: ['remarks', 'obfsparam', 'protoparam', 'group']
*/

type OutShadowSocksR struct {
	Address  string
	Port     int
	Method   string
	Password string
	OBFS     string
	Proto    string

	OBFSParam  string
	ProtoParam string
	Group      string
	Remarks    string
}
