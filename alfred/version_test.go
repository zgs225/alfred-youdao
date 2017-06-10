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

func TestAfter(t *testing.T) {
	v1 := &Version{1, 2, 3}
	v2 := &Version{0, 3, 10}
	if !v1.After(v2) {
		t.Error("Version after error 1")
	}

	v2 = &Version{1, 1, 19}
	if !v1.After(v2) {
		t.Error("Version after error 2")
	}

	v2 = &Version{1, 2, 2}
	if !v1.After(v2) {
		t.Error("Version after error 3")
	}

	v2 = &Version{1, 2, 4}
	if v1.After(v2) {
		t.Error("Version after error 4")
	}
}

func TestBefore(t *testing.T) {
	v1 := &Version{1, 2, 3}
	v2 := &Version{0, 3, 10}
	if v1.Before(v2) {
		t.Error("Version after error 1")
	}

	v2 = &Version{1, 1, 19}
	if v1.Before(v2) {
		t.Error("Version after error 2")
	}

	v2 = &Version{1, 2, 2}
	if v1.Before(v2) {
		t.Error("Version after error 3")
	}

	v2 = &Version{1, 2, 4}
	if !v1.Before(v2) {
		t.Error("Version after error 4")
	}
}
