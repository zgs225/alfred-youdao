package main

import (
	"testing"
)

func TestToYoudaoDictUrl(t *testing.T) {
	{
		v := toYoudaoDictUrl("hello world")
		if v != "http://dict.youdao.com/w/hello%20world/#keyfrom=dict2.top" {
			t.Error("toYoudaoDictUrl error 1")
		}
	}

	{
		v := toYoudaoDictUrl("hello world!")
		if v != "http://dict.youdao.com/w/hello%20world%21/#keyfrom=dict2.top" {
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
