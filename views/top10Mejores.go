package views

import (
	"fmt"
	"github.com/JhMateo/tallerStruct/reader"
	"github.com/JhMateo/tallerStruct/taller"
	"html/template"
	"net/http"
)

// Top10MejoresPorCursoHandler maneja la solicitud para mostrar el top 10 de estudiantes con mejores notas por curso.
func Top10MejoresPorCursoHandler(w http.ResponseWriter, r *http.Request, estudiantes []reader.Estudiante) {
	top10MejoresPorCurso := taller.Top10MejoresPorCurso(estudiantes)

	template, err := template.ParseFiles("templates/top10_mejores.html")
	if err != nil {
		fmt.Fprintf(w, "<h1>Página no encontrada!</h1>")
		panic(err)
	} else {
		template.Execute(w, top10MejoresPorCurso)
	}
}
