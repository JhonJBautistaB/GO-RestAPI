package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type users struct {
	Id        int    `json: "id"`
	Nombres   string `json: "nombres"`
	Apellidos string `json: "apellidos"`
	Documento string `json: "documento"`
	Movil     int    `json: "movil"`
	CreateAt  int    `json: "create_at"`
}

/**
** Conections DataBas
 */

func dbConnection() (db *sql.DB) {
	db, err := sql.Open("sqlite3", "./jikkoDemo.db")
	db.Exec("CREATE TABLE IF NOT EXISTS users(id INTEGER PRIMARY KEY AutoIncrement, nombres VARCHAR(60) NOT NULL, apellidos VARCHAR(60) NOT NULL, documento INT NOT NULL UNIQUE, movil INT, create_at INTEGER)")
	if err != nil {
		fmt.Println("Error de conexión a BD")
		panic(err.Error())
	}
	log.Println("Coexión a BD exitosa")
	return db
}

func (u *users) getUsers() ([]users, error) {
	db := dbConnection()
	q := `SELECT * FROM users`
	fmt.Println(db.Query(q))
	rows, err := db.Query(q)

	if err != nil {
		fmt.Println("Error al realizar consulta")
		panic(err.Error())
	}
	defer rows.Close()

	res := []users{}

	for rows.Next() {
		rows.Scan(&u.Id, &u.Nombres, &u.Apellidos, &u.Documento, &u.Movil, &u.CreateAt)
		res = append(res, *u)
	}
	return res, nil

}
