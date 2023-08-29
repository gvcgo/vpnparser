package xray

/*
https://xtls.github.io/config/outbounds/vless.html#outboundconfigurationobject

{
	"vnext": [
	  {
		"address": "example.com",
		"port": 443,
		"users": [
		  {
			"id": "5783a3e7-e373-51cd-8642-c83782b807c5",
			"encryption": "none",
			"flow": "xtls-rprx-vision",
			"level": 0
		  }
		]
	  }
	]
}

Flow:
流控模式，用于选择 XTLS 的算法。
目前出站协议中有以下流控模式可选：
无 flow，空字符或者 none：使用普通 TLS 代理
xtls-rprx-vision：使用新 XTLS 模式 包含内层握手随机填充 支持 uTLS 模拟客户端指纹
xtls-rprx-vision-udp443：同 xtls-rprx-vision, 但是放行了目标为 443 端口的 UDP 流量
此外，目前 XTLS 仅支持 TCP、mKCP、DomainSocket 这三种传输方式。

*/

var XrayVless string = `{
	"vnext": [
	  {
		"address": "example.com",
		"port": 443,
		"users": [
		  {
			"id": "5783a3e7-e373-51cd-8642-c83782b807c5",
			"encryption": "none",
			"flow": "xtls-rprx-vision",
			"level": 0
		  }
		]
	  }
	]
}`
