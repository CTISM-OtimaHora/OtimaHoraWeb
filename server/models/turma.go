package models

import "fmt"


const (
    YES = 2
    MAYBE = 1
    NO = 0
)

type Turma struct {
    Id      int
    Nome    string
    // Disciplinas string
    Horarios [][]uint8 
    Dias int
    Periodos int
    Disciplinas []string
    Professores []string
}

func NewTurma (Id int, Nome string, dias, periodos int) Turma {
    h := make([][]uint8, dias)
    for i := range h {
        h[i] = make([]uint8, periodos)
        for j := range h[i] {
            h[i][j] = YES
        }
    }

    return Turma {
        Id: Id,
        Nome: Nome,
        Horarios: h,
        Dias: dias,
        Periodos: periodos,
        Disciplinas: make([]string, 0),
        Professores: make([]string, 0),
    }
} 

func (t Turma) ToString () string {
    s := ""
    for i := range t.Horarios {
        for j := range t.Horarios[i] {
            s += fmt.Sprintf("%v ", t.Horarios[i][j])
        }
    }
    return fmt.Sprintf("{id: %v, nome: %v, horarios: %vx%v: %v}", t.Id, t.Nome, t.Dias, t.Periodos, s)
}
