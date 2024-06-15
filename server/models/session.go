package models

import (
	"encoding/binary"
	"encoding/json"
	"slices"

	"github.com/google/uuid"
)

type Session struct {
    Id              int
    Cursos          map[int]Curso
    Professores     map[int]Professor
    Disciplinas     map[int]Disciplina
    Contratos       map[int]Contrato
    Recursos        map[int]Recurso
}

func NewSession (id int) Session {
    return Session{
        Id: id,
        Cursos:         make(map[int]Curso),
        Professores:    make(map[int]Professor),
        Disciplinas:    make(map[int]Disciplina),
        Contratos:      make(map[int]Contrato),
        Recursos:       make(map[int]Recurso),
    }
}

func AddCurso (s * Session, c Curso) int {
    c.Dispo = NewDisponibilidade()

    c.Id = len(s.Cursos)
    s.Cursos[c.Id] = c
    return c.Id
}

func AddTurma (s * Session, t Turma) int {
    curso, _  := s.Cursos[t.Curso_id] // must exist
    return curso.AddTurma(t.Etapa_idx, t)

} 

func AddProfessor (s * Session, p Professor) int {
    p.Dispo = NewDisponibilidade()

    // gera uuid, pega o final, converte pra binario 32, converte pra 32bit sinalizado
    p.Id = int(binary.BigEndian.Uint32([]byte(uuid.NewString())[:4]))
    s.Professores[p.Id] = p
    return p.Id
}

func AddRecurso (s * Session, r Recurso) int {
    r.Dispo = NewDisponibilidade()

    r.Id = int(binary.BigEndian.Uint32([]byte(uuid.NewString())[:4]))
    s.Recursos[r.Id] = r
    return r.Id
}

func AddDisciplina (s * Session, d Disciplina) int {
    d.Dispo = NewDisponibilidade()

    d.Id = int(binary.BigEndian.Uint32([]byte(uuid.NewString())[:4]))
    s.Disciplinas[d.Id] = d
    return d.Id
}

func (s * Session) AddContrato (c Contrato) int {
    c.Id = int(binary.BigEndian.Uint32([]byte(uuid.NewString())[:4]))
    s.Contratos[c.Id] = c
    return c.Id
}

func (s * Session) UpdateContratos (changed Entidade) {
    for i, c := range s.Contratos {
        if !slices.ContainsFunc(c.Participantes, func(e SearchEntidade)bool{return e.Tipo == changed.GetTipo() && e.GetId() == changed.GetId()}) {
            continue
        }
        
        new_c := NewContrato(0, make([]Entidade, 0)) // empty

        c.Participantes[slices.IndexFunc(c.Participantes, func(e SearchEntidade)bool {
            return e.Tipo == changed.GetTipo() && e.GetId() == changed.GetId()
        })] = ToSearch(changed)

        new_c.Participantes = c.Participantes
        new_c.Id = c.Id
        new_c.Dispo = AndDisp(GetEntidadesOrNilSlice(c.Participantes, s)) // garanteed to not be nil 

        delete (s.Contratos, i)
        s.Contratos[i] = new_c
    }
}

func (s * Session) UpdateSessionFromDelete (changed Entidade) {
    for i, c := range s.Contratos {
        if !slices.ContainsFunc(c.Participantes, func(e SearchEntidade)bool{return e.Id == changed.GetId() && e.Tipo == changed.GetTipo()}) {
            continue
        }
        
        delete (s.Contratos, i)
    }

    // deve também alterar currículos
    if changed.GetTipo() == "disciplina" {
        for _, c := range s.Cursos {
            _, ok := c.Curriculo[changed.GetId()]
            if ok {
                delete(c.Curriculo, changed.GetId())
            }
        }
    }
}


func ProfessorGeter(s * Session) map[int]Professor {
    return s.Professores
}
func DisciplinaGeter(s * Session) map[int]Disciplina {
    return s.Disciplinas
}
func CursoGetter(s * Session) map[int]Curso {
    return s.Cursos
}
func ContratoGetter(s * Session) map[int]Contrato {
    return s.Contratos
}
func RecursoGetter(s * Session) map[int]Recurso {
    return s.Recursos
}

func map_to_slice[T Entidade, A comparable](m map[A]T) []T {
    slice := make([]T, 0, len(m))
    for _, v := range m {
        slice = append(slice, v)
    }
    return slice
}

func slice_to_map[T Entidade](s []T) map[int]T {
    m := make(map[int]T)
    for _, v := range s {
        m[v.GetId()] = v
    }
    return m 
}

type json_session struct {
    Id              int
    Cursos          []Curso
    Professores     []Professor
    Disciplinas     []Disciplina
    Contratos       []Contrato
    Recursos        []Recurso
}


func (s Session) MarshalJSON() ([]byte, error) {
    a := json_session{
        s.Id,
        map_to_slice(s.Cursos),
        map_to_slice(s.Professores),
        map_to_slice(s.Disciplinas),
        map_to_slice(s.Contratos),
        map_to_slice(s.Recursos),
    }
    return json.Marshal(a)
}

func (s * Session) UnmarshalJSON(bs []byte) error {
    var a json_session
    err := json.Unmarshal(bs, &a)
    if err != nil {
        return err
    }
    *s = Session{
        a.Id,
        slice_to_map(a.Cursos),
        slice_to_map(a.Professores),
        slice_to_map(a.Disciplinas),
        slice_to_map(a.Contratos),
        slice_to_map(a.Recursos),
    }
    return nil
}
