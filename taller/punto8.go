package taller

import (
	"fmt"
	"github.com/JhMateo/tallerStruct/reader"
)

/*
	8. Reporte de estudiantes que se matricularon en el año pasado (2022).
*/

func Matriculados2022(estudiantes []reader.Estudiante) map[string]string {
	// Crear un mapa para almacenar estudiantes matriculados en 2022 con sus fechas de matriculación
	matriculados := make(map[string]string)

	for _, estudiante := range estudiantes {
		if estudiante.Matriculado.Year() == 2022 {
			nombreCompleto := estudiante.Nombre + " " + estudiante.Apellido
			// Guardar fecha con el formato AAAA-MM-DD
			matriculados[nombreCompleto] = estudiante.Matriculado.Format("2006-01-02")
		}
	}

	if len(matriculados) == 0 {
		fmt.Println("No se encontraron estudiantes matriculados en 2022.")
	} else {
		fmt.Println("Estudiantes matriculados en 2022:")
		i := 1
		for nombre, fecha := range matriculados {
			fmt.Printf("%d. %s - Fecha de Matriculación: %s\n", i, nombre, fecha)
			i++
		}
	}

	return matriculados
}
