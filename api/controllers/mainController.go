package controllers

import (
	"net/http"
	"strconv"

	"github.com/Gubaz311/myOwnUnila_server/api/middleware"
	"github.com/Gubaz311/myOwnUnila_server/api/models"
)

func (s *Server) GetAllMahasiswa(w http.ResponseWriter, r *http.Request) {
	tahunAwal, err := strconv.Atoi(r.URL.Query().Get("tahun_awal"))
	if err != nil {
		middleware.Jsoncontenterror(w, http.StatusInternalServerError, err)
		return
	}

	tahunAkhir, err := strconv.Atoi(r.URL.Query().Get("tahun_akhir"))
	if err != nil {
		middleware.Jsoncontenterror(w, http.StatusInternalServerError, err)
		return
	}

	data, err := models.GetAllData(tahunAwal, tahunAkhir)

	if err != nil {
		middleware.Jsoncontenterror(w, http.StatusInternalServerError, err)
	}

	middleware.Jsoncontent(w, http.StatusOK, data)
}
