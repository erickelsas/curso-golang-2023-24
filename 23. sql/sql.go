package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	urlConexao := "root:root@/golang?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", urlConexao)
	if erro != nil {
		log.Fatal(erro)
		fmt.Println("erro")
	}
	defer db.Close()

	if erro = db.Ping(); erro != nil {
		fmt.Println(erro)
	}

	linhas, erro := db.Query("SELECT * from user")
	if erro != nil {
		log.Fatal(erro)
	}
	defer linhas.Close()

	fmt.Println(linhas)
}
