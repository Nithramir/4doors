package main

import (
	"database/sql"
	"mime/multipart"
)

func password_ok(password string, id string, db *sql.DB) int {
	request := "SELECT * WHERE FROM " + table_name + "WHERE id='" + id + "' AND passwd=MD5('" + password + "');"
	rows, err := db.Query(request)
	handle_err(err)
	for rows.Next() {
		return 1
	}
	return 0
}

func edit_room(titre string, id string, db *sql.DB, file_name string, file multipart.File, xav_doc string) {
	var alone int = 0
	request := "UPDATE " + table_name + " SET "
	if titre != "" {
		request += "Titre='" + titre + "' "
		alone = 1
	}
	if file_name != "" {
		if alone != 0 {
			request += ", "
		}
		request += "image='" + file_name + "'"
		alone = 1
	}
	if file_name != "" {
		if alone != 0 {
			request += ", "
		}
		request += "xav_doc='" + xav_doc + "'"
		alone = 1
	}
	if alone == 1 {
		request += " WHERE id='" + id + "';"
		print(request)
		db.Query(request)
	}
}
