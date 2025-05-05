package utils

import (
	"encoding/hex"
	"io"
	"os"

	"lukechampine.com/blake3"
)

// 获取某个文件的 blake3 hash 值
func GetBlake3HashFromFile(filePath string) (string, error) {
	hasher := blake3.New(32, nil)

	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(hasher, file)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
