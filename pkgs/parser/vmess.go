package parser

/*
vmess: ['v', 'ps', 'add', 'port', 'aid', 'scy', 'net', 'type', 'tls', 'id', 'sni', 'host', 'path', 'alpn', 'security', 'skip-cert-verify', 'fp', 'test_name', 'serverPort', 'nation']
*/

type OutVmess struct {
	Address string
	Port    int
	UUID    string

	AID            string
	ALPN           string
	FP             string
	Host           string
	Net            string
	Nation         string
	Path           string
	PS             string
	SCY            string
	Security       string
	ServerPort     string
	SkipCertVerify string
	SNI            string
	TLS            string
	Type           string
	TestName       string
	V              string
}
