package routes

import (
    "net/http"
    "strconv"
    "fmt"
    "encoding/json"

    ."github.com/CTISM-OtimaHora/OtimaHora/models"
)

func Add_turma_to_curso(w http.ResponseWriter, r * http.Request) {
    if r.Method != "POST" {
        w.WriteHeader(http.StatusNotFound)
        return
    }
    s := Session_or_nil(r)
    if s == nil {
        w.Write([]byte("No session or Session expired"))
        w.WriteHeader(http.StatusUnauthorized)
        return 
    }

    curso_idx, conv_err := strconv.Atoi(r.PathValue("id_curso"))

    var turma Turma
    if err := json.NewDecoder(r.Body).Decode(&turma); conv_err != nil || err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    
    w.Write([]byte(fmt.Sprint(s.Cursos[curso_idx].AddTurma(turma))))
    return
}


func Get_curso(w http.ResponseWriter, r * http.Request) {
    s := Session_or_nil(r)
    if s == nil {
        w.Write([]byte("No session or Session expired"))
        w.WriteHeader(http.StatusUnauthorized)
        return 
    }

    curso_idx, conv_err := strconv.Atoi(r.PathValue("id_curso"))
    if conv_err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    json.NewEncoder(w).Encode(s.Cursos[curso_idx])
    return
}
