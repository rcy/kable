package avatar

import (
	"fmt"
)

const (
	AvataaarsStyle = "avataaars"
	RingsStyle     = "rings"
)

type Avatar struct {
	Style string
	Seed  string
}

func New(seed string, style string) Avatar {
	return Avatar{Style: style, Seed: seed}
}

func (a Avatar) URL() string {
	if a.Style == "" {
		a.Style = AvataaarsStyle
	}
	return fmt.Sprintf("https://api.dicebear.com/9.x/%s/svg?seed=%s", a.Style, a.Seed)
}
