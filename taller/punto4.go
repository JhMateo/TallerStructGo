package taller

import (
	"github.com/JhMateo/tallerStruct/reader"
)

/*
	4. Top 10 estudiantes con peores notas de cada curso.
*/

func Top10PeoresPorCurso(estudiantes []reader.Estudiante) map[string][]NotaEstudiante {
	return procesarNotas(estudiantes, func(nota1, nota2 float64) bool {
		return nota1 < nota2
	})
}
