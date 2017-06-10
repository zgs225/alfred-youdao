package alfred

import (
	"testing"
)

func TestCanUpdate(t *testing.T) {
	p := &GithubProvider{"zgs225", "alfred-youdao"}
	v := &Version{1, 1, 0}
	u := &Updater{
		V: v,
		P: p,
	}

	if !u.CanUpdate() {
		t.Error("test error")
	}

	u.V = &Version{1, 2, 0}
	if u.CanUpdate() {
		t.Error("version should not can update")
	}
}
