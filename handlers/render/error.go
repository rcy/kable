package render

import (
	_ "embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var (
	//go:embed "error.gohtml"
	tContent string
	t        = template.Must(template.New("").Parse(tContent))
)

func Error(w http.ResponseWriter, err error, code int) {
	log.Printf("%d: %s", code, err)
	w.WriteHeader(code)
	Execute(w, t, struct {
		Message string
		Code    int
	}{
		Message: err.Error(),
		Code:    code,
	})
}

func NotFound(w http.ResponseWriter) {
	Error(w, fmt.Errorf("Oops, we couldn't find that page!"), http.StatusNotFound)
}
