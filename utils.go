package main

import (
	"bytes"
	"fmt"
	"net/url"
)

func toYoudaoDictUrl(q string) string {
	v := fmt.Sprintf("http://dict.youdao.com/w/%s/#keyfrom=dict2.top", q)
	u, err := url.Parse(v)
	if err != nil {
		panic(err)
	}
	return u.String()
}

func joinPhonetic(phonetic, uk, us string) string {
	buf := new(bytes.Buffer)
	empty := true
	if len(phonetic) > 0 {
		buf.WriteString(phonetic)
		empty = false
	}

	if len(uk) > 0 {
		if !empty {
			buf.WriteString("; ")
		}
		buf.WriteString("[UK] ")
		buf.WriteString(uk)
		empty = false
	}

	if len(us) > 0 {
		if !empty {
			buf.WriteString("; ")
		}
		buf.WriteString("[US] ")
		buf.WriteString(us)
	}

	return buf.String()
}
