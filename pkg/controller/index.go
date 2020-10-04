package controller

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, _ = w.Write([]byte(fmt.Sprintf(
		"Success for [%v] [%v]", r.Host, r.RequestURI,
	)))
}
