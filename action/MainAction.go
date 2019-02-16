package action 

import (
	"fmt"
	"net/http"
)

func Index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "this is index page")
}