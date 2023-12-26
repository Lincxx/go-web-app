package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// create a template cache
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	//get requested template from cahce
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	err = t.Execute(buf, nil)

	if err != nil {
		log.Println(err)
	}

	//render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

	//parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.gohtml")
	//err := parsedTemplate.Execute(w, nil)
	//
	//if err != nil {
	//	fmt.Println("error parsing template", err)
	//}

}

func createTemplateCache() (map[string]*template.Template, error) {
	//myCache := make(map[string]*template.Template)
	//these are the same
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.gohtml from ./templates
	pages, err := filepath.Glob("./templates/*.page.gohtml")

	if err != nil {
		return myCache, err
	}

	// range through all file ending with *.page.gohtml
	for _, page := range pages {
		name := filepath.Base(page)
		//parse file
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		//look for layout
		matches, err := filepath.Glob("./templates/*.layout.gohtml")

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}

//var tc = make(map[string]*template.Template)
//
//func RenderTemplate(w http.ResponseWriter, t string) {
//	var tmpl *template.Template
//	var err error
//
//	//check to see if we already have the templage in our cache
//	_, inMap := tc[t]
//	if !inMap {
//		//need to create the template
//		log.Println("creating template and adding to cache")
//		err = createTemplateCache(t)
//		if err != nil {
//			log.Println(err)
//		}
//	} else {
//		//we have the template in the cache
//		log.Println("using cached template")
//	}
//
//	tmpl = tc[t]
//
//	err = tmpl.Execute(w, nil)
//	if err != nil {
//		log.Println(err)
//	}
//}
//
//func createTemplateCache(t string) error {
//	templates := []string{
//		fmt.Sprintf("./templates/%s", t),
//		"./templates/base.layout.gohtml",
//	}
//
//	//parse the template
//	tmpl, err := template.ParseFiles(templates...)
//
//	if err != nil {
//		return err
//	}
//
//	//add template to cache (map)
//	tc[t] = tmpl
//
//	return nil
//}
