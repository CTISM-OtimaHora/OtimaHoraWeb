package models

import (
)

type ParticipanteContrato interface {
    GetId() int
    GetNome() string
    GetDispo() Disponibilidade
}

type Participante_to_json struct {
    Id int
    Nome string
    Tipo string
}

type Contrato struct {
    Id int
    Participantes []ParticipanteContrato
    Tipo_por_participante []string      // Tipo_por_participante[i] é o tipo do participante Participantes[i]
    Dispo       Disponibilidade
}

func (c Contrato) GetId() int{
    return c.Id
}

func NewContrato(id int, entidades []ParticipanteContrato) Contrato {
    return Contrato {
        Id: id,
        Participantes: entidades,
        Dispo: AndDisp(entidades),
    }
}

