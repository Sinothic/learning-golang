package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/base_producao?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	var (
		id_usuario int
		ds_nome    string
	)

	{
		if err := db.QueryRow("select id_usuario, ds_nome from usuarios where id_usuario = 1").Scan(&id_usuario, &ds_nome); err != nil {
			log.Fatal(err)
		}

		fmt.Println(id_usuario, ds_nome)
	}

}
