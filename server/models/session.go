package models

type Session struct {
    Id              int
    Cursos          []Curso
    Professores     []Professor
    Disciplinas     []Disciplina
    Contratos       []Contrato
    Recursos        []Recurso
}

func NewSession (id int) Session {
    return Session{Id: id}
}

func AddCurso (s * Session, c Curso) int {
    c.Dispo = NewDisponibilidade()

    c.Id = len(s.Cursos)
    s.Cursos = append(s.Cursos, c)
    return c.Id
}

func AddProfessor (s * Session, p Professor) int {
    p.Dispo = NewDisponibilidade()

    p.Id = len(s.Professores)
    s.Professores = append(s.Professores, p)
    return p.Id
}

func AddRecurso (s * Session, r Recurso) int {
    r.Dispo = NewDisponibilidade()

    r.Id = len(s.Recursos)
    s.Recursos = append(s.Recursos, r)
    return r.Id
}

func AddDisciplina (s * Session, d Disciplina) int {
    d.Dispo = NewDisponibilidade()

    d.Id = len(s.Disciplinas)
    s.Disciplinas = append(s.Disciplinas, d)
    return d.Id
}

func (s * Session) AddContrato (c Contrato) int {
    c.Id = len(s.Contratos)
    s.Contratos = append(s.Contratos, c)
    return c.Id
}


func ProfessorGeter(s * Session) []Professor {
    return s.Professores
}
func DisciplinaGeter(s * Session) []Disciplina {
    return s.Disciplinas
}
func CursoGetter(s * Session) []Curso {
    return s.Cursos
}
func ContratoGetter(s * Session) []Contrato {
    return s.Contratos
}
func RecursoGetter(s * Session) []Recurso {
    return s.Recursos
}
