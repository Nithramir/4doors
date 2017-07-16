package main

import (
	"database/sql"
	"errors"
	"fmt"
)

func init_database(db_user string, db_passw string) *sql.DB {
	db, err := sql.Open("mysql", db_user+":"+db_passw+"@tcp(127.0.0.1:3306)/")
	handle_err(err)
	db.Query("DROP TABLE " + table_name)
	_, err = db.Exec("CREATE DATABASE if not exists " + ddb_name)
	handle_err(err)
	_, err = db.Exec("USE " + ddb_name)
	handle_err(err)
	fmt.Println("Connected to database " + ddb_name)
	rows, err := db.Query("CREATE TABLE if not exists " + table_name + "(ID int NOT NULL AUTO_INCREMENT, Visite int DEFAULT 0, Good int DEFAULT 0, Bad int DEFAULT 0, xav_doc TEXT(5000), date DATETIME, passwd CHAR(255), salle1 int DEFAULT 0, salle2 int DEFAULT 0, salle3 int DEFAULT 0, salle4 int DEFAULT 0, primary key (ID));")
	handle_err(err)
	fmt.Println("Connected to " + table_name)
	defer rows.Close()
	return db
}

func create_room(db *sql.DB, id_act string, xav_doc string, num_salle string) {
	request := "INSERT INTO " + table_name + "(date, xav_doc) VALUES(NOW(), '" + xav_doc + "');"
	_, err := db.Query(request)
	handle_err(err)
}

func get_room(db *sql.DB, id_need string) (string, error) {
	request := "SELECT xav_doc FROM " + table_name + " WHERE id = '" + id_need + "';"
	rows, err := db.Query(request)
	handle_err(err)
	defer rows.Close()
	for rows.Next() {
		var code string
		err := rows.Scan(&code)
		handle_err(err)
		return code, nil
	}
	return "", errors.New("id not found")
}
