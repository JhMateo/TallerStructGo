package main

import (
	"github.com/JhMateo/tallerStruct/reader"
	"github.com/JhMateo/tallerStruct/views"
)

/*	Integrantes:
	Jhon Mauro Guevara - 160004315
	Joshep Mateo Granada - 160004314	*/

func main() {
	estudiantes, err := reader.ReadJson("generated.json")
	if err != nil {
		panic(err)
	}

	views.ExecuteWeb(estudiantes)
}
