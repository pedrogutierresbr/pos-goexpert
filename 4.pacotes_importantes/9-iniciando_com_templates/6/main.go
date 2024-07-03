// Mapeando funções nos templates --> uma forma de registrar funções para serem utilizadas dentro dos templates
// cuidado ao usar o "html/template" --> ele sabe que a info vai ser parseada em html, e uma vez que ela sabe disso, ela vai tentar se blindar ao máximo de alguns tipos de ataque
// use quando vai realmente exibir algo no browser
package main

import (
	"html/template"
	"os"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	t := template.New("content.html")
	t.Funcs(template.FuncMap{"ToUpper": ToUpper})
	t = template.Must(t.ParseFiles(templates...))
	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Java", 20},
		{"Python", 10},
	})
	if err != nil {
		panic(err)
	}

}