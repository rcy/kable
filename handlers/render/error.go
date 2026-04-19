package render

import (
	"fmt"
	"log"
	"net/http"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func Error(w http.ResponseWriter, err error, code int) {
	log.Printf("render.Error: %d: %s", code, err)
	w.WriteHeader(code)
	errorPage(code, err.Error()).Render(w)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Println("NotFound:", r.Header)
	Error(w, fmt.Errorf("Oops, we couldn't find that page!"), http.StatusNotFound)
}

func errorPage(code int, message string) g.Node {
	var containerClass, textClass, sprite, btnClass, btnText string
	if code == 404 {
		containerClass = "nes-container"
		textClass = "nes-text"
		sprite = "nes-bulbasaur"
		btnClass = "nes-btn is-primary"
		btnText = "go back to a real page"
	} else {
		containerClass = "nes-container is-dark"
		textClass = "nes-text is-error"
		sprite = "nes-charmander"
		btnClass = "nes-btn is-error"
		btnText = "go back to safety"
	}

	return h.HTML(
		h.Head(
			h.Meta(h.Charset("utf-8")),
			h.Meta(h.Name("viewport"), h.Content("width=device-width, initial-scale=1")),
			h.Link(h.Href("https://unpkg.com/nes.css@latest/css/nes.min.css"), h.Rel("stylesheet")),
			h.Link(h.Href("https://fonts.googleapis.com/css?family=Press+Start+2P"), h.Rel("stylesheet")),
			h.StyleEl(g.Raw(`html, body, pre, code, kbd, samp { font-family: 'Press Start 2P'; }`)),
		),
		h.Body(h.Style("max-width: 960px; margin: 0 auto"),
			h.Div(h.Style("display:flex; flex-direction: column; gap:1em"),
				h.Div(h.Class(containerClass),
					h.Div(h.Style("display:flex; gap: 1em"),
						h.Div(h.Style("flex:1"),
							h.H2(h.Class(textClass), g.Text(fmt.Sprintf("ERROR: %d", code))),
							h.P(g.Text(message)),
						),
						h.I(h.Class(sprite)),
					),
				),
				h.A(h.Class(btnClass), h.Href("/"), g.Text(btnText)),
			),
		),
	)
}
