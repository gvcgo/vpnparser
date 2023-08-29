package sing

/*
http://sing-box.sagernet.org/zh/configuration/outbound/shadowsocks/

{
  "type": "shadowsocks",
  "tag": "ss-out",

  "server": "127.0.0.1",
  "server_port": 1080,
  "method": "2022-blake3-aes-128-gcm",
  "password": "8JCsPssfgS8tiRwiMlhARg==",
  "plugin": "",
  "plugin_opts": "",
  "network": "udp",
  "udp_over_tcp": false | {},
  "multiplex": {},

  ... // 拨号字段
}

Method:
2022-blake3-aes-128-gcm
2022-blake3-aes-256-gcm
2022-blake3-chacha20-poly1305
none
aes-128-gcm
aes-192-gcm
aes-256-gcm
chacha20-ietf-poly1305
xchacha20-ietf-poly1305
aes-128-ctr
aes-192-ctr
aes-256-ctr
aes-128-cfb
aes-192-cfb
aes-256-cfb
rc4-md5
chacha20-ietf
xchacha20
*/

var SingSS string = `{
	"type": "shadowsocks",
	"tag": "ss-out",
	"server": "127.0.0.1",
	"server_port": 1080,
	"method": "2022-blake3-aes-128-gcm",
	"password": "8JCsPssfgS8tiRwiMlhARg==",
	"plugin": "",
	"plugin_opts": "",
	"network": "udp",
	"udp_over_tcp": false | {},
	"multiplex": {}
}`
