package routes

import (
    "net/http"
    "encoding/json"
    "strconv"

    ."github.com/CTISM-OtimaHora/OtimaHora/models"
)

func Get_turma(w http.ResponseWriter, r * http.Request) {
    s := Session_or_nil(w, r)
    if s == nil {
        w.Write([]byte("No session or Session expired"))
        w.WriteHeader(http.StatusUnauthorized)
        return 
    }

    curso_idx, conv_err := strconv.Atoi(r.PathValue("id_curso"))
    turma_idx, conv_err2 := strconv.Atoi(r.PathValue("id_turma"))
    if conv_err != nil || conv_err2 != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    json.NewEncoder(w).Encode(s.Cursos[curso_idx].Turmas[turma_idx])
    return 
}



func Set_Horario_Turma(w http.ResponseWriter, r * http.Request) {
    // estrutura do horario ainda tem q ser definida
    // opções:  


    // representação uni-dimensional de matriz
    //  "horarios": {"<n-dias>X<n-periodos>" : ["SIM", "NAO", "TALVEZ", "NAO", "SIM", "SIM"]}

    // representação de matriz
    // "horarios": [["SIM", "NAO", "TALVEZ"], ["NAO", "SIM", "SIM]]

    // representação por dias
    // "horarios": {"segunda": ["SIM", "NAO", "TALVEZ"], "terca": ["NAO", "SIM", "SIM"]}

}
