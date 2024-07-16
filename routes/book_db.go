package routes

import (
	"errors"
	"fmt"

	"github.com/astaxie/beego/orm"
	"github.com/harshad-salunke2002/books_api/models"
)

type BookDB struct {
	OrmDAo orm.Ormer
}

func (db *BookDB) InsertBook(book *models.Books) (bool, error) {
	_, err := db.OrmDAo.Insert(book)
	if err != nil {
		return false, err
	}
	return true, nil

}

func (db *BookDB) ReadBook(id int) (*models.Books, error) {
	book := models.Books{Id: id}
	err := db.OrmDAo.Read(&book)
	return &book, err
}

func (db *BookDB) UpdateBookDB(id int, book *models.Books) error {

	userToUpdate, err := db.ReadBook(id)
	fmt.Println(userToUpdate)

	if err != nil {
		return errors.New("book not found")
	}

	userToUpdate.Name = book.Name
	userToUpdate.Pages = book.Pages
	userToUpdate.Writer = book.Writer

	fmt.Println(userToUpdate)

	_, error := db.OrmDAo.Update(&userToUpdate)

	return error

}

func (db *BookDB) DeleteBookbyId(id int) error {
	book := models.Books{Id: id}
	if _, err := db.OrmDAo.Delete(&book); err == nil {
		return nil
	}
	return fmt.Errorf("book not found")
}

func (db *BookDB) ReadAllBook() ([]models.Books, error) {
	var books []models.Books
	_, err := db.OrmDAo.QueryTable(new(models.Books)).All(&books)
	return books, err
}
