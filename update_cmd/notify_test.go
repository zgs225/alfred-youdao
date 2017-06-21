package main

import (
	"testing"
)

func TestNotify(t *testing.T) {
	notify("Hello world", "Title", "with subtitle", "default")
}
