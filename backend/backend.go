package backend

import (
	"fmt"
	"database/sql"
	// "html/template"
)

import _ "github.com/go-sql-driver/mysql"

type data struct{
	Key string
	Data_1 string
	Data_2 string
	Data_3 string
}
var result []data

func connect()(*sql.DB,error)  {
	db, err := sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/db_golang")
	if err != nil {
		return nil,err
	}

	return db,err
}
func GetData()([]data){
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	defer db.Close()

	rows, err := db.Query("select * from tb_api")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	defer rows.Close()

	for rows.Next(){
		var each = data{}
		var err = rows.Scan(&each.Key,&each.Data_1,&each.Data_2,&each.Data_3)
		if err != nil{
			fmt.Println(err.Error())
			return nil
		}
		result = append(result,each)
	}
	if err = rows.Err();err != nil{
		fmt.Println(err.Error())
		return nil
	}
	return result
}