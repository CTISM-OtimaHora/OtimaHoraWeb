package models

type Curso struct {
    Id int
    Nome string
    Dispo Disponibilidade
    Turmas []Turma

}

func (c  Curso) GetId() int {
    return c.Id
}
func (c  Curso) GetNome() string {
    return c.Nome
}
func (c  Curso) GetDisponibilidade() *Disponibilidade {
    return &c.Dispo
}
func (c  Curso) GetTipo() string {
    return "curso"
}

func NewCurso (id int, nome string) Curso {
    return Curso{Id: id, Nome: nome, Turmas: make([]Turma, 0)}
}

func (c * Curso) AddTurma (t Turma) int {
    t.Id = len(c.Turmas)
    c.Turmas = append(c.Turmas, t)
    return t.Id
}
