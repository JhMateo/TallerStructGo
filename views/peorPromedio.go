package views

import (
	"fmt"
	"github.com/JhMateo/tallerStruct/reader"
	"github.com/JhMateo/tallerStruct/taller"
	"html/template"
	"net/http"
)

// PeorPromedioHandler maneja la solicitud para mostrar el estudiante con peor promedio en la página web.
func PeorPromedioHandler(w http.ResponseWriter, r *http.Request, estudiantes []reader.Estudiante) {
	peorEstudiante, peorPromedio := taller.PeorPromedio(estudiantes)

	data := struct {
		Nombre   string
		Apellido string
		Promedio float64
	}{
		Nombre:   peorEstudiante.Nombre,
		Apellido: peorEstudiante.Apellido,
		Promedio: peorPromedio,
	}

	template, err := template.ParseFiles("templates/peor_promedio.html")
	if err != nil {
		fmt.Fprintf(w, "<h1>Página no encontrada!</h1>")
		panic(err)
	} else {
		template.Execute(w, data)
	}
}
