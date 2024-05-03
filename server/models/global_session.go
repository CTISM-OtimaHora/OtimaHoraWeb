package models

import (
    "net/http"
    "strconv"
)



var Sessions = []*Session{}

var next_id int = 0

func Session_or_nil(w http.ResponseWriter, r * http.Request) *Session {
    session_cookie, err := r.Cookie("id")
    if err != nil {
        return nil
    }

    idx, _ := strconv.Atoi(session_cookie.Value)
    if idx < len(Sessions) {
        return (Sessions)[idx]
    }
    return nil
}

func CreateSession () *Session {
    s := NewSession(next_id)
    Sessions = append(Sessions, &s)
    next_id += 1
    return &s
}

func AddSession (sess * Session) {
    Sessions = append(Sessions, sess)
}


