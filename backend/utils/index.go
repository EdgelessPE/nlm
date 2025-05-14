package utils

import (
	"errors"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type NepNameParsed struct {
	Name    string
	Version string
	Author  string
	Flags   string
}

func ParseNepFileName(fileName string) (NepNameParsed, error) {
	// 去掉 .nep 后缀
	if !strings.HasSuffix(fileName, ".nep") {
		return NepNameParsed{}, errors.New("invalid nep extension name: " + fileName)
	}
	fileName = strings.TrimSuffix(fileName, ".nep")

	// 检查 flags 位是否存在
	// 检查文件名是否为 .跟随大写字母
	var flags string
	regex := regexp.MustCompile(`\.([A-Z]+)$`)
	if regex.MatchString(fileName) {
		flags = regex.FindStringSubmatch(fileName)[1]
		fileName = strings.TrimSuffix(fileName, "."+flags)
	}

	// 分割
	sp := strings.Split(fileName, "_")
	if len(sp) != 3 {
		return NepNameParsed{}, errors.New("invalid nep stem name: " + fileName)
	}
	name := sp[0]
	version := sp[1]
	author := sp[2]

	return NepNameParsed{
		Name:    name,
		Version: version,
		Author:  author,
		Flags:   flags,
	}, nil
}

func GetUUIDSubDir(uuid string) string {
	return uuid[:2]
}

func JoinUrl(base string, paths ...string) string {
	final := base
	for _, component := range paths {
		if len(component) == 0 {
			continue
		}

		lastFinalChar := ""
		if len(final) > 0 {
			lastFinalChar = string(final[len(final)-1])
		}

		if lastFinalChar == "/" && string(component[0]) == "/" {
			final = final[:len(final)-1] + component
		} else if lastFinalChar == "/" || string(component[0]) == "/" {
			final = final + component
		} else {
			final = final + "/" + component
		}
	}

	return final
}

func GetMajorVersion(version string) int {
	sp := strings.Split(version, ".")
	major, err := strconv.Atoi(sp[0])
	if err != nil {
		return 0
	}
	return major
}

func CleanBotTaskName(name string) string {
	if strings.Contains(name, "_") {
		return strings.Split(name, "_")[0]
	}
	return name
}

func SortFlags(flags string) string {
	sp := strings.Split(flags, "")
	sort.Slice(sp, func(i, j int) bool {
		return sp[i] > sp[j]
	})
	return strings.Join(sp, "")
}

func PointerBool(b bool) *bool {
	return &b
}
