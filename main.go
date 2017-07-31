package main

import (
	_ "bytes"
	"database/sql"
	"encoding/json"
	"fmt"
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

	_ "github.com/go-sql-driver/mysql"
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
	img, header, _ := r.FormFile("uploadfile")
	fmt.Println(header.Filename)
	fmt.Println(reflect.TypeOf(img))
	fmt.Println(img)
	f, err := ioutil.ReadAll(img)
	fmt.Println(string(f))
	handle_err(err)
	if (xav_code == "") || (id_act == "") || (num_room == "") || (titre == "") || (img == nil) {
		print("Erreur les formulaires envoyés")
	} else {
		print("formulaires bien rentrés, je les insère magueul\n")
		pass := create_room(database, id_act, xav_code, num_room, titre, img, header.Filename)
		fmt.Fprintf(w, pass)
	}
}

func get_salle(w http.ResponseWriter, r *http.Request) {
	print("get_salle")
	id_need := r.FormValue("id")
	var room room_type
	room, _ = get_room(database, id_need)
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

func can_modify(w http.ResponseWriter, r *http.Request) {

	password := r.FormValue("password")
	id_room := r.FormValue("id_room")
	if password_ok(id_room, password, database) == 1 {
		fmt.Fprintf(w, "ok")
	} else {
		fmt.Fprintf(w, "pasok")
	}

}

func main() {

	database = init_database("root", "1234")
	create_room(database, "3", "swagman", "2", "bat-room", nil, "name")

	http.HandleFunc("/get_salle/", get_salle)
	http.HandleFunc("/can_modify/", can_modify)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/create_salle/", create_salle)
	http.HandleFunc("/img/", imgHandler)
	http.ListenAndServe(":8080", nil)
}
