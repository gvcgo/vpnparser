package xray

/*
https://xtls.github.io/config/outbounds/shadowsocks.html#serverobject

{
	"servers": [
	  {
		"email": "love@xray.com",
		"address": "127.0.0.1",
		"port": 1234,
		"method": "加密方式",
		"password": "密码",
		"uot": true,
		"UoTVersion": 2,
		"level": 0
	  }
	]
}

Method:
2022-blake3-aes-128-gcm
2022-blake3-aes-256-gcm
2022-blake3-chacha20-poly1305
aes-256-gcm
aes-128-gcm
chacha20-poly1305 或称 chacha20-ietf-poly1305
xchacha20-poly1305 或称 xchacha20-ietf-poly1305
none 或 plain

UoTVersion:
UDP over TCP 的实现版本。
当前可选值：1, 2

*/

var XraySS string = `{
	"servers": [
	  {
		"email": "love@xray.com",
		"address": "127.0.0.1",
		"port": 1234,
		"method": "加密方式",
		"password": "密码",
		"uot": true,
		"UoTVersion": 2,
		"level": 0
	  }
	]
}`
