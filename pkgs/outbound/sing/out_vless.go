package sing

/*
http://sing-box.sagernet.org/zh/configuration/outbound/vless/

{
  "type": "vless",
  "tag": "vless-out",

  "server": "127.0.0.1", # 必填
  "server_port": 1080, # 必填
  "uuid": "bf000d23-0752-40b4-affe-68f7707a9661", # 必填
  "flow": "xtls-rprx-vision",
  "network": "tcp",
  "tls": {},
  "packet_encoding": "",
  "transport": {},

  ... // 拨号字段
}

*/

var SingVless string = `{
	"type": "vless",
	"tag": "vless-out",
	"server": "127.0.0.1",
	"server_port": 1080,
	"uuid": "bf000d23-0752-40b4-affe-68f7707a9661",
	"flow": "xtls-rprx-vision",
	"network": "tcp",
	"tls": {},
	"packet_encoding": "",
	"transport": {},
}`
