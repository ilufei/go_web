package action 

import (
	"fmt"
	"encoding/json"
	"net/http"
)

type Post struct {
	User string
	Threads []string
}

func ResponseWrite(writer http.ResponseWriter, request *http.Request) {
	str := `<html><head><title>Go Web Programming</title></head><body><h1>Hello World Go Web Programming</h1></body></html>`
	writer.Write([]byte(str))
}

func ResponseWriteHeader(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(501)
	fmt.Fprintf(writer, "no such service, try next door")
}

func ResponseLocation(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Location", "http://www.kankan.com")
	writer.WriteHeader(302)
}

func ResponseJson(writer http.ResponseWriter, request *http.Request) {
	data := &Post{User : "lufei", Threads : []string{"first", "second", "Third"},}
	json, _ := json.Marshal(data)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(json)
}

func ResponseSetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie1 := http.Cookie{
		Name : "lufei",
		Value : "trust your self one",
		HttpOnly : true,
	}
	
	cookie2 := http.Cookie{
		Name : "jack",
		Value : "trust your self two",
		HttpOnly : true,
	}
	
	cookie3 := http.Cookie{
		Name : "pain",
		Value : "trust your self thrid",
		HttpOnly : true,
	}
	
	writer.Header().Set("Set-Cookie", cookie1.String())
	writer.Header().Add("Set-Cookie", cookie2.String())
	http.SetCookie(writer, &cookie3)
	writer.Write([]byte("set cookie sucess"))
}