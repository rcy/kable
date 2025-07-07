package avatar

import (
	"fmt"
)

type Avatar struct {
	Style string
	Seed  string
}

func (a Avatar) URL() string {
	if a.Style == "" {
		a.Style = "avataaars"
	}
	if a.Seed == "" {
		a.Seed = "2"
	}
	return fmt.Sprintf("https://api.dicebear.com/7.x/%s/svg?seed=%s", a.Style, a.Seed)
}
