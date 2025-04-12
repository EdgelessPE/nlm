package utils

import (
	"errors"
	"regexp"
	"strings"
)

type NepNameParsed struct {
	Name    string
	Version string
	Author  string
	Flags   string
}

func ParseNepFileName(fileName string) (NepNameParsed, error) {
	// 第一次分割，去掉拓展名并尝试获取 flags
	sp1 := strings.Split(fileName, ".")
	steamName := sp1[0]
	var flags string
	if len(sp1) > 1 && regexp.MustCompile(`^[A-Z]+$`).MatchString(sp1[1]) {
		flags = sp1[1]
	}

	// 第二次分割
	sp2 := strings.Split(steamName, "_")
	if len(sp2) != 3 {
		return NepNameParsed{}, errors.New("invalid nep stem name: " + fileName)
	}
	name := sp2[0]
	version := sp2[1]
	author := sp2[1]

	return NepNameParsed{
		Name:    name,
		Version: version,
		Author:  author,
		Flags:   flags,
	}, nil
}
