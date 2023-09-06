package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"path/filepath"
	"strings"
	"time"
)

func GetFileHashName(filename string) string {
	ext := filepath.Ext(filename)
	name := strings.TrimSuffix(filename, ext)
	timeStamp := time.Now().UTC().String()

	hasher := sha1.New()
	_, _ = hasher.Write([]byte(name + timeStamp))

	hashedName := hex.EncodeToString(hasher.Sum(nil)) + ext
	return hashedName
}
