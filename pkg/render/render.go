package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/cornelia247/luna/pkg/config"
	"github.com/cornelia247/luna/pkg/models"
)

var app *config.AppConfig

// NewTemplates set the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a

}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate using the http/template package
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {

		tc = app.TemplateCache

	} else {
		tc, _ = CreateTemplateCached()
	}

	// get template cache from the AppConfig

	//Get requested template from cache

	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("could not get templates from templates cache")
	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td)
	_ = t.Execute(buf, td)

	//render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
}

func CreateTemplateCached() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	// get all the files naem *.page.tmpl frpm the ./templates folder

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	//range through the pages
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts

	}

	return myCache, nil
}

// var tc = make(map[string]*template.Template)
// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error
// 	// check and see if we already have template in cache
// 	_, inMap := tc[t]
// 	if !inMap {
// 		log.Println("Creating Template and adding to cache")
// 		// create template
// 		err = createTemplateCached(t)
// 		if err  != nil {
// 			log.Println(err)
// 		}

// 	} else {
// 		// template in cache
// 		log.Println("Using cached template")
// 	}

// 	tmpl = tc[t]
// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}

// }
// func createTemplateCached(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t), "./templates/base.layout.tmpl",

// 	}

// 	//Parse the templates

// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}
// 	// add template to cache

// 	tc[t] = tmpl
// 	return  nil

// }
