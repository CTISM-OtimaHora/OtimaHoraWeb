package models

import "fmt"

type Curso struct {
    Id int
    Nome string
    Turmas []Turma
}

func NewCurso (id int, nome string) Curso {
    return Curso{Id: id, Nome: nome, Turmas: make([]Turma, 0)}
}

func (c * Curso) Add (t Turma) int {
    t.Id = len(c.Turmas)
    c.Turmas = append(c.Turmas, t)
    return t.Id
}

func (c * Curso) ToString() string {
    turmas := "["
    for _, t := range c.Turmas {
        turmas += t.ToString() + ", "
    }
    turmas += "]"
    return fmt.Sprintf("{id: %v, nomes: %v turmas: %v}", c.Id, c.Nome, turmas)
    
}