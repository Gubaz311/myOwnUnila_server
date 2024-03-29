package controllers

import (
	"net/http"
	"strconv"

	"github.com/Gubaz311/myOwnUnila_server/api/middleware"
	"github.com/Gubaz311/myOwnUnila_server/api/models"
)

// Get all data from all tables
func (server *Server) GetAllMahasiswa(w http.ResponseWriter, r *http.Request) {
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
	db := server.DB
	mahasiswa := models.Mahasiswa{}
	data, err := mahasiswa.GetAllData(db, tahunAwal, tahunAkhir)

	if err != nil {
		middleware.Jsoncontenterror(w, http.StatusInternalServerError, err)
	}

	middleware.Jsoncontent(w, http.StatusOK, data)
}
