package utils

import (
	"encoding/base64"
)

var coder1 = base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789._")

func DecodeBase641(base64Str string) []byte {
	b, _ := coder1.DecodeString(base64Str)
	return b
}
func EncodeBase641(b []byte) string {
	return coder1.EncodeToString(b)
}
