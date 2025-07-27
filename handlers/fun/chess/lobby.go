package chess

import (
	"net/http"
	"oj/handlers/layout"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func (s *service) HandleLobby(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(ctx)
	layout.Layout(l, "chess club", lobbyEl()).Render(w)
}

func lobbyEl() g.Node {
	return h.Div(
		h.H1(g.Text("Chess Club")),
		// g.Map(friends, func(f api.User) g.Node {
		// 	return h.Div(
		// 		h.Div(h.Class("nes-container ghost"),
		// 			h.Button(h.Class("nes-btn is-primary"), g.Text("New Game"))),
		// 	)
		// }),
	)
}
