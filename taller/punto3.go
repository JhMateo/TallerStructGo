package taller

import (
	"fmt"
	"github.com/JhMateo/tallerStruct/reader"
	"sort"
)

/*
	3. Top 10 estudiantes con mejores notas de cada curso.
*/

type NotaEstudiante struct {
	Nombre   string
	Apellido string
	Nota     float64
}

type NotasPorCurso map[string][]NotaEstudiante

func Top10MejoresPorCurso(estudiantes []reader.Estudiante) {
	procesarNotas(estudiantes, func(nota1, nota2 float64) bool {
		return nota1 > nota2
	})
}

func procesarNotas(estudiantes []reader.Estudiante, obtenerTop func(float64, float64) bool) {
	// Crea un mapa para almacenar las notas de los estudiantes por curso
	notasPorCurso := make(NotasPorCurso)

	// Recorre los estudiantes y calcula las notas por curso
	for _, estudiante := range estudiantes {
		for _, curso := range estudiante.Cursos {
			cursoNombre := curso.Curso
			nota := curso.Nota

			if _, ok := notasPorCurso[cursoNombre]; !ok {
				notasPorCurso[cursoNombre] = []NotaEstudiante{}
			}

			notasPorCurso[cursoNombre] = append(notasPorCurso[cursoNombre], NotaEstudiante{
				Nombre:   estudiante.Nombre,
				Apellido: estudiante.Apellido,
				Nota:     nota,
			})
		}
	}

	// Ordena las notas de cada curso
	for curso, notas := range notasPorCurso {
		sort.SliceStable(notas, func(i, j int) bool {
			return obtenerTop(notas[i].Nota, notas[j].Nota)
		})
		notasPorCurso[curso] = notas
	}

	// Muestra el top 10 de notas y estudiantes por curso
	for curso, notas := range notasPorCurso {
		fmt.Printf("Curso: %s\n", curso)
		for i, notaEstudiante := range notas {
			if i < 10 {
				fmt.Printf("%d. %s %s - Nota: %.2f\n", i+1, notaEstudiante.Nombre, notaEstudiante.Apellido, notaEstudiante.Nota)
			} else {
				break
			}
		}
		fmt.Println()
	}
}
