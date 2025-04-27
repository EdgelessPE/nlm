package utils

import (
	"fmt"
	"io"
	"os"

	"github.com/klauspost/compress/zstd"
)

// CompressFileWithZstd compresses a file using zstd
func CompressFileWithZstd(inputPath string, outputPath string) error {
	// Open input file
	input, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open input file: %v", err)
	}
	defer input.Close()

	// Create output file
	output, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer output.Close()

	// Create zstd encoder
	enc, err := zstd.NewWriter(output)
	if err != nil {
		return fmt.Errorf("failed to create zstd encoder: %v", err)
	}
	defer enc.Close()

	// Copy and compress data
	if _, err := io.Copy(enc, input); err != nil {
		return fmt.Errorf("failed to compress data: %v", err)
	}

	return nil
}

// DecompressFileWithZstd decompresses a file using zstd
func DecompressFileWithZstd(inputPath string, outputPath string) error {
	// Open input file
	input, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open input file: %v", err)
	}
	defer input.Close()

	// Create output file
	output, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer output.Close()

	// Create zstd decoder
	dec, err := zstd.NewReader(input)
	if err != nil {
		return fmt.Errorf("failed to create zstd decoder: %v", err)
	}
	defer dec.Close()

	// Copy and decompress data
	if _, err := io.Copy(output, dec); err != nil {
		return fmt.Errorf("failed to decompress data: %v", err)
	}

	return nil
}
