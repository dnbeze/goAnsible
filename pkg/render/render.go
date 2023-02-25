package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tc = make(map[string]*template.Template) // this is creating variable that will hold template cache

//this is a simple template render function replaced below with template cache
/*
func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.html")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}
*/

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	//check to see if we have the template in our cache which is var tc
	_, inMap := tc[t] //inMap will try to pull the value that is at key position t from the map tc. Can use inMap as value or bool now. Ignoring index aka key
	if !inMap {
		//need to create the template
		log.Println("creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		//we have the template in the cache
		log.Println("using cached template")
	}

	tmpl = tc[t] // set var tmpl to the value of tc[t] which is of type template.Template because the map you made when initialized tc

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}

}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.html",
	}
	//parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	//add template to cache (map)
	tc[t] = tmpl

	return nil
}
