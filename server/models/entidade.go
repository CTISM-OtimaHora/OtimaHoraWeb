package models


// é necessário definir dois tipos de entidades, uma interface genérica para as funcionalidaes de back-end e uma Entidade de Busca concreta
// que é o tipo utilizado pelo front-end esta é jsonificavel e permite conversão simples para generica

type Entidade interface {
    GetId() int
    GetNome() string
    GetDisponibilidade() Disponibilidade 
    GetTipo() string
}

func ToSearch (e Entidade) SearchEntidade {
    return SearchEntidade{
        Id: e.GetId(),
        Nome: e.GetNome(),
        Tipo: e.GetTipo(),
    }
}

func ToSearchSlice (e []Entidade) []SearchEntidade {
    se := make ([]SearchEntidade, len(e))
    for i := range len(e) {
        se[i] = ToSearch(e[i])
    }
    return se
}

type SearchEntidade struct {
    Id int
    Nome string
    Tipo string
}

func (e SearchEntidade) GetId() int {
    return e.Id
}
func (e SearchEntidade) GetNome() string {
    return e.Nome
}

func (e SearchEntidade) GetEntidadeOrNil(s * Session) Entidade {
    switch e.Tipo {
        case "professor":
            return Entidade(&s.Professores[e.Id])
        case "disciplina":
            return Entidade(&s.Disciplinas[e.Id])
        case "curso":
            return Entidade(&s.Cursos[e.Id])
        case "recurso":
            // TODO
            // return Entidade(&s.Recursos[e.Id])
    }
    return nil
}

func GetEntidadesOrNilSlice(es []SearchEntidade, s * Session) []Entidade {
    es2 := make([]Entidade, len(es))
    for i := range es {
        a := es[i].GetEntidadeOrNil(s)
        if a == nil {
            return nil
        }
        es2[i] = a
    }
    return es2
}
