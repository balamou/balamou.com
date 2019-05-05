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
	case "playr":
		return loadPage("playr")
	case "gen_tree":
		return loadPage("gen_tree")
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

		if len(port) == 0 {
			fmt.Println("Global enviroment variable $PORT is empty. Using default port 80. ")
			port = "80"
		}

	  http.Handle("/", fs)
    http.HandleFunc("/indulge", handler)
		http.HandleFunc("/playr", handler)
		http.HandleFunc("/gen_tree", handler)
    log.Fatal(http.ListenAndServe(":" + port, nil))
}
