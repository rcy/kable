package attempt

import (
	_ "embed"
	"fmt"
	"net/http"
	"oj/api"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/handlers/u/quizzes/view"
	"oj/internal/middleware/auth"
	"strconv"
	"strings"

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
			render.Error(w, fmt.Errorf("GetAttemptID: %w", err), http.StatusNotFound)
			return
		}
		render.Error(w, fmt.Errorf("GetAttemptID: %w", err), http.StatusInternalServerError)
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

	questionCount, err := s.Queries.QuestionCount(ctx, quiz.ID)
	if err != nil {
		render.Error(w, fmt.Errorf("QuestionCount: %w", err), http.StatusInternalServerError)
		return
	}

	responseCount, err := s.Queries.ResponseCount(ctx, attempt.ID)
	if err != nil {
		render.Error(w, fmt.Errorf("ResponseCount: %w", err), http.StatusInternalServerError)
		return
	}

	question, err := s.Queries.AttemptNextQuestion(ctx, api.AttemptNextQuestionParams{
		QuizID:    attempt.QuizID,
		AttemptID: attempt.ID,
	})
	if err != nil {
		if err == pgx.ErrNoRows {
			url := fmt.Sprintf("/u/%d/quizzes/%d/attempts/%d/done", quiz.UserID, quiz.ID, attempt.ID)
			if r.Header.Get("HX-Request") == "true" {
				w.Header().Add("HX-Redirect", url)
				w.WriteHeader(http.StatusSeeOther)
				return
			}
			http.Redirect(w, r, url, http.StatusSeeOther)
			return
		}
		render.Error(w, fmt.Errorf("AttemptNextQuestion: %w", err), http.StatusInternalServerError)
		return
	}

	layout.Layout(l,
		quiz.Name,
		h.Div(
			h.Class("page nes-container ghost"),
			view.QuizTitleEl(author, quiz),
			//g.Text(fmt.Sprintf("%d = %d/%d", attempt.ID, responseCount, questionCount)),
			h.Progress(
				//h.Class("nes-progress"),
				h.Style("width:100%"),
				h.Max(fmt.Sprint(questionCount)),
				h.Value(fmt.Sprint(responseCount)),
			),
			h.Div(
				h.Style("margin-top: 2em"),
				h.Div(g.Text(question.Text)),
				h.Form(
					g.Attr("hx-post", fmt.Sprintf("/u/%d/quizzes/%d/attempts/%d/question/%d/response",
						quiz.UserID, quiz.ID, attempt.ID, question.ID)),
					g.Attr("hx-select", ".page"),
					g.Attr("hx-target", ".page"),
					g.Attr("hx-swap", "outerHTML"),
					h.Input(
						h.Class("nes-input"),
						h.AutoFocus(),
						h.Type("text"),
						h.Name("response"),
					),
					h.Button(
						h.Class("nes-btn"),
						g.Text("submit"),
					),
				),
			),
		)).Render(w)
}

func (s *service) PostResponse(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user := auth.FromContext(ctx)

	attemptID, err := strconv.Atoi(chi.URLParam(r, "attemptID"))
	if err != nil {
		render.Error(w, fmt.Errorf("Atoi: %w", err), http.StatusInternalServerError)
		return
	}

	attempt, err := s.Queries.GetAttemptByID(ctx, int64(attemptID))
	if err != nil {
		render.Error(w, fmt.Errorf("GetAttemptByID: %w", err), http.StatusInternalServerError)
		return
	}

	questionID, err := strconv.Atoi(chi.URLParam(r, "questionID"))
	if err != nil {
		render.Error(w, fmt.Errorf("Atoi: %w", err), http.StatusInternalServerError)
		return
	}

	text := strings.TrimSpace(r.FormValue("response"))

	if text != "" {
		_, err := s.Queries.CreateResponse(ctx, api.CreateResponseParams{
			QuizID:     attempt.QuizID,
			UserID:     user.ID,
			AttemptID:  int64(attemptID),
			QuestionID: int64(questionID),
			Text:       text,
		})

		if err != nil {
			render.Error(w, fmt.Errorf("CreateResponse: %w", err), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Add("HX-Location", fmt.Sprintf("/u/%d/quizzes/%d/attempts/%d", user.ID, attempt.QuizID, attemptID))
	w.WriteHeader(201)
}

func (s *service) CompletedPage(w http.ResponseWriter, r *http.Request) {
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
