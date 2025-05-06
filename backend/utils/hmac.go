package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// 验证 GitHub Webhook 的签名
func VerifySignature(secret string, compactJSONPayload string, signature string) (bool, error) {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(compactJSONPayload))
	expectedMac := mac.Sum(nil)
	actualMac, err := hex.DecodeString(signature)
	if err != nil {
		return false, err
	}
	return hmac.Equal(actualMac, expectedMac), nil
}
