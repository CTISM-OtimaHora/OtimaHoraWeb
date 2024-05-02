package main

import (
    "fmt"
	"net/http"
	"time"

    ."github.com/CTISM-OtimaHora/OtimaHora/models"
    ."github.com/CTISM-OtimaHora/OtimaHora/routes"
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
    // debug code
    last_len := len(Sessions)
    last_time := time.Now()
    go func () {
        for {
            if new_len := len(Sessions); new_len != last_len || time.Since(last_time) > time.Second * 7 {
                last_len = new_len
                last_time = time.Now()
                sessions := "["
                for _, s := range Sessions {
                    sessions += s.ToString() + ", "
                } 
                sessions += "]"
                fmt.Println(sessions)
            }
        }
    } ()
    // end of debug code

    r := http.NewServeMux()

    r.HandleFunc("/add-session", New_session)
    r.HandleFunc("/add-curso", Add_curso_to_session)
    r.HandleFunc("/add-turma/{id_curso}", Add_turma_to_curso)
    r.HandleFunc("/session", Get_session)
    r.HandleFunc("/session/{id_curso}", Get_curso)
    r.HandleFunc("/session/{id_curso}/{id_turma}", Get_turma)

    with_cors := corsMiddleware(r)
    http.ListenAndServe("localhost:3000", with_cors)
}
