package routes

import (
    "net/http"
    "encoding/json"
    "strconv"
    "fmt"

	. "github.com/CTISM-OtimaHora/OtimaHora/models"
)

func GetBuilder[T Entidade] (slice_geter func(*Session) []T) func(http.ResponseWriter, * http.Request) {
    return func (w http.ResponseWriter, r * http.Request) {
        s := Session_or_nil(r)
        if s == nil {
            w.WriteHeader(http.StatusUnauthorized)
            w.Write([]byte("No session or Session expired"))
            return 
        }

        id, conv_err := strconv.Atoi(r.PathValue("id"))
        if conv_err != nil {
            w.WriteHeader(http.StatusBadRequest)
            return
        }

        if json.NewEncoder(w).Encode(slice_geter(s)[id]) != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        return   
    }
}

func AddBuilder[T Entidade] (slice_adder func(sess *Session, en T) int) func (http.ResponseWriter, *http.Request) {
    return func (w http.ResponseWriter, r * http.Request) {
        s := Session_or_nil(r)
        if s == nil {
            w.WriteHeader(http.StatusUnauthorized)
            w.Write([]byte("No session or Session expired"))
            return 
        }

        var e T 

        if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte("malformed body 1: " + err.Error()))
            return
        }

        w.Write([]byte(fmt.Sprint(slice_adder(s, e))))
    }
}

func SliceGetBuilder[T Entidade] (slice_geter func(*Session) []T) func(http.ResponseWriter, * http.Request) {
    return func (w http.ResponseWriter, r * http.Request) {
        s := Session_or_nil(r)
        if s == nil {
            w.WriteHeader(http.StatusUnauthorized)
            w.Write([]byte("No session or Session expired"))
            return 
        }

        if json.NewEncoder(w).Encode(slice_geter(s)) != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        return   
    }
}

func DispoGetBuilder[T Entidade] (slice_geter func(*Session) []T) func(http.ResponseWriter, * http.Request) {
    return func (w http.ResponseWriter, r * http.Request) {
        s := Session_or_nil(r)
        if s == nil {
            w.WriteHeader(http.StatusUnauthorized)
            w.Write([]byte("No session or Session expired"))
            return 
        }

        id, conv_err := strconv.Atoi(r.PathValue("id"))
        if conv_err != nil {
            w.WriteHeader(http.StatusBadRequest)
            return
        }

        if json.NewEncoder(w).Encode(slice_geter(s)[id].GetDisponibilidade()) != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        return   
    }
}

func DispoSetBuilder[T Entidade] (slice_geter func(*Session) []T) func(http.ResponseWriter, * http.Request) {
    return func (w http.ResponseWriter, r * http.Request) {
        s := Session_or_nil(r)
        if s == nil {
            w.WriteHeader(http.StatusUnauthorized)
            w.Write([]byte("No session or Session expired"))
            return 
        }

        id, conv_err := strconv.Atoi(r.PathValue("id"))
        if conv_err != nil {
            w.WriteHeader(http.StatusBadRequest)
            return
        }
        
        var d Disponibilidade

        if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte("malformed body 1: " + err.Error()))
            return
        }

        p := slice_geter(s)[id].GetDisponibilidade()

        for i := range len(d) {
            for j := range len(d[0]) {
                (*p)[i][j] = d[i][j]           
            }
        }

        return   
    }
}
