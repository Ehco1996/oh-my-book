package models

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Name            string `gorm:"unique_index"`
	Author          string
	Summary         string
	ChapterListLink string
	LastChapterName string
	LastChapterLink string
}

func GetOrCreateBookByName(name string) (*Book, error) {
	book := &Book{}
	err := db.Where("Name = ?", name).First(book).Error
	if err == nil {
		return book, err
	}

	err, resp := getBookDetailFromApi(name)
	if err != nil {
		return nil, err
	}

	book.Name = resp.Name
	book.Author = resp.Author
	book.Summary = resp.Summary
	book.ChapterListLink = resp.ChapterLink
	db.Create(book)
	return book, nil
}

func (book *Book) Save(){
	db.Save(book)
}