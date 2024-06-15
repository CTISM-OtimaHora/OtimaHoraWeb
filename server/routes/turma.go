package routes

import (
	"encoding/json"
	"net/http"
	"slices"
	"strconv"

	. "github.com/CTISM-OtimaHora/OtimaHora/models"
)

func GetTurma(w http.ResponseWriter, r * http.Request) {
    s := Session_or_nil(r)
    if s == nil {
        w.WriteHeader(http.StatusUnauthorized)
        w.Write([]byte("No session or Session expired"))
        return 
    }

    curso_id, conv_err1 := strconv.Atoi(r.PathValue("curso_id"))
    etapa_id, conv_err2 := strconv.Atoi(r.PathValue("etapa_id"))
    turma_id, conv_err3 := strconv.Atoi(r.PathValue("id"))
    if conv_err1 != nil || conv_err2 != nil || conv_err3 != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("Conversion error from path parameter"))
        return
    }
    
    idx := -1
    for i, t := range s.Cursos[curso_id].Etapas[etapa_id] {
        if t.Id == turma_id {
            idx = i
            break
        }
    }

    if idx == -1 {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("Entity not found"))
        return
    }

    if err := json.NewEncoder(w).Encode(s.Cursos[curso_id].Etapas[etapa_id][idx]); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(err.Error()))
        return
    }
}

func SetTurma(w http.ResponseWriter, r * http.Request) {
    s := Session_or_nil(r)
    if s == nil {
        w.WriteHeader(http.StatusUnauthorized)
        w.Write([]byte("No session or Session expired"))
        return 
    }

    curso_id, conv_err1 := strconv.Atoi(r.PathValue("curso_id"))
    etapa_id, conv_err2 := strconv.Atoi(r.PathValue("etapa_id"))
    turma_id, conv_err3 := strconv.Atoi(r.PathValue("id"))
    if conv_err1 != nil || conv_err2 != nil || conv_err3 != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("Conversion error from path parameter"))
        return
    }
    
    idx := -1
    for i, t := range s.Cursos[curso_id].Etapas[etapa_id] {
        if t.Id == turma_id {
            idx = i
            break
        }
    }

    if idx == -1 {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("Entity not found"))
        return
    }

    var j_t Turma
    if err := json.NewDecoder(r.Body).Decode(&j_t); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(err.Error()))
        return
    }
    
    s.Cursos[curso_id].Etapas[etapa_id][idx] = j_t
}

func DeleteTurma(w http.ResponseWriter, r * http.Request) {
    s := Session_or_nil(r)
    if s == nil {
        w.WriteHeader(http.StatusUnauthorized)
        w.Write([]byte("No session or Session expired"))
        return 
    }

    curso_id, conv_err1 := strconv.Atoi(r.PathValue("curso_id"))
    etapa_id, conv_err2 := strconv.Atoi(r.PathValue("etapa_id"))
    turma_id, conv_err3 := strconv.Atoi(r.PathValue("id"))
    if conv_err1 != nil || conv_err2 != nil || conv_err3 != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("Conversion error from path parameter"))
        return
    }
    
    idx := -1
    for i, t := range s.Cursos[curso_id].Etapas[etapa_id] {
        if t.Id == turma_id {
            idx = i
            break
        }
    }

    if idx == -1 {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("Entity not found"))
        return
    }

    s.Cursos[curso_id].Etapas[etapa_id] = slices.Delete(s.Cursos[curso_id].Etapas[etapa_id], idx, idx+1)
}
