package views

import (
	"fmt"
	"github.com/JhMateo/tallerStruct/reader"
	"github.com/JhMateo/tallerStruct/taller"
	"html/template"
	"net/http"
)

// HombreMayorHandler maneja la solicitud para mostrar al estudiante masculino de mayor edad.
func HombreMayorHandler(w http.ResponseWriter, r *http.Request, estudiantes []reader.Estudiante) {
	estudiantesMasculinosMayorEdad := taller.EncontrarEstudiantesMayorEdadPorGenero(estudiantes, "male")

	taller.ImprimirEstudianteMasculino(estudiantes)

	template, err := template.ParseFiles("templates/hombre_mayor.html")
	if err != nil {
		fmt.Fprintf(w, "<h1>Página no encontrada!</h1>")
		panic(err)
	} else {
		template.Execute(w, estudiantesMasculinosMayorEdad)
	}
}
