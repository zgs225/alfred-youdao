package alfred

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

const (
	Mods_Shift = "shift"
	Mods_Ctrl  = "ctrl"
	Mods_Cmd   = "cmd"
	Mods_Alt   = "alt"
	Mods_Fn    = "fn"
)

type IconElement struct {
	Type string `json:"type,omitempty"`
	Path string `json:"path,omitempty"`
}

type ModElement struct {
	Valid     bool              `json:"valid,omitempty"`
	Arg       string            `json:"arg,omitempty"`
	Subtitle  string            `json:"subtitle,omitempty"`
	Variables map[string]string `json:"variables,omitempty"`
}

type TextElement struct {
	Copy      string `json:"copy,omitempty"`
	LargeType string `json:"largetype,omitempty"`
}

type ResultElement struct {
	Valid        bool                   `json:"valid,omitempty"`
	Uid          string                 `json:"uid,omitempty"`
	Type         string                 `json:"type,omitempty"`
	Title        string                 `json:"title,omitempty"`
	Subtitle     string                 `json:"subtitle,omitempty"`
	Arg          string                 `json:"arg,omitempty"`
	Autocomplete string                 `json:"autocomplete,omitempty"`
	QuickLookUrl string                 `json:"quicklookurl,omitempty"`
	Mods         map[string]*ModElement `json:"mods,omitempty"`
	Icon         *IconElement           `json:"icon,omitempty"`
	Text         *TextElement           `json:"text,omitempty"`
}

type Result struct {
	Items *[]*ResultElement `json:"items"`
}

func NewResult() *Result {
	i := make([]*ResultElement, 0)
	return &Result{
		Items: &i,
	}
}

// Append new item
func (r *Result) Append(items ...*ResultElement) {
	*r.Items = append(*r.Items, items...)
}

func (r *Result) Count() int {
	return len(*r.Items)
}

func (r *Result) End() {
	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(r); err != nil {
		panic(err)
	}
	fmt.Print(b.String())
	os.Exit(0)
}
