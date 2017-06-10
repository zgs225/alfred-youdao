package alfred

import (
	"testing"
)

func TestLatest(t *testing.T) {
	p := &GithubProvider{
		Owner: "zgs225",
		Repo:  "not-exists-repo",
	}

	_, err := p.Latest()
	if err != ErrRepoNotExists {
		t.Log(err)
		t.Error("Latest error")
	}

	p.Repo = "alfred-youdao"
	obj, err := p.Latest()

	t.Log(obj)
	if obj == nil {
		t.Log(err)
		t.Error("Latest error")
	}
}
