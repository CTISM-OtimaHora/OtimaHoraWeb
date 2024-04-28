package main

import "fmt"

type Session struct {
    Id int
    Cursos []Curso
}

func (s * Session) Add (c Curso) int {
    c.Id = len(s.Cursos)
    s.Cursos = append(s.Cursos, c)
    return c.Id
}

func (s * Session) ToString() string {
    cursos := "["
    for _, c := range s.Cursos {
        cursos += c.ToString() + ", "
    }
    cursos += "]"
    
    return fmt.Sprintf("{id: %v, cursos: %v}", s.Id, cursos)
}   
