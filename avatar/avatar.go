package avatar

import (
	"fmt"
)

const (
	AdventurerStyle = "adventurer"
	AvataaarsStyle  = "avataaars"
	RingsStyle      = "rings"
	DefaultStyle    = AdventurerStyle
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
		a.Style = DefaultStyle
	}
	return fmt.Sprintf("https://api.dicebear.com/7.x/%s/svg?seed=%s", a.Style, a.Seed)
}
