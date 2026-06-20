package gradients

import (
	"fmt"
	"net/http"
	"oj/api"
	"oj/gradient"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/internal/middleware/auth"

	"github.com/jackc/pgx/v5/pgxpool"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type service struct {
	Conn    *pgxpool.Pool
	Queries *api.Queries
}

func NewService(q *api.Queries, conn *pgxpool.Pool) *service {
	return &service{Queries: q, Conn: conn}
}

func Index(w http.ResponseWriter, r *http.Request) {
	l := layout.FromContext(r.Context())
	layout.Layout(l, "Gradients", gradientsPage(l.BackgroundGradient)).Render(w)
}

func Picker(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		render.Error(w, fmt.Errorf("ParseForm: %w", err), 500)
	}

	grad, err := gradient.NewFromURLValues(r.PostForm)
	if err != nil {
		render.Error(w, fmt.Errorf("gradientFromUrlValues: %w", err), 500)
		return
	}

	pickerEl(grad).Render(w)
}

func (s *service) SetBackground(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := auth.FromContext(ctx)

	err := r.ParseForm()
	if err != nil {
		render.Error(w, fmt.Errorf("ParseForm: %w", err), 500)
	}

	grad, err := gradient.NewFromURLValues(r.PostForm)
	if err != nil {
		render.Error(w, fmt.Errorf("gradientFromUrlValues: %w", err), 500)
		return
	}

	tx, err := s.Conn.Begin(ctx)
	if err != nil {
		render.Error(w, err, http.StatusInternalServerError)
		return
	}
	defer tx.Rollback(ctx)

	qtx := s.Queries.WithTx(tx)
	_, err = qtx.InsertGradient(ctx, api.InsertGradientParams{
		UserID:   user.ID,
		Gradient: grad,
	})
	if err != nil {
		render.Error(w, fmt.Errorf("InsertGradient: %w", err), 500)
		return
	}

	_, err = qtx.UpdateUserGradient(ctx, api.UpdateUserGradientParams{
		UserID:   user.ID,
		Gradient: grad,
	})
	if err != nil {
		render.Error(w, fmt.Errorf("UpdateUserGradient: %w", err), 500)
		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		render.Error(w, fmt.Errorf("Commit: %w", err), 500)
		return
	}

	style := fmt.Sprintf("body { background: %s; }", grad.Render())
	w.Write([]byte(style))
}

func gradientsPage(grad gradient.Gradient) g.Node {
	return g.Group{
		h.StyleEl(g.Raw(`
.hidden { visibility: hidden; }
.rotation { background: transparent; -webkit-appearance: none; -moz-appearance: none; appearance: none; }
.slider { margin: 0px; padding: 0; border: 0; outline: none; background: transparent; -webkit-appearance: none; -moz-appearance: none; appearance: none; width: 100%; height: 0; pointer-events: none; }
.slider::-webkit-slider-runnable-track { cursor: default; height: 1px; outline: 0; -webkit-appearance: none; }
.slider::-moz-range-track { cursor: default; height: 1px; outline: 0; -moz-appearance: none; }
.slider::-webkit-slider-thumb { -webkit-appearance: none; pointer-events: auto; width: 32px; height: 32px; margin-top: -15px; border-radius: 100%; background: radial-gradient(rgba(255,255,255,0) 32%,#000000 46%,#ffffff 54%,#000000 60%); }
.slider::-moz-range-thumb { -moz-appearance: none; pointer-events: auto; width: 32px; height: 32px; border-radius: 100%; background: radial-gradient(rgba(255,255,255,0) 32%,#000000 46%,#ffffff 54%,#000000 60%); }
.slider { margin-top: 16px; }
.track { width: 100%; height: 32px; }
.preview { width: 100%; height: 480px; border: 1px solid black; }
.picker { background: #ddd; padding: 1em; }
`)),
		pickerEl(grad),
	}
}

func pickerEl(grad gradient.Gradient) g.Node {
	rotationClass := "rotation"
	if grad.Type == "radial" {
		rotationClass = "rotation hidden"
	}

	return h.Form(h.Class("xpicker"),
		g.Attr("hx-post", "/fun/gradients/picker"),
		g.Attr("hx-swap", "innerHTML"),
		h.Div(h.Class("nes-container is-dark"),
			h.Button(h.Style("display: none"), h.ID("submit")),

			h.Div(h.Style("display:flex; justify-content:space-between; gap: 1em;align-items:center"),
				h.Div(h.Style("width:100%; margin-bottom: 1em"),
					h.Div(h.Class("nes-select is-dark"),
						g.El("select",
							h.Name("gradientType"),
							h.ID("gradient-type"),
							g.Attr("_", "on change log 'changed' then click() the #submit"),
							gradientOption("linear", "Linear Gradient", grad.Type),
							gradientOption("radial", "Radial Gradient", grad.Type),
							gradientOption("conic", "Conic Gradient", grad.Type),
						),
					),
				),
				h.Label(h.Style("white-space: nowrap"),
					h.Input(h.Class("nes-checkbox is-dark"),
						h.ID("repeatingCheckbox"),
						h.Type("checkbox"),
						h.Name("repeat"),
						g.Attr("_", "on change log 'changed' then click() the #submit"),
						g.If(grad.Repeat, g.Attr("checked", "")),
					),
					h.Span(g.Text("Repeating")),
				),
			),

			h.Div(h.Style("display:flex;justify-items:stretch"),
				g.Map(grad.Colors, func(color string) g.Node {
					return h.Input(h.Name("color"), h.Type("color"), h.Value(color),
						h.Style("width: 100%"),
						g.Attr("_", "on change log 'changed' then click() the #submit"),
					)
				}),
			),

			h.Div(h.Style("position:relative"),
				colorStopSliders(grad),
				h.Div(h.Class("track"), h.Style("background: "+string(grad.RenderBar()))),
			),

			h.Div(h.Style("display:flex"),
				h.Input(h.Class(rotationClass),
					h.Style("flex:1"),
					g.If(grad.Type == "radial", g.Attr("hidden", "")),
					h.Name("degrees"),
					h.Type("range"),
					g.Attr("min", "0"),
					g.Attr("max", "180"),
					h.Value(fmt.Sprint(grad.Degrees)),
					g.Attr("_", "on change click() the #submit"),
				),
			),

			h.Div(h.Class("preview"), h.Style("background: "+string(grad.Render()))),

			h.Div(h.Class("f-row justify-content:space-between"),
				h.Button(h.Class("nes-btn"),
					g.Attr("hx-post", "/fun/gradients/set-background"),
					g.Attr("hx-target", "#user-style"),
					g.Attr("hx-swap", "innerHTML"),
					g.Text("set as background"),
				),
			),
		),
	)
}

func gradientOption(value, label, selectedType string) g.Node {
	return g.El("option",
		g.If(selectedType == value, g.Attr("selected", "")),
		g.Attr("value", value),
		g.Text(label),
	)
}

func colorStopSliders(grad gradient.Gradient) g.Node {
	var nodes []g.Node
	for i, pos := range grad.Positions {
		_ = grad.Colors[i]
		nodes = append(nodes, h.Div(h.Class("group"),
			h.Input(h.Name("percent"),
				h.Style("position:absolute"),
				h.Type("range"),
				g.Attr("min", "0"),
				g.Attr("max", "100"),
				h.Value(fmt.Sprint(pos)),
				h.Class("slider"),
				g.Attr("_", "on click remove .selected from .selected add .selected to the closest .group\n      on change log 'changed' then click() the #submit"),
			),
		))
	}
	return g.Group(nodes)
}
