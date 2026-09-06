package ui

import (
	"os"
	"regexp"
	"testing"
)

func TestCharmStackMajorVersions(t *testing.T) {
	gomod, err := os.ReadFile("../../go.mod")
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct{ dep, floor string }{
		{"bubbles", "v1.0"},
		{"bubbletea", "v1.3"},
		{"lipgloss", "v1.1"},
	} {
		re := regexp.MustCompile(`github\.com/charmbracelet/` + tc.dep + ` (v[0-9.]+)`)
		m := re.FindSubmatch(gomod)
		if m == nil {
			t.Errorf("%s missing from go.mod", tc.dep)
			continue
		}
		if string(m[1]) < tc.floor {
			t.Errorf("charmbracelet/%s version = %s, below floor %s", tc.dep, m[1], tc.floor)
		}
	}
}
