package header

import (
	"html/template"
	"net/http"
	"oj/handlers/layout"
	"oj/handlers/render"
)

var t = template.Must(template.ParseFiles(layout.File))

func Header(w http.ResponseWriter, r *http.Request) {
	l, err := layout.GetData(r)
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}

	render.ExecuteNamed(w, t, "header", struct{ Layout layout.Data }{l})
}