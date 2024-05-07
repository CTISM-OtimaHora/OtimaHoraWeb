package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/CTISM-OtimaHora/OtimaHora/models"
)

func Add_Professor_to_turma(w http.ResponseWriter, r * http.Request) {
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

    var prof Professor

    if err := json.NewDecoder(r.Body).Decode(&prof); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    
    w.Write([]byte(fmt.Sprint(s.AddProfessor(prof))))
}
