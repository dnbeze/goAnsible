package render

import (
	"bytes"
	"goAnsible/pkg/config"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// package level variable for template cache
//var tc = make(map[string]*template.Template) // this is creating variable that will hold template cache

//this is a simple template render function replaced below with template cache
/*
func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.html.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}
*/
var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	var tc map[string]*template.Template

	if app.UseCache {
		// get template cache from app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok { //check if html page was found in template cache
		log.Fatal("could not get template from template cache")
	}
	// create buffer for finer grained error checking
	buf := new(bytes.Buffer)

	err := t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	//myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	//get all the files named *.html
	pages, err := filepath.Glob("./templates/*.html")
	if err != nil {
		return myCache, err
	}

	//range through all files *.html
	for _, page := range pages { //pages is a slice of strings that is the full path to all files *.html in the templates folder
		name := filepath.Base(page)                    //set name to the last element of the file path
		ts, err := template.New(name).ParseFiles(page) // then use name to populate the template name and then parse it which parses the filename aka var name
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
