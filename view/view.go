// Package view provides utilities that simplify rendering
// HTML templates to a http.ResponseWriter.
package view

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// A factory for creating Views. The factory should be
// given defaults tht views may fallback to if none other
// are given.
type Factory struct {
	defaultSiteTitle string
	templatePath string
}

// Create a new factory.
func NewFactory(defaultSiteTitle, templatePath string) Factory {
	return Factory{
		defaultSiteTitle: defaultSiteTitle,
		templatePath: templatePath,
	}
}

// Create a new View for the supplied template.
func (f Factory) New(template string) View {
	return View {
		templatePath: f.templatePath,
		template: template,
		title:    f.defaultSiteTitle,
		data:     nil,
	}
}

// View represents a HTML page that can be rendered given a
// http.ResponseWriter. View should contain all data that
// is necessary to render a template. If a template
// requires some specific data it have to be supplied using
// the method Data(interface{}). The rendered pages title
// can be altered using the method Title(string).
type View struct {
	template string
	title    string
	data     interface{}
	templatePath string

}

// Change the title of the rendered page to title.
func (p View) Title(title string) View {
	p.title = title
	return p
}

// Data(interface{}) is used to supply data for view
// templates that require it. If the template do not
// require any extra, calling this method will not do
// anything.
func (p View) Data(data interface{}) View {
	p.data = data
	return p
}

// Get the data supplied. Mainly used to render the view.
func (p View) GetData() interface{} {
	return p.data
}

// Get the title the view will be rendered with. Mainly
// used to render the view.
func (p View) GetTitle() string {
	return p.title
}

// Internal helper method to load the a complete template.
// A complete template consists of the common _base.gohtml
// template and a view specific sub-template.
func (p View) getTemplate() (*template.Template, error) {
	templates := []string{"_base.gohtml", p.template}
	for i, t := range templates {
		templates[i] = p.templatePath + string(filepath.Separator) + t
	}

	return template.ParseFiles(templates...)
}

// Render this view to the supplied response writer. If
// anything goes wrong while rendering the view, an HTTP
// 500 Internal Server Error will be written instead.
func (p View) Render(w http.ResponseWriter) {
	t, err := p.getTemplate()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	err = t.ExecuteTemplate(w, "_base.gohtml", p)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
}