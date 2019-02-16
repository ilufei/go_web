package action

import (
	"fmt"
	"net/http"
	"os"
)

type OSAction struct {}
var OS = &OSAction{}

func (OS *OSAction) Test(writer http.ResponseWriter, request *http.Request) {
    goPath := os.Getenv("GOPATH")
	currentpath, _ := os.Getwd()
	
	
	fmt.Fprintln(writer, "GOPATH : " + goPath)
	fmt.Fprintln(writer, "currentpath : " + currentpath)
}