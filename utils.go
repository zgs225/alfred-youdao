package main

import (
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
