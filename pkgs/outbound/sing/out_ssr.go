package sing

/*
http://sing-box.sagernet.org/zh/configuration/outbound/shadowsocksr/#_3
{
  "type": "shadowsocksr",
  "tag": "ssr-out",

  "server": "127.0.0.1", # 必填
  "server_port": 1080, # 必填
  "method": "aes-128-cfb", # 必填
  "password": "8JCsPssfgS8tiRwiMlhARg==", # 必填
  "obfs": "plain",
  "obfs_param": "",
  "protocol": "origin",
  "protocol_param": "",
  "network": "udp",

  ... // 拨号字段
}

Protocal:
origin
verify_sha1
auth_sha1_v4
auth_aes128_md5
auth_aes128_sha1
auth_chain_a
auth_chain_b

OBFS:
plain
http_simple
http_post
random_head
tls1.2_ticket_auth

*/

var SingSSR string = `{
	"type": "shadowsocksr",
	"tag": "ssr-out",
	"server": "127.0.0.1",
	"server_port": 1080,
	"method": "aes-128-cfb",
	"password": "8JCsPssfgS8tiRwiMlhARg==",
	"obfs": "plain",
	"obfs_param": "",
	"protocol": "origin",
	"protocol_param": "",
	"network": "udp"
}`
