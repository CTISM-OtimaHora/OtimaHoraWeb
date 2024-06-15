package models

import (
    "github.com/google/uuid"
    "encoding/binary"
)

type HorasAula struct {
    Horas int
    Formato string // exemplo 2+2, 4 1+3
}

type Curso struct {
    Id int
    Nome string
    Dispo Disponibilidade
    Etapas [][]Turma
    Curriculo map[int]HorasAula // ID para Horas aula
}

func (c  Curso) GetId() int {
    return c.Id
}
func (c  Curso) GetNome() string {
    return c.Nome
}
func (c  Curso) GetDispo() *Disponibilidade {
    return &c.Dispo
}
func (c  Curso) GetTipo() string {
    return "curso"
}

func NewCurso (id int, nome string) Curso {
    return Curso{
        Id: id,
        Nome: nome,
        Dispo: NewDisponibilidade(),
        Etapas:  make([][]Turma, 0),
        Curriculo: make(map[int]HorasAula),
    }
}

func (c * Curso) AddTurma (etapa int, t Turma) int {
    for etapa >= len(c.Etapas) {
        c.Etapas = append(c.Etapas, make([]Turma, 0))
    }

    t.Id = int(binary.BigEndian.Uint32([]byte(uuid.NewString())[:4]))
    t.Curso_id = c.Id
    t.Etapa_idx = etapa
    c.Etapas[etapa] = append(c.Etapas[etapa], t)
    return t.Id
}
