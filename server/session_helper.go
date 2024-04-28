package main

import (
    "net/http"
    "strconv"
)

func Session_or_error(w http.ResponseWriter, r * http.Request) *Session {
    session_cookie, err := r.Cookie("id")
    if err != nil {
        w.Write([]byte("No session or Session expired"))
        w.WriteHeader(http.StatusUnauthorized)
        return nil
    }
    idx, _ := strconv.Atoi(session_cookie.Value)
    return Sessions[idx]
}
