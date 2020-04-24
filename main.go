package main

import (
	"fmt"
	"html/template"
	"net/http"
	"encoding/json"
	"database/sql"
	// "runtime"
)
import "WEB_DEVELOPMENT/web_theme/backend"
import _ "github.com/go-sql-driver/mysql"


func connect()(*sql.DB, error){
	db,err := sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/db_golang")
	if err != nil {
		return nil,err
	}
	return db,err
}
type person struct{
	Key string
	Data1 string
	Data2 string
	Data3 string
}
func main()  {
	http.HandleFunc("/",HandleIndex)
	http.HandleFunc("/request",HandleRequest)
	http.HandleFunc("/doc",HandleDoc)
	http.HandleFunc("/proccess",HandleSave)
	http.HandleFunc("/api",RequestApi)

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	fmt.Println("localhost:8000")
	http.ListenAndServe(":8000",nil)

}

func HandleIndex(w http.ResponseWriter, r *http.Request)  {
	var tmpl = template.Must(template.ParseFiles("html/index.html"))
	if err := tmpl.Execute(w,nil);err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandleRequest(w http.ResponseWriter, r *http.Request)  {
	var tmpl = template.Must(template.ParseFiles("html/request.html"))
	if err := tmpl.Execute(w,nil);err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandleDoc(w http.ResponseWriter, r *http.Request)  {
	var tmpl = template.Must(template.ParseFiles("html/doc.html"))
	if err := tmpl.Execute(w,nil);err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func HandleSave(w http.ResponseWriter, r *http.Request)  {
	// runtime.GOMAXPROCS("2")
	if r.Method == "POST"{
		decoder := json.NewDecoder(r.Body)
		payload := struct{
			Key  string `json:"key"`
			Data1  string `json:"data1"`
			Data2  string `json:"data2"`
			Data3  string `json:"data3"`
		}{}
		if err := decoder.Decode(&payload); err != nil {
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
		}
		// fmt.Println(payload.Key)
	    db, err := connect()
		if err != nil {
			fmt.Println(err.Error())
		}
		defer db.Close()
	     _, err = db.Exec("insert into tb_api values(?,?,?,?)",payload.Key,payload.Data1,payload.Data2,payload.Data3)
		if err != nil{
			fmt.Println(err.Error())
			return
		}
		fmt.Println("success")
		message := fmt.Sprintf("Insert Sucess your key %s",
		payload.Key,
		)
		w.Write([]byte(message))
		return
	}
	http.Error(w, "Only accept POST request",http.StatusBadRequest)
}
func RequestApi(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")
	var result = backend.GetData()
	if r.Method == "POST"{
		var result, err = json.Marshal(result)
		if err != nil{
			http.Error(w, err.Error(),http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w,"",http.StatusBadRequest)
}