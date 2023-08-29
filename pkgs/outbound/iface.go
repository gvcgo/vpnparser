package outbound

type IOutbound interface {
	Parse(string)
	Addr() string
	Host() string
	Scheme() string
	GetOutboundStr() string
}
