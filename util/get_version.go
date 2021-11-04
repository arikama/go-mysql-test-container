package util

import (
	"regexp"
	"strconv"
	"strings"
)

func GetVersion(migrationFullPath string) int {
	splits := strings.Split(migrationFullPath, "/")
	r, _ := regexp.Compile(`V\d+`)
	s := splits[len(splits)-1]
	found := r.FindString(s)
	if found == "" {
		return 0
	}
	v, _ := strconv.Atoi(found[1:])
	return v
}
