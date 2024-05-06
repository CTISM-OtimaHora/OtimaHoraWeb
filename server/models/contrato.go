package models

type Contrato struct {
    Id int
    Participantes []Entidade
    Dispo       Disponibilidade
}

func NewContrato(id int, entidades []Entidade) Contrato {
    return Contrato {
        Id: id,
        Participantes: entidades,
        Dispo: AndDisp(entidades),
    }
}
