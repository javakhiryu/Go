package main

import (
	"os"
	"text/template"
)

func main() {
	// Создание нового шаблона t1 и его парсинг
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	// Выполнение шаблона t1 с передачей строки "some text"
	t1.Execute(os.Stdout, "some text")

	// Функция для создания шаблона и обработки ошибок
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	// Создание и выполнение шаблона t2 с использованием структуры
	t2 := Create("t2", "Name: {{.Name}}\n")

	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	// Выполнение шаблона t2 с использованием карты
	t2.Execute(os.Stdout, map[string]string{
		"Name": "John Doe",
	})

	// Создание и выполнение шаблона t3 с условным оператором
	t3 := Create("t3", "{{if . -}} Not empty: {{.}} {{else -}} empty {{end}}\n")
	t3.Execute(os.Stdout, "")
	t3.Execute(os.Stdout, "apple")

	// Создание и выполнение шаблона t4 с циклом range
	t4 := Create("t4", "Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout, []string{"apple", "banana", "orange"})
}

//Этот пример демонстрирует мощные возможности шаблонов в Go, включая вставку значений, использование структур и карт, условные операторы и циклы. Шаблоны полезны для форматирования текста, генерации HTML и других сценариев, требующих динамического создания текстовых данных.
