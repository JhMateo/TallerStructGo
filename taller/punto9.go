package taller

import (
	"fmt"
	"github.com/JhMateo/tallerStruct/reader"
)

/*
	9. Promedio de estudiantes por cada rango de edad (20-29, 30-39, 40-49,...)
*/

func PromediosPorRangoEdad(estudiantes []reader.Estudiante) {
	rangosEdad := map[string]int{
		"20-29":    0,
		"30-39":    0,
		"40-49":    0,
		"50-59":    0,
		"60 o más": 0,
	}

	totalEstudiantes := len(estudiantes)

	for _, estudiante := range estudiantes {
		rango := obtenerRangoEdad(estudiante.Edad)
		if rango != "" {
			rangosEdad[rango]++
		}
	}

	if totalEstudiantes == 0 {
		fmt.Println("No hay estudiantes para calcular promedios.")
		return
	}

	fmt.Println("Promedio de estudiantes por rango de edad:")
	for rango, cantidad := range rangosEdad {
		promedio := float64(cantidad) / float64(totalEstudiantes)
		fmt.Printf("Rango de edad(%s): %.2f -> %.1f%% estudiantes\n", rango, promedio, promedio*100)
	}
}

func obtenerRangoEdad(edad int) string {
	switch {
	case edad >= 20 && edad <= 29:
		return "20-29"
	case edad >= 30 && edad <= 39:
		return "30-39"
	case edad >= 40 && edad <= 49:
		return "40-49"
	case edad >= 50 && edad <= 59:
		return "50-59"
	case edad >= 60:
		return "60 o más"
	default:
		return ""
	}
}
