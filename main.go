package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/rakyll/statik/fs"
	_ "github.com/warawara28/test_statik/statik"
)

const (
	filename        = "/index.html"
	defaultBodyText = "Hello World!"
	address         = ":8080"
)

func main() {
	index, err := loadTemplate(filename)
	if err != nil {
		log.Fatalf("failed to load template file %s:%s", filename, err.Error())
		os.Exit(1)
	}

	log.Printf("listen to port %s", address)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		bodyText := r.URL.Query().Get("body")
		if bodyText == "" {
			bodyText = defaultBodyText
		}

		var body bytes.Buffer
		if err := index.Execute(&body, bodyText); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(body.String()))
	})
	http.ListenAndServe(address, nil)
}

func loadTemplate(filepath string) (*template.Template, error) {
	statikFS, err := fs.New()
	if err != nil {
		return nil, err
	}

	r, err := statikFS.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	contents, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	tpl, err := template.New(filepath).Parse(string(contents))
	if err != nil {
		return nil, fmt.Errorf("Failed to parse %s file: %w", filepath, err)
	}
	return tpl, nil
}
