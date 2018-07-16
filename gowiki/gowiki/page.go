package gowiki

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
)

var FILE_PATH = "./templates/"
var RES_PATH = "./resources/"
var templates = template.Must(template.ParseFiles(FILE_PATH+"edit.html", FILE_PATH+"view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) Save() error {
	filename := RES_PATH + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func LoadPage(title string) (*Page, error) {
	filename := RES_PATH + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func RenderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
