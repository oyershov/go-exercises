package main

import (
	gw "./gowiki"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/view/", gw.MakeHandler(gw.ViewHandler))
	http.HandleFunc("/edit/", gw.MakeHandler(gw.EditHandler))
	http.HandleFunc("/save/", gw.MakeHandler(gw.SaveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
