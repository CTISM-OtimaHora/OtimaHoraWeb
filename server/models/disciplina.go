package models

type Disciplina struct {
    Id int
    Nome string
    Dispo Disponibilidade
}

func NewDisciplina(id int, nome string) Disciplina {
    return Disciplina {Id: id, Nome: nome, Dispo: NewDisponibilidade()}
}

func (d Disciplina) GetId() int {
    return d.Id
}
func (d  Disciplina) GetNome() string {
    return d.Nome
}
func (d  Disciplina) GetDisponibilidade() *Disponibilidade {
    return &d.Dispo
}
func (d  Disciplina) GetTipo() string {
    return "disciplina"
}

