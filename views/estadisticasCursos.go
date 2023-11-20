package views

import (
	"fmt"
	"github.com/JhMateo/tallerStruct/reader"
	"github.com/JhMateo/tallerStruct/taller"
	"html/template"
	"net/http"
)

// EstadisticasPorCursoHandler maneja la solicitud para mostrar las estadísticas por curso.
func EstadisticasPorCursoHandler(w http.ResponseWriter, r *http.Request, estudiantes []reader.Estudiante) {
	estadisticasPorCurso := taller.CalcularEstadisticasPorCurso(estudiantes)

	taller.ImprimirEstadisticasPorCurso(estudiantes)

	template, err := template.ParseFiles("templates/estadisticas_cursos.html")
	if err != nil {
		fmt.Fprintf(w, "<h1>Página no encontrada!</h1>")
		panic(err)
	} else {
		template.Execute(w, estadisticasPorCurso)
	}
}
