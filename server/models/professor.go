package models

type Professor struct {
    Id int
    Nome string
    Dispo Disponibilidade
}


func NewProfessor(id int, nome string) Professor {
    return Professor {Id: id, Nome: nome, Dispo: NewDisponibilidade()}
}

func (p* Professor) GetId() int {
    return p.Id
}
func (p * Professor) GetNome() string {
    return p.Nome
}
func (p * Professor) GetDisponibilidade() Disponibilidade {
    return p.Dispo
}
func (p * Professor) GetTipo() string {
    return "professor"
}
