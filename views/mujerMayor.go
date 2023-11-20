package views

import (
	"fmt"
	"github.com/JhMateo/tallerStruct/reader"
	"github.com/JhMateo/tallerStruct/taller"
	"html/template"
	"net/http"
)

// MujerMayorHandler maneja la solicitud para mostrar a la estudiante femenina de mayor edad.
func MujerMayorHandler(w http.ResponseWriter, r *http.Request, estudiantes []reader.Estudiante) {
	estudiantesFemeninosMayorEdad := taller.EncontrarEstudiantesMayorEdadPorGenero(estudiantes, "female")

	taller.ImprimirEstudianteFemenino(estudiantes)

	template, err := template.ParseFiles("templates/mujer_mayor.html")
	if err != nil {
		fmt.Fprintf(w, "<h1>Página no encontrada!</h1>")
		panic(err)
	} else {
		template.Execute(w, estudiantesFemeninosMayorEdad)
	}
}
