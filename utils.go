package main

import (
	"bytes"
	"fmt"
	"net/url"
	"regexp"
	"strings"
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

var (
	langPattern = regexp.MustCompile(`^([a-zA-Z\-]+)=>([a-zA-Z\-]+)$`)
)

// 解析传入进来的参数
// 返回值是 查询字符串，源语言，目标语言，是否设置语言
func parseArgs(args []string) (q string, from string, to string, lang bool) {
	if len(args) < 2 {
		return
	}
	if v := langPattern.FindAllStringSubmatch(strings.TrimSpace(args[1]), -1); len(v) > 0 {
		lang = true
		from = v[0][1]
		to = v[0][2]
		q = strings.TrimSpace(strings.Join(args[2:], " "))
	} else {
		q = strings.TrimSpace(strings.Join(args[1:], " "))
	}
	return
}
