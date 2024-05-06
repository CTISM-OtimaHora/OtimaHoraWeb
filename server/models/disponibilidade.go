package models

const (
    YES = 1
    MAYBE = 0
    NO = -1

    DIAS = 5
    PERIODOS = 5
)

type Disponibilidade [][]int8

func NewDisponibilidade() Disponibilidade {
    d := make([][]int8, DIAS)
    for i := range DIAS {
        d[i] = make([]int8, PERIODOS)
        for j := range PERIODOS {
            d[i][j] = YES
        }
    }
    return Disponibilidade(d)
}

func AndDisp (ents []Entidade) Disponibilidade {
    dispos := make([]Disponibilidade, len(ents))
    for i := range ents {
        dispos[i] = ents[i].GetDisponibilidade()
    }

    d := NewDisponibilidade()
    for i := range DIAS {
        for j := range PERIODOS {
            var dispo_do_dia int8
            dispo_do_dia = YES
            for _, disp := range dispos {
                if disp[i][j] < dispo_do_dia {
                    dispo_do_dia = disp[i][j]
                }
            }
            d[i][j] = dispo_do_dia
        }
    }
    return d
}


