package routes

import (
	"encoding/json"
	"net/http"
	"slices"
	"strconv"

	. "github.com/CTISM-OtimaHora/OtimaHora/models"
)

// "GET /etapa/get/{curso_id}/{id}"
func GetEtapa(w http.ResponseWriter, r *http.Request) {
	s := Session_or_nil(r)
	if s == nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("No session or Session expired"))
		return
	}

	curso_id, conv_err1 := strconv.Atoi(r.PathValue("curso_id"))
	etapa_id, conv_err2 := strconv.Atoi(r.PathValue("id"))
	if conv_err1 != nil || conv_err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Conversion error from path parameter"))
		return
	}

	etapa := s.Cursos[curso_id].Etapas[etapa_id]

	if err := json.NewEncoder(w).Encode(etapa); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
}

// "GET /etapa/get/{curso_id}/{id}"
func SetEtapa(w http.ResponseWriter, r *http.Request) {
	s := Session_or_nil(r)
	if s == nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("No session or Session expired"))
		return
	}

	curso_id, conv_err1 := strconv.Atoi(r.PathValue("curso_id"))
	etapa_id, conv_err2 := strconv.Atoi(r.PathValue("id"))
	if conv_err1 != nil || conv_err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Conversion error from path parameter"))
		return
	}

	var etapa Etapa

	if err := json.NewDecoder(r.Body).Decode(&etapa); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	s.Cursos[curso_id].Etapas[etapa_id] = etapa
}

func DeleteEtapa(w http.ResponseWriter, r *http.Request) {
	s := Session_or_nil(r)
	if s == nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("No session or Session expired"))
		return
	}

	curso_id, conv_err1 := strconv.Atoi(r.PathValue("curso_id"))
	etapa_id, conv_err2 := strconv.Atoi(r.PathValue("id"))
	if conv_err1 != nil || conv_err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Conversion error from path parameter"))
		return
	}

	curso, ok := s.Cursos[curso_id]
	if ok {
		curso.Etapas = slices.Delete(curso.Etapas, etapa_id, etapa_id+1)
	}
}
