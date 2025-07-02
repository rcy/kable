package templatehelpers

import (
	"html/template"
	"oj/md"
	"time"

	"github.com/hako/durafmt"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rcy/durfmt"
)

var FuncMap = template.FuncMap{
	"fromNow": func(t time.Time) string {
		return durafmt.Parse(time.Now().Sub(t)).LimitFirstN(1).String()
	},
	"odd": func(i, j int) int {
		return (i + j) % 2
	},
	"html": func(str string) template.HTML {
		return md.RenderString(str)
	},
	"markdown": func(str string) template.HTML {
		return md.Markdown(str)
	},
	"ago": func(t pgtype.Timestamptz) string {
		dur := time.Now().Sub(t.Time)
		if dur < time.Minute {
			return "just now"
		}
		return durfmt.Format(dur) + " ago"
	},
}
