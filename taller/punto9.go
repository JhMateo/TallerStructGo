package taller

import (
	"fmt"
	"github.com/JhMateo/tallerStruct/reader"
)

/*
	9. Promedio de estudiantes por cada rango de edad (20-29, 30-39, 40-49,...)
*/

func PromediosPorRangoEdad(estudiantes []reader.Estudiante) map[string]float64 {
	rangosEdad := map[string][]float64{
		"20-29":    []float64{},
		"30-39":    []float64{},
		"40-49":    []float64{},
		"50-59":    []float64{},
		"60 o más": []float64{},
	}

	totalEstudiantes := len(estudiantes)

	if totalEstudiantes == 0 {
		fmt.Println("No hay estudiantes para calcular promedios.")
		return nil
	}

	for _, estudiante := range estudiantes {
		rango := obtenerRangoEdad(estudiante.Edad)
		if rango != "" {
			rangosEdad[rango] = append(rangosEdad[rango], calcularPromedio(estudiante.Cursos))
		}
	}

	resultados := make(map[string]float64)

	fmt.Println("Promedio de estudiantes por rango de edad:")
	for rango, notas := range rangosEdad {
		promedio := 0.0
		if len(notas) != 0 {
			promedio = calcularPromedioANotas(notas)
		}
		resultados[rango] = promedio
		fmt.Printf("Rango de edad(%s): %.2f\n", rango, promedio)
	}

	return resultados
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

func calcularPromedioANotas(notas []float64) float64 {
	suma := 0.0
	for _, nota := range notas {
		suma += nota
	}
	return suma / float64(len(notas))
}
