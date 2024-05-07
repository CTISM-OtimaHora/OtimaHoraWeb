package routes

import (
    "net/http"
    "time"
    "fmt"
    "encoding/json"

    ."github.com/CTISM-OtimaHora/OtimaHora/models"
)

func New_session(w http.ResponseWriter, r * http.Request) {
    s := Session_or_nil(w, r)
    if s != nil {
        http.SetCookie(w, &http.Cookie{Name: "id", Value: fmt.Sprint(s.Id), Path: "/", Expires: time.Now().Add(2*time.Hour)})
        return
    }

    new_s := CreateSession() 
    http.SetCookie(w, &http.Cookie{Name: "id", Value: fmt.Sprint(new_s.Id), Path: "/", Expires: time.Now().Add(2*time.Hour)})
    return
}

func Add_curso_to_session(w http.ResponseWriter, r * http.Request) {
    if r.Method != "POST" {
        w.WriteHeader(http.StatusNotFound)
        return
    }

    s := Session_or_nil(w, r)
    if s == nil {
        w.Write([]byte("No session or Session expired"))
        w.WriteHeader(http.StatusUnauthorized)
        return // errors already in w
    }

    var curso Curso
    if err := json.NewDecoder(r.Body).Decode(&curso); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("Malformed body"))
        return
    }
    w.Write([]byte(fmt.Sprint(s.AddCurso(curso))))
    return
}

func Get_session(w http.ResponseWriter, r * http.Request) {
    s := Session_or_nil(w, r)
    if s == nil {
        w.Write([]byte("No session or Session expired"))
        w.WriteHeader(http.StatusUnauthorized)
        return 
    }

    json.NewEncoder(w).Encode(s)
    return
}
