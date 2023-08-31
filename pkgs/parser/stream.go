package parser

type StreamField struct {
	Network          string
	StreamSecurity   string
	Path             string
	Host             string
	TCPHeaderType    string
	GRPCServiceName  string
	GRPCMultiMode    string
	ServerName       string
	TLSALPN          string
	Fingerprint      string
	RealityShortId   string
	RealitySpiderX   string
	RealityPublicKey string
	PacketEncoding   string
}
