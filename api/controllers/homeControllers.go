package controllers

import (
	"net/http"

	"github.com/Gubaz311/myOwnUnila_server/api/utils"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	utils.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
