package main

import (
	"os"
	"text/template"
)

func main() {
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	t1.Execute(os.Stdout, "some text")

	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	t2 := Create("t2", "Name: {{.Name}}\n")

	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	t2.Execute(os.Stdout, map[string]string{
		"Name": "John Doe",
	})

	t3 := Create("t3", "{{if . -}} Not empty: {{.}} {{else -}} empty {{end}}\n")
	t3.Execute(os.Stdout, "")
	t3.Execute(os.Stdout, "apple")

	t4 := Create("t4", "Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout, []string{"apple", "banana", "orange"})
}
