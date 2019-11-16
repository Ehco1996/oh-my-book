package models

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"time"
)

func getJson(url string, target interface{}) error {
	c := &http.Client{Timeout: 10 * time.Second}
	r, err := c.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

const (
	BookDetailUrl = "http://qingmo.zohar.space/2333/新笔趣阁/search?"
)

type bookDetailResp struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	Summary     string `json:"summary"`
	ChapterLink string `json:"link"`
}

type chapterListResp struct {
	Catalogs []struct {
		ChapterName string `json:"chapterName"`
		ChapterLink string `json:"chapterlink"`
	} `json:"catalogs"`
}

type chapterTextResp struct {
	Text string `json:"text"`
}

func getBookDetailFromApi(name string) (error, *bookDetailResp) {
	params := url.Values{}
	params.Add("key", name)
	respList := []*bookDetailResp{}
	err := getJson(BookDetailUrl+params.Encode(), &respList)
	if err != nil || len(respList) < 1 || respList[0].Name != name {
		return errors.New("search failed"), nil
	}
	return nil, respList[0]
}

func GetLastChapterNameAndLinkFromApi(url string) (string,string) {
	resp := &chapterListResp{}
	err := getJson(url, resp)
	if err != nil {
		panic(err)
	}
	return resp.Catalogs[len(resp.Catalogs)-1].ChapterName,resp.Catalogs[len(resp.Catalogs)-1].ChapterLink
}

func GetChapterTextFromApi(url string) string {
	resp := &chapterTextResp{}
	err := getJson(url, resp)
	if err != nil {
		panic(err)
	}


	return resp.Text
}
