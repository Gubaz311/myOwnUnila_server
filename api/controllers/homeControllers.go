package controllers

import (
	"net/http"

	"github.com/Gubaz311/myOwnUnila_server/api/middleware"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	middleware.Jsoncontent(w, http.StatusOK, "You are Connected to unila api, get query for more")
}
