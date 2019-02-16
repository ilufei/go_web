package action 

import (
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
	"bytes"
	"strings"
)

func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	header := request.Header
	fmt.Fprintln(writer, header)

	acceptLanguage := request.Header["Accept-Language"]
	fmt.Fprintln(writer, acceptLanguage)

	userAgent := request.Header.Get("User-Agent")
	fmt.Fprintln(writer, userAgent)
}

func RequestCookie(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "RequestCookie")
}

func RequestGet(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "RequestGet")
}

func RequestPostValue(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "RequestPostValue")
}

func RequestPostJson(writer http.ResponseWriter, request *http.Request) {
	len := request.ContentLength
	body := make([]byte, len)
	request.Body.Read(body)
	
	fmt.Fprintf(writer, string(body))
}

func RequestUploadFile(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "RequestUploadFile")
}

func RequestClientGet(writer http.ResponseWriter, request *http.Request) {
	response, err := http.Get("http://www.kankan.com")
	
	if err != nil {
		fmt.Fprintf(writer, "get the url Response error")
	}else{
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Fprintln(writer, string(body))
	}
}

func RequestClientPostJson(writer http.ResponseWriter, request *http.Request) {
	body := "{\"action\" : 20}"
	response, err := http.Post("http://lufei.me:9999/request/post/json", "application/json;charset=utf-8", bytes.NewBuffer([]byte(body)))
	
	if err != nil {
		fmt.Fprintf(writer, "post data error")
	}else{
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Fprintln(writer, string(body))
	}	
}

func RequestClientPostValue(writer http.ResponseWriter, request *http.Request) {
	v := url.Values{}
	v.Set("username", "lufei")
	v.Set("password", "myPassword")
	
	body := ioutil.NopCloser(strings.NewReader(v.Encode()))
	client := &http.Client{}
	
	request, err := http.NewRequest("POST", "http://lufei.me:9999/request/post/value", body)
	if err != nil {
		fmt.Fprintln(writer, "http newRequest failed")
		return
	}
	
	request.Header.Set("Content-Type", "application/x-www-form-urlencode;param=value")
	response, err := client.Do(request)
	defer response.Body.Close()
	
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Fprintln(writer, "http response get failed")
		return
	}
	
	fmt.Fprintln(writer, string(content))
}



