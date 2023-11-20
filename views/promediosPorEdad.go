package views

import (
	"fmt"
	"github.com/JhMateo/tallerStruct/reader"
	"github.com/JhMateo/tallerStruct/taller"
	"html/template"
	"net/http"
)

// PromediosPorRangoEdadHandler maneja la solicitud para mostrar el promedio de estudiantes por rango de edad.
func PromediosPorRangoEdadHandler(w http.ResponseWriter, r *http.Request, estudiantes []reader.Estudiante) {
	resultados := taller.PromediosPorRangoEdad(estudiantes)

	template, err := template.ParseFiles("templates/promedios_edad.html")
	if err != nil {
		fmt.Fprintf(w, "<h1>Página no encontrada!</h1>")
		panic(err)
	} else {
		template.Execute(w, resultados)
	}
}
