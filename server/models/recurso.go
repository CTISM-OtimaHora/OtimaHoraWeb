package models

type Recurso struct {
    Id                  int
    Nome                string 
    Dispo               Disponibilidade
}

func NewRecurso(id int, nome string) Recurso {
    return Recurso {Id: id, Nome: nome, Dispo: NewDisponibilidade()}
}

func (r Recurso) GetId() int {
    return r.Id
}
func (r Recurso) GetNome() string {
    return r.Nome
}
func (r Recurso) GetDisponibilidade() *Disponibilidade {
    return &r.Dispo
}
func (r Recurso) GetTipo() string {
    return "recurso"
}
