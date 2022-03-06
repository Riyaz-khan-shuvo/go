package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func init() {

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err = sql.Open("mysql", "root:123456Ri@tcp(127.0.0.1:3306)/hosting_db")
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	// defer db.Close()

	// insert, err := db.Query("INSERT INTO `request` (`id`, `name`, `company`, `email`, `status`) VALUES (NULL, 'Riyaz Hossain', 'riyaz dot com', 'mdriyaz5965@gmail.com', '1');")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// defer insert.Close()
	fmt.Println("Database connected Successful!!!")
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/features", features)
	http.HandleFunc("/docs", docs)
	http.HandleFunc("/request", request)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8888", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {

	//pToTem = pointer to template

	pToTem, err := template.ParseFiles("./template/base.gohtml")

	if err != nil {
		fmt.Println(err.Error())
	}
	pToTem.Execute(w, nil)

	fmt.Fprintf(w, "Welcome , to the home page ")
}

func features(w http.ResponseWriter, r *http.Request) {
	pToTem, err := template.ParseFiles("./template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	pToTem, err = pToTem.ParseFiles("./wpage/features.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	pToTem.Execute(w, nil)
}

func docs(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("wpage/docs.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp.Execute(w, nil)
}

func request(w http.ResponseWriter, r *http.Request) {

	// method - 1 is not good

	// name := r.FormValue("name")
	// company := r.FormValue("company")
	// email := r.FormValue("email")

	// fmt.Println("name	: ", name, "\nCompany	:", company, "\nEmail	:", email)
	// fmt.Fprintf(w, "Received")

	// method - 2 advance way

	r.ParseForm()
	for key, value := range r.Form {
		fmt.Println(key, value)
	}

	// work with database
	name := r.FormValue("name")
	company := r.FormValue("company")
	email := r.FormValue("email")

	qs := "INSERT INTO `request` (`id`, `name`, `company`, `email`, `status`) VALUES (NULL, '%s', '%s', '%s', '1');"
	sql := fmt.Sprintf(qs, name, company, email)

	fmt.Println(sql)

	insert, err := db.Query(sql)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}
