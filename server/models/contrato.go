package models

import "fmt"

type Contrato struct {
    Id int
    Participantes []SearchEntidade
    Dispo       Disponibilidade
}

// é necessario recever entidades e não SearchEntidades pois seach entidades não possuem disponibilidades
// armazena-se as entidades em forma simplificada pois sua conversão para forma complexa é trivial
func NewContrato(id int, entidades []Entidade) Contrato {
    return Contrato {
        Id: id,
        Participantes: ToSearchSlice(entidades),
        Dispo: AndDisp(entidades),
    }
}

func (c Contrato) GetId() int {
    return c.Id
}

func (c Contrato) GetNome() string {
    return fmt.Sprintf("Contrato %d", c.Id)
}

func (c Contrato) GetDispo() *Disponibilidade {
    return  &c.Dispo
}

func (c Contrato) GetTipo() string {
    return "contrato"
}
