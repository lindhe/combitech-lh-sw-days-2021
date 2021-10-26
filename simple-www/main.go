package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

const (
	defaultHTTPPort = "80"
	defaultColor    = "red"
)

// CSSStyles fileds must be exported, to let Go Templates parse them.
type CSSStyles struct {
	BackgroundColor string
}

func getHTML(w http.ResponseWriter, r *http.Request) {

	// Set background color
	style := CSSStyles{BackgroundColor: defaultColor}
	if color := os.Getenv("BACKGROUND_COLOR"); color != "" {
		style = CSSStyles{BackgroundColor: color}
	}

	if html, err := template.ParseFiles("index.html.gotmpl"); err == nil {
		if err = html.Execute(w, style); err != nil {
			panic(err)
		}
	} else {
		panic(err)
	}

	log.Printf("Served a %v page!", style.BackgroundColor)

}

func main() {

	// Set HTTP port
	httpPort := defaultHTTPPort
	if p := os.Getenv("HTTP_PORT"); p != "" {
		httpPort = p
	}

	http.HandleFunc("/", getHTML)

	log.Fatal(http.ListenAndServe(":"+httpPort, nil))

}
