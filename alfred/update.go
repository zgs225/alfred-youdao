package alfred

import (
	"log"
	"net/url"
)

type updateEntity struct {
	Name        string
	Description string
	V           *Version
	Url         *url.URL
}

type UpdateProvider interface {
	Latest() (*updateEntity, error)
}

type Updater struct {
	V *Version
	P UpdateProvider
	e *updateEntity
}

func (u *Updater) CanUpdate() bool {
	if u.e == nil {
		v, err := u.P.Latest()
		if err != nil {
			log.Println(err)
			return false
		}
		u.e = v
	}

	return u.e.V.After(u.V)
}
