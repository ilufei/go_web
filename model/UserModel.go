package model

import (
	lib "http_server/library"
	_ "fmt"
)

type UserModel struct {}
var User = &UserModel{}

func (user *UserModel) SelectOneUser(uid int) (map[string]string) {
	db, _ := lib.NewMysqlDBClient()
	rows, _ := db.Query("select * from userinfo where uid = ?", uid)
	column, _ := rows.Columns()              //读出查询出的列字段名
	
	values := make([][]byte, len(column))     //values是每个列的值，这里获取到byte里
	scans := make([]interface{}, len(column)) //因为每次查询出来的列是不定长的，用len(column)定住当次查询的长度
	for i := range values {                   //让每一行数据都填充到[][]byte里面
		scans[i] = &values[i]
	}
	
	row := make(map[string]string) 
	for rows.Next() { //循环，让游标往下移动
		if err := rows.Scan(scans...); err != nil { //query.Scan查询出来的不定长值放到scans[i] = &values[i],也就是每行都放在values里
			return row
		}
		for k, v := range values {     //每行数据是放在values里面，现在把它挪到row里
			key := column[k]
			row[key] = string(v)
		}
		break;
	}
	
	return row;
}

func (user *UserModel) SelectUser() ([]map[string]string) {
	db, _ := lib.NewMysqlDBClient()
	rows, _ := db.Query("select * from userinfo")
	column, _ := rows.Columns()              //读出查询出的列字段名
	
	values := make([][]byte, len(column))     //values是每个列的值，这里获取到byte里
	scans := make([]interface{}, len(column)) //因为每次查询出来的列是不定长的，用len(column)定住当次查询的长度
	for i := range values {                   //让每一行数据都填充到[][]byte里面
		scans[i] = &values[i]
	}
	
	results := []map[string]string{}
	i := 0
	for rows.Next() { //循环，让游标往下移动
		if err := rows.Scan(scans...); err != nil { //query.Scan查询出来的不定长值放到scans[i] = &values[i],也就是每行都放在values里
			return results
		}
		row := make(map[string]string) //每行数据
		for k, v := range values {     //每行数据是放在values里面，现在把它挪到row里
			key := column[k]
			row[key] = string(v)
		}
		//results[i] = row //装入结果集中
		results = append(results, row)
		i++
	}
 
	return results
}

func (user *UserModel) AddUser(username, department, created string) (lastInsertId int64, err error) {
	lastInsertId = 0
	db, err := lib.NewMysqlDBClient()
	if err != nil {
		return
	}
	
	stmt, err := db.Prepare("insert into userinfo(username, department, created) values(?, ?, ?)")
	if err != nil {
		return
	}
	res, err := stmt.Exec(username, department, created)
	if err != nil {
		return
	}
	lastInsertId, err = res.LastInsertId()
	if err != nil {
		return
	}
	
	return
}

func (user *UserModel) UpdateUser(uid int, username string, department string) (affect int64, err error) {
	db, err := lib.NewMysqlDBClient()
	if err != nil {
		return
	}
	
	stmt, err := db.Prepare("update userinfo set username = ?, department = ? where uid = ?")
	if err != nil {
		return
	}
	
	res, err := stmt.Exec(username, department, uid)
	if err != nil {
		return
	}
	affect, err = res.RowsAffected()
	if err != nil {
		return
	}
	
	return
}

func (user *UserModel) DeleteUser(uid int) (affect int64, err error) {
	db, err := lib.NewMysqlDBClient()
	if err != nil {
		return
	}
	
	stmt, err := db.Prepare("delete from userinfo where uid = ?")
	if err != nil {
		return
	}
	
	res, err := stmt.Exec(uid)
	if err != nil {
		return
	}
	
	affect, err = res.RowsAffected()
	if err != nil {
		return
	}
	
	return
}