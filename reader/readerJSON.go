package reader

import (
	"encoding/json"
	"os"
	"time"
)

type Curso struct {
	ID    int     `json:"id"`
	Curso string  `json:"curso"`
	Nota  float64 `json:"nota"`
}

type Estudiante struct {
	Index       int        `json:"index"`
	Nombre      string     `json:"nombre"`
	Apellido    string     `json:"apellido"`
	Edad        int        `json:"edad"`
	Gender      string     `json:"gender"`
	Email       string     `json:"email"`
	Phone       string     `json:"phone"`
	Address     string     `json:"address"`
	About       string     `json:"about"`
	Matriculado CustomTime `json:"matriculado"`
	Cursos      []Curso    `json:"cursos"`
}

type CustomTime struct {
	time.Time
}

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = s[1 : len(s)-1] // Remove the double quotes around the string
	t, err := time.Parse("2006-01-02T15:04:05 -07:00", s)
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}

func ReadJson(filename string) ([]Estudiante, error) {
	estudiantesJson, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer estudiantesJson.Close()

	var estudiantes []Estudiante
	jsonParser := json.NewDecoder(estudiantesJson)
	if err = jsonParser.Decode(&estudiantes); err != nil {
		return nil, err
	}

	return estudiantes, nil
}
