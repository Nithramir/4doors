package main

import (
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"strconv"
)

type bouton struct {
	ID    int
	Titre string
}
type room_type struct {
	ID      int
	Xav_doc string
	Titre   string
	Good    int
	Bad     int
	Bouton1 bouton
	Bouton2 bouton
	Bouton3 bouton
	Bouton4 bouton
	Image   []byte
}

func init_database(db_user string, db_passw string) *sql.DB {
	db, err := sql.Open("mysql", db_user+":"+db_passw+"@tcp(127.0.0.1:3306)/")
	handle_err(err)
	_, err = db.Exec("CREATE DATABASE if not exists " + ddb_name)
	handle_err(err)
	_, err = db.Exec("USE " + ddb_name)
	handle_err(err)
	fmt.Println("Connected to database " + ddb_name)
	rows, err := db.Query("CREATE TABLE if not exists " + table_name + "(ID int NOT NULL AUTO_INCREMENT, Visite int DEFAULT 0, Good int DEFAULT 0, Bad int DEFAULT 0, Titre CHAR(255), image CHAR(255), Xav_doc TEXT(5000), date DATETIME, passwd CHAR(255), salle1 int DEFAULT 0, salle2 int DEFAULT 0, salle3 int DEFAULT 0, salle4 int DEFAULT 0, primary key (ID));")
	handle_err(err)
	fmt.Println("Connected to " + table_name)
	defer rows.Close()
	return db
}

func create_room(db *sql.DB, id_act string, Xav_doc string, num_salle string, Titre string, img multipart.File, image_name string) string {
	if img != nil {
		file, _ := ioutil.ReadAll(img)
		err := ioutil.WriteFile("./img/"+image_name, file, 0644)
		handle_err(err)
		fmt.Println(string(file))
	}
	pass := string(randSeq(10))
	request := "INSERT INTO " + table_name + "(date, Xav_doc, Titre, image, passwd) VALUES(NOW(), '" + Xav_doc + "', '" + Titre + "', './img/" + image_name + "',  MD5('" + pass + "') );"
	print(request)
	print("\n")
	req, err := db.Query(request)
	handle_err(err)
	req.Close()
	handle_err(err)
	request = "UPDATE " + table_name + " SET salle" + num_salle + " = (SELECT ID FROM (SELECT MAX(ID) AS ID from " + table_name + ") AS ID) WHERE ID = " + id_act + ";"
	print(request)
	req, err = db.Query(request)
	handle_err(err)
	defer req.Close()
	return pass
}

func get_room(db *sql.DB, id_need string) (room_type, error) {

	var room room_type
	var file_name string
	request := "SELECT ID, image, Good, Bad, Titre, Xav_doc, salle1, salle2, salle3, salle4 FROM " + table_name + " WHERE id = '" + id_need + "';"
	print(request)
	rows, err := db.Query(request)
	handle_err(err)
	for rows.Next() {
		err := rows.Scan(&room.ID, &file_name, &room.Good, &room.Bad, &room.Titre, &room.Xav_doc, &room.Bouton1.ID, &room.Bouton2.ID, &room.Bouton3.ID, &room.Bouton4.ID)
		rows.Close()
		handle_err(err)
		room.Bouton1.Titre = get_Titre(db, room.Bouton1.ID)
		room.Bouton2.Titre = get_Titre(db, room.Bouton2.ID)
		room.Bouton3.Titre = get_Titre(db, room.Bouton3.ID)
		room.Bouton4.Titre = get_Titre(db, room.Bouton4.ID)
		room.Image, err = ioutil.ReadFile(file_name)
		handle_err(err)
		fmt.Println(string(room.Image))
		print("\nroom\n")
		fmt.Println(room)
		return room, nil
	}
	return room, errors.New("id not found")
}

func get_Titre(db *sql.DB, id int) string {
	print(strconv.Itoa(id))
	request := "SELECT Titre from " + table_name + " WHERE id = '" + strconv.Itoa(id) + "';"
	rows, err := db.Query(request)
	handle_err(err)
	var Titre string
	for rows.Next() {
		rows.Scan(&Titre)
		rows.Close()
	}
	print(Titre)
	return Titre
}
