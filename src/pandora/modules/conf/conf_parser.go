package conf

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	re  *regexp.Regexp
	pat string = "[#].*\\n|\\s+\\n|\\S+[=]|.*\n"
)

func init() {
	re, _ = regexp.Compile(pat)
}

type Configuration map[string]string

func (cfg Configuration) IsSet(key string) bool {
	_, ok := cfg[key]
	return ok
}

func (cfg Configuration) GetInt64(key string) int64 {
	if !cfg.IsSet(key) {
		return 0
	}
	v, err := strconv.ParseInt(cfg[key], 10, 64)
	if err != nil {
		return 0
	}
	return v

}
func (cfg Configuration) GetFloat64(key string) float64 {
	if !cfg.IsSet(key) {
		return 0
	}
	v, err := strconv.ParseFloat(cfg[key], 10)
	if err != nil {
		return 0
	}
	return v
}
func (cfg Configuration) GetString(key string) string {
	if !cfg.IsSet(key) {
		return ""
	}
	return cfg[key]
}
func (cfg Configuration) GetBool(key string) bool {
	if !cfg.IsSet(key) {
		return false
	}
	flag, err := strconv.ParseBool(cfg[key])
	if err != nil {
		return false
	}
	return flag

}

func LoadConfig(filePath string) (Configuration, error) {

	_, er := os.Stat(filePath)
	//fmt.Println("config file info", filePath, fi.Size(), fi.ModTime(), fi.Sys(), er)
	fmt.Println("loading config:", filePath)
	if er != nil {
		return nil, er
	}

	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	buf := bufio.NewReader(file)
	dest := Configuration{}

	for {
		line, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		}
		str := strings.TrimSpace(string(line))

		if strings.HasPrefix(str, "#") || (strings.HasPrefix(str, "[") && strings.HasSuffix(str, "]")) {
			continue
		} else {

			if strings.Contains(str, "=") {
				sl := strings.SplitN(str, "=", 2)
				dest[sl[0]] = sl[1]
			}

		}

	}
	return dest, nil

}
