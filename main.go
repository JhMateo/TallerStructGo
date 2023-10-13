package main

import (
	"fmt"
	"github.com/JhMateo/tallerStruct/reader"
	"github.com/JhMateo/tallerStruct/taller"
)

func main() {
	estudiantes, err := reader.ReadJson("generated.json")
	if err != nil {
		panic(err)
	}

	// 1. Estudiante con mejor promedio:
	fmt.Println("---- MEJOR PROMEDIO ----")
	taller.MejorPromedio(estudiantes)

	// 2. Estudiante con peor promderio
	fmt.Println("\n---- PEOR PROMEDIO ----")
	taller.PeorPromedio(estudiantes)

	// 3. Top 10 estudiantes con mejores notas de cada curso
	fmt.Println("\n---- TOP 10 ESTUDIANTES CON MEJORES NOTAS DE CADA CURSO ----")
	taller.Top10MejoresPorCurso(estudiantes)

	// 4. Top 10 estudiantes con peores notas de cada curso
	fmt.Println("\n---- TOP 10 ESTUDIANTES CON PEORES NOTAS DE CADA CURSO ----")
	taller.Top10PeoresPorCurso(estudiantes)

	// 8. Reporte de estudiantes que se matricularon en el año pasado (2022).
	fmt.Println("\n---- REPORTE DE ESTUDIANTES QUE SE MATRICULARON EN EL AÑO PASADO (2022) ----")
	taller.Matriculados2022(estudiantes)
}
