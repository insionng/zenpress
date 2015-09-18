package handler

import (
	//"github.com/insionng/toropress/helper"
	"net/http"
	//"github.com/insionng/toropress/libs"
	//"github.com/insionng/toropress/models"
	"github.com/insionng/vodka"
)

func SigninHandler(self *vodka.Context) error {
	return self.Render(http.StatusOK, "index.html", nil)
}
