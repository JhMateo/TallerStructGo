package taller

import (
	"fmt"
	"github.com/JhMateo/tallerStruct/reader"
	"sort"
)

/*
	3. Top 10 estudiantes con mejores notas de cada curso.
*/

func Top10Estudantes(estudiantes []reader.Estudiante) {
	cursoEstudiantes := make(map[string][]reader.Estudiante)

	// Iterar sobre los estudiantes y agruparlos por curso.
	for _, estudiante := range estudiantes {
		for _, curso := range estudiante.Cursos {
			cursoNombre := curso.Curso
			cursoEstudiantes[cursoNombre] = append(cursoEstudiantes[cursoNombre], estudiante)
		}
	}

	// Iterar sobre los cursos y mostrar los 10 mejores estudiantes de cada uno.
	for curso, estudiantes := range cursoEstudiantes {
		fmt.Printf("Top 10 estudiantes en %s:\n", curso)
		// Ordenar los estudiantes por nota en orden descendente.
		sort.Slice(estudiantes, func(i, j int) bool {
			return estudiantes[i].Cursos[0].Nota > estudiantes[j].Cursos[0].Nota
		})
		// Mostrar los 10 mejores estudiantes.
		for i, estudiante := range estudiantes[:10] {
			fmt.Printf("%d. %s %s - Nota: %.2f\n", i+1, estudiante.Nombre, estudiante.Apellido, estudiante.Cursos[0].Nota)
		}
		fmt.Println()
	}
}
