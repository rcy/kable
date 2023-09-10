package friends

import (
	"html/template"
	"net/http"
	"oj/db"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/models/gradients"
	"oj/models/users"
	"oj/templatehelpers"
	"sort"
)

var MyPageTemplate = template.Must(template.New("layout.gohtml").Funcs(templatehelpers.FuncMap).ParseFiles(layout.File, "handlers/me/friends/page.gohtml", "handlers/me/card.gohtml"))

type Unread struct {
	SenderID int64 `db:"sender_id"`
	Count    int
}

type Friend struct {
	users.User
	Role        string
	UnreadCount int
	GradientCSS template.CSS
}

func Page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l, err := layout.FromContext(ctx)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	friends, err := getFriends(l.User.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	d := struct {
		Layout  layout.Data
		User    users.User
		Friends []*Friend
	}{
		Layout:  l,
		User:    l.User,
		Friends: friends,
	}

	render.Execute(w, MyPageTemplate, d)
}

func getFriends(userID int64) ([]*Friend, error) {
	var friends []*Friend
	err := db.DB.Select(&friends, `
select users.*, fi.b_role role
from users
join friends fi on fi.b_id = users.id and fi.a_id = $1
join friends fo on fo.a_id = users.id and fo.b_id = $1
where fi.b_role = 'friend'
`, userID)
	if err != nil {
		return nil, err
	}

	err = addGradients(friends)
	if err != nil {
		return nil, err
	}

	unreads, err := getUnreads(userID)
	if err != nil {
		return nil, err
	}

	addUnreadCounts(friends, unreads)

	return friends, nil
}

func getUnreads(userID int64) ([]Unread, error) {
	var unreads []Unread

	err := db.DB.Select(&unreads, `
	  select sender_id, count(*) count
          from deliveries
          where recipient_id = ? and sent_at is null
          group by sender_id`, userID)
	if err != nil {
		return nil, err
	}
	return unreads, nil
}

func addUnreadCounts(friends []*Friend, unreads []Unread) {
	for _, friend := range friends {
		for _, unread := range unreads {
			if unread.SenderID == friend.ID {
				friend.UnreadCount = unread.Count
				break
			}
		}
	}

	sort.Slice(friends, func(i, j int) bool {
		return friends[j].UnreadCount < friends[i].UnreadCount
	})
}

func addGradients(friends []*Friend) error {
	for _, friend := range friends {
		bg, err := gradients.UserBackground(friend.User.ID)
		if err != nil {
			return err
		}

		friend.GradientCSS = bg.Render()
	}
	return nil
}