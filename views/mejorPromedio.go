package views

import (
	"fmt"
	"github.com/JhMateo/tallerStruct/reader"
	"github.com/JhMateo/tallerStruct/taller"
	"html/template"
	"net/http"
)

// MejorPromedioHandler maneja la solicitud para mostrar el estudiante con mejor promedio en la página web.
func MejorPromedioHandler(w http.ResponseWriter, r *http.Request, estudiantes []reader.Estudiante) {
	mejorEstudiante, mejorPromedio := taller.MejorPromedio(estudiantes)

	data := struct {
		Nombre   string
		Apellido string
		Promedio float64
	}{
		Nombre:   mejorEstudiante.Nombre,
		Apellido: mejorEstudiante.Apellido,
		Promedio: mejorPromedio,
	}

	template, err := template.ParseFiles("templates/mejor_promedio.html")
	if err != nil {
		fmt.Fprintf(w, "<h1>Página no encontrada!</h1>")
		panic(err)
	} else {
		template.Execute(w, data)
	}
}
