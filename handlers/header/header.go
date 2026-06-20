package header

import (
	"net/http"
	"oj/handlers/layout"
)

func Header(w http.ResponseWriter, r *http.Request) {
	l := layout.FromContext(r.Context())
	layout.HeaderEl(l.UnreadCount, l.User).Render(w)
}
