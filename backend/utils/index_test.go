package utils

import (
	"testing"
)

func TestParseNepFileName(t *testing.T) {
	// 普通包名
	parsed, err := ParseNepFileName("Visual Studio Code_1.80.0_Cno.nep")
	if err != nil {
		t.Fatal(err)
	}
	if parsed.Name != "Visual Studio Code" {
		t.Fatal("name is not correct: " + parsed.Name)
	}
	if parsed.Version != "1.80.0" {
		t.Fatal("version is not correct: " + parsed.Version)
	}
	if parsed.Author != "Cno" {
		t.Fatal("author is not correct: " + parsed.Author)
	}
	if parsed.Flags != "" {
		t.Fatal("flags is not correct: " + parsed.Flags)
	}

	// 带 flags 的
	parsed, err = ParseNepFileName("搜狗拼音_15.2.0.0_Cno.IE.nep")
	if err != nil {
		t.Fatal(err)
	}
	if parsed.Name != "搜狗拼音" {
		t.Fatal("name is not correct: " + parsed.Name)
	}
	if parsed.Version != "15.2.0.0" {
		t.Fatal("version is not correct: " + parsed.Version)
	}
	if parsed.Author != "Cno" {
		t.Fatal("author is not correct: " + parsed.Author)
	}
	if parsed.Flags != "IE" {
		t.Fatal("flags is not correct: " + parsed.Flags)
	}

	// 非法文件名
	_, err = ParseNepFileName("Visual Studio Code_1.80.0_Cno.nep.meta")
	if err == nil {
		t.Fatal("should be error")
	}
}
