package routes

import (
    "net/http"
    "strconv"
    "encoding/json"

    ."github.com/CTISM-OtimaHora/OtimaHora/models"
)

func Set_entidade_generic(w http.ResponseWriter, r * http.Request) {
    if r.Method != "POST" {
        w.WriteHeader(404)
        return
    }

    s := Session_or_nil(r)   
    if s == nil {
        w.WriteHeader(http.StatusUnauthorized)
        return 
    }

    tipo := r.PathValue("tipo")
    if tipo == "" {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    id, conv_err := strconv.Atoi(r.PathValue("id"))
    if conv_err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    
    var disp Disponibilidade

    if err := json.NewDecoder(r.Body).Decode(&disp); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }


    switch tipo {
        case "professor":
            s.Professores[id].Dispo = disp
        case "disciplina":
            s.Disciplinas[id].Dispo = disp
        case "curso":
            s.Cursos[id].Dispo = disp
    }
}


func Get_entidade_generic(w http.ResponseWriter, r * http.Request) {
    if r.Method != "GET" {
        w.WriteHeader(404)
        return
    }

    s := Session_or_nil(r)   
    if s == nil {
        w.WriteHeader(http.StatusUnauthorized)
        return 
    }

    tipo := r.PathValue("tipo")
    if tipo == "" {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    id, conv_err := strconv.Atoi(r.PathValue("id"))
    if conv_err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    
    var disp Disponibilidade

    switch tipo {
        case "professor":
            disp = s.Professores[id].Dispo
        case "disciplina":
            disp = s.Disciplinas[id].Dispo
        case "curso":
            disp = s.Cursos[id].Dispo
    }

    if err := json.NewEncoder(w).Encode(disp); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
}


