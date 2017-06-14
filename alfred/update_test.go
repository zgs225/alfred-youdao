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

	u.V = &Version{10, 2, 0}
	if u.CanUpdate() {
		t.Error("version should not can update")
	}
}

func TestUpdate(t *testing.T) {
	p := &GithubProvider{"zgs225", "alfred-youdao"}
	v := &Version{1, 1, 0}
	u := &Updater{
		V: v,
		P: p,
	}
	err := u.Update()
	if err != nil {
		t.Errorf("update version error: %v", err)
	}
}
