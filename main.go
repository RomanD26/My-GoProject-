package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var PORT = ":8080"
var basa = []User{}

type Data struct {
	Name string
	Age  uint16
}

type User struct {
	Id   uint16 `json:"id"`
	Name string `json:"name"`
	Age  uint16 `json:"age"`
}

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}

func (d *Data) setNewName(newName string) {
	d.Name = newName
}

//Функция обработчик генерирует стартовую страницу.
func homePage(w http.ResponseWriter, r *http.Request) {
	roman := Data{"Roman", 38}
	roman.setNewName("RoMaN")
	t, err := template.ParseFiles("./html/index.html", "./html/header.html", "./html/footer.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	db, _ := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/gobd")
	defer db.Close()
	// Выборка данных
	res, err := db.Query("SELECT * FROM `users`")
	if err != nil {
		panic(err)
	}
	basa = []User{}
	for res.Next() {
		var post User
		err = res.Scan(&post.Id, &post.Name, &post.Age)
		if err != nil {
			panic(err)
		}
		basa = append(basa, post)
	}
	t.ExecuteTemplate(w, "index", roman)
}

func regSql(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./html/form.html", "./html/header.html", "./html/footer.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	t.ExecuteTemplate(w, "form", nil)
}

func saveSql(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	age := r.FormValue("age")

	if name == "" || age == "" {
		fmt.Fprintf(w, "Не все данные заполнены")
	} else {
		db, _ := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/gobd")
		defer db.Close()

		// Установку данных
		insert, err := db.Query(fmt.Sprintf("INSERT INTO `users` (`name`, `age`) VALUES('%s', '%v')", name, age))
		if err != nil {
			panic(err)
		}
		defer insert.Close()

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func resumePage(w http.ResponseWriter, r *http.Request) {
	//создаем html-шаблон
	html := template.Must(template.ParseFiles("./html/user.html"))

	//выводим шаблон клиенту в браузер
	buf := &bytes.Buffer{}
	err := html.Execute(buf, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	buf.WriteTo(w)
}

//Функция обработчик генерирует динамическую страницу зарегистрированную по адресу /time, которая отображает текущую дату и время.
func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	Body := "Текущее время:"
	fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", Body)
	fmt.Fprintf(w, "<h2 align=\"center\">%s</h2>\n", t)
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served time for: %s\n", r.Host)
	path := "./html/time.html"
	http.ServeFile(w, r, path)
}

func HandleFunc() {
	r := mux.NewRouter()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog.Printf("Запуск сервера на порту: %s\n", PORT)

	r.HandleFunc("/", homePage).Methods("GET")
	r.HandleFunc("/reg", regSql).Methods("GET")
	r.HandleFunc("/regSql", saveSql).Methods("POST")
	r.HandleFunc("/resume", resumePage).Methods("GET")
	r.HandleFunc("/time", timeHandler).Methods("GET")

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./static/")})
	r.Handle("/static/", http.NotFoundHandler())
	r.Handle("/static/", http.StripPrefix("/static/", fileServer))

	//Запускаем веб-сервер.
	srv := &http.Server{
		Handler:      r,
		Addr:         PORT,
		ErrorLog:     errorLog,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}

func main() {
	HandleFunc()
}
