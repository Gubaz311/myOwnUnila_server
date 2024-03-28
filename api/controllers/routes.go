package controllers

import (
	"net/http"

	"github.com/Gubaz311/myOwnUnila_server/api/middleware"
)

func (s *Server) StartApi() {

	HandlerIndex := func(w http.ResponseWriter, r *http.Request) { // Making some dummy response for mapping routes
		w.Write([]byte("hello"))
	}

	//home -> send data info that api is online
	s.Router.HandleFunc("/", middleware.Jsonheader(s.Home)).Methods("GET")

	//feb for get data from feb table
	s.Router.HandleFunc("/feb", HandlerIndex)

	//fh for get data from fh table
	s.Router.HandleFunc("/fh", HandlerIndex)

	//fisip for get data from fisip table
	s.Router.HandleFunc("/fisip", HandlerIndex)

	//fk for get data from fk table
	s.Router.HandleFunc("/fk", HandlerIndex)

	//fkip for get data from fkip table
	s.Router.HandleFunc("/fkip", HandlerIndex)

	//fmipa for get data from fmipa table
	s.Router.HandleFunc("/fmipa", HandlerIndex)

	//fp for get data from fp table
	s.Router.HandleFunc("/fp", HandlerIndex)

	//ft for get data from ft table
	s.Router.HandleFunc("/ft", HandlerIndex)

	//all for get data from all table
	s.Router.HandleFunc("/all", HandlerIndex)
}
