package alfred

import (
	"errors"
	"log"
	"net/url"
	"os"
	"os/exec"
	"path"
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

func (u *Updater) Update() error {
	if u.CanUpdate() {
		file := u.filename()
		Download(file, u.e.Url.String())
		cmd := exec.Command("open", file)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		return cmd.Run()
	}
	return errors.New("Can't update")
}

func (u *Updater) filename() string {
	if u.e == nil {
		return ""
	}
	tmpdir := os.TempDir()
	base := path.Base(u.e.Url.Path)
	return path.Join(tmpdir, base)
}
