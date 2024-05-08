package models

type Session struct {
    Id              int
    Cursos          []Curso
    Professores     []Professor
    Disciplinas     []Disciplina
    Contratos       []Contrato
}

func NewSession (id int) Session {
    return Session{Id: id}
}

func (s * Session) AddCurso (c Curso) int {
    c.Dispo = NewDisponibilidade()

    c.Id = len(s.Cursos)
    s.Cursos = append(s.Cursos, c)
    return c.Id
}

func (s * Session) AddProfessor (p Professor) int {
    p.Dispo = NewDisponibilidade()

    p.Id = len(s.Professores)
    s.Professores = append(s.Professores, p)
    return p.Id
}

func (s * Session) AddDisciplina (d Disciplina) int {
    d.Dispo = NewDisponibilidade()

    d.Id = len(s.Disciplinas)
    s.Disciplinas = append(s.Disciplinas, d)
    return d.Id
}

func (s * Session) AddContrato (c Contrato) int {
    c.Dispo = NewDisponibilidade()

    c.Id = len(s.Contratos)
    s.Contratos = append(s.Contratos, c)
    return c.Id
}

