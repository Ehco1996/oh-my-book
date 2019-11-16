package controllers

import (
	"github.com/Ehco1996/oh-my-book/models"
	"github.com/astaxie/beego"
)

type Api struct {
	beego.Controller
}

type bookSearchResp struct {
	ErrMsg string
	models.Book
}

func (c *Api) Get() {
	bookName := c.GetString("name")

	resp := bookSearchResp{}
	msg := ""

	if bookName == "" {
		c.Ctx.ResponseWriter.WriteHeader(404)
		msg = "必须有name的query"
	} else {
		book, err := models.GetOrCreateBookByName(bookName)
		if err != nil {
			msg = err.Error()
			c.Ctx.ResponseWriter.WriteHeader(400)
		} else {
			cName, cLink := models.GetLastChapterNameAndLinkFromApi(book.ChapterListLink)
			book.LastChapterName = cName
			book.LastChapterLink = cLink
			book.Save()
			resp.Book = *book
		}
		resp.ErrMsg = msg
		c.Data["json"] = &resp
		c.ServeJSON()
	}
}
