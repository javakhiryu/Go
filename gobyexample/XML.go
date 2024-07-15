// Этот код на Go демонстрирует использование пакета encoding/xml для сериализации (маршалинга) и десериализации (анмаршалинга) структур в XML формат и из него.
package main

import (
	"encoding/xml"
	"fmt"
)

// Plant представляет собой структуру растения, которая будет сериализована в XML.
type Plant struct {
	XMLName xml.Name `xml:"plant`     // Название XML-элемента, соответствующего этой структуре.
	Id      int      `xml:"id, attr"` // Поле 'Id' будет представлено как атрибут XML.
	Name    string   `xml: "name"`    // Поле 'Name' будет элементом XML.
	Origin  []string `xml:"origin"`   // Поле 'Origin' будет элементом XML.
}

// Метод String для структуры Plant, который возвращает строковое представление структуры.
func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v", p.Id, p.Name, p.Origin)
}

func main() {
	// Создание объекта Plant.
	coffee := &Plant{
		Id:   1,
		Name: "Coffee",
	}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	// Сериализация объекта Plant в XML.
	out, _ := xml.MarshalIndent(coffee, " ", " ")
	fmt.Printf("%v\n", string(out))

	// Десериализация XML обратно в объект Plant.
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	// Создание ещё одного объекта Plant.
	tomato := &Plant{Id: 81, Name: "Tamoto"}
	tomato.Origin = []string{"Mexico", "California"}

	// Определение структуры, включающей вложенные элементы.
	type Nexting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	// Создание объекта Nesting, содержащего два Plant.
	nesting := &Nexting{}
	nesting.Plants = []*Plant{coffee, tomato}

	// Сериализация объекта Nesting в XML.
	out, _ = xml.MarshalIndent(nesting, " ", " ")
	fmt.Println(string(out))
}
