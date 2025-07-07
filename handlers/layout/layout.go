package layout

import (
	"context"
	_ "embed"
	"fmt"
	"html/template"
	"oj/api"
	"oj/element/gradient"
	"oj/templatehelpers"

	"github.com/georgysavva/scany/v2/pgxscan"
)

var (
	//go:embed "layout.gohtml"
	layoutContent string
)

func MustParse(templateContent ...string) *template.Template {
	tpl := template.New("layout").Funcs(templatehelpers.FuncMap)

	for i, content := range append([]string{layoutContent}, templateContent...) {
		var err error
		tpl, err = tpl.Parse(content)
		if err != nil {
			fmt.Println(i, content)
			panic(err)
		}
	}
	return tpl
}

type Data struct {
	User               api.User
	BackgroundGradient gradient.Gradient
	UnreadCount        int
}

func (s *service) FromUser(ctx context.Context, user api.User) (Data, error) {
	var unreadCount int
	err := pgxscan.Get(ctx, s.Conn, &unreadCount, `select count(*) from deliveries where recipient_id = $1 and sent_at is null`, user.ID)
	if err != nil {
		return Data{}, fmt.Errorf("pgxscan: %w", err)
	}

	return Data{
		User:               user,
		BackgroundGradient: user.Gradient,
		UnreadCount:        unreadCount,
	}, nil
}
