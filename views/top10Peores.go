package views

import (
	"fmt"
	"github.com/JhMateo/tallerStruct/reader"
	"github.com/JhMateo/tallerStruct/taller"
	"html/template"
	"net/http"
)

// Top10PeoresPorCursoHandler maneja la solicitud para mostrar el top 10 de estudiantes con peores notas por curso.
func Top10PeoresPorCursoHandler(w http.ResponseWriter, r *http.Request, estudiantes []reader.Estudiante) {
	top10PeoresPorCurso := taller.Top10PeoresPorCurso(estudiantes)

	template, err := template.ParseFiles("templates/top10_peores.html")
	if err != nil {
		fmt.Fprintf(w, "<h1>Página no encontrada!</h1>")
		panic(err)
	} else {
		template.Execute(w, top10PeoresPorCurso)
	}
}
