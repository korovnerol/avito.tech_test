package main

import (
	"fmt"
	"net/http"
	"log"
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "22092003" //Пароль вместо your_password
)


func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users/{id}", infoUser).Methods("GET")
	/*router.HandleFunc("/users/{new_slug}/{del_slug}/{id}", updateUser).Methods("PUT")*/

	router.HandleFunc("/segments/create/{slug}", createSegment).Methods("PUT")
	router.HandleFunc("/segments/delete/{slug}", deleteSegment).Methods("DELETE")

	fmt.Println("Server started on port 80")
	http.ListenAndServe(":80", router)
}

func infoUser(w http.ResponseWriter, r *http.Request) {
	var result string
	params := mux.Vars(r)
	id := params["id"]
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=users sslmode=disable", host, port, user, password)
	// Устанавливаем подключение к базе данных
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Получаем данные
	res := db.QueryRow("SELECT segments FROM users WHERE id=?", id)
	res.Scan(result)
	fmt.Fprintf(w, result)
	// Информация о сегментах пользователя
}

/*func updateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	new_slugs := params["new_slug"]
	// Изменение сегментов пользователя
}*/

func createSegment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	slug := params["slug"]
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=segments sslmode=disable", host, port, user, password)
	// Устанавливаем подключение к базе данных
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Обновляем данные в таблице
	_, err = db.Exec("INSERT INTO segments($1)", slug)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Creating a new segment")
	// Создание нового сегмента
}

func deleteSegment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	slug := params["slug"]
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=segments sslmode=disable", host, port, user, password)
	// Устанавливаем подключение к базе данных
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Обновляем данные в таблице
	_, err = db.Exec("DELETE FROM segments WHERE name_of_segment=$1", slug)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleting a segment")
	// Удаление сегмента
}