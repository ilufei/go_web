package action 

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"
	"time"
	"http_server/model"
)

type UserAction struct {}
var User = &UserAction{}

func (User *UserAction) Select(writer http.ResponseWriter, request *http.Request) {
    //animationId, _ := request.FormValue('animation_id')
	result.Data = "mysql select"
	
	jsonStr, _ := json.Marshal(result)
	fmt.Fprintf(writer, string(jsonStr))
}

func (User *UserAction) AddUser(writer http.ResponseWriter, request *http.Request) {
	username := "pain"
	department := "develop"
	created := time.Now().Format("2006-01-02 15:04:05")

	lastInsertId, err := model.User.AddUser(username, department, created)
	if err != nil {
		fmt.Fprintf(writer, "add user failed")
		return
	}
	
	fmt.Fprintf(writer, "add user sucess lastInsertId id " + strconv.FormatInt(lastInsertId,10))
}

func (User *UserAction) UpdateUser(writer http.ResponseWriter, request *http.Request) {
	uid := 7
	username := "pain-test"
	department := "develop-test"

	affect, err := model.User.UpdateUser(uid, username, department)
	if err != nil {
		fmt.Fprintf(writer, "update user failed")
		return
	}
	
	fmt.Fprintf(writer, "update user sucess affect rows is " + strconv.FormatInt(affect,10))
}

func (User *UserAction) DeleteUser(writer http.ResponseWriter, request *http.Request) {
	uid := 7

	affect, err := model.User.DeleteUser(uid)
	if err != nil {
		fmt.Fprintf(writer, "delete user failed")
		return
	}
	
	fmt.Fprintf(writer, "delete user sucess affect rows is " + strconv.FormatInt(affect,10))
}

func (User *UserAction) SelectUser(writer http.ResponseWriter, request *http.Request) {	
	users := model.User.SelectUser()
	result.Data = users
	
	jsonStr, _ := json.Marshal(result)
	fmt.Fprintf(writer, string(jsonStr))
}

func (User *UserAction) SelectByUid(writer http.ResponseWriter, request *http.Request) {
	uid := 6
	user := model.User.SelectOneUser(uid)
	result.Data = user
	
	jsonStr, _ := json.Marshal(result)
	fmt.Fprintf(writer, string(jsonStr))

}