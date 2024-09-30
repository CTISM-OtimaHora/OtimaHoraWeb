package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	. "github.com/CTISM-OtimaHora/OtimaHora/models"
)

func AddContrato (w http.ResponseWriter, r * http.Request) {
    s := Session_or_nil(r)
    if s == nil {
        w.Write([]byte("No session or Session expired"))
        w.WriteHeader(http.StatusUnauthorized)
        return 
    }

    var ents_real []Participante_to_json

    if err := json.NewDecoder(r.Body).Decode(&ents_real); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("malformed body 1: " + err.Error()))
        return
    }
    
    ents := make([]ParticipanteContrato, len(ents_real))
    tipos := make([]string, len(ents_real))

    for i, p := range ents_real {
        found := false
        var e ParticipanteContrato
        switch p.Tipo {
        case "disciplina":
            item, ok := s.Disciplinas[p.Id]
            e = item
            if !ok{
                w.WriteHeader(http.StatusBadRequest)
                w.Write([]byte("malformed body 2"))
                return
            }
            found = true
        case "professor":
            item, ok := s.Professores[p.Id]
            e = item
            if !ok{
                w.WriteHeader(http.StatusBadRequest)
                w.Write([]byte("malformed body 2"))
                return
            }
            found = true
        case "turma":
            for _, c := range s.Cursos {
                for _, et := range c.Etapas {
                    for _, t := range  et.Turmas {
                        if t.Id == p.Id && t.Nome == p.Nome {
                            e = t 
                            found = true
                            goto out
                        }
                    }
                }
            }
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte("malformed body 2"))
            return
        case "recurso":
            item, ok := s.Recursos[p.Id]
            e = item
            if !ok{
                w.WriteHeader(http.StatusBadRequest)
                w.Write([]byte("malformed body 2"))
                return
            }
            found = true
        }
        out:
        if !found {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte("malformed body 2"))
            return
        }
        fmt.Println("found")
        fmt.Println(e)
        ents[i] = e
        tipos[i] = p.Tipo
    }
    
    w.Write([]byte(fmt.Sprint(s.AddContrato(Contrato{Id:0, Participantes: ents, Tipo_por_participante: tipos, Dispo: AndDisp(ents)}))))
}

func GetContrato(w http.ResponseWriter, r * http.Request) {
    s := Session_or_nil(r)
    if s == nil {
        w.Write([]byte("No session or Session expired"))
        w.WriteHeader(http.StatusUnauthorized)
        return 
    }

    id, conv_err := strconv.Atoi(r.PathValue("id"))
    if conv_err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    
    if json.NewEncoder(w).Encode(s.Contratos[id]) != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    
    return   
}
