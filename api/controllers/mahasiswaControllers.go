package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Gubaz311/myOwnUnila_server/api/middleware"
	"github.com/Gubaz311/myOwnUnila_server/api/models"
	"github.com/gorilla/mux"
)

// Get data from specific table
func (server *Server) GetData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println("Mengambil data specific table")
	fmt.Println("Nilai dari fakultas: ", vars["idFakultas"])
	fmt.Println("Nilai tahunAwal:", vars["tahunAwal"])
	fmt.Println("Nilai tahunAkhir:", vars["tahunAkhir"])

	idFakultas, err := strconv.ParseUint(vars["idFakultas"], 10, 64)
	if err != nil {
		middleware.Jsoncontenterror(w, http.StatusBadRequest, err)
		return
	}
	tahunAwal, err := strconv.ParseUint(vars["tahunAwal"], 10, 64)
	if err != nil {
		middleware.Jsoncontenterror(w, http.StatusBadRequest, err)
		return
	}
	tahunAkhir, err := strconv.ParseUint(vars["tahunAkhir"], 10, 64)
	if err != nil {
		middleware.Jsoncontenterror(w, http.StatusBadRequest, err)
		return
	}

	mahasiswa := models.Mahasiswa{}
	data, err := mahasiswa.GetDataByRange(server.DB, tahunAwal, tahunAkhir, (idFakultas))
	if err != nil {
		middleware.Jsoncontenterror(w, http.StatusBadRequest, err)
		return
	}
	middleware.Jsoncontent(w, http.StatusOK, data)
}

// Get all data
func (server *Server) GetAll(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println("Mengambil data all table")
	fmt.Println("Nilai tahunAwal:", vars["tahunAwal"])
	fmt.Println("Nilai tahunAkhir:", vars["tahunAkhir"])

	tahunAwal, err := strconv.ParseUint(vars["tahunAwal"], 10, 64)
	if err != nil {
		middleware.Jsoncontenterror(w, http.StatusBadRequest, err)
		return
	}
	tahunAkhir, err := strconv.ParseUint(vars["tahunAkhir"], 10, 64)
	if err != nil {
		middleware.Jsoncontenterror(w, http.StatusBadRequest, err)
		return
	}

	mahasiswa := models.Mahasiswa{}
	data, err := mahasiswa.GetDataByRange(server.DB, tahunAwal, tahunAkhir, 0)
	if err != nil {
		middleware.Jsoncontenterror(w, http.StatusBadRequest, err)
		return
	}

	middleware.Jsoncontent(w, http.StatusOK, data)
}

// Get data for tglKeluar != nil
// GetDataLulus
// Get data from specific table

func (server *Server) GetDataLulus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println("Mengambil data specific table lulus only")
	fmt.Println("Nilai dari fakultas: ", vars["idFakultas"])
	fmt.Println("Nilai tahunAwal:", vars["tahunAwal"])
	fmt.Println("Nilai tahunAkhir:", vars["tahunAkhir"])

	idFakultas, err := strconv.ParseUint(vars["idFakultas"], 10, 64)
	if err != nil {
		middleware.Jsoncontenterror(w, http.StatusBadRequest, err)
		return
	}
	tahunAwal, err := strconv.ParseUint(vars["tahunAwal"], 10, 64)
	if err != nil {
		middleware.Jsoncontenterror(w, http.StatusBadRequest, err)
		return
	}

	tahunAkhir, err := strconv.ParseUint(vars["tahunAkhir"], 10, 64)
	if err != nil {
		middleware.Jsoncontenterror(w, http.StatusBadRequest, err)
		return
	}

	mahasiswa := models.Mahasiswa{}
	data, err := mahasiswa.GetDataLulus(server.DB, tahunAwal, tahunAkhir, (idFakultas))
	if err != nil {
		middleware.Jsoncontenterror(w, http.StatusBadRequest, err)
		return
	}
	middleware.Jsoncontent(w, http.StatusOK, data)
}

// Get all data
func (server *Server) GetAllLulus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println("Mengambil data all table lulus only")
	fmt.Println("Nilai dari fakultas: ", vars["idFakultas"])
	fmt.Println("Nilai tahunAwal:", vars["tahunAwal"])
	fmt.Println("Nilai tahunAkhir:", vars["tahunAkhir"])

	tahunAwal, err := strconv.ParseUint(vars["tahunAwal"], 10, 64)
	if err != nil {
		middleware.Jsoncontenterror(w, http.StatusBadRequest, err)
		return
	}
	fmt.Println(tahunAwal)
	tahunAkhir, err := strconv.ParseUint(vars["tahunAkhir"], 10, 64)
	if err != nil {
		middleware.Jsoncontenterror(w, http.StatusBadRequest, err)
		return
	}
	fmt.Println(tahunAkhir)

	mahasiswa := models.Mahasiswa{}
	data, err := mahasiswa.GetDataLulus(server.DB, tahunAwal, tahunAkhir, 0)
	if err != nil {
		middleware.Jsoncontenterror(w, http.StatusBadRequest, err)
		return
	}

	middleware.Jsoncontent(w, http.StatusOK, data)
}
