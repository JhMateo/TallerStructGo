package main

import (
	"fmt"
	"github.com/JhMateo/tallerStruct/reader"
	"github.com/JhMateo/tallerStruct/taller"
)

/*	Integrantes:
	Jhon Mauro Guevara - 160004315
	Joshep Mateo Granada - 160004314	*/

func main() {
	estudiantes, err := reader.ReadJson("generated.json")
	if err != nil {
		panic(err)
	}
	// 1. Estudiante con mejor promedio:
	fmt.Println("---- 1. MEJOR PROMEDIO ----")
	taller.MejorPromedio(estudiantes)

	// 2. Estudiante con peor promderio
	fmt.Println("\n---- 2. PEOR PROMEDIO ----")
	taller.PeorPromedio(estudiantes)

	// 3. Top 10 estudiantes con mejores notas de cada curso
	fmt.Println("\n---- 3. TOP 10 ESTUDIANTES CON MEJORES NOTAS DE CADA CURSO ----")
	taller.Top10MejoresPorCurso(estudiantes)

	// 4. Top 10 estudiantes con peores notas de cada curso
	fmt.Println("\n---- 4. TOP 10 ESTUDIANTES CON PEORES NOTAS DE CADA CURSO ----")
	taller.Top10PeoresPorCurso(estudiantes)

	// 5. Estudiante Masculino con mayor edad
	fmt.Println("\n---- 5. ESTUDIANTE MASCULINO DE MAYOR EDAD ----")
	taller.ImprimirEstudianteMasculino(estudiantes)

	// 6. Estudiante Femenino con mayor edad
	fmt.Println("\n---- 6. ESTUDIANTE FEMENINO DE MAYOR EDAD ----")
	taller.ImprimirEstudianteFemenino(estudiantes)

	fmt.Println("\n---- 7. PROMEDIO, RANGO, VARIANZA Y DESVIACIÓN ESTANDAR PARA CADA CURSO ----")
	taller.ImprimirEstadisticasPorCurso(estudiantes)

	// 8. Reporte de estudiantes que se matricularon en el año pasado (2022).
	fmt.Println("\n---- 8. REPORTE DE ESTUDIANTES QUE SE MATRICULARON EN EL AÑO PASADO (2022) ----")
	taller.Matriculados2022(estudiantes)

	// 9. Promedio de estudiantes por cada rango de edad (20-29, 30-39, 40-49,...)
	fmt.Println("\n---- 9. PROMEDIO DE ESTUDIANTES POR CADA RANGO DE EDAD (20-29, 30-39, 40-49,...) ----")
	taller.PromediosPorRangoEdad(estudiantes)
}
