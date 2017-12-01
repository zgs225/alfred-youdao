package main

import (
	"testing"
)

func TestToYoudaoDictUrl(t *testing.T) {
	{
		v := toYoudaoDictUrl("hello world")
		if v != "http://dict.youdao.com/search?q=hello+world&keyfrom=fanyi.smartResult" {
			t.Error("toYoudaoDictUrl error 1")
		}
	}

	{
		v := toYoudaoDictUrl("hello world!")
		if v != "http://dict.youdao.com/search?q=hello+world%21&keyfrom=fanyi.smartResult" {
			t.Log(v)
			t.Error("toYoudaoDictUrl error 2")
		}
	}
}

func TestParseArgs(t *testing.T) {
	args := []string{"./alfred-youdao", "zh-CHS=>ja", "你好"}
	q, from, to, lang := parseArgs(args)
	t.Log(q, from, to, lang)
	if lang {
		if q != "你好" || from != "zh-CHS" || to != "ja" {
			t.Error("parseArgs error: value got false")
		}
	} else {
		t.Error("parseArgs error: lang got false")
	}
}
