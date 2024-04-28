package main

import "fmt"

type Turma struct {
    Id      int
    Nome    string
    // Disciplinas string
    // Horarios string
}

func (t Turma) ToString () string {
   return fmt.Sprintf("{id: %v, nome: %v}", t.Id, t.Nome)
}
