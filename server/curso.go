package main

import "fmt"

type Curso struct {
    Id int
    Nome string
    turmas []Turma
}

func (c * Curso) Add (t Turma) int {
    t.Id = len(c.turmas)
    c.turmas = append(c.turmas, t)
    return t.Id
}

func (c * Curso) ToString() string {
    turmas := "["
    for _, t := range c.turmas {
        turmas += t.ToString() + ", "
    }
    turmas += "]"
    return fmt.Sprintf("{id: %v, nomes: %v turmas: %v}", c.Id, c.Nome, turmas)
    
}
