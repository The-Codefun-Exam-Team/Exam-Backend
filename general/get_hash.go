package general

import (
	"crypto/md5"
	"strings"
)

func GetHash(rawstr string) string {
	s := strings.ToLower(strings.TrimSpace(rawstr))
	hash := md5.Sum([]byte(s))
	return string(hash[:])
}
