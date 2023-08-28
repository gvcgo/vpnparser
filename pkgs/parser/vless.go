package parser

/*
vless: ['security', 'type', 'sni', 'path', 'encryption', 'headerType', 'packetEncoding', 'serviceName', 'mode', 'flow', 'alpn', 'host', 'fp', 'pbk', 'sid', 'spx']
*/

type ParserVless struct {
	Address string
	Port    int
	UUID    string

	ALPN           string
	Encryption     string
	Flow           string
	FP             string
	HeaderType     string
	Host           string
	Mode           string
	PacketEncoding string
	Path           string
	PBK            string
	Security       string
	ServiceName    string
	SID            string
	SNI            string
	SPX            string
	Type           string
}
