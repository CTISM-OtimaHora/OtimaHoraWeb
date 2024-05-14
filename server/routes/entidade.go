package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	. "github.com/CTISM-OtimaHora/OtimaHora/models"
)

func Set_dispo_generic(w http.ResponseWriter, r * http.Request) {
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


func Get_dispo_generic(w http.ResponseWriter, r * http.Request) {
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
        case "contrato":
            disp = s.Contratos[id].Dispo
    }

    if err := json.NewEncoder(w).Encode(disp); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
}

func Get_generic(w http.ResponseWriter, r * http.Request) {
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
       

    switch tipo {
        case "professor":
            if json.NewEncoder(w).Encode(s.Professores[id]) != nil {
                w.WriteHeader(http.StatusInternalServerError)
                return
            }
            return
        case "disciplina":
            if json.NewEncoder(w).Encode(s.Disciplinas[id]) != nil {
                w.WriteHeader(http.StatusInternalServerError)
                return
            }
            return
        case "curso":
            if json.NewEncoder(w).Encode(s.Cursos[id]) != nil {
                w.WriteHeader(http.StatusInternalServerError)
                return
            }
            return
        case "contrato":
            if json.NewEncoder(w).Encode(s.Contratos[id]) != nil {
                w.WriteHeader(http.StatusInternalServerError)
                return
            }
            return
    }
}

func Add_generic(w http.ResponseWriter, r * http.Request) {
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

    switch tipo {
        case "professor":
            var prof Professor

            if json.NewDecoder(r.Body).Decode(&prof) != nil {
                w.WriteHeader(http.StatusInternalServerError)
                return
            }

            w.Write([]byte(fmt.Sprint(s.AddProfessor(prof))))
            return
        case "disciplina":
            var dis Disciplina 

            if json.NewDecoder(r.Body).Decode(&dis) != nil {
                w.WriteHeader(http.StatusInternalServerError)
                return
            }

            w.Write([]byte(fmt.Sprint(s.AddDisciplina(dis))))
            return
        case "curso":
            var cur Curso 

            if json.NewDecoder(r.Body).Decode(&cur) != nil {
                w.WriteHeader(http.StatusInternalServerError)
                return
            }

            w.Write([]byte(fmt.Sprint(s.AddCurso(cur))))
            return
        case "contrato":
            var ents []SearchEntidade 

            if json.NewDecoder(r.Body).Decode(&ents) != nil {
                w.WriteHeader(http.StatusInternalServerError)
                return
            }

            a := GetEntidadesOrNilSlice(ents, s)
            if a == nil {
                w.WriteHeader(http.StatusInternalServerError)
                return
            }

            con := NewContrato(0, a)

            w.Write([]byte(fmt.Sprint(s.AddContrato(con))))
            return
    }
}
