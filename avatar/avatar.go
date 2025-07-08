package avatar

import (
	"fmt"
)

const (
	AvataaarsStyle = "avataaars"
)

type Avatar struct {
	Style string
	Seed  string
}

func New(seed string) Avatar {
	return Avatar{Style: AvataaarsStyle, Seed: seed}
}

func (a Avatar) URL() string {
	if a.Style == "" {
		a.Style = AvataaarsStyle
	}
	return fmt.Sprintf("https://api.dicebear.com/7.x/%s/svg?seed=%s", a.Style, a.Seed)
}
