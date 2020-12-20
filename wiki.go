package main

import (
	"log"
	"net/http"

	page "github.com/masayuki-0319/gowiki/page"
)

func main() {
	http.HandleFunc("/view/", page.MakeHandler(page.ViewHandler))
	http.HandleFunc("/edit/", page.MakeHandler(page.EditHandler))
	http.HandleFunc("/save/", page.MakeHandler(page.SaveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
