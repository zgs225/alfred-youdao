package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
)

// Wrap a quote to a string
func wrapInQuote(s string) string {
	return strconv.Quote(s)
}

// options are title, subtitle, sound.
// Sounds: Basso, Blow, Bottle, Frog, Funk, Glass, Hero, Morse, Ping, Pop, Purr,
// Sosumi, Submarine, Tink
func notify(text string, options ...string) {
	var title, subtitle, sound string

	if len(options) > 0 && options[0] != "" {
		title = options[0]
	}
	if len(options) > 1 && options[1] != "" {
		subtitle = options[1]
	}
	if len(options) > 2 && options[2] != "" {
		sound = options[2]
	}

	ascript := fmt.Sprintf(`display notification %s`, wrapInQuote(text))
	if len(title) > 0 {
		ascript = ascript + fmt.Sprintf(` with title %s`, wrapInQuote(title))
	}
	if len(subtitle) > 0 {
		ascript = ascript + fmt.Sprintf(` subtitle %s`, wrapInQuote(subtitle))
	}
	if len(sound) > 0 {
		ascript = ascript + ` sound name ` + wrapInQuote(sound)
	}

	log.Println(ascript)
	cmd := exec.Command("osascript", "-e", ascript)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
