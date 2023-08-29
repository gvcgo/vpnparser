package xray

/*
https://xtls.github.io/config/outbounds/vmess.html#outboundconfigurationobject

{
	"vnext": [
	  {
		"address": "127.0.0.1",
		"port": 37192,
		"users": [
		  {
			"id": "5783a3e7-e373-51cd-8642-c83782b807c5",
			"security": "auto",
			"level": 0
		  }
		]
	  }
	]
}

Security:
"aes-128-gcm" | "chacha20-poly1305" | "auto" | "none" | "zero"
加密方式，客户端将使用配置的加密方式发送数据，服务器端自动识别，无需配置。

"aes-128-gcm"：推荐在 PC 上使用
"chacha20-poly1305"：推荐在手机端使用
"auto"：默认值，自动选择（运行框架为 AMD64、ARM64 或 s390x 时为 aes-128-gcm 加密方式，其他情况则为 Chacha20-Poly1305 加密方式）
"none"：不加密
"zero"：不加密，也不进行消息认证 (v1.4.0+)
提示:
推荐使用"auto"加密方式，这样可以永久保证安全性和兼容性。
"none" 伪加密方式会计算并验证数据包的校验数据，由于认证算法没有硬件支持，在部分平台可能速度比有硬件加速的 "aes-128-gcm" 还慢。
"zero" 伪加密方式不会加密消息也不会计算数据的校验数据，因此理论上速度会高于其他任何加密方式。实际速度可能受到其他因素影响。
不推荐在未开启 TLS 加密并强制校验证书的情况下使用 "none" "zero" 伪加密方式。 如果使用 CDN 或其他会解密 TLS 的中转平台或网络环境建立连接，不建议使用 "none" "zero" 伪加密方式。
无论使用哪种加密方式， VMess 的包头都会受到加密和认证的保护。

*/

var XrayVmess string = `{
	"vnext": [
	  {
		"address": "127.0.0.1",
		"port": 37192,
		"users": [
		  {
			"id": "5783a3e7-e373-51cd-8642-c83782b807c5",
			"alterId": 0,
			"security": "auto",
			"level": 0
		  }
		]
	  }
	]
}`
