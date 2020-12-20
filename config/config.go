package config

import (
	"log"
	"net/http"

	"github.com/masayuki-0319/gowiki/page"
)

func SetRoute() {
	http.HandleFunc("/view/", page.MakeHandler(page.ViewHandler))
	http.HandleFunc("/edit/", page.MakeHandler(page.EditHandler))
	http.HandleFunc("/save/", page.MakeHandler(page.SaveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
