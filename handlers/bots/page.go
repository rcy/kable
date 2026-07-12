package bots

import (
	"fmt"
	"net/http"
	"oj/api"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/internal/ai"
	"oj/internal/middleware/auth"
	"oj/md"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgtype"
	openai "github.com/sashabaranov/go-openai"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type Resource struct {
	Model *api.Queries
	AI    *ai.AI
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

	bot, err := rs.Model.CreateBot(ctx, api.CreateBotParams{
		OwnerID:     user.ID,
		Name:        name,
		Description: instructions,
		Model:       rs.AI.Model,
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
	model := r.FormValue("model")
	if model == "" {
		model = rs.AI.Model
	}

	bot, err := rs.Model.UpdateBotDescription(ctx, api.UpdateBotDescriptionParams{
		OwnerID:     user.ID,
		ID:          bot.ID,
		Name:        name,
		Description: instructions,
		Model:       model,
	})
	if err != nil {
		render.Error(w, fmt.Errorf("UpdateBotDescription: %w", err), http.StatusInternalServerError)
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

	threads, err := rs.Model.BotThreads(ctx, api.BotThreadsParams{
		UserID: user.ID,
		BotID:  pgtype.Int8{Int64: bot.ID, Valid: true},
	})
	if err != nil {
		render.Error(w, fmt.Errorf("BotThreads: %w", err), http.StatusInternalServerError)
		return
	}

	var threadID string
	if len(threads) > 0 {
		threadID = threads[0].ThreadID
	} else {
		threadID = fmt.Sprintf("thread_%d_%d", bot.ID, user.ID)
		_, err = rs.Model.CreateThread(ctx, api.CreateThreadParams{
			BotID:    pgtype.Int8{Int64: bot.ID, Valid: true},
			ThreadID: threadID,
			UserID:   user.ID,
		})
		if err != nil {
			render.Error(w, fmt.Errorf("CreateThread: %w", err), http.StatusInternalServerError)
			return
		}
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

	messages, err := rs.Model.ThreadMessages(ctx, pgtype.Int8{Int64: userThread.ID, Valid: true})
	if err != nil {
		render.Error(w, fmt.Errorf("ThreadMessages: %w", err), http.StatusInternalServerError)
		return
	}

	layout.Layout(l, bot.Name, botsChatPage(bot, l.User, userThread, messages)).Render(w)
}

func (rs Resource) postMessage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	bot := botFromContext(ctx)
	l := layout.FromContext(ctx)

	content := strings.TrimSpace(r.FormValue("message"))
	if content == "" {
		http.Error(w, "empty message", http.StatusBadRequest)
		return
	}

	threadID := chi.URLParam(r, "threadID")

	userThread, err := rs.Model.UserThreadByID(ctx, api.UserThreadByIDParams{
		UserID:   l.User.ID,
		ThreadID: threadID,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("UserThreadByID: %s", err), http.StatusInternalServerError)
		return
	}

	_, err = rs.Model.CreateBotMessage(ctx, api.CreateBotMessageParams{
		ThreadID: pgtype.Int8{Int64: userThread.ID, Valid: true},
		Role:     "user",
		Content:  content,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("CreateBotMessage: %s", err), http.StatusInternalServerError)
		return
	}

	messages, err := rs.Model.ThreadMessages(ctx, pgtype.Int8{Int64: userThread.ID, Valid: true})
	if err != nil {
		http.Error(w, fmt.Sprintf("ThreadMessages: %s", err), http.StatusInternalServerError)
		return
	}

	chatMessages := []openai.ChatCompletionMessage{
		{Role: openai.ChatMessageRoleSystem, Content: bot.Description},
	}
	for _, msg := range messages {
		chatMessages = append(chatMessages, openai.ChatCompletionMessage{
			Role:    msg.Role,
			Content: msg.Content,
		})
	}

	model := bot.Model
	if model == "" {
		model = rs.AI.Model
	}

	resp, err := rs.AI.Client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:    model,
		Messages: chatMessages,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("CreateChatCompletion: %s", err), http.StatusInternalServerError)
		return
	}

	if len(resp.Choices) > 0 {
		assistantContent := resp.Choices[0].Message.Content
		_, err = rs.Model.CreateBotMessage(ctx, api.CreateBotMessageParams{
			ThreadID: pgtype.Int8{Int64: userThread.ID, Valid: true},
			Role:     "assistant",
			Content:  assistantContent,
		})
		if err != nil {
			http.Error(w, fmt.Sprintf("CreateBotMessage: %s", err), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Add("HX-Trigger", "messagesUpdated")
	botsInputEl(bot, userThread).Render(w)
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
				h.Div(h.Class("nes-field"),
					h.Label(g.Attr("for", "model"), g.Text("Model")),
					h.Input(h.ID("model"), h.Class("nes-input"), h.Name("model"), g.Attr("placeholder", "deepseek-chat"), h.Value(bot.Model)),
				),
				h.Button(h.Class("nes-btn is-primary"), g.Text("Save")),
				h.A(h.Href(fmt.Sprintf("/bots/%d", bot.ID)), h.Class("nes-btn"), g.Text("Cancel")),
			),
		),
	)
}

func botsChatPage(bot api.Bot, user api.User, thread api.Thread, messages []api.BotMessage) g.Node {
	return h.Div(h.Style("height:100%; display: flex; flex-direction: column; padding-bottom: 1em"),
		h.Div(h.Class("nes-container ghost"),
			h.Img(h.Height("64px"), h.Width("64px"), h.Src(botAvatarURL(bot.ID))),
			g.Text(bot.Name),
		),
		botsMessagesEl(bot, user, thread, messages),
		h.Div(h.Style("background: white"),
			botsInputEl(bot, thread),
		),
		h.Script(g.Raw("document.getElementById('messages-container').scrollTop = document.getElementById('messages-container').scrollHeight")),
	)
}

func botsMessagesEl(bot api.Bot, user api.User, thread api.Thread, messages []api.BotMessage) g.Node {
	return h.Div(
		h.ID("messages-container"),
		h.Class("nes-container ghost"),
		h.Style("flex: 1; min-height: 0; overflow-y: auto; display: flex; flex-direction: column; gap:4em"),
		g.Attr("hx-on::after-settle", "this.scrollTop = this.scrollHeight"),
		g.Attr("hx-get", thread.ThreadID),
		g.Attr("hx-trigger", "messagesUpdated from:body"),
		g.Attr("hx-select", "#messages-container"),
		g.Attr("hx-swap", "outerHTML"),
		g.Map(messages, func(msg api.BotMessage) g.Node {
			return botsMessageEl(bot, user, msg)
		}),
	)
}

func botsMessageEl(bot api.Bot, user api.User, msg api.BotMessage) g.Node {
	avatarSrc := user.Avatar.URL()
	if msg.Role == "assistant" {
		avatarSrc = botAvatarURL(bot.ID)
	}
	return h.Div(h.Style("display:flex; gap:2em"),
		h.Div(h.Img(h.Height("64px"), h.Width("64px"), h.Src(avatarSrc))),
		h.Div(botMarkdown(msg.Content)),
	)
}

func botsInputEl(bot api.Bot, thread api.Thread) g.Node {
	return h.Section(h.Style("background: black"),
		h.Div(h.Style("float:right"), g.Text("\u00a0")),
		h.Input(h.Class("nes-input"),
			g.Attr("autofocus", ""),
			h.Type("text"),
			h.Name("message"),
			g.Attr("placeholder", "Message "+bot.Name),
			g.Attr("hx-post", thread.ThreadID+"/messages"),
			g.Attr("hx-swap", "outerHTML"),
			g.Attr("hx-target", "closest section"),
		),
	)
}
