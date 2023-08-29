package xray

/*
https://xtls.github.io/config/outbounds/trojan.html#outboundconfigurationobject

{
	"servers": [
	  {
		"address": "127.0.0.1",
		"port": 1234,
		"password": "password",
		"email": "love@xray.com",
		"level": 0
	  }
	]
}
*/

var XrayTrojan string = `{
	"servers": [
	  {
		"address": "127.0.0.1",
		"port": 1234,
		"password": "password",
		"email": "love@xray.com",
		"level": 0
	  }
	]
}`
