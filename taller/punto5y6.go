package taller

import (
	"fmt"

	"github.com/JhMateo/tallerStruct/reader"
)

func EncontrarEstudiantesMayorEdadPorGenero(estudiantes []reader.Estudiante, genero string) []reader.Estudiante {
	var estudiantesMayoresEdad []reader.Estudiante
	var mayorEdad int

	for _, estudiante := range estudiantes {
		if estudiante.Gender == genero && estudiante.Edad > mayorEdad {
			mayorEdad = estudiante.Edad
		}
	}

	for _, estudiante := range estudiantes {
		if estudiante.Gender == genero && estudiante.Edad == mayorEdad {
			estudiantesMayoresEdad = append(estudiantesMayoresEdad, estudiante)
		}
	}

	return estudiantesMayoresEdad
}

func ImprimirEstudiantesMayorEdadNombreEdad(estudiantes []reader.Estudiante, genero string) {
	estudiantesMayoresEdad := EncontrarEstudiantesMayorEdadPorGenero(estudiantes, genero)

	fmt.Printf("Estudiantes %s de mayor edad:\n", genero)
	for _, estudiante := range estudiantesMayoresEdad {
		fmt.Printf("Nombre: %s %s\n", estudiante.Nombre, estudiante.Apellido)
		fmt.Printf("Edad: %d\n", estudiante.Edad)
		fmt.Println("-----")
	}
}

func ImprimirEstudianteMasculino(estudiantes []reader.Estudiante) {
	ImprimirEstudiantesMayorEdadNombreEdad(estudiantes, "male")
}

func ImprimirEstudianteFemenino(estudiantes []reader.Estudiante) {
	ImprimirEstudiantesMayorEdadNombreEdad(estudiantes, "female")
}
