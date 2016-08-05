package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"net"
)

//post 数据
func PostData(url string, paramMap map[string]string) string {
	fmt.Println("posting ", url)
	var pstr = ""
	for k, v := range paramMap {
		pstr += k + "=" + v + "&"
	}

	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader(pstr))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(string(body))
	return string(body)
}
func GetMyIp() string{
	conn, err := net.Dial("udp", "google.com:80")
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	defer conn.Close()
	return strings.Split(conn.LocalAddr().String(), ":")[0]
}