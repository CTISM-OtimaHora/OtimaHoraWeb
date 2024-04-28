package main

import (
	"encoding/json"
    "fmt"
	"net/http"
	"strconv"
	"time"
)

var Sessions []*Session = []*Session{}
var next_id int = 0

func new_session(w http.ResponseWriter, r * http.Request) {
    s := Session_or_nil(w, r)
    if s != nil {
        http.SetCookie(w, &http.Cookie{Name: "id", Value: fmt.Sprint(s.Id), Path: "/", Expires: time.Now().Add(2*time.Hour)})
        return
    }

    new_s := Session{next_id, []Curso{}}
    http.SetCookie(w, &http.Cookie{Name: "id", Value: fmt.Sprint(next_id), Path: "/", Expires: time.Now().Add(2*time.Hour)})
    Sessions = append(Sessions, &new_s)
    w.WriteHeader(200)
    next_id += 1
    return
}

func add_curso(w http.ResponseWriter, r * http.Request) {
    if r.Method != "POST" {
        w.WriteHeader(http.StatusNotFound)
        return
    }

    s := Session_or_nil(w, r)
    if s == nil {
        w.Write([]byte("No session or Session expired"))
        w.WriteHeader(http.StatusUnauthorized)
        return // errors already in w
    }

    var curso Curso
    if err := json.NewDecoder(r.Body).Decode(&curso); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("Malformed body"))
        return
    }
    w.Write([]byte(fmt.Sprint(s.Add(curso))))
    return
}

func add_turma(w http.ResponseWriter, r * http.Request) {
    if r.Method != "POST" {
        w.WriteHeader(http.StatusNotFound)
        return
    }
    s := Session_or_nil(w, r)
    if s == nil {
        w.Write([]byte("No session or Session expired"))
        w.WriteHeader(http.StatusUnauthorized)
        return // errors already on w
    }

    curso_idx, conv_err := strconv.Atoi(r.PathValue("id_curso"))

    var turma Turma
    if err := json.NewDecoder(r.Body).Decode(&turma); conv_err != nil || err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    
    w.Write([]byte(fmt.Sprint(s.Cursos[curso_idx].Add(turma))))
    return
}

func get_session(w http.ResponseWriter, r * http.Request) {
    s := Session_or_nil(w, r)
    if s == nil {
        w.Write([]byte("No session or Session expired"))
        w.WriteHeader(http.StatusUnauthorized)
        return // errors already on w
    }

    json.NewEncoder(w).Encode(s)
    return
}

func get_curso(w http.ResponseWriter, r * http.Request) {
    s := Session_or_nil(w, r)
    if s == nil {
        w.Write([]byte("No session or Session expired"))
        w.WriteHeader(http.StatusUnauthorized)
        return // errors already on w
    }

    curso_idx, conv_err := strconv.Atoi(r.PathValue("id_curso"))
    if conv_err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    json.NewEncoder(w).Encode(s.Cursos[curso_idx])
    return
}

func get_turma(w http.ResponseWriter, r * http.Request) {
    s := Session_or_nil(w, r)
    if s == nil {
        w.Write([]byte("No session or Session expired"))
        w.WriteHeader(http.StatusUnauthorized)
        return // errors already on w
    }

    curso_idx, conv_err := strconv.Atoi(r.PathValue("id_curso"))
    turma_idx, conv_err2 := strconv.Atoi(r.PathValue("id_turma"))
    if conv_err != nil || conv_err2 != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    json.NewEncoder(w).Encode(s.Cursos[curso_idx].Turmas[turma_idx])
    return 
}

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

func has_session(w http.ResponseWriter, r * http.Request) {
}

func main() {
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

    r := http.NewServeMux()
    r.HandleFunc("/add-session", new_session)
    r.HandleFunc("/add-curso", add_curso)
    r.HandleFunc("/add-turma/{id_curso}", add_turma)
    r.HandleFunc("/session", get_session)
    r.HandleFunc("/session/{id_curso}", get_curso)
    r.HandleFunc("/session/{id_curso}/{id_turma}", get_turma)
    with_cors := corsMiddleware(r)
    http.ListenAndServe("localhost:3000", with_cors)
}
