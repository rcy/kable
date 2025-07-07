package layout

import (
	"fmt"
	"oj/api"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func Layout(data Data, title string, main g.Node) g.Node {
	return h.HTML(
		h.Style("height:100%"),
		h.Head(
			h.TitleEl(g.Text(title)),
			h.Meta(
				h.Name("htmx-config"),
				h.Content(`{"disableInheritance": true}`),
			),
			h.Script(
				h.Src("https://unpkg.com/htmx.org@2.0.4/dist/htmx.min.js"),
			),
			h.Script(
				h.Src("https://unpkg.com/hyperscript.org@0.9.8"),
			),
			h.Link(
				h.Href("https://unpkg.com/nes.css@latest/css/nes.min.css"),
				h.Rel("stylesheet"),
			),
			h.Link(
				h.Href("https://unpkg.com/nes.icons@latest/css/nes-icons.min.css"),
				h.Rel("stylesheet"),
			),
			h.Link(
				h.Href("https://fonts.googleapis.com/css?family=Press+Start+2P"),
				h.Rel("stylesheet"),
			),
			h.Meta(
				h.Charset("utf-8"),
			),
			h.Meta(
				h.Name("viewport"),
				h.Content("width=device-width, initial-scale=1"),
			),
		),
		h.Body(
			h.Style("height: 100%"),
			h.StyleEl(g.Raw(style(string(data.BackgroundGradient.Render())))),
			h.Script(g.Raw(fmt.Sprintf(`
(function(){
  const es = new EventSource("/es/user-%d");
  es.addEventListener("USER_UPDATE", (e) => {
    document.getElementsByTagName('body')[0].dispatchEvent(new Event("USER_UPDATE"))
  });
  window.addEventListener('beforeunload', function(e) {
    es.close();
  });
})();`, data.User.ID)),
			),
			h.Audio(
				h.ID("beeper"),
				h.Src("/assets/chat-alert.mp3"),
			),
			h.Div(
				h.Style("height: 100%;display:flex; flex-direction:column; gap:1em"),
				header(data.UnreadCount, data.User),
				h.Div(
					h.Style("flex:auto; height:100%; overflow: auto;"),
					h.Main(
						h.Style("max-width: 960px; margin: 0 auto;"),
						main,
					),
				),
			),
		),
	)
}

func style(background string) string {
	return fmt.Sprintf(`
html, body, pre, code, kbd, samp {
         font-family: 'Press Start 2P';
}
body {
         background: %s;
}
.ghost {
	background:rgba(255,255,255,1.0);
}
body {
         height: 100vh;
}
header a {
         color: inherit;
}
header a:hover {
         color: cyan;
}
`, background)
}

func header(unreadCount int, user api.User) g.Node {
	return h.HTML(
		h.Head(),
		h.Body(
			h.Header(
				g.Attr("hx-get", "/header"),
				g.Attr("hx-trigger", "USER_UPDATE from:body"),
				g.Attr("hx-swap", "outerHTML"),
				h.Style("background: rgba(0,0,0,.9); color:white;"),
				h.Div(
					h.Style("max-width: 960px; margin: 0 auto;"),
					h.Div(
						h.Style("padding: 1em 0 1em"),
						h.Div(
							h.Style("display:flex;gap:4px;align-items:center;justify-content:space-between"),
							h.A(
								h.Href("/me/humans"),
								g.Text("Humans"),
							),
							h.A(
								h.Href("/bots"),
								g.Text("Robots"),
							),
							h.A(
								h.Href("/fun"),
								g.Text("Activities"),
							),
							h.A(
								h.Href("/me"),
								h.Style("display:flex; align-items: center; gap:6px"),
								g.If(unreadCount > 0,
									h.Span(
										h.Style("background:red; padding: 0px 4px"),
										h.Img(
											h.Style("width: 1em"),
											h.Src("/assets/126567_mail_email_send_contact_icon.svg"),
										),
									),
								),
								h.Span(
									g.Text(user.Username),
								),
								h.Img(
									h.Src(user.AvatarURL),
									h.Height("24px"),
								),
							),
						),
					),
				),
			),
		),
	)
}
