package parser

/*
shadowsocks: ['plugin', 'obfs', 'obfs-host', 'mode', 'path', 'mux', 'host']
*/

type OutShadowSocks struct {
	Address  string
	Port     int
	Method   string
	Password string

	Host     string
	Mode     string
	Mux      string
	Path     string
	Plugin   string
	OBFS     string
	OBFSHost string
}
