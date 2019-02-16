package library

import (
	"os"
	"fmt"
	"encoding/json"
	"http_server/utils"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "strconv"
	"time"
)

type MysqlConfiguration struct {
    Address      string
	User		 string
    Password     string
	Database	 string
}

var mysqlConfig MysqlConfiguration

func init() {
    loadDBConfig()
}

func loadDBConfig() {
	file, _ := os.Open("config/db.config.json")
	defer file.Close()

	jsonDecoder := json.NewDecoder(file)
	mysqlConfig = MysqlConfiguration{}
	err := jsonDecoder.Decode(&mysqlConfig)

	if err == nil {
    	utils.Slog("db config content decode sucess")
    } else {
    	utils.Slog("db config content decode failed")
    }
}

func NewMysqlDBClient() (*sql.DB, error) {
	mysqlDSN := fmt.Sprintf("%s:%s@%s/%s?charset=utf8", mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Address, mysqlConfig.Database)
	db, err := sql.Open("mysql", mysqlDSN)

	if err != nil {
		return db, err
	}
	
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	db.SetConnMaxLifetime(time.Hour)
	
	return db, err
}

func initMysqlClient() {
	mysqlDSN := fmt.Sprintf("%s:%s@%s/%s?charset=utf8", mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Address, mysqlConfig.Database)
	MysqlClient, err := sql.Open("mysql", mysqlDSN)

	MysqlClient.SetMaxOpenConns(2000)
	MysqlClient.SetMaxIdleConns(1000)
	MysqlClient.SetConnMaxLifetime(time.Hour)

	if err == nil {
		utils.Slog("mysql client init sucess")
	} else {
		utils.Slog("mysql client init failed")
	}
	
	err = MysqlClient.Ping()
	if err == nil {
		utils.Slog("mysql ping sucess")
	} else {
		utils.Slog("mysql ping failed")
	}
	
	/*
	//query
	rows, err := MysqlClient.Query("select * from userinfo")
	if err != nil {
		utils.Slog("mysql query failed")
	}
	
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		if err != nil {
			utils.Slog("mysql scan failed")
		} else {
			rowStr := fmt.Sprintf("uid:%d , username:%s, department:%s, created:%s", uid, username, department, created)
			utils.Slog(rowStr)
		}
	}
	*/
}