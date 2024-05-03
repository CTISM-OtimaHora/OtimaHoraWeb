package models

import (
	"encoding/json"
	"bytes"
)


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
    Disciplinas []Disciplina
    Professores []Professor
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
        Disciplinas: make([]Disciplina, 0),
        Professores: make([]Professor, 0),
    }
} 

func (t * Turma) AddProfessor (p  Professor) {
    t.Professores = append(t.Professores, p)
}
func (t * Turma) AddDisciplina (d Disciplina) {
    t.Disciplinas = append(t.Disciplinas, d)
}

func (t Turma) ToString () string {
    b := bytes.NewBuffer(make([]byte, 2000))
    _ = json.NewEncoder(b).Encode(t)
    return b.String()
}

func (t * Turma) SetHorario (h [][]uint8) {
    for i := range t.Horarios {
        for j := range t.Horarios[i] {
            t.Horarios[i][j] = h[i][j]
        }
    }
}
