package outbound

type IOutbound interface {
	Parse(string)
	Addr() string
	Port() int
	Scheme() string
	GetOutboundStr() string
}
