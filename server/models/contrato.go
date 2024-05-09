package models

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
