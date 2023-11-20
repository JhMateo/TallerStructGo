package views

import (
	"fmt"
	"github.com/JhMateo/tallerStruct/reader"
	"github.com/JhMateo/tallerStruct/taller"
	"html/template"
	"net/http"
)

// MatriculadosHandler maneja la solicitud para mostrar los estudiantes que se matricularon en el año pasado (2022).
func MatriculadosHandler(w http.ResponseWriter, r *http.Request, estudiantes []reader.Estudiante) {
	estudiantesMatriculados := taller.Matriculados2022(estudiantes)

	var resultados []struct {
		Nombre             string
		FechaMatriculacion string
	}

	for nombre, fecha := range estudiantesMatriculados {
		resultado := struct {
			Nombre             string
			FechaMatriculacion string
		}{nombre, fecha}
		resultados = append(resultados, resultado)
	}

	template, err := template.ParseFiles("templates/matriculados.html")
	if err != nil {
		fmt.Fprintf(w, "<h1>Página no encontrada!</h1>")
		panic(err)
	} else {
		template.Execute(w, resultados)
	}
}
