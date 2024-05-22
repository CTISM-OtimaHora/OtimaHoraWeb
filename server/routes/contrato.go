package routes

import (
	"encoding/json"
	"fmt"
    "strconv"
	"net/http"

	. "github.com/CTISM-OtimaHora/OtimaHora/models"
)

func AddContrato (w http.ResponseWriter, r * http.Request) {
    s := Session_or_nil(r)
    if s == nil {
        w.Write([]byte("No session or Session expired"))
        w.WriteHeader(http.StatusUnauthorized)
        return 
    }

    var ents_real []SearchEntidade

    if err := json.NewDecoder(r.Body).Decode(&ents_real); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("malformed body 1: " + err.Error()))
        return
    }
    
    ents := make([]Entidade, len(ents_real))

    for i := range ents_real {
        e := ents_real[i].GetEntidadeOrNil(s)
        if e == nil {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte("malformed body 2"))
            return
        }
        ents[i] = e
    }
    
    w.Write([]byte(fmt.Sprint(s.AddContrato(NewContrato(len(s.Contratos), ents)))))
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
