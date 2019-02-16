package action 

import (
	"fmt"
	"net/http"
	"strconv"
)

func Strconv(writer http.ResponseWriter, request *http.Request) {
	num := 113
	fmt.Fprintln(writer, "num 113 strconv to string is " + strconv.Itoa(num))
	
	str := "114"
	strNum, _ := strconv.Atoi(str)
	fmt.Fprintln(writer, "string 114 strconv to num is " + strconv.Itoa(strNum))
	
	/*
	#string到int
	int,err:=strconv.Atoi(string)
	
	#string到int64
	int64, err := strconv.ParseInt(string, 10, 64)
	
	#int到string
	string:=strconv.Itoa(int)
	
	#int64到string
	string:=strconv.FormatInt(int64,10)
	*/
}