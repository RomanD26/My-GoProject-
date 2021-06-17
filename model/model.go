package model

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

type User struct {
	Id   uint16 `json:"id"`
	Name string `json:"name"`
	Age  uint16 `json:"age"`
}

func RegSql(w http.ResponseWriter, r *http.Request) {
	// Получение данных
	name := r.FormValue("name")
	age := r.FormValue("age")

	// Выполняем проверку данных
	if name == "" || age == "" {
		fmt.Fprintf(w, "Не все данные заполнены")
	}
	// Подключение к БД
	db, err := sql.Open("mysql", "mysql:mysql@tcp(127.0.0.1:8889)/gobd")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Добавление данных в таблицу users
	insert, err := db.Queri(fmt.Sprintf("INSERT INTO 'users' ('name', 'age') VALUES('%s', '%d')", name, age))
	if err != nil {
		panic(err)
	}
	defer insert.Close()

	// Переадресация
	http.Redirect(w, r, "/resume", http.StatusSeeOther)
}
