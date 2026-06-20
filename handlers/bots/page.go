package bots

import (
	"fmt"
	"net/http"
	"oj/api"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/internal/middleware/auth"
	"oj/md"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/sashabaranov/go-openai"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func strptr(str string) *string {
	return &str
}

type Resource struct {
	Model *api.Queries
	AI    *openai.Client
}

func (rs Resource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", rs.listPage)
	r.Get("/create", rs.createPage)
	r.Post("/create", rs.postCreate)
	r.Route("/{botID}", func(r chi.Router) {
		r.Use(rs.provideBot)
		r.Get("/", rs.assistantPage)
		r.Get("/chat", rs.chatRedirectPage)
		r.Get("/edit", rs.editPage)
		r.Post("/edit", rs.postEdit)
		r.Get("/chat/{threadID}", rs.chatPage)
		r.Post("/chat/{threadID}/messages", rs.postMessage)
		r.Get("/chat/{threadID}/runstatus/{runID}", rs.getRunStatus)
	})
	return r
}

func (rs Resource) listPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(ctx)

	botRows, err := rs.Model.AllBots(ctx)
	if err != nil {
		render.Error(w, fmt.Errorf("AllBots: %w", err), http.StatusInternalServerError)
		return
	}

	layout.Layout(l, "Bots", botsListPage(botRows)).Render(w)
}

func (rs Resource) createPage(w http.ResponseWriter, r *http.Request) {
	l := layout.FromContext(r.Context())
	layout.Layout(l, "Create Bot", botsCreatePage()).Render(w)
}

func (rs Resource) postCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := auth.FromContext(ctx)

	name := r.FormValue("name")
	if name == "" {
		http.Redirect(w, r, "/bots/create", http.StatusSeeOther)
		return
	}
	instructions := r.FormValue("instructions")

	model := openai.GPT4oMini

	asst, err := rs.AI.CreateAssistant(ctx, openai.AssistantRequest{
		Model:        model,
		Name:         &name,
		Instructions: &instructions,
	})
	if err != nil {
		render.Error(w, fmt.Errorf("CreateAssistant: %w", err), http.StatusInternalServerError)
		return
	}

	bot, err := rs.Model.CreateBot(ctx, api.CreateBotParams{
		OwnerID:     user.ID,
		AssistantID: asst.ID,
		Name:        name,
		Description: instructions,
	})
	if err != nil {
		render.Error(w, fmt.Errorf("CreateBot: %w", err), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/bots/%d", bot.ID), http.StatusSeeOther)
}

func (rs Resource) editPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(ctx)
	bot := botFromContext(ctx)

	layout.Layout(l, "Edit Bot", botsEditPage(bot)).Render(w)
}

func (rs Resource) postEdit(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	bot := botFromContext(ctx)
	user := auth.FromContext(ctx)

	name := r.FormValue("name")
	if name == "" {
		http.Redirect(w, r, r.URL.Path, http.StatusSeeOther)
		return
	}
	instructions := r.FormValue("instructions")
	if instructions == "" {
		http.Redirect(w, r, r.URL.Path, http.StatusSeeOther)
		return
	}

	bot, err := rs.Model.UpdateBotDescription(ctx, api.UpdateBotDescriptionParams{
		OwnerID:     user.ID,
		ID:          bot.ID,
		Name:        name,
		Description: instructions,
	})
	if err != nil {
		render.Error(w, fmt.Errorf("UpdateBotDescription: %w", err), http.StatusInternalServerError)
		return
	}

	_, err = rs.AI.ModifyAssistant(ctx, bot.AssistantID, openai.AssistantRequest{
		Name:         &name,
		Instructions: &instructions,
	})
	if err != nil {
		render.Error(w, fmt.Errorf("ModifyAssistant: %w", err), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/bots/%d", bot.ID), http.StatusSeeOther)
}

func (rs Resource) assistantPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(ctx)
	bot := botFromContext(ctx)

	layout.Layout(l, bot.Name, botsAssistantPage(bot, bot.OwnerID == l.User.ID)).Render(w)
}

func (rs Resource) chatRedirectPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := auth.FromContext(ctx)
	bot := botFromContext(ctx)

	threads, err := rs.Model.AssistantThreads(ctx, api.AssistantThreadsParams{
		UserID:      user.ID,
		AssistantID: bot.AssistantID,
	})
	if err != nil {
		render.Error(w, fmt.Errorf("AssistantThreads: %w", err), http.StatusInternalServerError)
		return
	}

	var threadID string
	if len(threads) > 0 {
		threadID = threads[0].ThreadID
	} else {
		thread, err := rs.AI.CreateThread(ctx, openai.ThreadRequest{})
		if err != nil {
			render.Error(w, fmt.Errorf("CreateThread: %w", err), http.StatusInternalServerError)
			return
		}

		_, err = rs.Model.CreateThread(ctx, api.CreateThreadParams{
			AssistantID: bot.AssistantID,
			ThreadID:    thread.ID,
			UserID:      user.ID,
		})
		if err != nil {
			render.Error(w, fmt.Errorf("CreateThread: %w", err), http.StatusInternalServerError)
			return
		}

		threadID = thread.ID
	}

	http.Redirect(w, r, r.URL.Path+"/"+threadID, http.StatusSeeOther)
}

func (rs Resource) chatPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(ctx)
	bot := botFromContext(ctx)

	userThread, err := rs.Model.UserThreadByID(ctx, api.UserThreadByIDParams{
		UserID:   l.User.ID,
		ThreadID: chi.URLParam(r, "threadID"),
	})
	if err != nil {
		render.Error(w, fmt.Errorf("UserThreadByID: %w", err), http.StatusInternalServerError)
		return
	}

	thread, err := rs.AI.RetrieveThread(ctx, userThread.ThreadID)
	if err != nil {
		render.Error(w, fmt.Errorf("RetrieveThread: %w", err), http.StatusInternalServerError)
		return
	}

	messagesList, err := rs.AI.ListMessage(ctx, thread.ID, nil, strptr("desc"), nil, nil, nil)
	if err != nil {
		render.Error(w, fmt.Errorf("ListMessage: %w", err), http.StatusInternalServerError)
		return
	}

	layout.Layout(l, bot.Name, botsChatPage(bot, l.User, thread, messagesList.Messages)).Render(w)
}

func (rs Resource) postMessage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	bot := botFromContext(ctx)

	content := strings.TrimSpace(r.FormValue("message"))
	if content == "" {
		http.Error(w, "empty message", http.StatusBadRequest)
		return
	}

	threadID := chi.URLParam(r, "threadID")

	_, err := rs.AI.CreateMessage(ctx, threadID, openai.MessageRequest{
		Role:    openai.ChatMessageRoleUser,
		Content: content,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("CreateMessage: %s", err), http.StatusInternalServerError)
		return
	}

	run, err := rs.AI.CreateRun(ctx, threadID, openai.RunRequest{
		AssistantID: bot.AssistantID,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("CreateRun: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Add("HX-Trigger", "messagesUpdated")
	botsThinkingEl(bot, run).Render(w)
}

func (rs Resource) getRunStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	bot := botFromContext(ctx)

	run, err := rs.AI.RetrieveRun(ctx, chi.URLParam(r, "threadID"), chi.URLParam(r, "runID"))
	if err != nil {
		render.Error(w, fmt.Errorf("RetrieveRun: %w", err), http.StatusInternalServerError)
		return
	}

	switch run.Status {
	case openai.RunStatusQueued, openai.RunStatusInProgress:
		botsThinkingEl(bot, run).Render(w)
	default:
		w.Header().Add("HX-Trigger", "messagesUpdated")

		thread, err := rs.AI.RetrieveThread(ctx, run.ThreadID)
		if err != nil {
			render.Error(w, fmt.Errorf("RetrieveThread: %w", err), http.StatusInternalServerError)
			return
		}

		botsInputEl(bot, thread).Render(w)
	}
}

func botAvatarURL(botID int64) string {
	return fmt.Sprintf("https://api.dicebear.com/7.x/bottts/svg?seed=%d", botID)
}

func botMarkdown(text string) g.Node {
	return g.Raw(string(md.RenderString(text)))
}

func botsListPage(botRows []api.AllBotsRow) g.Node {
	return h.Div(h.Class("nes-container ghost"),
		h.Div(h.Style("display:flex; justify-content:space-between"),
			h.Div(h.Class("nes-text is-error"), h.Style("font-size: 2em; padding-bottom: 2em;"), g.Text("Kable Bots")),
			h.Div(h.A(h.Href("/bots/create"), h.Class("nes-btn is-primary"), g.Text("Create Bot"))),
		),
		h.Div(h.Style("display:flex; flex-direction:column; gap:2em"),
			g.Map(botRows, func(row api.AllBotsRow) g.Node {
				return h.Div(h.Style("display:flex; gap:1em; align-items:center"),
					h.Img(h.Height("100px"), h.Width("100px"), h.Src(botAvatarURL(row.Bot.ID))),
					h.Div(h.Style("display:flex; flex-direction:column"),
						h.H2(h.A(h.Href(fmt.Sprintf("/bots/%d", row.Bot.ID)), g.Text(row.Bot.Name))),
						h.Div(h.Style("display:flex; gap:1em"),
							h.Span(h.Style("color:gray"), g.Text("Programmed by ")),
							h.Div(h.Style("display:flex"),
								h.Span(g.Text(row.User.Username)),
								h.Img(h.Width("20px"), h.Height("20px"), h.Src(row.User.Avatar.URL())),
							),
						),
					),
				)
			}),
		),
	)
}

func botsAssistantPage(bot api.Bot, isOwner bool) g.Node {
	return h.Div(h.Class("nes-container ghost"),
		h.Div(h.Style("display: flex; flex-direction: column; gap: 2em"),
			h.Div(h.Style("display: flex; align-items: center; gap: 1em"),
				h.Img(h.Height("100px"), h.Width("100px"), h.Src(botAvatarURL(bot.ID))),
				h.H1(h.Class("nes-text is-error"), g.Text(bot.Name)),
			),
			h.P(g.Text(bot.Description)),
			h.Div(h.Style("display:flex; justify-content:space-between"),
				h.A(h.Href(fmt.Sprintf("/bots/%d/chat", bot.ID)), h.Class("nes-btn is-success"), g.Text("chat with this bot")),
				g.If(isOwner,
					h.A(h.Href(fmt.Sprintf("/bots/%d/edit", bot.ID)), h.Class("nes-btn"), g.Text("edit instructions")),
				),
			),
		),
	)
}

func botsCreatePage() g.Node {
	return h.Div(h.Class("nes-container ghost"),
		h.H1(g.Text("Welcome to the Bot Workshop")),
		h.P(g.Text("Here you can build a new bot.")),
		h.P(g.Text("Let's get started.")),
		h.Hr(),
		h.Form(h.Method("post"),
			h.Div(h.Style("display:flex; flex-direction:column; gap: 2em"),
				h.Div(h.Class("nes-field"),
					h.Label(g.Attr("for", "name"), g.Text("Bot's Name")),
					h.Input(h.ID("name"), h.Class("nes-input"), h.Name("name"), g.Attr("placeholder", "ROBO-123"), g.Attr("required", "")),
				),
				h.Div(h.Class("nes-field"),
					h.Label(g.Attr("for", "instructions"), g.Text("Instructions")),
					h.Textarea(h.ID("instructions"), h.Class("nes-textarea"), h.Name("instructions"), g.Attr("rows", "10"), g.Attr("placeholder", "You are a helpful assistant")),
				),
				h.Button(h.Class("nes-btn is-primary"), g.Text("Create")),
				h.A(h.Href("/bots"), h.Class("nes-btn"), g.Text("Cancel")),
			),
		),
	)
}

func botsEditPage(bot api.Bot) g.Node {
	return h.Div(h.Class("nes-container ghost"),
		h.H1(g.Text("Welcome to the Bot Workshop")),
		h.P(g.Text("Let's edit your bot.")),
		h.Hr(),
		h.Form(h.Method("post"),
			h.Div(h.Style("display:flex; flex-direction:column; gap: 2em"),
				h.Div(h.Class("nes-field"),
					h.Label(g.Attr("for", "name"), g.Text("Bot's Name")),
					h.Input(h.ID("name"), h.Class("nes-input"), h.Name("name"), g.Attr("placeholder", "ROBO-123"), g.Attr("required", ""), h.Value(bot.Name)),
				),
				h.Div(h.Class("nes-field"),
					h.Label(g.Attr("for", "instructions"), g.Text("Instructions")),
					h.Textarea(h.ID("instructions"), h.Class("nes-textarea"), h.Name("instructions"), g.Attr("rows", "10"), g.Attr("placeholder", "You are a helpful assistant"),
						g.Text(bot.Description),
					),
				),
				h.Button(h.Class("nes-btn is-primary"), g.Text("Save")),
				h.A(h.Href(fmt.Sprintf("/bots/%d", bot.ID)), h.Class("nes-btn"), g.Text("Cancel")),
			),
		),
	)
}

func botsChatPage(bot api.Bot, user api.User, thread openai.Thread, messages []openai.Message) g.Node {
	return h.Div(h.Style("height:100%; display: flex; flex-direction: column; padding-bottom: 1em"),
		h.Div(h.Class("nes-container ghost"),
			h.Img(h.Height("64px"), h.Width("64px"), h.Src(botAvatarURL(bot.ID))),
			g.Text(bot.Name),
		),
		botsMessagesEl(bot, user, thread, messages),
		h.Div(h.Style("background: white"),
			botsInputEl(bot, thread),
		),
	)
}

func botsMessagesEl(bot api.Bot, user api.User, thread openai.Thread, messages []openai.Message) g.Node {
	return h.Div(
		h.ID("messages-container"),
		h.Class("nes-container ghost"),
		h.Style("overflow-y: scroll; display: flex; flex-direction: column-reverse; gap:4em"),
		g.Attr("hx-get", thread.ID),
		g.Attr("hx-trigger", "messagesUpdated from:body"),
		g.Attr("hx-select", "#messages-container"),
		g.Attr("hx-swap", "outerHTML"),
		g.Map(messages, func(msg openai.Message) g.Node {
			return botsMessageEl(bot, user, msg)
		}),
	)
}

func botsMessageEl(bot api.Bot, user api.User, msg openai.Message) g.Node {
	avatarSrc := user.Avatar.URL()
	if msg.Role == "assistant" {
		avatarSrc = botAvatarURL(bot.ID)
	}
	return h.Div(h.Style("display:flex; gap:2em"),
		h.Div(h.Img(h.Height("64px"), h.Width("64px"), h.Src(avatarSrc))),
		h.Div(
			g.Map(msg.Content, func(c openai.MessageContent) g.Node {
				if c.Text == nil {
					return g.Text("")
				}
				return botMarkdown(c.Text.Value)
			}),
		),
	)
}

func botsInputEl(bot api.Bot, thread openai.Thread) g.Node {
	return h.Section(h.Style("background: black"),
		h.Div(h.Style("float:right"), g.Text("\u00a0")),
		h.Input(h.Class("nes-input"),
			g.Attr("autofocus", ""),
			h.Type("text"),
			h.Name("message"),
			g.Attr("placeholder", "Message "+bot.Name),
			g.Attr("hx-post", thread.ID+"/messages"),
			g.Attr("hx-swap", "outerHTML"),
			g.Attr("hx-target", "closest section"),
		),
	)
}

func botsThinkingEl(bot api.Bot, run openai.Run) g.Node {
	return h.Section(h.Style("background: black; color: white"),
		h.Div(h.Style("float:right"),
			g.Attr("hx-get", fmt.Sprintf("%s/runstatus/%s", run.ThreadID, run.ID)),
			g.Attr("hx-trigger", "load delay:300ms"),
			g.Attr("hx-target", "closest section"),
			g.Text(string(run.Status)),
		),
		h.Input(h.Class("nes-input"), h.Type("text"), g.Attr("disabled", ""),
			g.Attr("placeholder", "Message "+bot.Name),
		),
	)
}
