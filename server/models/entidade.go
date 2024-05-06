package models

type Entidade interface {
    GetId() int
    GetNome() string
    GetDisponibilidade() Disponibilidade 
}

type EntidadeStruct struct {
    Id int
    Nome string
    Dispo Disponibilidade
}

func (e EntidadeStruct) GetId() int {
    return e.Id
}
func (e EntidadeStruct) GetNome() string {
    return e.Nome
}
func (e EntidadeStruct) GetDisponibilidade() Disponibilidade {
    return e.Dispo
}
