package alfred

import (
	"testing"
)

func TestParseVersion(t *testing.T) {
	{
		v, err := ParseVersion("1.2.3")
		if err != nil {
			t.Error("parse version error")
		}
		if v.Major != 1 {
			t.Error("parse version error")
		}
		if v.Minor != 2 {
			t.Error("parse version error")
		}
		if v.Patch != 3 {
			t.Error("parse version error")
		}
	}

	{
		v, err := ParseVersion("v1.2.3")
		if err != nil {
			t.Error("parse version error")
		}
		if v.Major != 1 {
			t.Error("parse version error")
		}
		if v.Minor != 2 {
			t.Error("parse version error")
		}
		if v.Patch != 3 {
			t.Error("parse version error")
		}
	}

	{
		_, err := ParseVersion("1.2")
		if err == nil {
			t.Error("parse version error")
		}
	}
}
