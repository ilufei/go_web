package main

import (
	"net/http"
	"http_server/action"
	"http_server/utils"
)

func main() {
    defer func() {
        if err := recover(); err != nil {
            utils.Slog("error catched : ", err)
        }
    }()

    config := utils.Config
	mux := http.NewServeMux()

	//处理静态文件
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	//首页
	mux.HandleFunc("/", action.Index)

	//Json操作
	mux.HandleFunc("/json/encode", action.JsonEncode)
	mux.HandleFunc("/json/decode", action.JsonDecode)
	mux.HandleFunc("/json/file/encode", action.JsonFileEncode)
	mux.HandleFunc("/json/file/decode", action.JsonFileDecode)

	//Request操作
	mux.HandleFunc("/request/client/get", action.RequestClientGet)
	mux.HandleFunc("/request/client/post/json", action.RequestClientPostJson)
	mux.HandleFunc("/request/client/post/value", action.RequestClientPostValue)
	mux.HandleFunc("/request/header", action.RequestHeader)
	mux.HandleFunc("/request/cookie", action.RequestCookie)
	mux.HandleFunc("/request/get", action.RequestGet)
	mux.HandleFunc("/request/post/value", action.RequestPostValue)
	mux.HandleFunc("/request/post/json", action.RequestPostJson)
	mux.HandleFunc("/request/upload/file", action.RequestUploadFile)
	
	//Response
	mux.HandleFunc("/response/write/header", action.ResponseWriteHeader)
	mux.HandleFunc("/response/location", action.ResponseLocation)
	mux.HandleFunc("/response/write", action.ResponseWrite)
	mux.HandleFunc("/response/json", action.ResponseJson)
	mux.HandleFunc("/response/cookie/set", action.ResponseSetCookie)

	//ioutil
	mux.HandleFunc("/io/ioutil/read/dir", action.IoutilReadDir)
	mux.HandleFunc("/io/ioutil/read/file", action.IoutilReadFile)
	mux.HandleFunc("/io/ioutil/write/file", action.IoutilWriteFile)
	mux.HandleFunc("/io/ioutil/temp", action.IoutilTempDirFile)
	mux.HandleFunc("/io/read/block", action.ReadFileBlock)
	
	//bufio
	mux.HandleFunc("/io/bufio/read/file/block", action.BufioReadBlock)
	mux.HandleFunc("/io/bufio/read/file/line", action.BufioReadLine)
	
	//Redis
	mux.HandleFunc("/redis/get", action.RedisGet)
	mux.HandleFunc("/redis/set", action.RedisSet)
	mux.HandleFunc("/redis/list/len", action.RedisListLen)
	
	//User 
	mux.HandleFunc("/user/select1", action.User.Select)
	mux.HandleFunc("/user/select", action.User.SelectUser)
	mux.HandleFunc("/user/select/6", action.User.SelectByUid)
	mux.HandleFunc("/user/add", action.User.AddUser)
	mux.HandleFunc("/user/update", action.User.UpdateUser)
	mux.HandleFunc("/user/delete", action.User.DeleteUser)
	
	//OS
	mux.HandleFunc("/os/test", action.OS.Test)
	
	//Strconv
	mux.HandleFunc("/strconv", action.Strconv)
	
	//time
	mux.HandleFunc("/time", action.PrintTime)
	
	server := &http.Server{
		Addr : config.Address,
		Handler : mux,
		MaxHeaderBytes : 1 << 20,
	}

	utils.Slog("system start..")
	server.ListenAndServe()
}