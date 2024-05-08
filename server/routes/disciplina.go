package routes

import (
    "net/http"
    "encoding/json"
    "fmt"

    ."github.com/CTISM-OtimaHora/OtimaHora/models"
)

func Add_Disciplina_to_turma(w http.ResponseWriter, r * http.Request) {
    if r.Method != "POST" {
        w.WriteHeader(404)
        return
    }   
    
    s := Session_or_nil(r)
    if s == nil {
        w.Write([]byte("No session or Session expired"))
        w.WriteHeader(http.StatusUnauthorized)
        return 
    }

    var dis Disciplina 

    if err := json.NewDecoder(r.Body).Decode(&dis); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    
    w.Write([]byte(fmt.Sprint(s.AddDisciplina(dis))))
}
