package controllers

import (
	"github.com/Gubaz311/myOwnUnila_server/api/middleware"
)

func (s *Server) StartApi() {

	//home -> send data info that api is online
	s.Router.HandleFunc("/", middleware.Jsonheader(s.Home)).Methods("GET")

	// Get data from specific table when tgl_keluar is.True
	s.Router.HandleFunc("/lulus/{tahunAwal}/{tahunAkhir}/{idFakultas}", middleware.Jsonheader(s.GetDataLulus)).Methods("GET")
	// Get data from all table when tgl_keluar is.True
	s.Router.HandleFunc("/lulus/{tahunAwal}/{tahunAkhir}", middleware.Jsonheader(s.GetAllLulus)).Methods("GET")

	// Get data from specific table
	s.Router.HandleFunc("/{tahunAwal}/{tahunAkhir}/{idFakultas}", middleware.Jsonheader(s.GetData)).Methods("GET")
	// Get data from all table
	s.Router.HandleFunc("/{tahunAwal}/{tahunAkhir}", middleware.Jsonheader(s.GetAll)).Methods("GET")

}
