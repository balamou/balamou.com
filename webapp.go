package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body []byte
}

func loadPage(title string) *Page {
    filename := "static/" + title + ".html"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil
    }
    return &Page{Title: title, Body: body}
}

func router(request string) (*Page) {
	switch request {
	case "indulge":
		return loadPage("indulge")
	case "":
		return loadPage("index")
	default:
		return nil
	}
}

func handler(writer http.ResponseWriter, request *http.Request) {
		page := router(request.URL.Path[1:])

		if page != nil {
			writer.Write(page.Body)
			fmt.Println(request.URL.Path[1:]) // debug
		}
}

func main() {
		port := os.Getenv("PORT")
		fs := http.FileServer(http.Dir("static"))
		
	  http.Handle("/", fs)
    http.HandleFunc("/indulge", handler)
    log.Fatal(http.ListenAndServe(":" + port, nil))
}
