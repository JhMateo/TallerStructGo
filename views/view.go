package views

import (
	"fmt"
	"github.com/JhMateo/tallerStruct/reader"
	"html/template" // Renderizado
	"net/http"      // Comunicación
)

func Index(w http.ResponseWriter, r *http.Request) {
	// w para escribir, r para recibir información de la petición
	template, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Fprintf(w, "<h1>Página no encontrada!</h1>")
		panic(err)
	} else {
		template.Execute(w, r)
	}
}

func ExecuteWeb(estudiantes []reader.Estudiante) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Index(w, r)
	})

	http.HandleFunc("/mejor_promedio", func(w http.ResponseWriter, r *http.Request) {
		MejorPromedioHandler(w, r, estudiantes)
	})

	http.HandleFunc("/peor_promedio", func(w http.ResponseWriter, r *http.Request) {
		PeorPromedioHandler(w, r, estudiantes)
	})

	http.HandleFunc("/top10_mejores", func(w http.ResponseWriter, r *http.Request) {
		Top10MejoresPorCursoHandler(w, r, estudiantes)
	})

	http.HandleFunc("/top10_peores", func(w http.ResponseWriter, r *http.Request) {
		Top10PeoresPorCursoHandler(w, r, estudiantes)
	})

	http.HandleFunc("/hombre_mayor", func(w http.ResponseWriter, r *http.Request) {
		HombreMayorHandler(w, r, estudiantes)
	})

	http.HandleFunc("/mujer_mayor", func(w http.ResponseWriter, r *http.Request) {
		MujerMayorHandler(w, r, estudiantes)
	})

	http.HandleFunc("/estadisticas_cursos", func(w http.ResponseWriter, r *http.Request) {
		EstadisticasPorCursoHandler(w, r, estudiantes)
	})

	fsStatic := http.FileServer(http.Dir("templates/static"))

	// Establece una ruta para servir archivos estáticos.
	http.Handle("/static/", http.StripPrefix("/static/", fsStatic))

	fmt.Println("Escuchando en http://localhost:8000...")
	http.ListenAndServe("localhost:8000", nil)
}
