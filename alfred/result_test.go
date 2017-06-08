package alfred

import (
	"testing"
)

func TestAppend(t *testing.T) {
	r := NewResult()
	i := 0

	for j := 0; j < 10; j++ {
		r.Append(&ResultElement{
			Valid: true,
			Title: "Hello world",
		})
		i++

		if i != r.Count() {
			t.Error("Append error")
		}
	}
}
