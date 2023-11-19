package taller

import (
	"fmt"
	"math"
	"sort"

	"github.com/JhMateo/tallerStruct/reader"
)

type EstadisticasCurso struct {
	Promedio      float64
	Rango         float64
	Varianza      float64
	DesviacionStd float64
	Curso         string
	Notas         []float64
}

func CalcularEstadisticasPorCurso(estudiantes []reader.Estudiante) map[string]EstadisticasCurso {
	// Mapa para almacenar estadísticas por curso
	estadisticasPorCurso := make(map[string]EstadisticasCurso)

	// Recorre todos los estudiantes
	for _, estudiante := range estudiantes {
		// Recorre todos los cursos del estudiante
		for _, curso := range estudiante.Cursos {
			// Verifica si ya existe el curso en el mapa
			_, ok := estadisticasPorCurso[curso.Curso]
			if !ok {
				// Si no existe, inicializa las estadísticas para el curso
				estadisticasPorCurso[curso.Curso] = EstadisticasCurso{
					Curso: curso.Curso,
					Notas: []float64{},
				}
			}

			// Copia el valor del mapa
			estadisticas := estadisticasPorCurso[curso.Curso]

			// Modifica la copia
			estadisticas.Notas = append(estadisticas.Notas, curso.Nota)

			// Asigna la copia de nuevo al mapa
			estadisticasPorCurso[curso.Curso] = estadisticas
		}
	}

	// Calcula estadísticas para cada curso
	for curso, estadisticas := range estadisticasPorCurso {
		estadisticasPorCurso[curso] = calcularEstadisticas(estadisticas)
	}

	return estadisticasPorCurso
}

func calcularEstadisticas(estadisticas EstadisticasCurso) EstadisticasCurso {
	// Calcula el promedio
	estadisticas.Promedio = calcularPromedios(estadisticas.Notas)

	// Calcula el rango
	estadisticas.Rango = calcularRango(estadisticas.Notas)

	// Calcula la varianza
	estadisticas.Varianza = calcularVarianza(estadisticas.Notas)

	// Calcula la desviación estándar
	estadisticas.DesviacionStd = calcularDesviacionEstandar(estadisticas.Varianza)

	return estadisticas
}

func calcularPromedios(notas []float64) float64 {
	suma := 0.0
	for _, nota := range notas {
		suma += nota
	}
	return suma / float64(len(notas))
}

func calcularRango(notas []float64) float64 {
	sort.Float64s(notas)
	return notas[len(notas)-1] - notas[0]
}

func calcularVarianza(notas []float64) float64 {
	promedio := calcularPromedios(notas)
	suma := 0.0

	for _, nota := range notas {
		suma += math.Pow(nota-promedio, 2)
	}

	return suma / float64(len(notas))
}

func calcularDesviacionEstandar(varianza float64) float64 {
	return math.Sqrt(varianza)
}

func ImprimirEstadisticasPorCurso(estudiantes []reader.Estudiante) {
	estadisticasPorCurso := CalcularEstadisticasPorCurso(estudiantes)

	for _, estadisticas := range estadisticasPorCurso {
		fmt.Printf("Curso: %s\n", estadisticas.Curso)
		fmt.Printf("Promedio: %.2f\n", estadisticas.Promedio)
		fmt.Printf("Rango: %.2f\n", estadisticas.Rango)
		fmt.Printf("Varianza: %.2f\n", estadisticas.Varianza)
		fmt.Printf("Desviación estándar: %.2f\n", estadisticas.DesviacionStd)
		fmt.Println("------------------------")
	}
}
