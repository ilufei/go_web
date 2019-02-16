package action 

import (
	"fmt"
	"net/http"
)

func JsonEncode(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "JsonEncode")
}

func JsonDecode(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "JsonDecode")
}

func JsonFileEncode(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "JsonFileEncode")
}

func JsonFileDecode(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "JsonFileDecode")
}