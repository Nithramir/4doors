package main

import (
	_ "bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	_ "os"
	"reflect"
	_ "strconv"
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
	print("create_salle")
	id_act := r.FormValue("id_act")
	num_room := r.FormValue("num_room")
	xav_code := r.FormValue("xav_code")
	titre := r.FormValue("titre")
	print("ok\n")
	r.ParseMultipartForm(32 << 20)
	img, header, err := r.FormFile("uploadfile")
	handle_err(err)
	fmt.Println(header.Filename)
	fmt.Println(reflect.TypeOf(num_room))
	if (xav_code == "") || (id_act == "") || (num_room == "") || (titre == "") || (img == nil) {
		print("Erreur les formulaires envoyés")
	} else {
		print("formulaires bien rentrés, je les insère magueul\n")
	}
	create_room(database, id_act, xav_code, num_room, titre)
	fmt.Fprintf(w, "ok")
}

func get_salle(w http.ResponseWriter, r *http.Request) {
	print("get_salle")
	//id_need := r.FormValue("id")
	var room room_type
	room, _ = get_room(database, "1")
	b, err := json.Marshal(room)
	handle_err(err)
	if err != nil {
		fmt.Fprintf(w, "Id error")
	}
	fmt.Fprintf(w, string(b))
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	//	title := r.URL.Path[len("/edit/"):]
	p, _ := ioutil.ReadFile("edit.html")
	print(p)

	fmt.Fprintf(w, string(p))
}

func imgHandler(w http.ResponseWriter, r *http.Request) {
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	blue := color.RGBA{0, 0, 255, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)
	file, err := ioutil.ReadFile("img.png")
	file = []byte("<a href=\"pageone.php\"><img src=\"https://www.w3schools.com/css/img_fjords.jpg\"  /></a>")
	handle_err(err)
	/*buffer := new(bytes.Buffer)
	var img image.Image = m
	if err := jpeg.Encode(buffer, img, nil); err != nil {
		log.Println("unable to encode image.")
	}*/

	//	w.Header().Set("Content-Type", "image/jpeg")
	//	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(file); err != nil {
		log.Println("unable to write image.")
	}
}

func main() {

	type Test struct {
		Chaine string
		Entier int
	}
	database = init_database("root", "1234")
	create_room(database, "3", "swagman", "2", "bat-room")

	http.HandleFunc("/get_salle/", get_salle)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/create_salle/", create_salle)
	http.HandleFunc("/img/", imgHandler)
	http.ListenAndServe(":8080", nil)
}
