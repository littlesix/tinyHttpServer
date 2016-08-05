package utils

import (
	"strconv"
	"fmt"
	"strings"
)

func GetInt64(str string) int64 {
	str=strings.Trim(str,"\n")
	str=strings.Trim(str," ")
	n, e := strconv.ParseInt(str, 10, 64)
	if e!=nil{
		fmt.Println(e.Error())
	}
	return n
}
