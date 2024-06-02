package models

import (
    "encoding/json"
)

type Session struct {
    Id              int
    Cursos          map[int]Curso
    Professores     map[int]Professor
    Disciplinas     map[int]Disciplina
    Contratos       map[int]Contrato
    Recursos        map[int]Recurso
}

func NewSession (id int) Session {
    return Session{
        Id: id,
        Cursos:         make(map[int]Curso),
        Professores:    make(map[int]Professor),
        Disciplinas:    make(map[int]Disciplina),
        Contratos:      make(map[int]Contrato),
        Recursos:       make(map[int]Recurso),
    }
}

func AddCurso (s * Session, c Curso) int {
    c.Dispo = NewDisponibilidade()

    c.Id = len(s.Cursos)
    s.Cursos[c.Id] = c
    return c.Id
}

func AddProfessor (s * Session, p Professor) int {
    p.Dispo = NewDisponibilidade()

    p.Id = len(s.Professores)
    s.Professores[p.Id] = p
    return p.Id
}

func AddRecurso (s * Session, r Recurso) int {
    r.Dispo = NewDisponibilidade()

    r.Id = len(s.Recursos)
    s.Recursos[r.Id] = r
    return r.Id
}

func AddDisciplina (s * Session, d Disciplina) int {
    d.Dispo = NewDisponibilidade()

    d.Id = len(s.Disciplinas)
    s.Disciplinas[d.Id] = d
    return d.Id
}

func (s * Session) AddContrato (c Contrato) int {
    c.Id = len(s.Contratos)
    s.Contratos[c.Id] = c
    return c.Id
}


func ProfessorGeter(s * Session) map[int]Professor {
    return s.Professores
}
func DisciplinaGeter(s * Session) map[int]Disciplina {
    return s.Disciplinas
}
func CursoGetter(s * Session) map[int]Curso {
    return s.Cursos
}
func ContratoGetter(s * Session) map[int]Contrato {
    return s.Contratos
}
func RecursoGetter(s * Session) map[int]Recurso {
    return s.Recursos
}

func map_to_slice[T Entidade, A comparable](m map[A]T) []T {
    slice := make([]T, 0, len(m))
    for _, v := range m {
        slice = append(slice, v)
    }
    return slice
}

func slice_to_map[T Entidade](s []T) map[int]T {
    m := make(map[int]T)
    for i, v := range s {
        m[i] = v
    }
    return m 
}

type json_session struct {
    Id              int
    Cursos          []Curso
    Professores     []Professor
    Disciplinas     []Disciplina
    Contratos       []Contrato
    Recursos        []Recurso
}


func (s Session) MarshalJSON() ([]byte, error) {
    a := json_session{
        s.Id,
        map_to_slice(s.Cursos),
        map_to_slice(s.Professores),
        map_to_slice(s.Disciplinas),
        map_to_slice(s.Contratos),
        map_to_slice(s.Recursos),
    }
    return json.Marshal(a)
}

func (s * Session) UnmarshalJSON(bs []byte) error {
    var a json_session
    err := json.Unmarshal(bs, &a)
    if err != nil {
        return err
    }
    *s = Session{
        a.Id,
        slice_to_map(a.Cursos),
        slice_to_map(a.Professores),
        slice_to_map(a.Disciplinas),
        slice_to_map(a.Contratos),
        slice_to_map(a.Recursos),
    }
    return nil
}
