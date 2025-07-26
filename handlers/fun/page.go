package fun

import (
	_ "embed"
	"net/http"
	"oj/handlers/layout"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func Page(w http.ResponseWriter, r *http.Request) {
	l := layout.FromContext(r.Context())

	layout.Layout(l, "Activities",
		h.Div(
			h.H1(
				g.Text("Here's some fun things to do!"),
			),
			h.Div(
				h.Style("display:flex; flex-direction:column; gap: 1em"),
				h.A(
					h.Href("/stickers"),
					h.Div(
						h.Class("nes-container ghost"),
						h.Style("display:flex; gap:1em"),
						h.I(
							h.Class("nes-kirby"),
						),
						h.H1(
							g.Text("Sticker Book"),
						),
					),
				),
				h.A(
					h.Href("fun/gradients"),
					h.Div(
						h.Class("nes-container ghost"),
						h.Style("display:flex; gap:1em"),
						h.Div(
							h.Style("height:96px; width:96px; background: "+string(l.BackgroundGradient.Render())),
						),
						h.H1(
							g.Text("Gradients"),
						),
					),
				),
				h.A(
					h.Href("fun/notes"),
					h.Div(
						h.Class("nes-container ghost"),
						h.Style("display:flex; gap:1em"),
						h.Img(
							h.Height("100px"),
							h.Width("100px"),
							h.Src("https://api.dicebear.com/7.x/icons/svg?seed=Jessica"),
						),
						h.H1(
							g.Text("Notebook"),
						),
					),
				),
			))).Render(w)
}
