package alfred

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrInvalidVersion = errors.New("Version format invalid")
)

type Version struct {
	Major int
	Minor int
	Patch int
}

// ParseVersion parsing string like 1.3.0, v1.3.0 to Version instance
func ParseVersion(s string) (*Version, error) {
	if len(s) == 0 {
		return nil, ErrInvalidVersion
	}
	ss := strings.Split(s, ".")
	if len(ss) != 3 {
		return nil, ErrInvalidVersion
	}
	if ss[0][0] == 'v' || ss[0][0] == 'V' {
		ss[0] = ss[0][1:]
	}
	major, err := strconv.Atoi(ss[0])
	if err != nil {
		return nil, ErrInvalidVersion
	}
	minor, err := strconv.Atoi(ss[1])
	if err != nil {
		return nil, ErrInvalidVersion
	}
	patch, err := strconv.Atoi(ss[2])
	if err != nil {
		return nil, ErrInvalidVersion
	}
	return &Version{major, minor, patch}, nil
}
