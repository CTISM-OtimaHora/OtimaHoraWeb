package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/CTISM-OtimaHora/OtimaHora/models"
)

func AddContrato (w http.ResponseWriter, r * http.Request) {
    if r.Method != "POST" {
        w.WriteHeader(404)
        return
    }   
    
    s := Session_or_nil(w, r)
    if s == nil {
        w.Write([]byte("No session or Session expired"))
        w.WriteHeader(http.StatusUnauthorized)
        return 
    }

    var ents_real []EntidadeStruct

    if err := json.NewDecoder(r.Body).Decode(&ents_real); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    
    ents := make([]Entidade, len(ents_real))

    for i := range ents_real {
        ents[i] = Entidade(ents_real[i])
    }
    
    
    w.Write([]byte(fmt.Sprint(s.AddContrato(NewContrato(len(s.Contratos), ents)))))
}
