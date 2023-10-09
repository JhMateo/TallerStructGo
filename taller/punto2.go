package taller

import (
	"fmt"
	"github.com/JhMateo/tallerStruct/reader"
)

/*
	2. Estudiante con peor promedio
*/

func PeorPromedio(estudiantes []reader.Estudiante) {
	peorEstudiante := estudiantes[0]
	peorPromedio := calcularPromedio(peorEstudiante.Cursos)

	// Iteramos sobre los demás estudiantes y comparamos los promedios.
	for _, estudiante := range estudiantes[1:] {
		promedio := calcularPromedio(estudiante.Cursos)
		if promedio < peorPromedio {
			peorPromedio = promedio
			peorEstudiante = estudiante
		}
	}

	fmt.Println("El estudiante con el peor promedio es:")
	fmt.Printf("Nombre: %s %s\n", peorEstudiante.Nombre, peorEstudiante.Apellido)
	fmt.Printf("Promedio: %.2f\n", peorPromedio)
}
