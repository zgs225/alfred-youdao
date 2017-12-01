package main

import (
	"bytes"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/zgs225/alfred-youdao/alfred"
	"github.com/zgs225/youdao"
)

func toYoudaoDictUrl(q string) string {
	v := fmt.Sprintf("http://dict.youdao.com/search?q=%s&keyfrom=%s", url.QueryEscape(q), url.QueryEscape("fanyi.smartResult"))
	return v
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

func copyModElementMap(m map[string]*alfred.ModElement) map[string]*alfred.ModElement {
	m2 := make(map[string]*alfred.ModElement)
	for k, v := range m {
		m2[k] = v
	}
	return m2
}

func wordsToSayCmdOption(q string, r *youdao.Result) string {
	ls := strings.Split(r.L, "2")
	if len(ls) >= 2 {
		l := languageToSayLanguage(ls[1])
		return fmt.Sprintf("-l %s %s", l, q)
	}
	return q
}

func languageToSayLanguage(l string) string {
	switch l {
	case "zh-CHS":
		return "zh_CN"
	case "ja":
		return "ja_JP"
	case "EN":
		return "en_US"
	case "ko":
		return "ko_KR"
	case "fr":
		return "fr_FR"
	case "ru":
		return "ru_RU"
	case "pt":
		return "pt_PT"
	case "es":
		return "es_ES"
	default:
		return l
	}
}
