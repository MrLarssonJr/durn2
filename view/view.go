package view

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type Factory struct {
	defaultSiteTitle string
	templatePath string
}

func NewFactory(defaultSiteTitle, templatePath string) Factory {
	return Factory{
		defaultSiteTitle: defaultSiteTitle,
		templatePath: templatePath,
	}
}

func (f Factory) New(template string) View {
	return View {
		templatePath: f.templatePath,
		template: template,
		title:    f.defaultSiteTitle,
		data:     nil,
	}
}

type View struct {
	template string
	title    string
	data     interface{}
	templatePath string

}

func (p View) Title(title string) View {
	p.title = title
	return p
}

func (p View) Data(data interface{}) View {
	p.data = data
	return p
}

func (p View) GetData() interface{} {
	return p.data
}

func (p View) GetTitle() string {
	return p.title
}

func (p View) getTemplate() (*template.Template, error) {
	templates := []string{"_base.gohtml", p.template}
	for i, t := range templates {
		templates[i] = p.templatePath + string(filepath.Separator) + t
	}

	return template.ParseFiles(templates...)
}

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