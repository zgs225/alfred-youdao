package main

import (
	"os"
	"testing"

	"github.com/patrickmn/go-cache"
	"github.com/zgs225/youdao"
)

func TestQuery(t *testing.T) {
	c := &youdao.Client{
		AppID:     APPID,
		AppSecret: APPSECRET,
	}
	q := "hello"
	a := agentClient{
		Client: c,
		Cache:  cache.New(CACHE_EXPIRES, CACHE_EXPIRES),
	}
	a.Query(q)

	_, ok := a.Cache.Get(q)
	if !ok {
		t.Error("query error")
	}

	a.Cache.SaveFile(".cache")

	c2 := cache.New(CACHE_EXPIRES, CACHE_EXPIRES)
	c2.LoadFile(".cache")
	_, ok = c2.Get(q)
	if !ok {
		t.Error("query error")
	}
	os.Remove(".cache")
}
