package models

import (
	"fmt"
)

// Participantes deve ser do tipo SearchEntidade pois quando convertido "de" ou "para" JSON
// é importante que a informação seja independende da Session, coisa que uma simples Entidade
// não é
type Contrato struct {
    Id int
    Participantes []SearchEntidade
    Dispo       Disponibilidade
}

func NewContrato(id int, entidades []Entidade) Contrato {
    return Contrato {
        Id: id,
        Participantes: ToSearchSlice(entidades),
        Dispo: AndDisp(entidades),
    }
}

type json_contrato struct {
    Id int
    Participantes []SearchEntidade
    Dispo Disponibilidade
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
