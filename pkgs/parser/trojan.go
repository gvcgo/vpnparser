package parser

/*
trojan: ['allowInsecure', 'peer', 'sni', 'type', 'path', 'security', 'headerType']
*/

type ParserTrojan struct {
	Address  string
	Port     int
	Password string

	AllowInsecure int
	HeaderType    string
	Path          string
	Peer          string
	Security      string
	SNI           string
	Type          string
}
