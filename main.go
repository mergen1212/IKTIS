package main

import (
	"Go+HTMX/dbInter"
	"Go+HTMX/handle"
	"fmt"
	"log"
	"net/http"
)

func main() {
	db, err := dbInter.Conn()
	if err != nil {
		log.Println(err)
	}
	//db, err = dbInter.Parse(db)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("--------------------pijioij------------------------------------")
	http.HandleFunc("/", handle.YourHandler(db))
	http.HandleFunc("/kkk", handle.YourHandle())
	http.ListenAndServe(":8080", nil)
}
