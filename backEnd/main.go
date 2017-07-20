package main

import (
	"log"
	"net/http"

	"fmt"

	"database/sql"

	"encoding/json"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type user struct {
	ID       int
	Name     string
	Password string
	Age      int
	Work     string
}
type tool struct {
	ID     int
	Name   string
	Ussage string
	Owner  string
}
type table struct {
	Name string
	Nums int
}

//GetTables :write database tables to client in json
func GetTables(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", "user=postgres host=localhost port=5432 dbname=dbcontrol ")
	if err != nil {
		log.Fatal(err)
		fmt.Fprintf(w, "error")
		return
	}
	rows, err := db.Query("SELECT table_name FROM information_schema.tables WHERE table_type = 'BASE TABLE' AND table_schema = 'public' ORDER BY table_name")
	var tables []*table
	if err != nil {
		log.Fatal(err)
		fmt.Fprintf(w, "error")
		return
	}
	for rows.Next() {
		var t table
		rows.Scan(&t.Name)
		tables = append(tables, &t)
	}
	for _, i := range tables {
		rows, err := db.Query("SELECT COUNT(*) FROM " + i.Name + ";")
		if err != nil {
			fmt.Println(err)
			fmt.Fprintf(w, "error")
			return
		}
		for rows.Next() {
			var num int
			rows.Scan(&num)
			i.Nums = num
		}
	}
	JSONval, err := json.Marshal(tables)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "error")
		return
	}
	fmt.Fprint(w, string(JSONval))
}

//GetTableContextUsers : write rquseted tables records
func GetTableContextUsers(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", "user=postgres host=localhost port=5432 dbname=dbcontrol ")
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "error")
		return
	}
	rows, err := db.Query("SELECT * FROM  users;")
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "error")
		return
	}
	var users []user
	for rows.Next() {
		var Tuser user
		rows.Scan(&Tuser.ID, &Tuser.Name, &Tuser.Password, &Tuser.Age, &Tuser.Work)
		users = append(users, Tuser)
	}
	JSONval, err := json.Marshal(users)
	fmt.Fprint(w, string(JSONval))
}
func GetTableContextTools(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", "user=postgres host=localhost port=5432 dbname=dbcontrol ")
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "error")
		return
	}
	rows, err := db.Query("SELECT * FROM  tools;")
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "error")
		return
	}
	var tools []tool
	for rows.Next() {
		var Ttool tool
		rows.Scan(&Ttool.ID, &Ttool.Name, &Ttool.Ussage, &Ttool.Owner)
		tools = append(tools, Ttool)
	}
	JSONval, err := json.Marshal(tools)
	fmt.Fprint(w, string(JSONval))
}

func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {

}
func main() {
	router := mux.NewRouter()
	router.Methods("GET").Path("/tables").HandlerFunc(GetTables)
	router.Methods("GET").Path("/tables/users").HandlerFunc(GetTableContextUsers)
	router.Methods("GET").Path("/tables/tools").HandlerFunc(GetTableContextTools)
	router.Methods("DELETE").Path("/delete/users/{id}").HandlerFunc(DeleteUsersHandler)
	http.ListenAndServe(":8080", router)
}
