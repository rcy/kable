package completed

import (
	_ "embed"
	"fmt"
	"net/http"
	"oj/api"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/handlers/u/quizzes/view"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type service struct {
	Queries *api.Queries
}

func NewService(q *api.Queries) *service {
	return &service{Queries: q}
}

func (s *service) Page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(ctx)

	attemptID, _ := strconv.Atoi(chi.URLParam(r, "attemptID"))
	attempt, err := s.Queries.GetAttemptByID(ctx, int64(attemptID))
	if err != nil {
		if err == pgx.ErrNoRows {
			render.Error(w, fmt.Errorf("GetAttemptByID: %w", err), http.StatusNotFound)
			return
		}
		render.Error(w, fmt.Errorf("GetAttemptByID: %w", err), http.StatusNotFound)
		return
	}

	quiz, err := s.Queries.Quiz(ctx, attempt.QuizID)
	if err != nil {
		if err == pgx.ErrNoRows {
			render.Error(w, fmt.Errorf("Quiz: %w", err), http.StatusNotFound)
			return
		}
		render.Error(w, fmt.Errorf("Quiz: %w", err), http.StatusInternalServerError)
		return
	}

	author, err := s.Queries.UserByID(ctx, quiz.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// questionCount, err := s.Queries.QuestionCount(ctx, quiz.ID)
	// if err != nil {
	// 	render.Error(w, fmt.Errorf("QuestionCount: %w", err), http.StatusInternalServerError)
	// 	return
	// }

	responses, err := s.Queries.Responses(ctx, attempt.ID)
	if err != nil {
		render.Error(w, fmt.Errorf("Responses: %w", err), http.StatusInternalServerError)
		return
	}

	layout.Layout(l,
		quiz.Name,
		h.Div(
			h.Class("page nes-container ghost"),
			h.Style("display: flex; flex-direction: column; gap: 2em"),
			h.Div(
				view.QuizTitleEl(author, quiz),
				h.Progress(
					h.Style("width:100%"),
					h.Max("1"),
					h.Value("1"),
				),
				h.H1(g.Text("Completed!")),
			),
			h.Div(
				h.Style("display:flex; flex-direction:column; gap: 1em;"),
				g.Map(responses, func(response api.ResponsesRow) g.Node {
					return h.Div(
						h.Div(
							g.Text("Question: "+response.QuestionText),
						),
						g.If(response.IsCorrect,
							h.Div(
								g.Text("Your answer: "+response.Text),
								h.Span(
									h.Class("nes-text is-success"),
									g.Text(" Correct!"),
								),
							)),
						g.If(!response.IsCorrect,
							h.Div(
								h.Div(
									g.Text("Your answer: "+response.Text),
									h.Span(
										h.Class("nes-text is-error"),
										g.Text(" Incorrect"),
									),
								),
								h.Div(
									g.Text("Correct answer: "+response.QuestionAnswer),
								),
							),
						),
					)
				}),
			),
			h.Div(
				h.Button(
					g.Attr("hx-post", fmt.Sprintf("/u/%d/quizzes/%d/view/attempt", quiz.UserID, quiz.ID)),
					h.Class("nes-btn is-success"),
					g.Text("Try Again"),
				),
				h.A(
					h.Href(fmt.Sprintf("/u/%d#quizzes", quiz.UserID)),
					h.Class("nes-btn"),
					g.Text("Back"),
				),
			),
		)).Render(w)
}
