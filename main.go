package main

import (
	"log"
	"net/http"

	"github.com/astaxie/beego/orm"
	"github.com/gorilla/mux"
	"github.com/harshad-salunke2002/books_api/models"
	"github.com/harshad-salunke2002/books_api/routes"

	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDataBase("default", "postgres", "postgres://postgres:harshad2002@localhost:5432/books?sslmode=disable")
	orm.RegisterModel(new(models.Books))
	orm.RunSyncdb("default", false, true)
}

func main() {

	bookDB := routes.BookDB{
		OrmDAo: orm.NewOrm(),
	}

	main_router := mux.NewRouter()

	book_router := main_router.PathPrefix("/book").Subrouter()
	book_router.HandleFunc("/addbook", bookDB.AddBook).Methods("POST")
	book_router.HandleFunc("/getbook/{bookId}", bookDB.GetBookByID).Methods("GET")
	book_router.HandleFunc("/books", bookDB.GetBooks).Methods("GET")
	book_router.HandleFunc("/update/{bookId}", bookDB.UpdateBook).Methods("PATCH")
	book_router.HandleFunc("/delete/{bookId}", bookDB.DeleteBook).Methods("DELETE")

	// auther_routes:=main_router.PathPrefix("author").Subrouter()
	// auther_routes.HandleFunc("/add")

	server := http.Server{
		Addr:    ":8080",
		Handler: main_router,
	}
	log.Println("Server starting")
	server.ListenAndServe()

}
