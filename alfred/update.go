package alfred

import (
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
}
