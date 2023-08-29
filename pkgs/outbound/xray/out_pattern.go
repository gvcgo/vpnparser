package xray

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
    "outbounds": [{
		"sendThrough": "0.0.0.0",
		"protocol": "协议名称",
    "tag": "标识",
		"settings": %s,
		"streamSettings": %s
	}]
}`
