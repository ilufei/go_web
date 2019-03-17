package main

import (
	"net/http"
	"http_server/utils"
	"http_server/router"
)

func main() {
    defer func() {
        if err := recover(); err != nil {
            utils.Slog("error catched : ", err)
        }
    }()

    config := utils.Config
	mux := http.NewServeMux()

	//加载路由
	router.Load(mux)
	
	server := &http.Server{
		Addr : config.Address,
		Handler : mux,
		MaxHeaderBytes : 1 << 20,
	}

	utils.Slog("system start..")
	server.ListenAndServe()
}