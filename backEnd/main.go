package main

import (
	"log"
	"net/http"

	"fmt"

	"database/sql"

	"encoding/json"

	"strings"

	"strconv"

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
	if r.Method == "GET" {
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
	} else if r.Method == "DELETE" {
		id := r.FormValue("ID")
		fmt.Println(id)
		_, err = db.Query("DELETE FROM users WHERE id=" + id + ";")
		if err != nil {
			fmt.Println(err)
			fmt.Fprintf(w, "error")
			return
		}
	} else if r.Method == "POST" {
		fmt.Println("post request")
		r.ParseForm()
		var tuser user
		tuser.ID, _ = strconv.Atoi(r.FormValue("ID"))
		tuser.Name = r.FormValue("Name")
		tuser.Password = r.FormValue("Password")
		tuser.Age, _ = strconv.Atoi(r.FormValue("Age"))
		tuser.Work = r.FormValue("Work")
		_, err := db.Query("UPDATE users SET name=$2,password=$3,age=$4,work=$5 WHERE id=$1", tuser.ID, tuser.Name, tuser.Password, tuser.Age, tuser.Work)
		if err != nil {
			fmt.Println(err)
			fmt.Fprint(w, "problem in update row")
		}
		fmt.Fprint(w, "updated successs fully")
	}
}
func GetTableContextTools(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", "user=postgres host=localhost port=5432 dbname=dbcontrol ")
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "error")
		return
	}
	if r.Method == "GET" {
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
	} else if r.Method == "DELETE" {
		id := r.URL.Query().Get("id")
		if strings.ContainsAny(id, "OR*/-''") {
			fmt.Fprint(w, "invalid id")
			return
		}
		fmt.Println(id)
		_, err = db.Query("DELETE FROM tools WHERE id=" + id + ";")
		if err != nil {
			fmt.Println(err)
			fmt.Fprintf(w, "error")
			return
		}
	} else if r.Method == "POST" {
		fmt.Println("post request")
		r.ParseForm()
		var ttool tool
		ttool.ID, _ = strconv.Atoi(r.FormValue("ID"))
		ttool.Name = r.FormValue("Name")
		ttool.Ussage = r.FormValue("Ussage")
		_, err := db.Query("UPDATE tools SET name=$2,ussage=$3 WHERE id=$1", ttool.ID, ttool.Name, ttool.Ussage)
		if err != nil {
			fmt.Println(err)
			fmt.Fprint(w, "problem in update row")
		}
		fmt.Fprint(w, "updated successs fully")
	}
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("requested")
	db, err := sql.Open("postgres", "user=postgres host=localhost port=5432 dbname=dbcontrol ")
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "error")
		return
	}
	m := mux.Vars(r)
	table := m["table"]
	id := r.FormValue("ID")
	fmt.Println(id)
	_, err = db.Query("DELETE FROM " + table + " WHERE id=" + id + ";")
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "error")
		return
	}
}
func main() {
	router := mux.NewRouter()
	router.Methods("GET").Path("/tables").HandlerFunc(GetTables)
	router.Methods("GET", "POST").Path("/tables/users").HandlerFunc(GetTableContextUsers)
	router.Methods("GET", "POST").Path("/tables/tools").HandlerFunc(GetTableContextTools)
	router.Methods("POST").Path("/delete/{table}").HandlerFunc(DeleteHandler)
	http.ListenAndServe(":8080", router)
}
