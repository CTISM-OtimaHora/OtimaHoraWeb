package models

type Turma struct {
    Id      int
    Nome    string
    Horarios Disponibilidade
}

func (t * Turma) GetId() int {
    return t.Id
}
func (t * Turma) GetNome() string {
    return t.Nome
}
func (t * Turma) GetDisponibilidade() Disponibilidade {
    return t.Horarios
}
func (t * Turma) GetTipo() string {
    return "turma"
}

func NewTurma (Id int, Nome string, dias, periodos int) Turma {
    return Turma {
        Id: Id,
        Nome: Nome,
        Horarios: NewDisponibilidade(),
    }
} 
