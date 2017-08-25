package handlers

import (
	"net/http"

	"github.com/TheoRev/apirest_go/utils"
)

func Index(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "application/index", nil)
}
