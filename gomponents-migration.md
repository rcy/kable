# Gomponents Migration Plan

Migrate all remaining Go template (`.gohtml`) code to Gomponents.

## Current State

- **24 `.gohtml` files** remain (~1,171 lines total)
- **Gomponents layout** (`gom.layout.go`) is already complete and production-ready — newer handlers use it
- **Both patterns coexist**: old handlers call `render.Execute(w, tmpl, data)`, new ones call `layout.Layout(...).Render(w)`
- `handlers/layout/context.go` and `FromContext()` already support both

The migration is purely mechanical: convert template HTML into Gomponents Go functions, then delete the `.gohtml` file and its `MustParse`/`Execute` wiring.

---

## Phase 1 — Low complexity (admin + small pages) ✅ DONE

Small, mostly static list/card templates with minimal logic.

| File | Lines | Notes |
|---|---|---|
| `handlers/admin/page.gohtml` | ~10 | User list |
| `handlers/postoffice/page.gohtml` | ~20 | Postoffice listing |
| `handlers/me/card.gohtml` | ~30 | User card |
| `handlers/render/error.gohtml` | ~50 | Error page |
| `internal/resources/parent/parent.gohtml` | ~80 | Parent profile |
| `internal/resources/stickers/page.gohtml` | ~50 | Sticker resource |

For each: create a `page.go` alongside the `.gohtml`, translate HTML to Gomponents, update the handler to call `layout.Layout(l, title, pageEl).Render(w)`, delete the template file and remove the `//go:embed` / `MustParse` wiring.

---

## Phase 2 — Medium complexity (welcome/onboarding + bots) ✅ DONE

Six welcome templates with conditional flows and form handling, plus bot chat list/edit/create.

| File | Notes |
|---|---|
| `handlers/welcome/welcome.gohtml` + siblings (5 files) | Step-based auth flow, shared `welcome/layout.gohtml` |
| `handlers/bots/chat.gohtml` | Bot chat interface |
| `handlers/admin/quizzes/*.gohtml` | Admin quiz CRUD |
| `handlers/fun/notebook/*.gohtml` | Notebook pages |
| `handlers/deliveries/*.gohtml` | Delivery list |

Migrate `welcome/layout.gohtml` to a Gomponents wrapper function first, then migrate each step page. The `welcome_route.go` currently embeds and parses templates — replace with direct `layout.Layout(...)` calls.

---

## Phase 3 — High complexity (chat + gradients)

Nested templates, SSE wiring, and significant inline CSS/JS.

| File | Lines | Complexity drivers |
|---|---|---|
| `handlers/chat/page.gohtml` | ~200 | Nested templates, SSE connection, per-message conditionals |
| `handlers/fun/gradients/page.gohtml` | ~220 | Inline CSS, range sliders, live preview JS |

- **Chat**: Extract sub-templates (message bubble, input box, SSE setup) into individual Gomponents functions. SSE and HTMX attributes translate cleanly to `g.Attr("hx-ext", "sse")` etc.
- **Gradients**: Inline CSS strings become `g.Raw(...)` or `h.StyleEl(g.Raw(...))`. Range sliders are straightforward `h.Input(h.Type("range"), ...)`.

---

## Cleanup (after all phases)

1. Delete `handlers/layout/layout.gohtml` and `handlers/layout/layout.go` (template-based MustParse infrastructure)
2. Delete `handlers/render/render.go` and `handlers/render/error.gohtml` if nothing remains using `render.Execute`
3. Delete `templatehelpers/templatehelpers.go` (FuncMap helpers — Gomponents calls Go functions directly)
4. Check `md/md.go` — still needed if markdown bodies use `template.HTML`, but the FuncMap use can be removed

---

## Key Patterns

```go
// Conditionals
g.If(condition, h.Span(g.Text("unread")))

// Iteration
g.Map(messages, func(m api.Message) g.Node {
    return h.Div(h.Class("message"), g.Text(m.Body))
})

// Custom/HTMX attributes
g.Attr("hx-ext", "sse"), g.Attr("sse-connect", "/events")

// Raw CSS/HTML fragments
g.Raw(`<style>...</style>`)

// Render to writer
layout.Layout(l, "Page Title", mainNode).Render(w)
```
