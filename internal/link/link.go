package link

import (
	"fmt"
	"oj/api"
	"strings"
)

func Link(resource string, id any, rest ...string) string {
	return strings.TrimSuffix(fmt.Sprintf("/%s/%v", resource, id)+"/"+strings.Join(rest, "/"), "/")
}

func User(id int64, rest ...string) string {
	return Link("u", id, rest...)
}

func ChessMatch(id int64, rest ...string) string {
	return Link("chess", id, rest...)
}

func ConnectFriend(id int64) string {
	return Link("connectkids/friend", id)
}

func Quiz(quiz api.Quiz, rest ...string) string {
	return fmt.Sprintf("/u/%d/quizzes/%d", quiz.UserID, quiz.ID)
}

func ParentKid(id int64, rest ...string) string {
	return Link("parent/kids", id, rest...)
}
