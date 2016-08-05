package utils
import (
	"strings"
	"fmt"
	"regexp"
)

func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

//替换文本中包含key=value 的字符串中value值并返回新的内容
func ReplacePairValue(content,key,newValue string)string{
	//re:= regexp.MustCompile("[\\S]+([\\s*]|)+=([\\s*]|)+[^\\s*]*")
	regStr:=key+"([\\s*]|)+=([\\s*]|)+[^\\s*]*"
	re:= regexp.MustCompile(regStr)
	cs := re.Find([]byte(content))
	fmt.Println(string(cs))
	if scs:=string(cs);scs!=""{
		v:=strings.Split(scs,"=")[1]
		v=strings.TrimSpace(v)
		if(strings.HasPrefix(v,"\"")||strings.HasPrefix(v,"'")){
			//fmt.Println(v,"is string",tarStr)
		}
		bcs1:=strings.Replace(scs,v,newValue,-1);
		content=strings.Replace(content,scs,bcs1,1);
	}
	return  content
}
//查找文本中包含key=value 的字符串中value值
func GetPairValue(content,key string)string{
	//re:= regexp.MustCompile("[\\S]+([\\s*]|)+=([\\s*]|)+[^\\s*]*")
	regStr:=key+"([\\s*]|)+=([\\s*]|)+[^\\s*]*"

	re:= regexp.MustCompile(regStr)
	cs := re.Find([]byte(content))
	//fmt.Println(string(cs))
	if scs:=string(cs);scs!=""{
		v:=strings.Split(scs,"=")[1]
		v=strings.TrimSpace(v)
		return v
	}
	return ""
}
//替换文本中包含key=value 的字符串中value值+step
func IncreasePairValue(content,key string,step int64) string{
	regStr:=key+"([\\s*]|)+=([\\s*]|)+[^\\s*]*"
	re:= regexp.MustCompile(regStr)
	cs := re.Find([]byte(content))
	fmt.Println(string(cs))
	if scs:=string(cs);scs!=""{
		v:=strings.Split(scs,"=")[1]
		v=strings.TrimSpace(v)
		newValue:="";
		if(strings.HasPrefix(v,"\"")||strings.HasPrefix(v,"'")){
			v=Substr(v,1,len(v)-2)
			vlen:=len(v)
			n:=GetInt64(v)
			n+=step
			sn:=fmt.Sprint(n)
			newValue=sn
			for i:=0;i<vlen-len(sn);i++{
				newValue="0"+newValue
			}
		}else{
			n1:=GetInt64(v)
			n1+=step
			newValue=fmt.Sprint(n1)
		}
		bcs1:=strings.Replace(scs,v,newValue,-1);
		content=strings.Replace(content,scs,bcs1,1);
	}
	return  content
}
