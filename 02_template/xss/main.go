package main

import (
	"html/template"
	"log"
	"os"
)

type Page struct {
	Title   string
	Heading string
	Input   string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	home := Page{
		Title:   "Escaped",
		Heading: "Danger is escaped with html/template",
		Input:   `<script>alert("Yeah!");</script>`,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.html", home)
	if err != nil {
		log.Fatalln(err)
	}
}
