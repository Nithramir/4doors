package main

import "database/sql"

func password_ok(password string, id string, db *sql.DB) int {
	request := "SELECT * WHERE FROM " + table_name + "WHERE id='" + id + "' AND passwd=MD5('" + password + "');"
	rows, err := db.Query(request)
	handle_err(err)
	for rows.Next() {
		return 1
	}
	return 0
}
