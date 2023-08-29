package sing

/*
http://sing-box.sagernet.org/zh/configuration/outbound/trojan/
{
  "type": "trojan",
  "tag": "trojan-out",

  "server": "127.0.0.1", # 必填
  "server_port": 1080, # 必填
  "password": "8JCsPssfgS8tiRwiMlhARg==", # 必填
  "network": "tcp",
  "tls": {},
  "multiplex": {},
  "transport": {},

  ... // 拨号字段
}

*/

var SingTrojan string = `{
	"type": "trojan",
	"tag": "trojan-out",
	"server": "127.0.0.1",
	"server_port": 1080,
	"password": "8JCsPssfgS8tiRwiMlhARg==",
	"network": "tcp",
	"tls": {},
	"multiplex": {},
	"transport": {}
}`
