package logger

import (
	"fmt"
	"pandora/modules/utils/datetimeutil"
	"log"
	"os"
	"runtime"
	"strings"
)

func D(args ...interface{}) {
	out("DEBUG", args)
}
func W(args ...interface{}) {
	out("WARN", args)
}
func E(args ...interface{}) {
	out("ERROR", args)
	//panic(args)
}
var mLog *log.Logger
var f=""
var sep="/"
var std = log.New(os.Stderr, "", 0)

func out(level string, args []interface{}) {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	if(strings.Contains(file,sep)){
		sr:=strings.Split(file,sep)
		file=sr[len(sr)-1]
	}
	fpathstr:="at:"+file+" "+fmt.Sprint(line)
	str := "[" + datetimeutil.FormatDateTimeMillsNow() + "] " + level + "/:"
	for _, s := range args {
		if s == nil {
			s = "nil"
		}
		str += fmt.Sprint(s) + " "
	}
	str+=" "+fpathstr

	f1:="logs"+sep+datetimeutil.FormatDateNow()+".log"
	if(f1!=f){
		f=f1
	}
	if(level=="ERROR"){
		std.Println(str)
	}else{
		fmt.Println(str)
	}


	saveToFile(f,str)
}
func saveToFile(filePath, str string) int64 {
	//fmt.Println("saving", filePath, str)
	userFile := filePath
	fout, err := os.OpenFile(userFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	defer fout.Close()
	if err != nil {
		fmt.Println(err)
	}
	n,er:=fout.WriteString(str+"\n")
	if(er!=nil){
		n=0;
		fmt.Println(er.Error())
	}
	return int64(n)
}

