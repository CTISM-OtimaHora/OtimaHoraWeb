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

	r.HandleFunc("/add-session", New_session)
	r.HandleFunc("/add-curso", Add_curso_to_session)
	r.HandleFunc("/add-turma/{id_curso}", Add_turma_to_curso)
	r.HandleFunc("/add-professor/", Add_Professor_to_turma)
	r.HandleFunc("/add-disciplina/", Add_Disciplina_to_turma)
	r.HandleFunc("/session", Get_session)
	r.HandleFunc("/session/{id_curso}", Get_curso)
	r.HandleFunc("/session/{id_curso}/{id_turma}", Get_turma)

	with_cors := corsMiddleware(r)
	http.ListenAndServe("localhost:3000", with_cors)
}
