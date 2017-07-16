package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

const (
	table_name = "table1"
	ddb_name   = "4doors"
)

var database *sql.DB

func handle_err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func create_salle(w http.ResponseWriter, r *http.Request) {
	id_act := r.FormValue("id_act")
	num_room := r.FormValue("num_room")
	xav_code := r.FormValue("xav_code")
	fmt.Println(reflect.TypeOf(num_room))
	if (xav_code == "") || (id_act == "") || (num_room == "") {
		print("Erreur les formulaires envoyés")
	} else {
		print("formulaires bien rentrés, je les insère magueul")
	}
	create_room(database, id_act, xav_code, num_room)
	fmt.Fprintf(w, "ok")
}

func get_salle(w http.ResponseWriter, r *http.Request) {
	id_need := r.FormValue("id")
	xav_code, err := get_room(database, id_need)
	if err != nil {
		fmt.Fprintf(w, "Id error")
	}
	fmt.Fprintf(w, xav_code)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	//	title := r.URL.Path[len("/edit/"):]
	p, _ := ioutil.ReadFile("edit.html")

	fmt.Fprintf(w, string(p))
}

func main() {
	database = init_database("root", "1234")
	http.HandleFunc("/get_salle/", get_salle)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/create_salle/", create_salle)
	http.ListenAndServe(":8080", nil)
}
