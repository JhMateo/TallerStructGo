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

	// 2. Top 10 estudiantes con mejores notas de cada curso
	fmt.Println("\n---- TOP 10 ESTUDIANTES CON MEJORES NOTAS DE CADA CURSO ----")
	taller.Top10Estudantes(estudiantes)
}
