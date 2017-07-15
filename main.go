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

/*func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}
*/
/*func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}*/

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
	create_room(database, "string", "yolo", "swag")
	fmt.Fprintf(w, "ok")
}

func get_salle(w http.ResponseWriter, r *http.Request) {
	id_need := r.FormValue("id")
	xav_code := get_room(database, id_need)
	fmt.Fprintf(w, string(xav_code))
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	//	title := r.URL.Path[len("/edit/"):]
	p, _ := ioutil.ReadFile("edit.html")

	fmt.Fprintf(w, string(p))
}

func main() {
	database = init_database("root", "1234")
	get_room(database, "1")
	http.HandleFunc("/view/", get_salle)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/create_salle/", create_salle)
	http.ListenAndServe(":8080", nil)
}
