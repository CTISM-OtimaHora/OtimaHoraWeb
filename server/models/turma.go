package models

type Turma struct {
    Id      int
    Curso_id int
    Etapa_idx int
    Nome    string
    Dispo Disponibilidade
}

func (t Turma) GetId() int {
    return t.Id
}
func (t Turma) GetNome() string {
    return t.Nome
}
func (t Turma) GetDispo() Disponibilidade {
    return t.Dispo
}

func NewTurma (Id int, Nome string, dias, periodos int) Turma {
    return Turma {
        Id: Id,
        Nome: Nome,
        Dispo: NewDisponibilidade(),
    }
} 
