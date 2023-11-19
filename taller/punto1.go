package taller

import (
	"fmt"
	"github.com/JhMateo/tallerStruct/reader"
)

/*
	1. Estudiante con mejor promedio
*/

func MejorPromedio(estudiantes []reader.Estudiante) (mejorEstudiante reader.Estudiante, mejorPromedio float64) {
	mejorEstudiante = estudiantes[0]
	mejorPromedio = calcularPromedio(mejorEstudiante.Cursos)

	// Iteramos sobre los demás estudiantes y comparamos los promedios.
	for _, estudiante := range estudiantes[1:] {
		promedio := calcularPromedio(estudiante.Cursos)
		if promedio > mejorPromedio {
			mejorPromedio = promedio
			mejorEstudiante = estudiante
		}
	}

	fmt.Println("El estudiante con el mejor promedio es:")
	fmt.Printf("Nombre: %s %s\n", mejorEstudiante.Nombre, mejorEstudiante.Apellido)
	fmt.Printf("Promedio: %.2f\n", mejorPromedio)

	return mejorEstudiante, mejorPromedio
}

func calcularPromedio(cursos []reader.Curso) float64 {
	if len(cursos) == 0 {
		return 0.0
	}
	sumaNotas := 0.0
	for _, curso := range cursos {
		sumaNotas += curso.Nota
	}
	return sumaNotas / float64(len(cursos))
}
