package parser

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/moqsien/goutils/pkgs/crypt"
)

var SSRMethod map[string]struct{} = map[string]struct{}{
	"aes-128-ctr":   {},
	"aes-192-ctr":   {},
	"aes-256-ctr":   {},
	"aes-128-cfb":   {},
	"aes-192-cfb":   {},
	"aes-256-cfb":   {},
	"rc4-md5":       {},
	"chacha20-ietf": {},
	"xchacha20":     {},
}

var SSROBFS map[string]struct{} = map[string]struct{}{
	"plain":              {},
	"http_simple":        {},
	"http_post":          {},
	"random_head":        {},
	"tls1.2_ticket_auth": {},
}

/*
shadowsocksr: ['remarks', 'obfsparam', 'protoparam', 'group']
*/

type ParserSSR struct {
	Address    string
	Port       int
	Method     string
	Password   string
	OBFS       string
	Proto      string
	OBFSParam  string
	ProtoParam string

	*StreamField
}

func (that *ParserSSR) Parse(rawUri string) {
	r := strings.ReplaceAll(rawUri, SchemeSSR, "")
	vList := strings.Split(r, "?")
	if len(vList) == 2 {
		that.parseMethod(vList[0])
		that.parseParams(vList[1])
	} else {
		vList = strings.Split(r, "remarks=")
		if len(vList) == 2 {
			that.parseMethod(vList[0])
			that.parseParams("remarks=" + vList[1])
		}
	}

	if _, ok := SSRMethod[that.Method]; !ok {
		that.Method = "rc4-md5"
	}
	if _, ok := SSROBFS[that.OBFS]; !ok {
		that.OBFS = "plain"
	}
	that.StreamField = &StreamField{}
}

func (that *ParserSSR) parseParams(s string) {
	s = strings.ReplaceAll(s, "+", "-")
	s = strings.ReplaceAll(s, "/", "_")
	testUrl := fmt.Sprintf("https://www.test.com/?%s", s)
	if u, err := url.Parse(testUrl); err == nil {
		that.OBFSParam = u.Query().Get("obfsparam")
		if that.OBFSParam != "" {
			that.OBFSParam = crypt.DecodeBase64(that.OBFSParam)
		}
		that.ProtoParam = u.Query().Get("protoparam")
		if that.ProtoParam != "" {
			that.ProtoParam = crypt.DecodeBase64(that.ProtoParam)
		}
	}
}

func (that *ParserSSR) parseMethod(s string) {
	vList := strings.Split(s, ":")
	if len(vList) == 6 {
		that.Address = vList[0]
		that.Port = gconv.Int(vList[1])
		that.Proto = vList[2]
		that.Method = vList[3]
		that.OBFS = vList[4]
		p := strings.TrimSuffix(vList[5], "/")
		that.Password = crypt.DecodeBase64(p)

	} else if len(vList) == 5 {
		that.Address = vList[0]
		that.Port = gconv.Int(vList[1])
		that.Proto = vList[2]
		that.Method = vList[3]
		obfs_pwd := vList[4]
		for obfs_name := range SSROBFS {
			if strings.Contains(obfs_pwd, obfs_name) {
				that.OBFS = obfs_name
				that.Password = crypt.DecodeBase64(strings.TrimPrefix(obfs_pwd, obfs_name))
			}
		}
	}
}

func (that *ParserSSR) GetAddr() string {
	return that.Address
}

func (that *ParserSSR) GetPort() int {
	return that.Port
}

func (that *ParserSSR) Show() {
	fmt.Printf("addr: %s, port: %d, method: %s, password: %s\n",
		that.Address,
		that.Port,
		that.Method,
		that.Password)
}

func SSRTest() {
	rawUri := "ssr://94.23.116.190:443:origin:aes-256-ctr:tls1.2_ticket_authSG93ZHlCeXBhc3NlcjIwMjI=remarks=MTJ8UmxKZmMzQmxaV1J1YjJSbFh6QXdNalUlM0Q=\u0026obfsparam=VG05dVpRJTNEJTNE\u0026protoparam=VG05dVpRJTNEJTNE"
	// rawUri := "ssr://94.23.116.190:443:origin:aes-256-ctr:tls1.2_ticket_auth:SG93ZHlCeXBhc3NlcjIwMjI=/?obfsparam=Tm9uJSXvv70lJe+/vVxceDFm\u0026protoparam=Tm9uJSXvv70lJe+/vVxceDFm\u0026remarks=5rOV5Zu9XzA4MjgwMDk\u0026group="
	p := &ParserSSR{}
	p.Parse(rawUri)
	p.Show()
}
