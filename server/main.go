package main

import (
	"fmt"
	"net/http"
	"time"

	m "github.com/CTISM-OtimaHora/OtimaHora/models"
	. "github.com/CTISM-OtimaHora/OtimaHora/routes"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// If it's an OPTIONS request, just return OK status
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	go func() {
		for {
			if len(m.Sessions) > 0 {
				fmt.Printf("Professores: %v\n", m.Sessions[0].Professores)
				fmt.Printf("Disciplinas: %v\n", m.Sessions[0].Disciplinas)
				fmt.Printf("Cursos: %v\n", m.Sessions[0].Cursos)
				time.Sleep(2 * time.Second)
			}
		}
	}()

	r := http.NewServeMux()

	r.HandleFunc("GET /add-session", AddSession)
	r.HandleFunc("GET /session", GetSession)

    // professor routes
	r.HandleFunc("POST /professor/add", AddBuilder(m.AddProfessor))
	r.HandleFunc("GET /professor/get/{id}", GetBuilder(m.ProfessorGeter))
	r.HandleFunc("PUT /professor/set/{id}", SetBuilder(m.ProfessorGeter)) // professor tem um setter pra suas disciplinas
    r.HandleFunc("GET /professor/slice", SliceGetBuilder(m.ProfessorGeter))
	r.HandleFunc("GET /professor/disp/get/{id}", DispoGetBuilder(m.ProfessorGeter))
	r.HandleFunc("PUT /professor/disp/set/{id}", DispoSetBuilder(m.ProfessorGeter))


    // disciplina routes
	r.HandleFunc("POST /disciplina/add", AddBuilder(m.AddDisciplina))
	r.HandleFunc("GET /disciplina/get/{id}", GetBuilder(m.DisciplinaGeter))
    r.HandleFunc("GET /disciplina/slice", SliceGetBuilder(m.DisciplinaGeter))
	r.HandleFunc("GET /disciplina/disp/get/{id}", DispoGetBuilder(m.DisciplinaGeter))
	r.HandleFunc("PUT /disciplina/disp/set/{id}", DispoSetBuilder(m.DisciplinaGeter))

    // recurso routes
	r.HandleFunc("POST /recurso/add", AddBuilder(m.AddRecurso))
	r.HandleFunc("GET /recurso/get/{id}", GetBuilder(m.RecursoGetter))
    r.HandleFunc("GET /recurso/slice", SliceGetBuilder(m.RecursoGetter))
	r.HandleFunc("GET /recurso/disp/get/{id}", GetBuilder(m.RecursoGetter))
	r.HandleFunc("PUT /recurso/disp/set/{id}", DispoSetBuilder(m.RecursoGetter))

    // curso routes
	r.HandleFunc("POST /curso/add", AddBuilder(m.AddCurso))
    r.HandleFunc("GET /curso/get/{id}", GetBuilder(m.CursoGetter))
    r.HandleFunc("GET /curso/slice", SliceGetBuilder(m.CursoGetter))
    r.HandleFunc("GET /curso/disp/get/{id}", GetBuilder(m.CursoGetter))
    r.HandleFunc("PUT /curso/disp/set/{id}", DispoSetBuilder(m.CursoGetter))

    // contrato routes
	r.HandleFunc("POST /contrato/add", AddContrato)
    r.HandleFunc("GET /contrato/get/{id}", GetContrato)

	with_cors := corsMiddleware(r)
	http.ListenAndServe("localhost:3000", with_cors)
}
