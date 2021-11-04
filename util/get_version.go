package util

import (
	"regexp"
	"strconv"
	"strings"
)

func GetVersion(migrationFullPath string) int {
	splits := strings.Split(migrationFullPath, "/")
	r, err := regexp.Compile(`V\d+`)
	if err != nil {
		return 0
	}
	s := splits[len(splits)-1]
	v, err := strconv.Atoi(r.FindString(s)[1:])
	if err != nil {
		return 0
	}
	return v
}
