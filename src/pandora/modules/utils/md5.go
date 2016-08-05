package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"pandora/modules/utils/datetimeutil"
	"time"

	"math/rand"
)

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
func Md5Bytes(b []byte) string {
	h := md5.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

func GenTicket(secretKey string) string {
	rand.Seed(time.Now().UnixNano())
	secretKey += fmt.Sprint(rand.Intn(1000000))
	return Md5(fmt.Sprint(datetimeutil.GetCurrentTimeMillis()) + secretKey)
}
