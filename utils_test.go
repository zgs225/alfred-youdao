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
