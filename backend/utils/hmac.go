package utils

import (
	"crypto/hmac"
	"crypto/sha256"
)

// 验证 GitHub Webhook 的签名
func VerifySignature(secret, payload, signature string) bool {
	// 使用 HMAC-SHA256 算法和 secret 生成签名
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(payload))
	expectedSignature := mac.Sum(nil)

	// 比较生成的签名与请求头中的签名
	return hmac.Equal(expectedSignature, []byte(signature))
}
