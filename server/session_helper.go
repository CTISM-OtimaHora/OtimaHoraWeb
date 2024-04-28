package main

import (
    "net/http"
    "strconv"
)

func Session_or_nil(w http.ResponseWriter, r * http.Request) *Session {
    session_cookie, err := r.Cookie("id")
    if err != nil {
        return nil
    }

    idx, _ := strconv.Atoi(session_cookie.Value)
    if idx < len(Sessions) {
        return Sessions[idx]
    }
    return nil
}
