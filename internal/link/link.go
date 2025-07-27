package link

import (
	"fmt"
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
