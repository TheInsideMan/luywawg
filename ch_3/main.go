package main

import (
	"html/template"
	"os"
)

type Product struct {
	Price    float64
	Quantity float64
}
type Article struct {
	Name       string
	AuthorName string
	Draft      bool
	Price      float64
}

func Multiply(a, b float64) float64 {
	return a * b
}
func main() {
	// goArticle := []Article{}
	// goArticle = append(goArticle, Article{"The Go html/template package", "Mal Curtis", true, 17.89})
	// goArticle = append(goArticle, Article{"The way of the Wolf", "Aaron Kenny", false, 12.3})
	// goArticle = append(goArticle, Article{"Magic is the way", "My Little Pony", true, 9.99})
	// tmpl, err := template.New("Foo").Parse(`
	// 	{{define "ArticleResource"}}
	// 	<p>{{.Name}} by {{.AuthorName}}{{if .Draft}} (DRAFT){{end}}. Price: ${{printf "%.2f" .Price}}</p>
	// 	{{end}}
	// 	{{define "ArticleLoop"}}
	// 		{{range .}}
	// 			{{template "ArticleResource" .}}
	// 		{{else}}
	// 			<p>No Published articles yet</p>
	// 		{{end}}
	// 	{{end}}
	// 	{{template "ArticleLoop" .}}
	// 	`)

	// if err != nil {
	// 	panic(err)
	// }
	// err = tmpl.Execute(os.Stdout, goArticle)
	// if err != nil {
	// 	panic(err)
	// }

	tmpl := template.New("Foo")
	// tmpl.Funcs accepts a map of functions
	tmpl.Funcs(template.FuncMap{
		"multiply": Multiply,
	})
	tmpl, err := tmpl.Parse(
		"Price: ${{multiply .Price .Quantity | printf \"%.2f\"}}\n")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, Product{12.3, 2})
	if err != nil {
		panic(err)
	}
}
