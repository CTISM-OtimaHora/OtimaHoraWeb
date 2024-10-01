package models

import "errors"

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

func (c Contrato) HasParticipante(find ParticipanteContrato) (int, error) {
    for i, p := range c.Participantes {
        if p.GetId() == find.GetId() && p.GetNome() == find.GetNome() {
            return i, nil
        }
    }
    return -1, errors.New("Cant find item in Contrato")
}



