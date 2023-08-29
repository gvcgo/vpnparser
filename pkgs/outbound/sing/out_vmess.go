package sing

/*
http://sing-box.sagernet.org/zh/configuration/outbound/vmess/
{
  "type": "vmess",
  "tag": "vmess-out",

  "server": "127.0.0.1", # 必填
  "server_port": 1080, # 必填
  "uuid": "bf000d23-0752-40b4-affe-68f7707a9661", # 必填
  "security": "auto",
  "alter_id": 0,
  "global_padding": false,
  "authenticated_length": true,
  "network": "tcp",
  "tls": {},
  "packet_encoding": "",
  "multiplex": {},
  "transport": {},

  ... // 拨号字段
}

Security:
auto
none
zero
aes-128-gcm
chacha20-poly1305
aes-128-ctr

alter_id:
0	禁用旧协议
1	启用旧协议
>1	未使用, 行为同 1

packet_encoding:
(空)		禁用
packetaddr	由 v2ray 5+ 支持
xudp		由 xray 支持

*/

var SingVmess string = `{
	"type": "vmess",
	"tag": "vmess-out",
	"server": "127.0.0.1",
	"server_port": 1080,
	"uuid": "bf000d23-0752-40b4-affe-68f7707a9661",
	"security": "auto",
	"alter_id": 0,
	"global_padding": false,
	"authenticated_length": true,
	"network": "tcp",
	"tls": {},
	"packet_encoding": "",
	"multiplex": {},
	"transport": {},
`
