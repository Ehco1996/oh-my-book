package controllers

import (
	"github.com/Ehco1996/oh-my-book/models"
	"github.com/astaxie/beego"
	"strings"
)

type BookRender struct {
	beego.Controller
}

func (c *BookRender) Get() {

	bookName := c.GetString("name")
	book, err := models.GetOrCreateBookByName(bookName)
	if err != nil {
		c.Abort("500")
	}
	c.Data["title"] = book.Name + book.LastChapterName
	c.Data["article_name"] = book.LastChapterName
	text := models.GetChapterTextFromApi(book.LastChapterLink)
	article_paragraphs := strings.Split(text, "\n\u003cbr\u003e\n\u003cbr\u003e\u0026nbsp;\u0026nbsp;\u0026nbsp;\u0026nbsp;")
	for idx, c := range article_paragraphs {
		article_paragraphs[idx] = strings.ReplaceAll(c,"&nbsp;","")
	}
	c.Data["article_paragraphs"] = article_paragraphs
	c.TplName = "render.html"
}
