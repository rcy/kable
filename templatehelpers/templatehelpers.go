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
	"fromNow":  FromNow,
	"odd":      Odd,
	"html":     HTML,
	"markdown": Markdown,
	"ago":      Ago,
}

func FromNow(t time.Time) string {
	return durafmt.Parse(time.Now().Sub(t)).LimitFirstN(1).String()
}
func Odd(i, j int) int {
	return (i + j) % 2
}
func HTML(str string) template.HTML {
	return md.RenderString(str)
}
func Markdown(str string) template.HTML {
	return md.Markdown(str)
}
func Ago(t pgtype.Timestamptz) string {
	dur := time.Now().Sub(t.Time)
	if dur < time.Minute {
		return "just now"
	}
	return durfmt.Format(dur) + " ago"
}
