package gradient

import (
	"net/url"
	"testing"
)

func TestNewFromURLValues(t *testing.T) {
	values := url.Values{
		"gradientType": {"linear"},
		"repeat":       {"on"},
		"degrees":      {"90"},
		"color":        {"#ff00ff", "#ffff00", "#00ffff"},
		"percent":      {"0", "50", "100"},
	}

	g, err := NewFromURLValues(values)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if g.Type != "linear" || !g.Repeat || g.Degrees != 90 {
		t.Errorf("Gradient properties do not match expected values: %v", g)
	}

	if len(g.Colors) != 3 || len(g.Positions) != 3 {
		t.Errorf("Expected 3 colors and positions, got %d and %d", len(g.Colors), len(g.Positions))
	}

	got := string(g.Render())
	want := "repeating-linear-gradient(90deg, #ff00ff 0px,#ffff00 50px,#00ffff 100px)"
	if got != want {
		t.Errorf("Expected rendered gradient to be\n'%s', got\n'%s'", want, got)
	}
}
