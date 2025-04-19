package utils

import (
	"fmt"
	"io"
	"os"

	"github.com/klauspost/compress/zstd"
)

// CompressFileWithZstd 使用zstd压缩文件
func CompressFileWithZstd(inputPath string, outputPath string) error {
	// 打开输入文件
	input, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("打开输入文件失败: %v", err)
	}
	defer input.Close()

	// 创建输出文件
	output, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("创建输出文件失败: %v", err)
	}
	defer output.Close()

	// 创建zstd编码器
	enc, err := zstd.NewWriter(output)
	if err != nil {
		return fmt.Errorf("创建zstd编码器失败: %v", err)
	}
	defer enc.Close()

	// 复制并压缩数据
	if _, err := io.Copy(enc, input); err != nil {
		return fmt.Errorf("压缩数据失败: %v", err)
	}

	return nil
}

// DecompressFileWithZstd 使用zstd解压缩文件
func DecompressFileWithZstd(inputPath string, outputPath string) error {
	// 打开输入文件
	input, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("打开输入文件失败: %v", err)
	}
	defer input.Close()

	// 创建输出文件
	output, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("创建输出文件失败: %v", err)
	}
	defer output.Close()

	// 创建zstd解码器
	dec, err := zstd.NewReader(input)
	if err != nil {
		return fmt.Errorf("创建zstd解码器失败: %v", err)
	}
	defer dec.Close()

	// 复制并解压数据
	if _, err := io.Copy(output, dec); err != nil {
		return fmt.Errorf("解压数据失败: %v", err)
	}

	return nil
}
