package xray

import (
	"fmt"

	"github.com/gogf/gf/encoding/gjson"
)

/*
https://xtls.github.io/config/outbound.html#outboundobject

{
  "outbounds": [
    {
      "sendThrough": "0.0.0.0",
      "protocol": "协议名称",
      "settings": {outbound设置},
      "tag": "标识",
      "streamSettings": {},
      "proxySettings": {
        "tag": "another-outbound-tag"
      },
      "mux": {}
    }
  ]
}
*/

var XrayOut string = `{
  "sendThrough": "0.0.0.0",
  "protocol": "协议名称",
  "tag": "标识",
  "settings": %s,
  "streamSettings": %s
}`

func GetPattern() string {
	x := fmt.Sprintf(XrayOut, "{}", "{}")
	j := gjson.New(x)
	j.Set("outbounds.0.protocol", "vmess")
	j.Set("outbounds.0.tag", "PROXY_OUT")
	return j.MustToJsonIndentString()
}
