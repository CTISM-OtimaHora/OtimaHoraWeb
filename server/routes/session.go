package routes

import (
    "net/http"
    "time"
    "fmt"
    "encoding/json"

    ."github.com/CTISM-OtimaHora/OtimaHora/models"
)

func AddSession(w http.ResponseWriter, r * http.Request) {
    s := Session_or_nil(r)
    if s != nil {
        http.SetCookie(w, &http.Cookie{Name: "id", Value: fmt.Sprint(s.Id), Path: "/", Expires: time.Now().Add(2*time.Hour)})
        return
    }

    new_s := CreateSession() 
    http.SetCookie(w, &http.Cookie{Name: "id", Value: fmt.Sprint(new_s.Id), Path: "/", Expires: time.Now().Add(2*time.Hour)})
    return
}

func GetSession(w http.ResponseWriter, r * http.Request) {
    s := Session_or_nil(r)
    if s == nil {
        w.WriteHeader(http.StatusUnauthorized)
        w.Write([]byte("No session or Session expired"))
        return 
    }

    json.NewEncoder(w).Encode(s)
    return
}
